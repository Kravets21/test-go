# EmailTemplateContentUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **NullableString** | The is a subject of the email notification. | [optional] 
**TextHtml** | Pointer to **NullableString** | This is an HTML part of the body of the email notification. | [optional] 
**TextPlain** | Pointer to **NullableString** | This is a TXT part of the body of the email notification. | [optional] 
**TextAmp** | Pointer to **NullableString** | This is an AMP part of the body of the email notification. | [optional] 
**BuilderHtml** | Pointer to **NullableString** | This is settings for UI Template Builder. Must be valid JSON. | [optional] 

## Methods

### NewEmailTemplateContentUpdateRequest

`func NewEmailTemplateContentUpdateRequest() *EmailTemplateContentUpdateRequest`

NewEmailTemplateContentUpdateRequest instantiates a new EmailTemplateContentUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateContentUpdateRequestWithDefaults

`func NewEmailTemplateContentUpdateRequestWithDefaults() *EmailTemplateContentUpdateRequest`

NewEmailTemplateContentUpdateRequestWithDefaults instantiates a new EmailTemplateContentUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *EmailTemplateContentUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailTemplateContentUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailTemplateContentUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailTemplateContentUpdateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *EmailTemplateContentUpdateRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *EmailTemplateContentUpdateRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTextHtml

`func (o *EmailTemplateContentUpdateRequest) GetTextHtml() string`

GetTextHtml returns the TextHtml field if non-nil, zero value otherwise.

### GetTextHtmlOk

`func (o *EmailTemplateContentUpdateRequest) GetTextHtmlOk() (*string, bool)`

GetTextHtmlOk returns a tuple with the TextHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextHtml

`func (o *EmailTemplateContentUpdateRequest) SetTextHtml(v string)`

SetTextHtml sets TextHtml field to given value.

### HasTextHtml

`func (o *EmailTemplateContentUpdateRequest) HasTextHtml() bool`

HasTextHtml returns a boolean if a field has been set.

### SetTextHtmlNil

`func (o *EmailTemplateContentUpdateRequest) SetTextHtmlNil(b bool)`

 SetTextHtmlNil sets the value for TextHtml to be an explicit nil

### UnsetTextHtml
`func (o *EmailTemplateContentUpdateRequest) UnsetTextHtml()`

UnsetTextHtml ensures that no value is present for TextHtml, not even an explicit nil
### GetTextPlain

`func (o *EmailTemplateContentUpdateRequest) GetTextPlain() string`

GetTextPlain returns the TextPlain field if non-nil, zero value otherwise.

### GetTextPlainOk

`func (o *EmailTemplateContentUpdateRequest) GetTextPlainOk() (*string, bool)`

GetTextPlainOk returns a tuple with the TextPlain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextPlain

`func (o *EmailTemplateContentUpdateRequest) SetTextPlain(v string)`

SetTextPlain sets TextPlain field to given value.

### HasTextPlain

`func (o *EmailTemplateContentUpdateRequest) HasTextPlain() bool`

HasTextPlain returns a boolean if a field has been set.

### SetTextPlainNil

`func (o *EmailTemplateContentUpdateRequest) SetTextPlainNil(b bool)`

 SetTextPlainNil sets the value for TextPlain to be an explicit nil

### UnsetTextPlain
`func (o *EmailTemplateContentUpdateRequest) UnsetTextPlain()`

UnsetTextPlain ensures that no value is present for TextPlain, not even an explicit nil
### GetTextAmp

`func (o *EmailTemplateContentUpdateRequest) GetTextAmp() string`

GetTextAmp returns the TextAmp field if non-nil, zero value otherwise.

### GetTextAmpOk

`func (o *EmailTemplateContentUpdateRequest) GetTextAmpOk() (*string, bool)`

GetTextAmpOk returns a tuple with the TextAmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAmp

`func (o *EmailTemplateContentUpdateRequest) SetTextAmp(v string)`

SetTextAmp sets TextAmp field to given value.

### HasTextAmp

`func (o *EmailTemplateContentUpdateRequest) HasTextAmp() bool`

HasTextAmp returns a boolean if a field has been set.

### SetTextAmpNil

`func (o *EmailTemplateContentUpdateRequest) SetTextAmpNil(b bool)`

 SetTextAmpNil sets the value for TextAmp to be an explicit nil

### UnsetTextAmp
`func (o *EmailTemplateContentUpdateRequest) UnsetTextAmp()`

UnsetTextAmp ensures that no value is present for TextAmp, not even an explicit nil
### GetBuilderHtml

`func (o *EmailTemplateContentUpdateRequest) GetBuilderHtml() string`

GetBuilderHtml returns the BuilderHtml field if non-nil, zero value otherwise.

### GetBuilderHtmlOk

`func (o *EmailTemplateContentUpdateRequest) GetBuilderHtmlOk() (*string, bool)`

GetBuilderHtmlOk returns a tuple with the BuilderHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderHtml

`func (o *EmailTemplateContentUpdateRequest) SetBuilderHtml(v string)`

SetBuilderHtml sets BuilderHtml field to given value.

### HasBuilderHtml

`func (o *EmailTemplateContentUpdateRequest) HasBuilderHtml() bool`

HasBuilderHtml returns a boolean if a field has been set.

### SetBuilderHtmlNil

`func (o *EmailTemplateContentUpdateRequest) SetBuilderHtmlNil(b bool)`

 SetBuilderHtmlNil sets the value for BuilderHtml to be an explicit nil

### UnsetBuilderHtml
`func (o *EmailTemplateContentUpdateRequest) UnsetBuilderHtml()`

UnsetBuilderHtml ensures that no value is present for BuilderHtml, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


