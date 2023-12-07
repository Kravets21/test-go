# SmsNotificationSendRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the template to use for the notification. The id or name parameter must be specified. | [optional] 
**Name** | Pointer to **string** | The name of the template to use for the notification. The id or name parameter must be specified. | [optional] 
**Locale** | Pointer to **string** | The locale of the template to use for the notification. | [optional] 
**Variables** | Pointer to **map[string]string** | The parameters used to replace placeholders of the Template Content Where key - is a placeholder name, and value - is a placeholder value. If the value is not provided, the placeholder will use the default value from the Template Content. | [optional] 

## Methods

### NewSmsNotificationSendRequestTemplate

`func NewSmsNotificationSendRequestTemplate() *SmsNotificationSendRequestTemplate`

NewSmsNotificationSendRequestTemplate instantiates a new SmsNotificationSendRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsNotificationSendRequestTemplateWithDefaults

`func NewSmsNotificationSendRequestTemplateWithDefaults() *SmsNotificationSendRequestTemplate`

NewSmsNotificationSendRequestTemplateWithDefaults instantiates a new SmsNotificationSendRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsNotificationSendRequestTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsNotificationSendRequestTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsNotificationSendRequestTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsNotificationSendRequestTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SmsNotificationSendRequestTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmsNotificationSendRequestTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmsNotificationSendRequestTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmsNotificationSendRequestTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocale

`func (o *SmsNotificationSendRequestTemplate) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SmsNotificationSendRequestTemplate) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SmsNotificationSendRequestTemplate) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SmsNotificationSendRequestTemplate) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetVariables

`func (o *SmsNotificationSendRequestTemplate) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SmsNotificationSendRequestTemplate) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SmsNotificationSendRequestTemplate) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SmsNotificationSendRequestTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


