# PushNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Notification. | [optional] 
**Locale** | Pointer to [**LocaleCode**](LocaleCode.md) |  | [optional] 
**Token** | Pointer to **string** | The recipient token of the Notification. | [optional] 
**RecipientId** | Pointer to **string** | The ID of recipient. | [optional] 
**Status** | Pointer to **int32** | The state of the Notification status.  - 1: Notification Accepted  - 2: Notification Blocked  - 3: Notification Sent  - 4: Notification Not Sent  - 5: Notification Delivered  - 6: Notification Not Delivered | [optional] 
**CustomParams** | Pointer to **map[string]interface{}** | The custom parameters of the Notification. | [optional] 
**DeliveredAt** | Pointer to **time.Time** | The datetime when the Notification was delivered. | [optional] 
**OpenedAt** | Pointer to **time.Time** | The datetime when the Notification was opened. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Notification was created. | [optional] 

## Methods

### NewPushNotification

`func NewPushNotification() *PushNotification`

NewPushNotification instantiates a new PushNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationWithDefaults

`func NewPushNotificationWithDefaults() *PushNotification`

NewPushNotificationWithDefaults instantiates a new PushNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PushNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PushNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PushNotification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PushNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocale

`func (o *PushNotification) GetLocale() LocaleCode`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PushNotification) GetLocaleOk() (*LocaleCode, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PushNotification) SetLocale(v LocaleCode)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PushNotification) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetToken

`func (o *PushNotification) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PushNotification) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PushNotification) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PushNotification) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetRecipientId

`func (o *PushNotification) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PushNotification) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PushNotification) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *PushNotification) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetStatus

`func (o *PushNotification) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PushNotification) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PushNotification) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PushNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomParams

`func (o *PushNotification) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *PushNotification) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *PushNotification) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *PushNotification) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetDeliveredAt

`func (o *PushNotification) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *PushNotification) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *PushNotification) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *PushNotification) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### GetOpenedAt

`func (o *PushNotification) GetOpenedAt() time.Time`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *PushNotification) GetOpenedAtOk() (*time.Time, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *PushNotification) SetOpenedAt(v time.Time)`

SetOpenedAt sets OpenedAt field to given value.

### HasOpenedAt

`func (o *PushNotification) HasOpenedAt() bool`

HasOpenedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PushNotification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PushNotification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PushNotification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PushNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


