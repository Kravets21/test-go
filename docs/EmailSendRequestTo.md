# EmailSendRequestTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The full name to which the notification is sent. | [optional] 
**EmailAddress** | Pointer to **string** | The email address to which the notification is sent. The email_address or recipient_id parameter must be specified. | [optional] 
**ExternalId** | Pointer to **string** | The external ID of recipient. This field creates a single recipient across different transports (email, push, SMS). | [optional] 
**RecipientId** | Pointer to **string** | The ID of recipient. This allows to send notification to existing recipient. If the recipient does not exist or does not have fields for sending email, then sending is blocked. The email_address or recipient_id parameter must be specified. | [optional] 

## Methods

### NewEmailSendRequestTo

`func NewEmailSendRequestTo() *EmailSendRequestTo`

NewEmailSendRequestTo instantiates a new EmailSendRequestTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendRequestToWithDefaults

`func NewEmailSendRequestToWithDefaults() *EmailSendRequestTo`

NewEmailSendRequestToWithDefaults instantiates a new EmailSendRequestTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailSendRequestTo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailSendRequestTo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailSendRequestTo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailSendRequestTo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *EmailSendRequestTo) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EmailSendRequestTo) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EmailSendRequestTo) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *EmailSendRequestTo) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetExternalId

`func (o *EmailSendRequestTo) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EmailSendRequestTo) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EmailSendRequestTo) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EmailSendRequestTo) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRecipientId

`func (o *EmailSendRequestTo) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EmailSendRequestTo) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EmailSendRequestTo) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EmailSendRequestTo) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


