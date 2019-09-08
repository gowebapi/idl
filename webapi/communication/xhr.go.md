# XMLHttpRequest Standard

.title = XMLHttpRequest Standard
.url = <https://xhr.spec.whatwg.org/>
.package = github.com/gowebapi/webapi/communication/xhr

@on "FormData*": .package = github.com/gowebapi/webapi/html
    @eventprop cancelable: false, bubbles: false

## XMLHttpRequest

@patch idlconst
@changetype responseXML rawjs
.constSuffix = ""
@event ReadyStateChange Event

## XMLHttpRequestEventTarget

@event LoadStart ProgressEvent
@event Progress ProgressEvent
@event Abort ProgressEvent
@event Error ProgressEvent
@event Load ProgressEvent
@event LoadEnd ProgressEvent
@event TimeOut ProgressEvent
