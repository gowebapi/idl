# Presentation API

.title = Presentation API
.url = <https://w3c.github.io/presentation-api/>
.package = github.com/gowebapi/webapi/graphics/presentation

@on ".": @replace .name "Presentation" ""

## Presentation

.name = "Presentation"

## PresentationAvailability

@event Change Event cancelable: false, bubbles:false

## PresentationConnection

@event Close PresentationConnectionCloseEvent cancelable: false, bubbles:false
@event Connect Event cancelable: false, bubbles:false
@event Message MessageEvent cancelable: false, bubbles:false
@event Terminate Event cancelable: false, bubbles:false

## PresentationConnectionCloseReason

.prefix = "Reason"
.suffix = ""

## PresentationConnectionList

@event ConnectionAvailable PresentationConnectionAvailableEvent cancelable: false, bubbles:false

## PresentationConnectionState

.prefix = "State"
.suffix = ""

## PresentationRequest

@event ConnectionAvailable PresentationConnectionAvailableEvent cancelable: false, bubbles:false
