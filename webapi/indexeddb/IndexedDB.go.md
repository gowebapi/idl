# Indexed Database API 3.0

.title = Indexed Database API 3.0
.url = <https://w3c.github.io/IndexedDB/>
.package = github.com/gowebapi/webapi/indexeddb

## IDBDatabase

@event Abort Event cancelable: false, bubbles: false
@event Close Event cancelable: false, bubbles: false
@event Error Event cancelable: false, bubbles: false
@event VersionChange IDBVersionChangeEvent cancelable: false, bubbles: false

## IDBOpenDBRequest

@event Blocked Event cancelable: false, bubbles: false
@event UpgradeNeeded Event cancelable: false, bubbles: false

## IDBRequest

@event Error Event cancelable: false, bubbles: false
@event Success Event cancelable: false, bubbles: false

## IDBTransaction

@event Abort Event cancelable: false, bubbles: false
@event Complete Event cancelable: false, bubbles: false
@event Error Event cancelable: false, bubbles: false
