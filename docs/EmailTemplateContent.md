# EmailTemplateContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to [**LocaleCode**](LocaleCode.md) |  | [optional] 
**Subject** | Pointer to **string** | The is a subject of the email notification. | [optional] 
**TextPlain** | Pointer to **string** | The is a text body of the email notification. | [optional] 
**TextHtml** | Pointer to **string** | The is a html body of the email notification. Must be valid HTML. | [optional] 
**TextAmp** | Pointer to **string** | The is a amp body of the email notification. | [optional] 
**BuilderHtml** | Pointer to **string** | This is settings for UI Template Builder. Must be valid JSON. | [optional] 
**VariablesSettings** | Pointer to [**[]EmailTemplateContentVariablesSettingsInner**](EmailTemplateContentVariablesSettingsInner.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Template Content was created. | [optional] 

## Methods

### NewEmailTemplateContent

`func NewEmailTemplateContent() *EmailTemplateContent`

NewEmailTemplateContent instantiates a new EmailTemplateContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateContentWithDefaults

`func NewEmailTemplateContentWithDefaults() *EmailTemplateContent`

NewEmailTemplateContentWithDefaults instantiates a new EmailTemplateContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *EmailTemplateContent) GetLocale() LocaleCode`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *EmailTemplateContent) GetLocaleOk() (*LocaleCode, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *EmailTemplateContent) SetLocale(v LocaleCode)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *EmailTemplateContent) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetSubject

`func (o *EmailTemplateContent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailTemplateContent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailTemplateContent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailTemplateContent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTextPlain

`func (o *EmailTemplateContent) GetTextPlain() string`

GetTextPlain returns the TextPlain field if non-nil, zero value otherwise.

### GetTextPlainOk

`func (o *EmailTemplateContent) GetTextPlainOk() (*string, bool)`

GetTextPlainOk returns a tuple with the TextPlain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextPlain

`func (o *EmailTemplateContent) SetTextPlain(v string)`

SetTextPlain sets TextPlain field to given value.

### HasTextPlain

`func (o *EmailTemplateContent) HasTextPlain() bool`

HasTextPlain returns a boolean if a field has been set.

### GetTextHtml

`func (o *EmailTemplateContent) GetTextHtml() string`

GetTextHtml returns the TextHtml field if non-nil, zero value otherwise.

### GetTextHtmlOk

`func (o *EmailTemplateContent) GetTextHtmlOk() (*string, bool)`

GetTextHtmlOk returns a tuple with the TextHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextHtml

`func (o *EmailTemplateContent) SetTextHtml(v string)`

SetTextHtml sets TextHtml field to given value.

### HasTextHtml

`func (o *EmailTemplateContent) HasTextHtml() bool`

HasTextHtml returns a boolean if a field has been set.

### GetTextAmp

`func (o *EmailTemplateContent) GetTextAmp() string`

GetTextAmp returns the TextAmp field if non-nil, zero value otherwise.

### GetTextAmpOk

`func (o *EmailTemplateContent) GetTextAmpOk() (*string, bool)`

GetTextAmpOk returns a tuple with the TextAmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAmp

`func (o *EmailTemplateContent) SetTextAmp(v string)`

SetTextAmp sets TextAmp field to given value.

### HasTextAmp

`func (o *EmailTemplateContent) HasTextAmp() bool`

HasTextAmp returns a boolean if a field has been set.

### GetBuilderHtml

`func (o *EmailTemplateContent) GetBuilderHtml() string`

GetBuilderHtml returns the BuilderHtml field if non-nil, zero value otherwise.

### GetBuilderHtmlOk

`func (o *EmailTemplateContent) GetBuilderHtmlOk() (*string, bool)`

GetBuilderHtmlOk returns a tuple with the BuilderHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderHtml

`func (o *EmailTemplateContent) SetBuilderHtml(v string)`

SetBuilderHtml sets BuilderHtml field to given value.

### HasBuilderHtml

`func (o *EmailTemplateContent) HasBuilderHtml() bool`

HasBuilderHtml returns a boolean if a field has been set.

### GetVariablesSettings

`func (o *EmailTemplateContent) GetVariablesSettings() []EmailTemplateContentVariablesSettingsInner`

GetVariablesSettings returns the VariablesSettings field if non-nil, zero value otherwise.

### GetVariablesSettingsOk

`func (o *EmailTemplateContent) GetVariablesSettingsOk() (*[]EmailTemplateContentVariablesSettingsInner, bool)`

GetVariablesSettingsOk returns a tuple with the VariablesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesSettings

`func (o *EmailTemplateContent) SetVariablesSettings(v []EmailTemplateContentVariablesSettingsInner)`

SetVariablesSettings sets VariablesSettings field to given value.

### HasVariablesSettings

`func (o *EmailTemplateContent) HasVariablesSettings() bool`

HasVariablesSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmailTemplateContent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailTemplateContent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailTemplateContent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmailTemplateContent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


