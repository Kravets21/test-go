# EmailNotificationDetails

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
**Opens** | Pointer to [**[]EmailNotificationOpen**](EmailNotificationOpen.md) |  | [optional] 
**Clicks** | Pointer to [**[]EmailNotificationClick**](EmailNotificationClick.md) |  | [optional] 

## Methods

### NewEmailNotificationDetails

`func NewEmailNotificationDetails() *EmailNotificationDetails`

NewEmailNotificationDetails instantiates a new EmailNotificationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationDetailsWithDefaults

`func NewEmailNotificationDetailsWithDefaults() *EmailNotificationDetails`

NewEmailNotificationDetailsWithDefaults instantiates a new EmailNotificationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailNotificationDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailNotificationDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailNotificationDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailNotificationDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocale

`func (o *EmailNotificationDetails) GetLocale() LocaleCode`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *EmailNotificationDetails) GetLocaleOk() (*LocaleCode, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *EmailNotificationDetails) SetLocale(v LocaleCode)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *EmailNotificationDetails) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetRecipientId

`func (o *EmailNotificationDetails) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EmailNotificationDetails) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EmailNotificationDetails) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EmailNotificationDetails) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetToName

`func (o *EmailNotificationDetails) GetToName() string`

GetToName returns the ToName field if non-nil, zero value otherwise.

### GetToNameOk

`func (o *EmailNotificationDetails) GetToNameOk() (*string, bool)`

GetToNameOk returns a tuple with the ToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToName

`func (o *EmailNotificationDetails) SetToName(v string)`

SetToName sets ToName field to given value.

### HasToName

`func (o *EmailNotificationDetails) HasToName() bool`

HasToName returns a boolean if a field has been set.

### GetToEmailAddress

`func (o *EmailNotificationDetails) GetToEmailAddress() string`

GetToEmailAddress returns the ToEmailAddress field if non-nil, zero value otherwise.

### GetToEmailAddressOk

`func (o *EmailNotificationDetails) GetToEmailAddressOk() (*string, bool)`

GetToEmailAddressOk returns a tuple with the ToEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEmailAddress

`func (o *EmailNotificationDetails) SetToEmailAddress(v string)`

SetToEmailAddress sets ToEmailAddress field to given value.

### HasToEmailAddress

`func (o *EmailNotificationDetails) HasToEmailAddress() bool`

HasToEmailAddress returns a boolean if a field has been set.

### GetFromName

`func (o *EmailNotificationDetails) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *EmailNotificationDetails) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *EmailNotificationDetails) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *EmailNotificationDetails) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetFromEmailAddress

`func (o *EmailNotificationDetails) GetFromEmailAddress() string`

GetFromEmailAddress returns the FromEmailAddress field if non-nil, zero value otherwise.

### GetFromEmailAddressOk

`func (o *EmailNotificationDetails) GetFromEmailAddressOk() (*string, bool)`

GetFromEmailAddressOk returns a tuple with the FromEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmailAddress

`func (o *EmailNotificationDetails) SetFromEmailAddress(v string)`

SetFromEmailAddress sets FromEmailAddress field to given value.

### HasFromEmailAddress

`func (o *EmailNotificationDetails) HasFromEmailAddress() bool`

HasFromEmailAddress returns a boolean if a field has been set.

### GetStatus

`func (o *EmailNotificationDetails) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailNotificationDetails) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailNotificationDetails) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailNotificationDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomParams

`func (o *EmailNotificationDetails) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *EmailNotificationDetails) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *EmailNotificationDetails) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *EmailNotificationDetails) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmailNotificationDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailNotificationDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailNotificationDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmailNotificationDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOpens

`func (o *EmailNotificationDetails) GetOpens() []EmailNotificationOpen`

GetOpens returns the Opens field if non-nil, zero value otherwise.

### GetOpensOk

`func (o *EmailNotificationDetails) GetOpensOk() (*[]EmailNotificationOpen, bool)`

GetOpensOk returns a tuple with the Opens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpens

`func (o *EmailNotificationDetails) SetOpens(v []EmailNotificationOpen)`

SetOpens sets Opens field to given value.

### HasOpens

`func (o *EmailNotificationDetails) HasOpens() bool`

HasOpens returns a boolean if a field has been set.

### GetClicks

`func (o *EmailNotificationDetails) GetClicks() []EmailNotificationClick`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *EmailNotificationDetails) GetClicksOk() (*[]EmailNotificationClick, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *EmailNotificationDetails) SetClicks(v []EmailNotificationClick)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *EmailNotificationDetails) HasClicks() bool`

HasClicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


