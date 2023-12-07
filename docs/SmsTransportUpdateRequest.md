# SmsTransportUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the Transport. | [optional] 
**Status** | Pointer to **NullableInt32** | The status of the Transport. Every digit matters: 1 - active, 2 - inactive, | [optional] 
**Credentials** | Pointer to **map[string]interface{}** | The credentials of the Transport for type. | [optional] 

## Methods

### NewSmsTransportUpdateRequest

`func NewSmsTransportUpdateRequest() *SmsTransportUpdateRequest`

NewSmsTransportUpdateRequest instantiates a new SmsTransportUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsTransportUpdateRequestWithDefaults

`func NewSmsTransportUpdateRequestWithDefaults() *SmsTransportUpdateRequest`

NewSmsTransportUpdateRequestWithDefaults instantiates a new SmsTransportUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SmsTransportUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmsTransportUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmsTransportUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmsTransportUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SmsTransportUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SmsTransportUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *SmsTransportUpdateRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsTransportUpdateRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsTransportUpdateRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsTransportUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SmsTransportUpdateRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SmsTransportUpdateRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCredentials

`func (o *SmsTransportUpdateRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SmsTransportUpdateRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SmsTransportUpdateRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SmsTransportUpdateRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *SmsTransportUpdateRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *SmsTransportUpdateRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


