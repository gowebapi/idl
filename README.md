# WebIDL

This repository contains WebIDL specifications and language transformation metadata used to generate Go WebAssembly bindings (`gowebapi/webapi`).

## Specification Sources

Files in `idl/` are sourced from:
- Curated W3C and WHATWG specifications: <https://github.com/w3c/webref> (formerly `tidoust/reffy-reports`).
- Extracted directly from WHATWG standard sources using `tools/extractor`.

A cross-reference browser: <https://dontcallmedom.github.io/webidlpedia/>

## Updating Specifications

To update specifications using the extractor tool, see [tools/extractor/README.md](tools/extractor/README.md).
