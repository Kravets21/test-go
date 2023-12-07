# EmailSenderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Sender. | 
**EmailAddress** | **string** | The email address of the Sender. | 
**RestrictCustomization** | Pointer to **bool** | Indicates to prevent the sender from applying custom values. Default value is false. | [optional] 

## Methods

### NewEmailSenderCreateRequest

`func NewEmailSenderCreateRequest(name string, emailAddress string, ) *EmailSenderCreateRequest`

NewEmailSenderCreateRequest instantiates a new EmailSenderCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSenderCreateRequestWithDefaults

`func NewEmailSenderCreateRequestWithDefaults() *EmailSenderCreateRequest`

NewEmailSenderCreateRequestWithDefaults instantiates a new EmailSenderCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailSenderCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailSenderCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailSenderCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmailAddress

`func (o *EmailSenderCreateRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EmailSenderCreateRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EmailSenderCreateRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetRestrictCustomization

`func (o *EmailSenderCreateRequest) GetRestrictCustomization() bool`

GetRestrictCustomization returns the RestrictCustomization field if non-nil, zero value otherwise.

### GetRestrictCustomizationOk

`func (o *EmailSenderCreateRequest) GetRestrictCustomizationOk() (*bool, bool)`

GetRestrictCustomizationOk returns a tuple with the RestrictCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictCustomization

`func (o *EmailSenderCreateRequest) SetRestrictCustomization(v bool)`

SetRestrictCustomization sets RestrictCustomization field to given value.

### HasRestrictCustomization

`func (o *EmailSenderCreateRequest) HasRestrictCustomization() bool`

HasRestrictCustomization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


