# Web Speech API

.title = Web Speech API
.url = <https://w3c.github.io/speech-api/>
.package = github.com/gowebapi/webapi/media/speech
.comment = Type changes in SpeechRecognitionEvent(Init)

## SpeechRecognition

@event AudioEnd Event cancelable: false, bubbles: false
@event AudioStart Event cancelable: false, bubbles: false
@event End Event cancelable: false, bubbles: false
@event Error SpeechRecognitionErrorEvent cancelable: false, bubbles: false
@event NoMatch SpeechRecognitionEvent cancelable: false, bubbles: false
@event Result SpeechRecognitionEvent cancelable: false, bubbles: false
@event SoundEnd Event cancelable: false, bubbles: false
@event SoundStart Event cancelable: false, bubbles: false
@event SpeechEnd Event cancelable: false, bubbles: false
@event SpeechStart Event cancelable: false, bubbles: false
@event Start Event cancelable: false, bubbles: false

## SpeechRecognitionEventInit

@changetype emma rawjs

## SpeechRecognitionEvent

@changetype emma rawjs

## SpeechSynthesis

@event VoicesChanged Event cancelable: false, bubbles: false

## SpeechSynthesisUtterance

@event Boundary SpeechSynthesisEvent cancelable: false, bubbles: false
@event End SpeechSynthesisEvent cancelable: false, bubbles: false
@event Error SpeechSynthesisErrorEvent cancelable: false, bubbles: false
@event Mark SpeechSynthesisEvent cancelable: false, bubbles: false
@event Pause SpeechSynthesisEvent cancelable: false, bubbles: false
@event Resume SpeechSynthesisEvent cancelable: false, bubbles: false
@event Start SpeechSynthesisEvent cancelable: false, bubbles: false
