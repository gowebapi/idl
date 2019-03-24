# DOM Standard

.package = github.com/gowebapi/webapi/dom
.title = DOM Standard
.url = <https://dom.spec.whatwg.org/>

## Node

@changetype ownerDocument rawjs
@patch idlconst
.constSuffix = ""

## NodeFilter

@patch idlconst
.constSuffix = ""

## Element

@changetype assignedSlot rawjs

## Range

@patch idlconst
.constSuffix = ""

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