# DOM Standard

.package = github.com/gowebapi/webapi/dom
.title = DOM Standard
.url = <https://dom.spec.whatwg.org/>

@on "DOMTokenList*": .package = github.com/gowebapi/webapi/dom/domcore

## Node

@changetype ownerDocument rawjs
@patch idlconst
.constSuffix = ""

## NodeFilter

@patch idlconst
.constSuffix = ""

## Element

@changetype assignedSlot rawjs

    fullscreen.idl
@event FullscreenChange Event cancelable: false, bubbles: true
@event FullscreenError Event cancelable: false, bubbles: true

## Range

@patch idlconst
.constSuffix = ""

## Text

@changetype assignedSlot rawjs

    Types that can be moved elsewhere:
    ## TreeWalker
    ## Range
    ## AbstractRange
    ## StaticRange
    ## CharacterData
    ## Text
    ## CDATASection
    ## ProcessingInstruction
    ## Comment
    ## DocumentType
