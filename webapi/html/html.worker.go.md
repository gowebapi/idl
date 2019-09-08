# Web Worker API

    https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API

## AbstractWorker

@event Error Event cancelable: false, bubbles: false

## Worker

.package = github.com/gowebapi/webapi/html/worker
    @event Error Event cancelable: false, bubbles: false
@event Message MessageEvent cancelable: false, bubbles: false
@event MessageError MessageEvent cancelable: false, bubbles: false

## WorkerLocation

.package = github.com/gowebapi/webapi/html/worker

## SharedWorker

.package = github.com/gowebapi/webapi/html/worker

## WorkerGlobalScope

.package = github.com/gowebapi/webapi/html/worker
@event Error Event cancelable: false, bubbles: false
@event LanguageChange Event cancelable: false, bubbles: false
@event Offline Event cancelable: false, bubbles: false
@event Online Event cancelable: false, bubbles: false
@event RejectionHandled PromiseRejectionEvent cancelable: false, bubbles: false
@event UnhandledRejection PromiseRejectionEvent cancelable: false, bubbles: false

## DedicatedWorkerGlobalScope

.package = github.com/gowebapi/webapi/html/worker
@event Message MessageEvent cancelable: false, bubbles: false
@event MessageError MessageEvent cancelable: false, bubbles: false

## SharedWorkerGlobalScope

.package = github.com/gowebapi/webapi/html/worker
@event Connect MessageEvent cancelable: false, bubbles: false

## WorkerNavigator

.package = github.com/gowebapi/webapi/html/worker
@notevent onLine

## WorkerOptions

.package = github.com/gowebapi/webapi/html/worker
