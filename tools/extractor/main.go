// Package main extracts WebIDL definitions from specification documents
// (Bikeshed .bs, HTML, or curated w3c/webref files) and normalizes them
// for the Go WebAssembly bindings generator.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var args struct {
	preIdl    string
	inputIdl  string
	output    string
	patched   string
	raw       bool
	normalize bool
	webrefDir string
	targetDir string
}

func main() {
	if msg := parseArgs(); msg != "" {
		fmt.Fprintln(os.Stderr, "command line error:", msg)
		flag.Usage()
		os.Exit(1)
	}
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func parseArgs() string {
	flag.StringVar(&args.preIdl, "pre-idl", "", "extract <pre class=idl> elements from Bikeshed/HTML specification file")
	flag.StringVar(&args.inputIdl, "input-idl", "", "input raw WebIDL file to normalize")
	flag.StringVar(&args.output, "output", "", "output file for extracted/normalized WebIDL")
	flag.BoolVar(&args.raw, "raw-token", false, "legacy compatibility flag (unused)")
	flag.StringVar(&args.patched, "patched", "", "saving raw extracted document before normalization")
	flag.BoolVar(&args.normalize, "normalize", true, "apply WebIDL syntax normalization for webidlparser compatibility")
	flag.StringVar(&args.webrefDir, "webref-dir", "", "path to w3c/webref repository or ed/idl folder to sync into idl/idl/")
	flag.StringVar(&args.targetDir, "target-dir", "../../idl", "target IDL directory when using -webref-dir")
	flag.Parse()

	if args.webrefDir != "" {
		return ""
	}
	if args.output == "" {
		return "missing output file"
	}
	if args.preIdl == "" && args.inputIdl == "" {
		return "missing input file (-pre-idl or -input-idl or -webref-dir)"
	}
	return ""
}

func run() error {
	if args.webrefDir != "" {
		return syncWebref(args.webrefDir, args.targetDir)
	}

	var content []byte
	var err error

	if args.preIdl != "" {
		content, err = extractIDLFromDoc(args.preIdl)
		if err != nil {
			return fmt.Errorf("failed to extract IDL from %s: %w", args.preIdl, err)
		}
	} else if args.inputIdl != "" {
		content, err = os.ReadFile(args.inputIdl)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", args.inputIdl, err)
		}
	}

	if args.patched != "" {
		if err := os.WriteFile(args.patched, content, 0664); err != nil {
			return fmt.Errorf("failed to write patched output: %w", err)
		}
	}

	if args.normalize {
		content = []byte(normalizeWebIDL(string(content)))
	}

	if err := os.WriteFile(args.output, content, 0664); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("Successfully wrote IDL to %s (%d bytes)\n", args.output, len(content))
	return nil
}

// extractIDLFromDoc parses HTML or Bikeshed specification files and extracts
// all normative WebIDL sections (<pre class="idl"> or <code class="idl">).
func extractIDLFromDoc(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	z := html.NewTokenizer(f)
	var out bytes.Buffer
	var currentTag string
	var inIDL bool
	var saved int

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			err := z.Err()
			if err == io.EOF {
				fmt.Printf("Extracted %d IDL fragments from %s\n", saved, filename)
				return out.Bytes(), nil
			}
			return nil, err

		case html.StartTagToken:
			tn, hasAttr := z.TagName()
			currentTag = string(tn)
			if (currentTag == "pre" || currentTag == "code") && hasAttr {
				isIDL := false
				isExtractOnly := false
				for {
					k, v, more := z.TagAttr()
					if string(k) == "class" {
						for c := range strings.FieldsSeq(string(v)) {
							if c == "idl" {
								isIDL = true
							}
							if c == "extract" {
								isExtractOnly = true
							}
						}
					}
					if !more {
						break
					}
				}
				if isIDL && !isExtractOnly {
					inIDL = true
					saved++
				}
			}

		case html.TextToken:
			if inIDL {
				out.Write(z.Text())
			}

		case html.EndTagToken:
			tn, _ := z.TagName()
			if inIDL && string(tn) == currentTag {
				inIDL = false
				out.WriteString("\n\n// ----------\n\n")
			}
		}
	}
}

// normalizeWebIDL transforms modern WebIDL (post-2019 standards) into
// syntax fully compatible with webidlparser and webidl-bind.
func normalizeWebIDL(content string) string {
	s := content

	// 1. Wildcard exposed: [Exposed=*] -> [Exposed=(Window,Worker)]
	s = strings.ReplaceAll(s, "Exposed=*", "Exposed=(Window,Worker)")

	// 2. Renamed extended attributes
	s = strings.ReplaceAll(s, "LegacyUnforgeable", "Unforgeable")
	s = strings.ReplaceAll(s, "LegacyNoInterfaceObject", "NoInterfaceObject")
	s = strings.ReplaceAll(s, "LegacyNullToEmptyString", "TreatNullAs=EmptyString")
	s = strings.ReplaceAll(s, "LegacyLenientThis", "")

	// 3. Strip HTML reflection attributes not used by Go bindings:
	// Reflect, ReflectURL, ReflectRange=(...), Reflect="..."
	reReflect := regexp.MustCompile(`,\s*Reflect\w*(?:=\([^)]*\)|="[^"]*"|=[^,\]]+)?`)
	s = reReflect.ReplaceAllString(s, "")
	reReflect2 := regexp.MustCompile(`\[\s*Reflect\w*(?:=\([^)]*\)|="[^"]*"|=[^,\]]+)?\s*,?\s*`)
	s = reReflect2.ReplaceAllString(s, "[")
	reEmptyAnnotation := regexp.MustCompile(`(?m)^\s*\[\s*\]\s*\r?\n?|\[\s*\]\s*([a-zA-Z])`)
	s = reEmptyAnnotation.ReplaceAllString(s, "$1")

	// 4. Strip attributes with numbers in value lists, e.g. [..., CustomAttr=(0, 8)]
	reNonIdentAttr := regexp.MustCompile(`,\s*\w+=\([^)]*[0-9][^)]*\)`)
	s = reNonIdentAttr.ReplaceAllString(s, "")
	reNonIdentAttr2 := regexp.MustCompile(`\[\s*\w+=\([^)]*[0-9][^)]*\)\s*,\s*`)
	s = reNonIdentAttr2.ReplaceAllString(s, "[")
	reNonIdentAttrSolo := regexp.MustCompile(`\[\s*\w+=\([^)]*[0-9][^)]*\)\s*\]`)
	s = reNonIdentAttrSolo.ReplaceAllString(s, "")
	s = reEmptyAnnotation.ReplaceAllString(s, "")

	// 5. Strip extended attributes placed inside union types or sequence/typedef types
	reUnionAnnot1 := regexp.MustCompile(`\(([^)]*?)\s*\[[^\]]*\]\s*([^)]*?\bor\b[^)]*?)\)`)
	s = reUnionAnnot1.ReplaceAllString(s, "($1 $2)")
	reUnionAnnot2 := regexp.MustCompile(`\(([^)]*?\bor\b[^)]*?)\s*\[[^\]]*\]\s*([^)]*?)\)`)
	s = reUnionAnnot2.ReplaceAllString(s, "($1 $2)")
	reSeqAnnot := regexp.MustCompile(`\bsequence<\s*\[[^\]]*\]\s*`)
	s = reSeqAnnot.ReplaceAllString(s, "sequence<")
	reTypedefAnnot := regexp.MustCompile(`typedef\s+\[[^\]]*\]\s*`)
	s = reTypedefAnnot.ReplaceAllString(s, "typedef ")

	// 6. Comment out namespace definitions and their annotations (unsupported by webidlparser)
	reNamespace := regexp.MustCompile(`(?s)(?:(\[[^\]]*\])\s*)?(?:partial\s+)?namespace\s+\w+\s*\{[^}]*\};`)
	s = reNamespace.ReplaceAllStringFunc(s, func(m string) string {
		idx := strings.Index(s, m)
		if idx >= 0 {
			lineStart := strings.LastIndex(s[:idx], "\n")
			if lineStart == -1 {
				lineStart = 0
			} else {
				lineStart++
			}
			precedingOnLine := strings.TrimSpace(s[lineStart:idx])
			if strings.HasPrefix(precedingOnLine, "//") {
				return m
			}
			lastOpen := strings.LastIndex(s[:idx], "/*")
			lastClose := strings.LastIndex(s[:idx], "*/")
			if lastOpen > lastClose {
				return m
			}
		}
		lines := strings.Split(m, "\n")
		allCommented := true
		for _, l := range lines {
			t := strings.TrimSpace(l)
			if t != "" && !strings.HasPrefix(t, "//") {
				allCommented = false
				break
			}
		}
		if allCommented {
			return m
		}
		for i, l := range lines {
			if !strings.HasPrefix(strings.TrimSpace(l), "//") {
				lines[i] = "// " + l
			}
		}
		return strings.Join(lines, "\n")
	})

	// 7. Comment out hyphenated attributes (e.g. margin-top in cssom.idl)
	reDashedAttr := regexp.MustCompile(`(?m)^(\s*)(attribute\s+[^;]*?[a-zA-Z]+-[a-zA-Z-]+;.*)$`)
	s = reDashedAttr.ReplaceAllString(s, "${1}// $2")

	// 8. undefined in unions and promises: (Foo or undefined) -> (Foo or any), Promise<undefined> -> Promise<void>
	reUnionUndefined := regexp.MustCompile(`\bor\s+undefined\b`)
	s = reUnionUndefined.ReplaceAllString(s, "or any")
	reUndefinedUnion := regexp.MustCompile(`\bundefined\s+or\b`)
	s = reUndefinedUnion.ReplaceAllString(s, "any or")
	s = strings.ReplaceAll(s, "Promise<undefined>", "Promise<void>")
	rePromiseUnion := regexp.MustCompile(`Promise<\([^)]+\)>`)
	s = rePromiseUnion.ReplaceAllString(s, "Promise<any>")
	s = strings.ReplaceAll(s, "ObservableArray<", "sequence<")

	// 9. inherit attribute -> attribute, [Exposed=...] stringifier; -> stringifier;
	s = strings.ReplaceAll(s, "inherit attribute", "attribute")
	reStringifier := regexp.MustCompile(`\[[^\]]*\]\s*stringifier;`)
	s = reStringifier.ReplaceAllString(s, "stringifier;")

	// Comment out includes for unsupported/external mixins and interfaces not defined in the workspace
	reUnsupportedIncludes := regexp.MustCompile(`(?m)^[ \t]*(?:MathMLElement\s+includes\s+[^;]+;|[^;]+includes\s+GenericTransformStream\s*;).*$`)
	s = reUnsupportedIncludes.ReplaceAllStringFunc(s, func(m string) string {
		trimmed := strings.TrimSpace(m)
		if strings.HasPrefix(trimmed, "//") {
			return m
		}
		return "// " + trimmed
	})

	// 10. record<K, V> unsupported in webidl-bind: strip from unions, replace standalone with object
	reOrRecord := regexp.MustCompile(`\s+or\s+record<[^>]+(?:<[^>]+>[^>]*)?>`)
	s = reOrRecord.ReplaceAllString(s, "")
	reRecordOr := regexp.MustCompile(`\brecord<[^>]+(?:<[^>]+>[^>]*)?>\s+or\s+`)
	s = reRecordOr.ReplaceAllString(s, "")
	for {
		idx := strings.Index(s, "record<")
		if idx == -1 {
			break
		}
		depth := 0
		end := -1
		for i := idx + 7; i < len(s); i++ {
			if s[i] == '<' {
				depth++
			} else if s[i] == '>' {
				if depth == 0 {
					end = i
					break
				}
				depth--
			}
		}
		if end == -1 {
			break
		}
		s = s[:idx] + "object" + s[end+1:]
	}

	s = strings.ReplaceAll(s, "async_iterable<", "iterable<")
	s = strings.ReplaceAll(s, "async iterable<", "iterable<")
	s = strings.ReplaceAll(s, "async_sequence<", "sequence<")

	// 11. undefined return types -> void (including setters, deleters, and static functions)
	reUndefinedReturn := regexp.MustCompile(`(?m)^(\s*(?:\[[^\]]*\]\s*)*(?:static\s+)?(?:(?:getter|setter|deleter)\s+)?)undefined\b`)
	s = reUndefinedReturn.ReplaceAllString(s, `${1}void`)
	reCallbackUndefined := regexp.MustCompile(`(?m)(callback\s+\w+\s*=\s*)undefined\b`)
	s = reCallbackUndefined.ReplaceAllString(s, `${1}void`)

	// 12. Modern types and unknown types:
	// bigint -> long long, Float16Array -> Float32Array, AllowSharedBufferSource -> BufferSource
	// TrustedHTML/Script/ScriptURL -> DOMString
	reBigInt := regexp.MustCompile(`\bbigint\b`)
	s = reBigInt.ReplaceAllString(s, "long long")
	reFloat16 := regexp.MustCompile(`\bFloat16Array\b`)
	s = reFloat16.ReplaceAllString(s, "Float32Array")
	reAllowShared := regexp.MustCompile(`\bAllowSharedBufferSource\b`)
	s = reAllowShared.ReplaceAllString(s, "BufferSource")
	reTrusted := regexp.MustCompile(`\b(?:TrustedHTML|TrustedScript|TrustedScriptURL|TrustedType)\b`)
	s = reTrusted.ReplaceAllString(s, "DOMString")

	// 13. Fix PaymentAddress -> ContactAddress in basic card
	s = strings.ReplaceAll(s, "PaymentAddress", "ContactAddress")

	// 14. Strip dictionary parameter default values unsupported by parser (= {})
	reEmptyDict := regexp.MustCompile(`=\s*\{\}`)
	s = reEmptyDict.ReplaceAllString(s, "")

	// 15. Transform in-body constructor(...) into [Constructor(...)] interface attributes
	s = transformConstructors(s)

	// 16. Permissions Policy ReportBody: ReportBody is an interface in reporting.idl, so violation body must be an interface
	if strings.Contains(s, "PermissionsPolicyViolationReportBody : ReportBody") {
		s = strings.ReplaceAll(s, "dictionary PermissionsPolicyViolationReportBody : ReportBody {", "interface PermissionsPolicyViolationReportBody : ReportBody {")
		reReportBody := regexp.MustCompile(`(?s)interface PermissionsPolicyViolationReportBody : ReportBody\s*\{([^}]+)\}`)
		s = reReportBody.ReplaceAllStringFunc(s, func(m string) string {
			lines := strings.Split(m, "\n")
			for i, l := range lines {
				trimmed := strings.TrimSpace(l)
				if trimmed == "" || strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "interface") || strings.HasPrefix(trimmed, "}") {
					continue
				}
				if !strings.HasPrefix(trimmed, "readonly attribute") {
					lines[i] = strings.Replace(l, trimmed, "readonly attribute "+trimmed, 1)
				}
			}
			return strings.Join(lines, "\n")
		})
	}

	// 17. Comment out duplicate AddressInit in payment handler (already defined in payment-request.idl)
	if strings.Contains(s, "PaymentRequestEvent") && strings.Contains(s, "dictionary AddressInit") {
		reDupAddressInit := regexp.MustCompile(`(?s)dictionary AddressInit\s*\{[^}]*\};`)
		s = reDupAddressInit.ReplaceAllString(s, "// AddressInit already defined in payment-request.idl")
	}

	// 18. Geolocation backward-compatibility aliases and EpochTimeStamp
	if strings.Contains(s, "interface GeolocationCoordinates") && !strings.Contains(s, "typedef GeolocationCoordinates Coordinates;") {
		s += "\n\ntypedef GeolocationCoordinates Coordinates;\ntypedef GeolocationPosition Position;\ntypedef GeolocationPositionError PositionError;\ntypedef unsigned long long EpochTimeStamp;\n"
	}

	// 19. WebCrypto enums (moved to webcrypto-modern-algos upstream)
	if strings.Contains(s, "interface SubtleCrypto") && !strings.Contains(s, "enum KeyUsage") {
		s = "enum KeyUsage { \"encrypt\", \"decrypt\", \"sign\", \"verify\", \"deriveKey\", \"deriveBits\", \"wrapKey\", \"unwrapKey\" };\n\n" +
			"enum KeyFormat { \"raw\", \"spki\", \"pkcs8\", \"jwk\" };\n\n" + s
	}

	// 20. WebGPU and WebCodecs enums (PredefinedColorSpace, BitrateMode)
	if strings.Contains(s, "interface GPUCanvasContext") && !strings.Contains(s, "enum PredefinedColorSpace") {
		s = "enum PredefinedColorSpace { \"srgb\", \"display-p3\" };\n\n" + s
	}
	if strings.Contains(s, "interface VideoEncoder") && !strings.Contains(s, "enum BitrateMode") {
		s = "enum BitrateMode { \"constant\", \"variable\" };\n\n" + s
	}

	return s
}

// transformConstructors moves in-body constructor(...) declarations into [Constructor(...)] annotations.
func transformConstructors(s string) string {
	reInterface := regexp.MustCompile(`(?s)(?:(\[[^\]]*\])\s*)?((?:partial\s+)?(?:callback\s+)?interface\s+(?:mixin\s+)?\w+(?:\s*:[^{]+)?)\s*\{([^}]*)\};`)
	reConstructor := regexp.MustCompile(`(?m)^\s*(\[HTMLConstructor\]\s*)?constructor\s*\(([\s\S]*?)\)\s*;[^\S\r\n]*(?:\r?\n)?`)

	return reInterface.ReplaceAllStringFunc(s, func(m string) string {
		sub := reInterface.FindStringSubmatch(m)
		if len(sub) < 4 {
			return m
		}
		annotations := strings.TrimSpace(sub[1])
		header := strings.TrimSpace(sub[2])
		body := sub[3]

		matches := reConstructor.FindAllStringSubmatch(body, -1)
		if len(matches) == 0 {
			return m
		}

		var constructors []string
		for _, match := range matches {
			isHTML := match[1] != ""
			params := strings.Join(strings.Fields(strings.TrimSpace(match[2])), " ")
			if isHTML {
				constructors = append(constructors, "HTMLConstructor")
			} else if params == "" {
				constructors = append(constructors, "Constructor")
			} else {
				constructors = append(constructors, fmt.Sprintf("Constructor(%s)", params))
			}
		}

		cleanedBody := reConstructor.ReplaceAllString(body, "")

		var newAnnotations string
		constructorAnnotation := strings.Join(constructors, ",\n ")
		if annotations != "" {
			inner := strings.TrimSpace(annotations[1 : len(annotations)-1])
			newAnnotations = fmt.Sprintf("[%s,\n %s]", constructorAnnotation, inner)
		} else {
			newAnnotations = fmt.Sprintf("[%s]", constructorAnnotation)
		}

		return fmt.Sprintf("%s\n%s {\n%s};", newAnnotations, header, cleanedBody)
	})
}

// syncWebref synchronizes and normalizes specifications from a w3c/webref repository clone.
func syncWebref(webrefDir, targetDir string) error {
	idlSrcDir := filepath.Join(webrefDir, "ed", "idl")
	if _, err := os.Stat(idlSrcDir); err != nil {
		// Maybe webrefDir was already pointing to ed/idl
		idlSrcDir = webrefDir
	}

	targetFiles, err := filepath.Glob(filepath.Join(targetDir, "*.idl"))
	if err != nil {
		return err
	}

	var synced, skipped int
	for _, tf := range targetFiles {
		base := filepath.Base(tf)
		// Skip local patches and additions
		if strings.HasSuffix(base, ".patch.idl") || strings.HasSuffix(base, ".addition.idl") {
			continue
		}

		srcPath := filepath.Join(idlSrcDir, base)
		if _, err := os.Stat(srcPath); err != nil {
			// Case-insensitive fallback
			entries, _ := os.ReadDir(idlSrcDir)
			for _, entry := range entries {
				if strings.EqualFold(entry.Name(), base) {
					srcPath = filepath.Join(idlSrcDir, entry.Name())
					break
				}
			}
			if _, err := os.Stat(srcPath); err != nil {
				// Normalize existing file in-place if not found upstream
				if raw, err := os.ReadFile(tf); err == nil {
					normalized := normalizeWebIDL(string(raw))
					_ = os.WriteFile(tf, []byte(normalized), 0664)
				}
				skipped++
				continue
			}
		}

		raw, err := os.ReadFile(srcPath)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", srcPath, err)
		}

		normalized := normalizeWebIDL(string(raw))
		if err := os.WriteFile(tf, []byte(normalized), 0664); err != nil {
			return fmt.Errorf("failed to write %s: %w", tf, err)
		}
		synced++
	}

	fmt.Printf("Sync complete: updated %d IDL files in %s (skipped %d unmatched/patch files)\n", synced, targetDir, skipped)
	return nil
}
