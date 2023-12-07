# EmailTransportCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Transport. | 
**Type** | **string** | The type of the Transport. | 
**Credentials** | **map[string]interface{}** | The credentials of the Transport for type. | 
**Headers** | Pointer to **map[string]interface{}** | The headers of the Transport. | [optional] 

## Methods

### NewEmailTransportCreateRequest

`func NewEmailTransportCreateRequest(name string, type_ string, credentials map[string]interface{}, ) *EmailTransportCreateRequest`

NewEmailTransportCreateRequest instantiates a new EmailTransportCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTransportCreateRequestWithDefaults

`func NewEmailTransportCreateRequestWithDefaults() *EmailTransportCreateRequest`

NewEmailTransportCreateRequestWithDefaults instantiates a new EmailTransportCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailTransportCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailTransportCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailTransportCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EmailTransportCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailTransportCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailTransportCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCredentials

`func (o *EmailTransportCreateRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *EmailTransportCreateRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *EmailTransportCreateRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.


### GetHeaders

`func (o *EmailTransportCreateRequest) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmailTransportCreateRequest) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmailTransportCreateRequest) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmailTransportCreateRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *EmailTransportCreateRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *EmailTransportCreateRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


