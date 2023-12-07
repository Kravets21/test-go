# SmsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Notification. | [optional] 
**Locale** | Pointer to **string** | The locale of the Notification. | [optional] 
**Price** | Pointer to **string** | The price of the SMS. | [optional] 
**SenderId** | Pointer to **string** | The ID of sender. | [optional] 
**RecipientId** | Pointer to **string** | The ID of recipient. | [optional] 
**MessageSid** | Pointer to **string** | The unique identifier of the SMS. | [optional] 
**ToNumber** | Pointer to **string** | The phone number to send the SMS to. | [optional] 
**Status** | Pointer to **int32** | The state of the Notification status.  - 1: Notification Accepted  - 2: Notification Blocked  - 3: Notification Sent  - 4: Notification Not Sent  - 5: Notification Delivered  - 6: Notification Not Delivered | [optional] 
**CustomParams** | Pointer to **map[string]interface{}** | The custom parameters of the notification. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Notification was created. | [optional] 

## Methods

### NewSmsNotification

`func NewSmsNotification() *SmsNotification`

NewSmsNotification instantiates a new SmsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsNotificationWithDefaults

`func NewSmsNotificationWithDefaults() *SmsNotification`

NewSmsNotificationWithDefaults instantiates a new SmsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsNotification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocale

`func (o *SmsNotification) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SmsNotification) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SmsNotification) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SmsNotification) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetPrice

`func (o *SmsNotification) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SmsNotification) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SmsNotification) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SmsNotification) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSenderId

`func (o *SmsNotification) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *SmsNotification) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *SmsNotification) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *SmsNotification) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetRecipientId

`func (o *SmsNotification) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *SmsNotification) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *SmsNotification) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *SmsNotification) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetMessageSid

`func (o *SmsNotification) GetMessageSid() string`

GetMessageSid returns the MessageSid field if non-nil, zero value otherwise.

### GetMessageSidOk

`func (o *SmsNotification) GetMessageSidOk() (*string, bool)`

GetMessageSidOk returns a tuple with the MessageSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSid

`func (o *SmsNotification) SetMessageSid(v string)`

SetMessageSid sets MessageSid field to given value.

### HasMessageSid

`func (o *SmsNotification) HasMessageSid() bool`

HasMessageSid returns a boolean if a field has been set.

### GetToNumber

`func (o *SmsNotification) GetToNumber() string`

GetToNumber returns the ToNumber field if non-nil, zero value otherwise.

### GetToNumberOk

`func (o *SmsNotification) GetToNumberOk() (*string, bool)`

GetToNumberOk returns a tuple with the ToNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToNumber

`func (o *SmsNotification) SetToNumber(v string)`

SetToNumber sets ToNumber field to given value.

### HasToNumber

`func (o *SmsNotification) HasToNumber() bool`

HasToNumber returns a boolean if a field has been set.

### GetStatus

`func (o *SmsNotification) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsNotification) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsNotification) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomParams

`func (o *SmsNotification) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *SmsNotification) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *SmsNotification) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *SmsNotification) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SmsNotification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmsNotification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmsNotification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SmsNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


