# Web Authentication

.title = Web Authentication
.url = <https://w3c.github.io/webauthn/>
.package = github.com/gowebapi/webapi/crypto/authentication

@on "PublicKey*": .package = github.com/gowebapi/webapi/crypto/credential
@on "Authenticator*": @replace .name "Authenticator" ""

## AuthenticatorAttachment

.suffix = ""

## AttestationConveyancePreference

.suffix = ""

## AuthenticatorTransport

.suffix = ""
.prefix = "Transport"

## UserVerificationRequirement

.suffix = ""
