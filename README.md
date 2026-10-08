# WebIDL

This repository contains WebIDL browser specifications and language transformation metadata used by [`webidl-bind`](file:///home/janpf/Projects/gowebapi/webidl-bind/) to generate Go WebAssembly bindings for [`gowebapi/webapi`](file:///home/janpf/Projects/gowebapi/webapi/).

## Repository Layout

- [`idl/`](file:///home/janpf/Projects/gowebapi/idl/idl/): Normalized WebIDL specification files (`*.idl`), including supplementary patches (`*.addition.idl`, `*.patch.idl`).
- [`webapi/`](file:///home/janpf/Projects/gowebapi/idl/webapi/): Go language transformation files (`*.go.md`) mapping WebIDL types to Go packages, names, and event handlers.
- [`tools/extractor/`](file:///home/janpf/Projects/gowebapi/idl/tools/extractor/): Go CLI tool to extract and normalize WebIDL from upstream standards.

---

## How to Run the Extraction

WebIDL specifications are primarily sourced from the W3C curated repository [`w3c/webref`](https://github.com/w3c/webref) or directly from WHATWG Bikeshed (`.bs`) and HTML source repositories.

All raw WebIDL files must pass through the normalization tool [`tools/extractor`](file:///home/janpf/Projects/gowebapi/idl/tools/extractor/main.go) to ensure compatibility with [`webidl-bind`](file:///home/janpf/Projects/gowebapi/webidl-bind/) (which includes the integrated `ast` and `parser` packages).

### A. Bulk Sync from `w3c/webref` (Recommended)

To sync all standard WebIDL specifications:

```bash
# 1. Clone the latest curated WebIDL release from W3C
git clone --depth 1 -b "@webref/idl@latest" https://github.com/w3c/webref.git /tmp/webref

# 2. Run the extractor tool to normalize and sync matching IDL files
cd tools/extractor
go run main.go -webref-dir /tmp/webref/ed/idl -target-dir ../../idl
```

### B. Extracting Directly from WHATWG Specification Sources

For specifications maintained in Bikeshed (`.bs`) or HTML format:

```bash
# WHATWG DOM
git clone --depth 1 https://github.com/whatwg/dom /tmp/dom
cd tools/extractor
go run main.go -pre-idl /tmp/dom/dom.bs -output ../../idl/dom.idl

# WHATWG HTML
git clone --depth 1 https://github.com/whatwg/html /tmp/html
cd tools/extractor
go run main.go -pre-idl /tmp/html/source -output ../../idl/html.idl
```

### C. Normalizing an Individual WebIDL File

To normalize an arbitrary or newly downloaded `.idl` file:

```bash
cd tools/extractor
go run main.go -input-idl /path/to/raw.idl -output ../../idl/<spec-name>.idl
```

### D. Normalizations Performed by the Extractor

The extractor tool automatically transforms modern WebIDL constructs into formats supported by [`webidl-bind`](file:///home/janpf/Projects/gowebapi/webidl-bind/):
1. **Constructors**: Converts in-body `constructor(...)` declarations into `[Constructor(...)]` interface annotations.
2. **Return Types**: Converts `undefined` returns to `void`.
3. **Exposed Annotations**: Converts `[Exposed=*]` wildcard attributes to `[Exposed=(Window,Worker)]`.
4. **Extended Attributes**: Renames legacy annotations (`[LegacyUnforgeable]`, `[LegacyNoInterfaceObject]`, `[LegacyNullToEmptyString]`) to their standard equivalents (`[Unforgeable]`, `[NoInterfaceObject]`, `[TreatNullAs=EmptyString]`).
5. **Reflection & Tags**: Strips non-standard HTML reflection attributes (`Reflect`, `ReflectURL`, `ReflectRange`) and HTML element tag annotations.
6. **Types & Collections**:
   - Replaces `record<K, V>` with `object` (or removes from union types).
   - Maps `async_iterable<...>` and `async iterable<...>` to `iterable<...>`.
   - Maps `async_sequence<...>` and `ObservableArray<...>` to `sequence<...>`.
   - Normalizes union return promises `Promise<(A or B)>` to `Promise<any>`.
   - Strips extended attributes inside union types and sequence/typedef definitions (e.g. `sequence<[EnforceRange] long>`).
7. **Namespaces & Unsupported Constructs**: Comments out `namespace` blocks and hyphenated attributes (e.g., CSS properties in `cssom.idl`.

---

## Change Log

### 2026-10 

This branch modernizes the IDL extraction pipeline (go1.27), synchronizes upstream specifications to the latest 
W3C/WHATWG releases, and adds modern Web specifications standardized since 2019.

- Integrated the WebIDL AST and parser packages directly into [`webidl-bind`](file:///home/janpf/Projects/gowebapi/webidl-bind/) (`github.com/gowebapi/webidl-bind/ast` and `github.com/gowebapi/webidl-bind/parser`) replacing the external `github.com/gowebapi/webidlparser` dependency, and modernized them with `modernize -fix`.
- Removed broken test `Dockerfile`.

#### Specification Renames & Consolidations

Several upstream W3C/WHATWG specifications changed file names or identifiers since 2019. These have been updated along with their corresponding `.go.md` mapping rules:

| Old Specification File | Updated Specification File | Go Package |
| :--- | :--- | :--- |
| `BackgroundSync.idl` | [`background-sync.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/background-sync.idl) | `serviceworker/backgroundsync` |
| `cookie-store.idl` | [`cookiestore.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/cookiestore.idl) | `cookie` |
| `geolocation-API.idl` | [`geolocation.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/geolocation.idl) | `device/sensor` |
| `InputDeviceCapabilities.idl` | [`input-device-capabilities.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/input-device-capabilities.idl) | `device/inputcapabilities` |
| `payment-handler.idl` | [`web-based-payment-handler.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/web-based-payment-handler.idl) | `payment/webbasedpaymenthandler` |
| `ResizeObserver.idl` | [`resize-observer.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/resize-observer.idl) | `css/resizeobserver` |
| `wake-lock.idl` | [`screen-wake-lock.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/screen-wake-lock.idl) | `device/wakelock` |
| `webappsec-feature-policy.idl` | [`permissions-policy.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/permissions-policy.idl) | `featurepolicy` |
| `WebCryptoAPI.idl` | [`webcrypto.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/webcrypto.idl) | `crypto` |

#### New Web APIs Added

24 modern browser specifications have been added to [`idl/idl/`](file:///home/janpf/Projects/gowebapi/idl/idl/) with corresponding Go language transformation definitions in [`idl/webapi/`](file:///home/janpf/Projects/gowebapi/idl/webapi/):

| Specification IDL | Target Go Package | Description |
| :--- | :--- | :--- |
| [`webgpu.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/webgpu.idl) | `graphics/webgpu` | WebGPU API (`GPUAdapter`, `GPUDevice`, `GPUBuffer`, `GPURenderPipeline`, etc.) |
| [`webnn.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/webnn.idl) | `ml/webnn` | Web Neural Network API (`MLContext`, `MLGraph`, `MLOperand`, `MLTensor`) |
| [`webcodecs.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/webcodecs.idl) | `media/webcodecs` | WebCodecs API (`VideoEncoder`, `VideoDecoder`, `AudioEncoder`, `AudioDecoder`, `VideoFrame`) |
| [`fs.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/fs.idl) | `file/fs` | WHATWG File System standard (`FileSystemHandle`, `FileSystemDirectoryHandle`, `FileSystemFileHandle`) |
| [`file-system-access.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/file-system-access.idl) | `file/fs` | W3C File System Access API (`showOpenFilePicker`, `showSaveFilePicker`, `showDirectoryPicker`) |
| [`web-locks.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/web-locks.idl) | `storage/weblocks` | Web Locks API (`LockManager`, `Lock`) |
| [`compression.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/compression.idl) | `compression` | Compression Streams API (`CompressionStream`, `DecompressionStream`) |
| [`storage-buckets.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/storage-buckets.idl) | `storage/buckets` | Storage Buckets API (`StorageBucketManager`, `StorageBucket`) |
| [`storage-access.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/storage-access.idl) | `storage/storageaccess` | Storage Access API (`hasStorageAccess`, `requestStorageAccess`) |
| [`badging.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/badging.idl) | `badging` | App Badging API (`setAppBadge`, `clearAppBadge`) |
| [`compute-pressure.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/compute-pressure.idl) | `device/computepressure` | Compute Pressure API (`PressureObserver`, `PressureRecord`) |
| [`serial.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/serial.idl) | `device/serial` | Web Serial API (`Serial`, `SerialPort`) |
| [`idle-detection.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/idle-detection.idl) | `device/idledetection` | Idle Detection API (`IdleDetector`) |
| [`virtual-keyboard.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/virtual-keyboard.idl) | `device/virtualkeyboard` | VirtualKeyboard API (`VirtualKeyboard`) |
| [`eyedropper-api.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/eyedropper-api.idl) | `dom/eyedropper` | EyeDropper API (`EyeDropper`) |
| [`css-anchor-position.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/css-anchor-position.idl) | `css/anchorposition` | CSS Anchor Positioning Module Level 1 |
| [`css-highlight-api.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/css-highlight-api.idl) | `css/highlight` | CSS Custom Highlight API (`HighlightRegistry`, `Highlight`) |
| [`css-view-transitions.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/css-view-transitions.idl) | `css/viewtransitions` | CSS View Transitions Levels 1 & 2 (`ViewTransition`, `ViewTransitionTypeSet`) |
| [`document-picture-in-picture.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/document-picture-in-picture.idl) | `webapi` | Document Picture-in-Picture API (`DocumentPictureInPicture`) |
| [`urlpattern.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/urlpattern.idl) | `url/urlpattern` | URL Pattern API (`URLPattern`) |
| [`window-management.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/window-management.idl) | `window/management` | Window Management API (`ScreenDetails`, `ScreenDetailed`) |
| [`digital-credentials.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/digital-credentials.idl) | `crypto/credential` | Digital Credentials API (`DigitalCredential`) |
| [`fedcm.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/fedcm.idl) | `crypto/credential` | Federated Credential Management API (`IdentityCredential`) |
| [`layout-instability.idl`](file:///home/janpf/Projects/gowebapi/idl/idl/layout-instability.idl) | `performance/layoutinstability` | Layout Instability API / CLS (`LayoutShift`, `LayoutShiftAttribution`) |