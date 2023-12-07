# EmailSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Sender. | [optional] 
**Name** | Pointer to **string** | The name of the Sender. | [optional] 
**EmailAddress** | Pointer to **string** | The email address of the Sender. | [optional] 
**RestrictCustomization** | Pointer to **bool** | Indicates to prevent the sender from applying custom values. Default value is false. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Sender was created. | [optional] 

## Methods

### NewEmailSender

`func NewEmailSender() *EmailSender`

NewEmailSender instantiates a new EmailSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSenderWithDefaults

`func NewEmailSenderWithDefaults() *EmailSender`

NewEmailSenderWithDefaults instantiates a new EmailSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailSender) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailSender) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailSender) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailSender) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EmailSender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailSender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailSender) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailSender) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *EmailSender) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EmailSender) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EmailSender) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *EmailSender) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetRestrictCustomization

`func (o *EmailSender) GetRestrictCustomization() bool`

GetRestrictCustomization returns the RestrictCustomization field if non-nil, zero value otherwise.

### GetRestrictCustomizationOk

`func (o *EmailSender) GetRestrictCustomizationOk() (*bool, bool)`

GetRestrictCustomizationOk returns a tuple with the RestrictCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictCustomization

`func (o *EmailSender) SetRestrictCustomization(v bool)`

SetRestrictCustomization sets RestrictCustomization field to given value.

### HasRestrictCustomization

`func (o *EmailSender) HasRestrictCustomization() bool`

HasRestrictCustomization returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmailSender) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailSender) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailSender) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmailSender) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


