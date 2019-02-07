# Extract IDL content from specification files

## DOM

To extract from DOM

```bash
git clone https://github.com/whatwg/dom
go run main.go -pre-idl dom/dom.bs -output dom.idl -patched dom.xml
```

## HTML

To extract from HTML

```bash
git clone https://github.com/whatwg/html
go run main.go -pre-idl html/source -output html.idl -patched html.xml
```
