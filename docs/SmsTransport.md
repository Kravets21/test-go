# SmsTransport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Transport. | [optional] 
**Name** | Pointer to **string** | The name of the Transport. | [optional] 
**Type** | Pointer to **string** | The type of the Transport. | [optional] 
**Status** | Pointer to **int32** | The status of the Transport. Every digit matters: 1 - active, 2 - inactive. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Transport was created. | [optional] 

## Methods

### NewSmsTransport

`func NewSmsTransport() *SmsTransport`

NewSmsTransport instantiates a new SmsTransport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsTransportWithDefaults

`func NewSmsTransportWithDefaults() *SmsTransport`

NewSmsTransportWithDefaults instantiates a new SmsTransport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsTransport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsTransport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsTransport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsTransport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SmsTransport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmsTransport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmsTransport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmsTransport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SmsTransport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmsTransport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmsTransport) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SmsTransport) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *SmsTransport) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsTransport) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsTransport) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsTransport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SmsTransport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmsTransport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmsTransport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SmsTransport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


