# HTML media

    this is extracting the html/media package

## AudioTrack

.package = github.com/gowebapi/webapi/html/media

## AudioTrackList

.package = github.com/gowebapi/webapi/html/media
@event AddTrack TrackEvent cancelable: false, bubbles: false
@event Change Event cancelable: false, bubbles: true
@event RemoveTrack TrackEvent cancelable: false, bubbles: false

## CanPlayTypeResult

.package = github.com/gowebapi/webapi/html/media

## HTMLAudioElement

.package = github.com/gowebapi/webapi/html/media

## HTMLMediaElement

@patch idlconst
.constSuffix = ""
.package = github.com/gowebapi/webapi/html/media

    encrypted-media.idl
@event Encrypted MediaEncryptedEvent cancelable: false, bubbles: false
@event WaitingForKey Event cancelable: false, bubbles: false

## HTMLTrackElement

@patch idlconst
.constSuffix = ""
.package = github.com/gowebapi/webapi/html/media

## HTMLVideoElement

.package = github.com/gowebapi/webapi/html/media

@event EnterPictureInPicture EnterPictureInPictureEvent cancelable: false, bubbles: false
@event LeavePictureInPicture Event cancelable: false, bubbles: true

## MediaError

@patch idlconst
.constSuffix = ""
.package = github.com/gowebapi/webapi/html/media

## TextTrack

.package = github.com/gowebapi/webapi/html/media
@event CueChange Event cancelable: false, bubbles: false

## TextTrackKind

.package = github.com/gowebapi/webapi/html/media

## TextTrackList

.package = github.com/gowebapi/webapi/html/media
@event AddTrack TrackEvent cancelable: false, bubbles: false
@event Change Event cancelable: false, bubbles: true
@event RemoveTrack TrackEvent cancelable: false, bubbles: false

## TextTrackCue

.package = github.com/gowebapi/webapi/html/media
@event Enter Event cancelable: false, bubbles: false
@event Exit Event cancelable: false, bubbles: false

## TextTrackCueList

.package = github.com/gowebapi/webapi/html/media

## TextTrackMode

.package = github.com/gowebapi/webapi/html/media

## VideoTrack

.package = github.com/gowebapi/webapi/html/media

## VideoTrackList

.package = github.com/gowebapi/webapi/html/media
@event AddTrack TrackEvent cancelable: false, bubbles: false
@event Change Event cancelable: false, bubbles: true
@event RemoveTrack TrackEvent cancelable: false, bubbles: false
