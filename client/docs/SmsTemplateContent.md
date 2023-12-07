# SmsTemplateContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to [**LocaleCode**](LocaleCode.md) |  | [optional] 
**Body** | Pointer to **string** | This is the body for SMS Notification. | [optional] 
**VariablesSettings** | Pointer to [**[]EmailTemplateContentVariablesSettingsInner**](EmailTemplateContentVariablesSettingsInner.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Template Content was created. | [optional] 

## Methods

### NewSmsTemplateContent

`func NewSmsTemplateContent() *SmsTemplateContent`

NewSmsTemplateContent instantiates a new SmsTemplateContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsTemplateContentWithDefaults

`func NewSmsTemplateContentWithDefaults() *SmsTemplateContent`

NewSmsTemplateContentWithDefaults instantiates a new SmsTemplateContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *SmsTemplateContent) GetLocale() LocaleCode`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SmsTemplateContent) GetLocaleOk() (*LocaleCode, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SmsTemplateContent) SetLocale(v LocaleCode)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SmsTemplateContent) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetBody

`func (o *SmsTemplateContent) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SmsTemplateContent) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SmsTemplateContent) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SmsTemplateContent) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetVariablesSettings

`func (o *SmsTemplateContent) GetVariablesSettings() []EmailTemplateContentVariablesSettingsInner`

GetVariablesSettings returns the VariablesSettings field if non-nil, zero value otherwise.

### GetVariablesSettingsOk

`func (o *SmsTemplateContent) GetVariablesSettingsOk() (*[]EmailTemplateContentVariablesSettingsInner, bool)`

GetVariablesSettingsOk returns a tuple with the VariablesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesSettings

`func (o *SmsTemplateContent) SetVariablesSettings(v []EmailTemplateContentVariablesSettingsInner)`

SetVariablesSettings sets VariablesSettings field to given value.

### HasVariablesSettings

`func (o *SmsTemplateContent) HasVariablesSettings() bool`

HasVariablesSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SmsTemplateContent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmsTemplateContent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmsTemplateContent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SmsTemplateContent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


