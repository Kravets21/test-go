# SmsTransportCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Transport. | 
**Type** | **string** | The type of the Transport. | 
**Credentials** | **map[string]interface{}** | The credentials of the Transport for type. | 

## Methods

### NewSmsTransportCreateRequest

`func NewSmsTransportCreateRequest(name string, type_ string, credentials map[string]interface{}, ) *SmsTransportCreateRequest`

NewSmsTransportCreateRequest instantiates a new SmsTransportCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsTransportCreateRequestWithDefaults

`func NewSmsTransportCreateRequestWithDefaults() *SmsTransportCreateRequest`

NewSmsTransportCreateRequestWithDefaults instantiates a new SmsTransportCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SmsTransportCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmsTransportCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmsTransportCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SmsTransportCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmsTransportCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmsTransportCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCredentials

`func (o *SmsTransportCreateRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SmsTransportCreateRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SmsTransportCreateRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


