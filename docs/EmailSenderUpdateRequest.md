# EmailSenderUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the Sender. | [optional] 
**EmailAddress** | Pointer to **NullableString** | The email address of the Sender. | [optional] 
**RestrictCustomization** | Pointer to **bool** | Indicates to prevent the sender from applying custom values. Default value is false. | [optional] 

## Methods

### NewEmailSenderUpdateRequest

`func NewEmailSenderUpdateRequest() *EmailSenderUpdateRequest`

NewEmailSenderUpdateRequest instantiates a new EmailSenderUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSenderUpdateRequestWithDefaults

`func NewEmailSenderUpdateRequestWithDefaults() *EmailSenderUpdateRequest`

NewEmailSenderUpdateRequestWithDefaults instantiates a new EmailSenderUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailSenderUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailSenderUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailSenderUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailSenderUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EmailSenderUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EmailSenderUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmailAddress

`func (o *EmailSenderUpdateRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EmailSenderUpdateRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EmailSenderUpdateRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *EmailSenderUpdateRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *EmailSenderUpdateRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *EmailSenderUpdateRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetRestrictCustomization

`func (o *EmailSenderUpdateRequest) GetRestrictCustomization() bool`

GetRestrictCustomization returns the RestrictCustomization field if non-nil, zero value otherwise.

### GetRestrictCustomizationOk

`func (o *EmailSenderUpdateRequest) GetRestrictCustomizationOk() (*bool, bool)`

GetRestrictCustomizationOk returns a tuple with the RestrictCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictCustomization

`func (o *EmailSenderUpdateRequest) SetRestrictCustomization(v bool)`

SetRestrictCustomization sets RestrictCustomization field to given value.

### HasRestrictCustomization

`func (o *EmailSenderUpdateRequest) HasRestrictCustomization() bool`

HasRestrictCustomization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


