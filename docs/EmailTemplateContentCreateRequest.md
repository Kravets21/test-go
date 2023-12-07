# EmailTemplateContentCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | The is a subject of the email notification. | 
**TextHtml** | **string** | This is an HTML part of the body of the email notification. | 
**TextPlain** | Pointer to **string** | This is a TXT part of the body of the email notification. | [optional] 
**TextAmp** | Pointer to **string** | This is an AMP part of the body of the email notification. | [optional] 
**BuilderHtml** | Pointer to **string** | This is settings for UI Template Builder. Must be valid JSON. | [optional] 

## Methods

### NewEmailTemplateContentCreateRequest

`func NewEmailTemplateContentCreateRequest(subject string, textHtml string, ) *EmailTemplateContentCreateRequest`

NewEmailTemplateContentCreateRequest instantiates a new EmailTemplateContentCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateContentCreateRequestWithDefaults

`func NewEmailTemplateContentCreateRequestWithDefaults() *EmailTemplateContentCreateRequest`

NewEmailTemplateContentCreateRequestWithDefaults instantiates a new EmailTemplateContentCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *EmailTemplateContentCreateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailTemplateContentCreateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailTemplateContentCreateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTextHtml

`func (o *EmailTemplateContentCreateRequest) GetTextHtml() string`

GetTextHtml returns the TextHtml field if non-nil, zero value otherwise.

### GetTextHtmlOk

`func (o *EmailTemplateContentCreateRequest) GetTextHtmlOk() (*string, bool)`

GetTextHtmlOk returns a tuple with the TextHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextHtml

`func (o *EmailTemplateContentCreateRequest) SetTextHtml(v string)`

SetTextHtml sets TextHtml field to given value.


### GetTextPlain

`func (o *EmailTemplateContentCreateRequest) GetTextPlain() string`

GetTextPlain returns the TextPlain field if non-nil, zero value otherwise.

### GetTextPlainOk

`func (o *EmailTemplateContentCreateRequest) GetTextPlainOk() (*string, bool)`

GetTextPlainOk returns a tuple with the TextPlain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextPlain

`func (o *EmailTemplateContentCreateRequest) SetTextPlain(v string)`

SetTextPlain sets TextPlain field to given value.

### HasTextPlain

`func (o *EmailTemplateContentCreateRequest) HasTextPlain() bool`

HasTextPlain returns a boolean if a field has been set.

### GetTextAmp

`func (o *EmailTemplateContentCreateRequest) GetTextAmp() string`

GetTextAmp returns the TextAmp field if non-nil, zero value otherwise.

### GetTextAmpOk

`func (o *EmailTemplateContentCreateRequest) GetTextAmpOk() (*string, bool)`

GetTextAmpOk returns a tuple with the TextAmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAmp

`func (o *EmailTemplateContentCreateRequest) SetTextAmp(v string)`

SetTextAmp sets TextAmp field to given value.

### HasTextAmp

`func (o *EmailTemplateContentCreateRequest) HasTextAmp() bool`

HasTextAmp returns a boolean if a field has been set.

### GetBuilderHtml

`func (o *EmailTemplateContentCreateRequest) GetBuilderHtml() string`

GetBuilderHtml returns the BuilderHtml field if non-nil, zero value otherwise.

### GetBuilderHtmlOk

`func (o *EmailTemplateContentCreateRequest) GetBuilderHtmlOk() (*string, bool)`

GetBuilderHtmlOk returns a tuple with the BuilderHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderHtml

`func (o *EmailTemplateContentCreateRequest) SetBuilderHtml(v string)`

SetBuilderHtml sets BuilderHtml field to given value.

### HasBuilderHtml

`func (o *EmailTemplateContentCreateRequest) HasBuilderHtml() bool`

HasBuilderHtml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


