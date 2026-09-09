package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNormalizeWebIDL(t *testing.T) {
	input := `[Exposed=*]
interface Example {
  constructor(DOMString name, optional EventInit dict = {});
  undefined doSomething();
  [LegacyUnforgeable] readonly attribute boolean isCool;
  [CEReactions, Reflect="data-foo", ReflectRange=(0, 8)] attribute DOMString foo;
  attribute (DOMString or [Annotation] long) bar;
  attribute CSSOMString margin-top;
};`

	normalized := normalizeWebIDL(input)

	if strings.Contains(normalized, "Exposed=*") {
		t.Errorf("Expected Exposed=* to be normalized, got: %s", normalized)
	}
	if !strings.Contains(normalized, "Constructor(DOMString name, optional EventInit dict)") {
		t.Errorf("Expected constructor to be converted to [Constructor(...)], got: %s", normalized)
	}
	if strings.Contains(normalized, "undefined doSomething") {
		t.Errorf("Expected undefined return to be converted to void, got: %s", normalized)
	}
	if !strings.Contains(normalized, "void doSomething") {
		t.Errorf("Expected void doSomething, got: %s", normalized)
	}
	if strings.Contains(normalized, "LegacyUnforgeable") {
		t.Errorf("Expected LegacyUnforgeable to be converted, got: %s", normalized)
	}
	if !strings.Contains(normalized, "Unforgeable") {
		t.Errorf("Expected Unforgeable, got: %s", normalized)
	}
	if strings.Contains(normalized, "Reflect") {
		t.Errorf("Expected Reflect annotations to be stripped, got: %s", normalized)
	}
	if strings.Contains(normalized, "[Annotation]") {
		t.Errorf("Expected annotation in union to be stripped, got: %s", normalized)
	}
	if !strings.Contains(normalized, "// attribute CSSOMString margin-top;") {
		t.Errorf("Expected hyphenated attribute to be commented out, got: %s", normalized)
	}
}

func TestExtractIDLFromDoc(t *testing.T) {
	tmpDir := t.TempDir()
	sampleDoc := filepath.Join(tmpDir, "spec.html")
	content := `<!DOCTYPE html>
<html>
<body>
<p>Some text with <a for=/>unclosed or void tag</a></p>
<pre class="idl">
[Exposed=Window]
interface TestDoc {
  readonly attribute DOMString name;
};
</pre>
<pre class="idl extract">
interface ExampleExtractOnly {};
</pre>
</body>
</html>`

	if err := os.WriteFile(sampleDoc, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	extracted, err := extractIDLFromDoc(sampleDoc)
	if err != nil {
		t.Fatalf("extractIDLFromDoc failed: %v", err)
	}

	s := string(extracted)
	if !strings.Contains(s, "interface TestDoc") {
		t.Errorf("Expected interface TestDoc in extracted output, got: %s", s)
	}
	if strings.Contains(s, "ExampleExtractOnly") {
		t.Errorf("Did not expect extract-only example to be included, got: %s", s)
	}
}
