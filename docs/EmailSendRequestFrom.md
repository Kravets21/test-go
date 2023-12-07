# EmailSendRequestFrom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The full name from which the notification is sent. | [optional] 
**EmailAddress** | Pointer to **string** | The email address from which the notification is sent. | [optional] 

## Methods

### NewEmailSendRequestFrom

`func NewEmailSendRequestFrom() *EmailSendRequestFrom`

NewEmailSendRequestFrom instantiates a new EmailSendRequestFrom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendRequestFromWithDefaults

`func NewEmailSendRequestFromWithDefaults() *EmailSendRequestFrom`

NewEmailSendRequestFromWithDefaults instantiates a new EmailSendRequestFrom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailSendRequestFrom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailSendRequestFrom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailSendRequestFrom) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailSendRequestFrom) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *EmailSendRequestFrom) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EmailSendRequestFrom) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EmailSendRequestFrom) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *EmailSendRequestFrom) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


