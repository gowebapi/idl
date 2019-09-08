# DOM classes in main webapi

    following types are refering to many other this and need to be
    moved into webapi package

@on "Mutation." : .package = github.com/gowebapi/webapi

## Document

.package = github.com/gowebapi/webapi

@event ReadyStateChange Event cancelable: false, bubbles: false

    fullscreen.idl
@event FullscreenChange Event cancelable: false, bubbles: true
@event FullscreenError Event cancelable: false, bubbles: true

    pointerlock.idl
@event PointerLockChange Event cancelable: false, bubbles: true
@event PointerLockError Event cancelable: false, bubbles: true

    page-visibility.idl
@event VisibilityChange Event cancelable: false, bubbles: true

## DOMImplementation

.package = github.com/gowebapi/webapi

## ElementCreationOptions

.package = github.com/gowebapi/webapi

## GlobalScope

    javascript global scope

.package = github.com/gowebapi/webapi
document = GetDocument
window = GetWindow

## XMLDocument

.package = github.com/gowebapi/webapi
