# SmsChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Channel. | [optional] 
**Name** | Pointer to **string** | The name of the Channel. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Channel was created. | [optional] 

## Methods

### NewSmsChannel

`func NewSmsChannel() *SmsChannel`

NewSmsChannel instantiates a new SmsChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsChannelWithDefaults

`func NewSmsChannelWithDefaults() *SmsChannel`

NewSmsChannelWithDefaults instantiates a new SmsChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SmsChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmsChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmsChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmsChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SmsChannel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmsChannel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmsChannel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SmsChannel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


