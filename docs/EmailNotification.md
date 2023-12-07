# EmailNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Notification. | [optional] 
**Locale** | Pointer to [**LocaleCode**](LocaleCode.md) |  | [optional] 
**RecipientId** | Pointer to **string** | The ID of recipient. | [optional] 
**ToName** | Pointer to **string** | The name of the recipient. | [optional] 
**ToEmailAddress** | Pointer to **string** | The email address of the recipient. | [optional] 
**FromName** | Pointer to **string** | The name of the sender. | [optional] 
**FromEmailAddress** | Pointer to **string** | The email address of the sender. | [optional] 
**Status** | Pointer to **int32** | The state of the Notification status.  - 1: Notification Accepted  - 2: Notification Blocked  - 3: Notification Sent  - 4: Notification Not Sent  - 5: Notification Delivered  - 6: Notification Not Delivered | [optional] 
**CustomParams** | Pointer to **map[string]interface{}** | The custom parameters of the notification. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Notification was created. | [optional] 

## Methods

### NewEmailNotification

`func NewEmailNotification() *EmailNotification`

NewEmailNotification instantiates a new EmailNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationWithDefaults

`func NewEmailNotificationWithDefaults() *EmailNotification`

NewEmailNotificationWithDefaults instantiates a new EmailNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailNotification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocale

`func (o *EmailNotification) GetLocale() LocaleCode`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *EmailNotification) GetLocaleOk() (*LocaleCode, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *EmailNotification) SetLocale(v LocaleCode)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *EmailNotification) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetRecipientId

`func (o *EmailNotification) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EmailNotification) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EmailNotification) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EmailNotification) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetToName

`func (o *EmailNotification) GetToName() string`

GetToName returns the ToName field if non-nil, zero value otherwise.

### GetToNameOk

`func (o *EmailNotification) GetToNameOk() (*string, bool)`

GetToNameOk returns a tuple with the ToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToName

`func (o *EmailNotification) SetToName(v string)`

SetToName sets ToName field to given value.

### HasToName

`func (o *EmailNotification) HasToName() bool`

HasToName returns a boolean if a field has been set.

### GetToEmailAddress

`func (o *EmailNotification) GetToEmailAddress() string`

GetToEmailAddress returns the ToEmailAddress field if non-nil, zero value otherwise.

### GetToEmailAddressOk

`func (o *EmailNotification) GetToEmailAddressOk() (*string, bool)`

GetToEmailAddressOk returns a tuple with the ToEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEmailAddress

`func (o *EmailNotification) SetToEmailAddress(v string)`

SetToEmailAddress sets ToEmailAddress field to given value.

### HasToEmailAddress

`func (o *EmailNotification) HasToEmailAddress() bool`

HasToEmailAddress returns a boolean if a field has been set.

### GetFromName

`func (o *EmailNotification) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *EmailNotification) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *EmailNotification) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *EmailNotification) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetFromEmailAddress

`func (o *EmailNotification) GetFromEmailAddress() string`

GetFromEmailAddress returns the FromEmailAddress field if non-nil, zero value otherwise.

### GetFromEmailAddressOk

`func (o *EmailNotification) GetFromEmailAddressOk() (*string, bool)`

GetFromEmailAddressOk returns a tuple with the FromEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmailAddress

`func (o *EmailNotification) SetFromEmailAddress(v string)`

SetFromEmailAddress sets FromEmailAddress field to given value.

### HasFromEmailAddress

`func (o *EmailNotification) HasFromEmailAddress() bool`

HasFromEmailAddress returns a boolean if a field has been set.

### GetStatus

`func (o *EmailNotification) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailNotification) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailNotification) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomParams

`func (o *EmailNotification) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *EmailNotification) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *EmailNotification) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *EmailNotification) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmailNotification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailNotification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailNotification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmailNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


