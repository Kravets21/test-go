# EmailTransport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Platform. | [optional] 
**Name** | Pointer to **string** | The name of the Platform. | [optional] 
**Type** | Pointer to **string** | The type of the Platform. | [optional] 
**Headers** | Pointer to **map[string]interface{}** | The headers of the Transport. | [optional] 
**Status** | Pointer to **int32** | The status of the Transport. Every digit matters: 1 - active, 2 - inactive, 3 - incomplete. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Platform was created. | [optional] 

## Methods

### NewEmailTransport

`func NewEmailTransport() *EmailTransport`

NewEmailTransport instantiates a new EmailTransport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTransportWithDefaults

`func NewEmailTransportWithDefaults() *EmailTransport`

NewEmailTransportWithDefaults instantiates a new EmailTransport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailTransport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailTransport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailTransport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailTransport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EmailTransport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailTransport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailTransport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailTransport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EmailTransport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailTransport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailTransport) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmailTransport) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHeaders

`func (o *EmailTransport) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmailTransport) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmailTransport) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmailTransport) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetStatus

`func (o *EmailTransport) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailTransport) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailTransport) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailTransport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmailTransport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailTransport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailTransport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmailTransport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


