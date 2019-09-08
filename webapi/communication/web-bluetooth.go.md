# Web Bluetooth

.title = Web Bluetooth
.url = <https://webbluetoothcg.github.io/web-bluetooth/>
.package = github.com/gowebapi/webapi/communication/bluetooth

@on ".": @replace .name "Bluetooth" ""

## Bluetooth

.name = "Bluetooth"

@event AvailabilityChanged ValueEvent

## CharacteristicEventHandlers

@event CharacteristicValueChanged Event bubbles: true

## BluetoothDeviceEventHandlers

@event AdvertisementReceived BluetoothAdvertisingEvent bubbles: true
@event GattServerDisconnected Event

## ServiceEventHandlers

@event ServiceAdded Event
@event ServiceChanged Event
@event ServiceRemoved Event

    ## BluetoothDevice

    @event AdvertisementReceived Event
    @event CharacteristicValueChanged Event
    @event GattServerDisconnected Event
    @event ServiceAdded Event
    @event ServiceChanged Event
    @event ServiceRemoved Event

    ## BluetoothRemoteGATTCharacteristic

    @event CharacteristicValueChanged

    ## BluetoothRemoteGATTService

    @event CharacteristicValueChanged Event
    @event ServiceAdded Event
    @event ServiceChanged Event
