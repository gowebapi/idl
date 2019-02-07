// this will take different documentations and is trying to
// extract their webidl parts
package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

var args struct {
	preIdl  string
	output  string
	raw     bool
	patched string
}

type patchMode int

const (
	scanningMode patchMode = iota
	tagNameMode
	attrWaitName
	attrInNameMode
	attrValueStartMode
	attrValueStringMode
	attrValuePatchMode
	commentMode
)

func main() {
	if msg := parseArgs(); msg != "" {
		fmt.Fprintln(os.Stderr, "command line error:", msg)
		os.Exit(1)
	}
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	content, err := ioutil.ReadFile(args.preIdl)
	if err != nil {
		return err
	}
	if content, err = patchAttributes(content); err != nil {
		return err
	}
	if args.patched != "" {
		if err = ioutil.WriteFile(args.patched, content, 0666); err != nil {
			return err
		}
	}
	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.Strict = false
	save := false
	saved, candidate := 0, 0
	endTag := ""
	var out bytes.Buffer
	for {
		token, err := nextToken(decoder)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		// fmt.Printf("DEBUG: %#v\n", token)
		switch t := token.(type) {
		case xml.StartElement:
			name := t.Name.Local
			if haveEqualAttr("class", "idl", t.Attr) {
				candidate++
			}
			if !save && (name == "pre" || name == "code") {
				// if len(t.Attr) > 0 {
				// 	fmt.Println("maybe: ", t.Attr)
				// }
				if haveEqualAttr("class", "idl", t.Attr) {
					save = true
					endTag = name
					saved++
				}
			}
		case xml.CharData:
			if save == true {
				out.Write(t)
			}
		case xml.EndElement:
			if t.Name.Local == endTag && save {
				save = false
				out.WriteString("\n\n// ----------\n\n")
			}
		case xml.Comment:
		case xml.Directive:
		default:
			return fmt.Errorf("internal error: unknown xml decoder type not handled in code: %#v", t)
		}
	}
	if err = ioutil.WriteFile(args.output, out.Bytes(), 0664); err != nil {
		return err
	}
	fmt.Println("saved", saved, "idl code pieces from spec", candidate)
	return nil
}

func haveEqualAttr(name, value string, attr []xml.Attr) bool {
	for _, a := range attr {
		if strings.ToLower(a.Name.Local) == name && strings.ToLower(a.Value) == value {
			return true
		}
	}
	return false
}

func nextToken(d *xml.Decoder) (xml.Token, error) {
	if args.raw {
		return d.RawToken()
	}
	return d.Token()
}

// patchAttributes is going over the document and add '"' to
// attributes inside xml nodes
func patchAttributes(in []byte) ([]byte, error) {
	// if in review, please ignore this hard coded hack...
	in = bytes.Replace(in, []byte("<code>]]></code>"), []byte("<code>]]&lt;</code>"), -1)
	in = bytes.Replace(in, []byte("<code data-x=\"\">]]></code>"), []byte("<code data-x=\"\">]]&lt;</code>"), -1)
	in = bytes.Replace(in, []byte("&lt;![CDATA[x&lt;y]]>"), []byte("&lt;![CDATA[x&lt;y]]&gt;"), -1)
	in = bytes.Replace(in, []byte("&lt;![CDATA[x&lt;y3]]>"), []byte("&lt;![CDATA[x&lt;y3]]&gt;"), -1)

	// process the remaning of the document
	var out bytes.Buffer
	var endQuote rune
	mode := scanningMode
	commentStart := []byte("<!--")
	commentEnd := []byte("-->")
	for i := 0; i < len(in); {
		ch, width := utf8.DecodeRune(in[i:])
		if ch == utf8.RuneError {
			return nil, fmt.Errorf("found invalid utf8 rune, what now?")
		}
		i += width
		// if mode != scanningMode && isSpace(ch) {
		// 	ch = ' '
		// }
		if (ch == '>' || ch == '/') && mode != scanningMode && mode != attrValueStringMode {
			// common closing of tag
			if mode == attrValuePatchMode {
				out.WriteRune('"')
			} else if mode == attrInNameMode {
				out.WriteString("=\"true\"")
			}
			mode = scanningMode
			out.WriteRune(ch)
			continue
		}
		switch mode {
		case scanningMode:
			if bytes.HasPrefix(in[(i-1):], commentStart) {
				mode = commentMode
			} else if ch == '<' {
				if i < len(in) && in[i] == '{' {
					// some other weird syntax of something ...e.g. <{area}>
					out.WriteString("&gt;")
					continue
				}
				mode = tagNameMode
			}
		case tagNameMode:
			if isSpace(ch) {
				mode = attrWaitName
			}
		case attrWaitName:
			if ch == '=' {
				out.WriteString(" foo")
				mode = attrValueStartMode
			} else if !isSpace(ch) {
				mode = attrInNameMode
			}
		case attrInNameMode:
			if ch == '=' {
				mode = attrValueStartMode
			} else if isSpace(ch) {
				// fixed unexpected attribute end
				out.WriteString("=\"true\"")
				mode = attrWaitName
			}
		case attrValueStartMode:
			if ch == '"' || ch == '\'' {
				mode = attrValueStringMode
				endQuote = ch
				ch = '"'
			} else {
				out.WriteRune('"')
				mode = attrValuePatchMode
			}
		case attrValueStringMode:
			if ch == endQuote {
				mode = attrWaitName
				ch = '"'
			} else if ch == '<' || ch == '>' {
				ch = '@'
			}
		case attrValuePatchMode:
			if isSpace(ch) {
				out.WriteRune('"')
				mode = attrWaitName
			}
		case commentMode:
			if bytes.HasPrefix(in[i:], commentEnd) {
				// i += len(commentEnd)
				mode = scanningMode
			}
		}
		if _, err := out.WriteRune(ch); err != nil {
			return nil, fmt.Errorf("patch failture: %s", err)
		}
	}
	// fmt.Println(string(out.Bytes()[0:300]))
	return out.Bytes(), nil
}

func isSpace(ch rune) bool {
	return unicode.IsSpace(ch) || ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func parseArgs() string {
	flag.StringVar(&args.preIdl, "pre-idl", "", "extract <pre class=idl> elements as webidl content")
	flag.StringVar(&args.output, "output", "", "output file")
	flag.BoolVar(&args.raw, "raw-token", false, "use RawToken() instead of Token() on xml decoder")
	flag.StringVar(&args.patched, "patched", "", "saving patched input document")
	flag.Parse()
	if args.output == "" {
		return "missing output file"
	}
	if args.preIdl == "" {
		return "missing input file"
	}
	return ""
}
