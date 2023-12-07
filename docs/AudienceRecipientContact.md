# AudienceRecipientContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the Contact. | [optional] 
**Address** | Pointer to **NullableString** | The address of the Contact. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Contact was created. | [optional] 

## Methods

### NewAudienceRecipientContact

`func NewAudienceRecipientContact() *AudienceRecipientContact`

NewAudienceRecipientContact instantiates a new AudienceRecipientContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceRecipientContactWithDefaults

`func NewAudienceRecipientContactWithDefaults() *AudienceRecipientContact`

NewAudienceRecipientContactWithDefaults instantiates a new AudienceRecipientContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AudienceRecipientContact) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AudienceRecipientContact) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AudienceRecipientContact) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AudienceRecipientContact) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *AudienceRecipientContact) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AudienceRecipientContact) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AudienceRecipientContact) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AudienceRecipientContact) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *AudienceRecipientContact) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *AudienceRecipientContact) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCreatedAt

`func (o *AudienceRecipientContact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AudienceRecipientContact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AudienceRecipientContact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AudienceRecipientContact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


