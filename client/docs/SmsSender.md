# SmsSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Transport Number. | [optional] 
**TransportId** | Pointer to **string** | The unique identifier Transport ID of the Sender. | [optional] 
**Name** | Pointer to **string** | The name of the Transport Number. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number of the Transport Number. | [optional] 
**Status** | Pointer to **int32** | The status of the Sender. Every digit matters: 1 - active, 2 - inactive. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Transport Number was created. | [optional] 

## Methods

### NewSmsSender

`func NewSmsSender() *SmsSender`

NewSmsSender instantiates a new SmsSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsSenderWithDefaults

`func NewSmsSenderWithDefaults() *SmsSender`

NewSmsSenderWithDefaults instantiates a new SmsSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsSender) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsSender) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsSender) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsSender) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTransportId

`func (o *SmsSender) GetTransportId() string`

GetTransportId returns the TransportId field if non-nil, zero value otherwise.

### GetTransportIdOk

`func (o *SmsSender) GetTransportIdOk() (*string, bool)`

GetTransportIdOk returns a tuple with the TransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportId

`func (o *SmsSender) SetTransportId(v string)`

SetTransportId sets TransportId field to given value.

### HasTransportId

`func (o *SmsSender) HasTransportId() bool`

HasTransportId returns a boolean if a field has been set.

### GetName

`func (o *SmsSender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmsSender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmsSender) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmsSender) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *SmsSender) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *SmsSender) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *SmsSender) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *SmsSender) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *SmsSender) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsSender) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsSender) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsSender) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SmsSender) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmsSender) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmsSender) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SmsSender) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


