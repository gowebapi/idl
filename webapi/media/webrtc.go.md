# WebRTC 1.0

.title = WebRTC 1.0
.url = <https://w3c.github.io/webrtc-pc/>
.package = github.com/gowebapi/webapi/media/webrtc

@on ".": @replace .name "RTC" ""

## RTCDTMFSender

@event ToneChange RTCDTMFToneChangeEvent cancelable: false, bubbles: false

## RTCDataChannel

@event BufferedAmountLow Event cancelable: false, bubbles: false
@event Close Event cancelable: false, bubbles: false
@event Error RTCErrorEvent cancelable: false, bubbles: false
@event Message MessageEvent cancelable: false, bubbles: false
@event Open Event cancelable: false, bubbles: false

## RTCDtlsTransport

@event Error RTCErrorEvent cancelable: false, bubbles: false
@event StateChange Event cancelable: false, bubbles: false

## RTCIceTransport

@event GatheringStateChange Event cancelable: false, bubbles: false
@event SelectedCandidatePairChange Event cancelable: false, bubbles: false
@event StateChange Event cancelable: false, bubbles: false

## RTCPeerConnection

@event ConnectionStateChange Event cancelable: false, bubbles: false
@event DataChannel RTCDataChannelEvent cancelable: false, bubbles: false
@event IceCandidate RTCPeerConnectionIceEvent cancelable: false, bubbles: false
@event IceCandidateError RTCPeerConnectionIceErrorEvent cancelable: false, bubbles: false
@event IceConnectionStateChange Event cancelable: false, bubbles: false
@event IceGatheringStateChange Event cancelable: false, bubbles: false
@event NegotiationNeeded Event cancelable: false, bubbles: false
@event SignalingStateChange Event cancelable: false, bubbles: false
@event StatsEnded RTCStatsEvent cancelable: false, bubbles: false
@event Track RTCTrackEvent cancelable: false, bubbles: false

## RTCSctpTransport

@event StateChange Event cancelable: false, bubbles: false
