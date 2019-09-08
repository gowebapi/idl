# Web NFC API

.title = Web NFC API
.url = <https://w3c.github.io/web-nfc/>
.package = github.com/gowebapi/webapi/communication/nfc

@on ".": @replace .name "NFC" ""

## NDEFCompatibility

.prefix = "Compatibility"
.suffix = ""

## NFCReader

@event Error NFCErrorEvent cancelable: false, bubbles: false
@event Reading NFCReadingEvent cancelable: false, bubbles: false

## NDEFRecordType

.prefix = "Record"
.suffix = ""

## NFCPushTarget

.prefix = "PushTarget"
.suffix = ""
