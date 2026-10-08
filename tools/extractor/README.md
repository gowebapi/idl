# WebIDL Extractor and Normalizer Tool

This tool extracts WebIDL parts from browser specification source files (Bikeshed `.bs`, HTML, or curated `w3c/webref` repositories) and normalizes them for the Go WebAssembly code generator (`webidl-bind`).

## Syncing all specifications from `w3c/webref`

The easiest and most comprehensive way to update all WebIDL files:

```bash
# Clone the curated WebIDL specifications from W3C
git clone --depth 1 -b "@webref/idl@latest" https://github.com/w3c/webref.git /tmp/webref

# Sync and normalize all matching IDL files into idl/idl/
cd tools/extractor
go run main.go -webref-dir /tmp/webref/ed/idl -target-dir ../../idl
```

## Extracting from WHATWG Bikeshed Specifications (DOM)

```bash
git clone --depth 1 https://github.com/whatwg/dom /tmp/dom
go run main.go -pre-idl /tmp/dom/dom.bs -output ../../idl/dom.idl
```

## Extracting from WHATWG HTML Specifications

```bash
git clone --depth 1 https://github.com/whatwg/html /tmp/html
go run main.go -pre-idl /tmp/html/source -output ../../idl/html.idl
```

## Normalizing an existing WebIDL file

```bash
go run main.go -input-idl /path/to/raw.idl -output /path/to/normalized.idl
```

## WebIDL Normalizations Performed

The tool automatically normalizes modern WebIDL syntax constructs into representations compatible with the Go bindings generator:
- In-body `constructor(...)` declarations are converted to interface `[Constructor(...)]` attributes.
- `undefined` return types are converted to `void`.
- `[Exposed=*]` wildcard exposed scopes are converted to `[Exposed=(Window,Worker)]`.
- Renamed extended attributes (`[LegacyUnforgeable]`, `[LegacyNoInterfaceObject]`, `[LegacyNullToEmptyString]`) are mapped to their standard counterparts (`[Unforgeable]`, `[NoInterfaceObject]`, `[TreatNullAs=EmptyString]`).
- HTML reflection attributes not used by Go (`Reflect`, `ReflectURL`, `ReflectRange`) and non-standard number attributes are stripped.
- Extended attributes inside union types (e.g. `(A or [Attr] B)`) are cleaned.
- Unsupported namespaces and hyphenated property names are commented out.
