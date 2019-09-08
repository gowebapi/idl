# HTML Standard

.title = HTML Standard
.url = <https://html.spec.whatwg.org/>
.package = github.com/gowebapi/webapi/html/htmlmisc
@on "HTML." : .package = github.com/gowebapi/webapi/html

## ApplicationCache

@patch idlconst
.constSuffix = ""
@event Cached Event cancelable: false, bubbles: false
@event Checking Event cancelable: false, bubbles: false
@event Downloading Event cancelable: false, bubbles: false
@event Error Event cancelable: false, bubbles: false
@event Noupdate Event cancelable: false, bubbles: false
@event Obsolete Event cancelable: false, bubbles: false
@event Progress ProgressEvent cancelable: false, bubbles: false
@event Updateready Event cancelable: false, bubbles: false

## BlobCallback

.package = github.com/gowebapi/webapi/file

## DocumentAndElementEventHandlers

@event Copy ClipboardEvent cancelable: true, bubbles: true
@event Cut ClipboardEvent cancelable: true, bubbles: true
@event Paste ClipboardEvent cancelable: true, bubbles: true

## EventSource

@patch idlconst
.constSuffix = ""
@event Error ErrorEvent cancelable: false, bubbles: false
@event Message MessageEvent cancelable: false, bubbles: false
@event Open Event cancelable: false, bubbles: false

## GlobalEventHandlers

@event Abort Event maybe: UIEvent, cancelable: false, bubbles: false
@event Auxclick MouseEvent cancelable: true, bubbles: true
@event Blur FocusEvent cancelable: false, bubbles: true
@event Cancel Event
@event CanPlay Event cancelable: false, bubbles: false
@event CanPlayThrough Event cancelable: false, bubbles: false
@event Change Event cancelable: false, bubbles: true
@event Click MouseEvent cancelable: true, bubbles: true
@event Close Event cancelable: false, bubbles: false
@event ContextMenu MouseEvent cancelable: true, bubbles: true
@event CueChange Event
@event DblClick MouseEvent cancelable: true, bubbles: true
@event Drag DragEvent cancelable: true, bubbles: true
@event DragEnd DragEvent cancelable: false, bubbles: true
@event DragEnter DragEvent cancelable: true, bubbles: true
@event DragExit DragEvent cancelable: false, bubbles: true
@event DragLeave DragEvent cancelable: false, bubbles: true
@event DragOver DragEvent cancelable: true, bubbles: true
@event DragStart DragEvent cancelable: true, bubbles: true
@event Drop DragEvent cancelable: true, bubbles: true
@event DurationChange Event cancelable: false, bubbles: false
@event Emptied Event cancelable: false, bubbles: false
@event Ended Event cancelable: false, bubbles: false
@event Error Event maybe: UIEvent
@event Focus FocusEvent cancelable: false, bubbles: true
@event FormData FormDataEvent cancelable: false, bubbles: true
@event Input InputEvent cancelable: false, bubbles: true
@event Invalid Event cancelable: true, bubbles: false
@event KeyDown KeyboardEvent cancelable: true, bubbles: true
@event KeyPress KeyboardEvent cancelable: true, bubbles: true
@event KeyUp KeyboardEvent cancelable: true, bubbles: true
@event Load Event maybe: UIEvent
@event LoadedData Event cancelable: false, bubbles: false
@event LoadedMetaData Event cancelable: false, bubbles: false
@event LoadEnd ProgressEvent cancelable: false, bubbles: false
@event LoadStart Event cancelable: false, bubbles: false
@event MouseDown MouseEvent cancelable: true, bubbles: true
@event MouseEnter MouseEvent cancelable: false, bubbles: false
@event MouseLeave MouseEvent cancelable: false, bubbles: false
@event MouseMove MouseEvent cancelable: true, bubbles: true
@event MouseOut MouseEvent cancelable: true, bubbles: true
@event MouseOver MouseEvent cancelable: true, bubbles: true
@event MouseUp MouseEvent cancelable: true, bubbles: true
@event Wheel WheelEvent cancelable: true, bubbles: true
@event Pause Event cancelable: false, bubbles: false
@event Play Event cancelable: false, bubbles: false
@event Playing Event cancelable: false, bubbles: false
@event Progress ProgressEvent cancelable: false, bubbles: false
@event RateChange Event cancelable: false, bubbles: false
@event Reset Event cancelable: true, bubbles: true
@event Resize UIEvent cancelable: false, bubbles: false
@event Scroll Event maybe: UIEvent, cancelable: false, bubbles: true
@event SecurityPolicyViolation SecurityPolicyViolationEvent cancelable: false, bubbles: true
@event Seeked Event cancelable: false, bubbles: false
@event Seeking Event cancelable: false, bubbles: false
@event Select Event maybe: UIEvent, cancelable: false, bubbles: true
@event Stalled Event cancelable: false, bubbles: false
@event Submit Event cancelable: true, bubbles: true
@event Suspend Event cancelable: false, bubbles: false
@event TimeUpdate Event cancelable: false, bubbles: false
@event Toggle Event cancelable: false, bubbles: false
@event VolumeChange Event cancelable: false, bubbles: false
@event Waiting Event cancelable: false, bubbles: false

    css-animations.idl
@event AnimationCancel AnimationEvent cancelable: false, bubbles: true
@event AnimationEnd AnimationEvent cancelable: false, bubbles: true
@event AnimationIteration AnimationEvent cancelable: false, bubbles: true
@event AnimationStart AnimationEvent cancelable: false, bubbles: true

    pointerevents.idl
@event GotPointerCapture PointerEvent cancelable: false, bubbles: false
@event LostPointerCapture PointerEvent cancelable: false, bubbles: false
@event PointerCancel PointerEvent cancelable: false, bubbles: true
@event PointerDown PointerEvent cancelable: true, bubbles: true
@event PointerEnter PointerEvent cancelable: false, bubbles: false
@event PointerLeave PointerEvent cancelable: false, bubbles: false
@event PointerMove PointerEvent cancelable: true, bubbles: true
@event PointerOut PointerEvent cancelable: true, bubbles: true
@event PointerOver PointerEvent cancelable: true, bubbles: true
@event PointerUp PointerEvent cancelable: true, bubbles: true

    selection-api.idl
@event SelectionChange Event cancelable: false, bubbles: false
@event SelectStart Event cancelable: true, bubbles: true

    touch-events.idl
@event TouchCancel TouchEvent cancelable: false, bubbles: true
@event TouchEnd TouchEvent cancelable: true, bubbles: true
@event TouchMove TouchEvent cancelable: true, bubbles: true
@event TouchStart TouchEvent cancelable: true, bubbles: true

    css-transitions.idl
@event TransitionCancel TransitionEvent cancelable: false, bubbles: true
@event TransitionEnd TransitionEvent cancelable: true, bubbles: true
@event TransitionRun TransitionEvent cancelable: true, bubbles: true
@event TransitionStart TransitionEvent cancelable: false, bubbles: true

## HTMLBodyElement

    compat.idl
@event OrientationChange Event cancelable: false, bubbles: false

## HTMLMarqueeElement

@event Bounce Event cancelable: false, bubbles: false
@event Finish Event cancelable: false, bubbles: false
@event Start Event cancelable: false, bubbles: false

## Navigator

@notevent onLine

## TimeRanges

.package = github.com/gowebapi/webapi/html

## FocusOptions

.package = github.com/gowebapi/webapi/html

## ValidityState

.package = github.com/gowebapi/webapi/html

## SelectionMode

.package = github.com/gowebapi/webapi/html

## AssignedNodesOptions

.package = github.com/gowebapi/webapi/html

## OffscreenRenderingContextId

.package = github.com/gowebapi/webapi/html

## ImageEncodeOptions

.package = github.com/gowebapi/webapi/html

## WindowEventHandlers

@event AfterPrint Event cancelable: false, bubbles: false
@event BeforePrint Event cancelable: false, bubbles: false
@event BeforeUnload BeforeUnloadEvent cancelable: true, bubbles: false
@event HashChange HashChangeEvent cancelable: false, bubbles: true
@event LanguageChange Event cancelable: false, bubbles: false
@event Message MessageEvent cancelable: false, bubbles: false
@event MessageError MessageEvent cancelable: false, bubbles: false
@event Offline Event cancelable: false, bubbles: false
@event Online Event cancelable: false, bubbles: false
@event PageHide PageTransitionEvent cancelable: false, bubbles: false
@event PageShow PageTransitionEvent cancelable: false, bubbles: false
@event PopState PopStateEvent cancelable: false, bubbles: true
@event RejectionHandled PromiseRejectionEvent cancelable: false, bubbles: false
@event Storage StorageEvent cancelable: false, bubbles: false
@event UnhandledRejection PromiseRejectionEvent cancelable: false, bubbles: false
@event Unload Event cancelable: false, bubbles: false
