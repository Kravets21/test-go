# PushNotificationSendRequestTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The external ID of recipient. This field creates a single recipient across different transports (email, push, SMS). | [optional] 
**RecipientId** | Pointer to **string** | The ID of recipient. This allows to send notification to existing recipient. If the recipient does not exist or does not have fields for sending push, then sending is blocked. The message.token or recipient_id parameter must be specified. | [optional] 
**Token** | Pointer to **string** | The Token of recipient. | [optional] 
**Platform** | Pointer to **string** | The Recipient push token platform. | [optional] 

## Methods

### NewPushNotificationSendRequestTo

`func NewPushNotificationSendRequestTo() *PushNotificationSendRequestTo`

NewPushNotificationSendRequestTo instantiates a new PushNotificationSendRequestTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationSendRequestToWithDefaults

`func NewPushNotificationSendRequestToWithDefaults() *PushNotificationSendRequestTo`

NewPushNotificationSendRequestToWithDefaults instantiates a new PushNotificationSendRequestTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *PushNotificationSendRequestTo) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PushNotificationSendRequestTo) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PushNotificationSendRequestTo) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PushNotificationSendRequestTo) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRecipientId

`func (o *PushNotificationSendRequestTo) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PushNotificationSendRequestTo) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PushNotificationSendRequestTo) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *PushNotificationSendRequestTo) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetToken

`func (o *PushNotificationSendRequestTo) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PushNotificationSendRequestTo) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PushNotificationSendRequestTo) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PushNotificationSendRequestTo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetPlatform

`func (o *PushNotificationSendRequestTo) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PushNotificationSendRequestTo) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PushNotificationSendRequestTo) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PushNotificationSendRequestTo) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


