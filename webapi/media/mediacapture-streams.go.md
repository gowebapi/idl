# Media Capture and Streams

.title = Media Capture and Streams
.url = <https://w3c.github.io/mediacapture-main/>
.package = github.com/gowebapi/webapi/media/capture/local

## ConstrainablePattern

@event OverConstrained Event cancelable: false, bubbles: false

## MediaDevices

@event DeviceChange Event cancelable: false, bubbles: false

## MediaStream

@event AddTrack MediaStreamTrackEvent cancelable: false, bubbles: false
@event RemoveTrack MediaStreamTrackEvent cancelable: false, bubbles: false

## MediaStreamTrack

@event Ended Event cancelable: false, bubbles: false
@event Mute Event cancelable: false, bubbles: false
@event OverConstrained Event cancelable: false, bubbles: false
@event Unmute Event cancelable: false, bubbles: false
