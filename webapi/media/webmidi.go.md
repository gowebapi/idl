# Web MIDI API

.title = Web MIDI API
.url = <http://webaudio.github.io/web-midi-api/>
.package = github.com/gowebapi/webapi/media/midi

@on ".": @replace .name "MIDI" ""

## MIDIAccess

@event StateChange MIDIConnectionEvent cancelable: false, bubbles: false

## MIDIInput

@event MIDIMessage MIDIMessageEvent cancelable: false, bubbles: false

## MIDIPort

@event StateChange MIDIConnectionEvent cancelable: false, bubbles: false

## MIDIPortConnectionState

.suffix = ""

## MIDIPortDeviceState

.suffix = ""

## MIDIPortType

.suffix = "Port"
