# File API

.title = File API
.url = <https://w3c.github.io/FileAPI/>
.package = github.com/gowebapi/webapi/file

## FileReader

@patch idlconst
.constSuffix = ""

@eventprop cancelable: false, bubbles: false
@event LoadStart ProgressEvent
@event Progress ProgressEvent
@event Abort ProgressEvent
@event Error ProgressEvent
@event Load ProgressEvent
@event LoadEnd ProgressEvent
