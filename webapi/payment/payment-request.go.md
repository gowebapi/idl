# Payment Request API

.title = Payment Request API
.url = <https://w3c.github.io/payment-request/>
.package = github.com/gowebapi/webapi/payment/request

## PaymentRequest

@event MerchantValidation MerchantValidationEvent cancelable: false, bubbles: false
@event PaymentMethodChange PaymentMethodChangeEvent cancelable: false, bubbles: false
@event ShippingAddressChange PaymentRequestUpdateEvent cancelable: false, bubbles: false
@event ShippingOptionChange PaymentRequestUpdateEvent cancelable: false, bubbles: false

## PaymentResponse

@event PayerDetailChange PaymentRequestUpdateEvent cancelable: false, bubbles: false
