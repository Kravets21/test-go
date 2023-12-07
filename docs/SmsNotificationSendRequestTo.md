# SmsNotificationSendRequestTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | The phone number to send the SMS to. The number or recipient_id parameter must be specified. | [optional] 
**ExternalId** | Pointer to **string** | The external ID of recipient. This field creates a single recipient across different transports (email, push, SMS). | [optional] 
**RecipientId** | Pointer to **string** | The ID of recipient. This allows to send notification to existing recipient. If the recipient does not exist or does not have fields for sending push, then sending is blocked. The number or recipient_id parameter must be specified. | [optional] 

## Methods

### NewSmsNotificationSendRequestTo

`func NewSmsNotificationSendRequestTo() *SmsNotificationSendRequestTo`

NewSmsNotificationSendRequestTo instantiates a new SmsNotificationSendRequestTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsNotificationSendRequestToWithDefaults

`func NewSmsNotificationSendRequestToWithDefaults() *SmsNotificationSendRequestTo`

NewSmsNotificationSendRequestToWithDefaults instantiates a new SmsNotificationSendRequestTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *SmsNotificationSendRequestTo) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SmsNotificationSendRequestTo) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SmsNotificationSendRequestTo) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *SmsNotificationSendRequestTo) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetExternalId

`func (o *SmsNotificationSendRequestTo) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SmsNotificationSendRequestTo) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SmsNotificationSendRequestTo) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SmsNotificationSendRequestTo) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRecipientId

`func (o *SmsNotificationSendRequestTo) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *SmsNotificationSendRequestTo) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *SmsNotificationSendRequestTo) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *SmsNotificationSendRequestTo) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


