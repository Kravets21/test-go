# EmailTemplateContentVariablesSettingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The unique variable path. | [optional] 
**DefaultValue** | Pointer to **string** | Optional default value for the variable. | [optional] 
**RecipientProperty** | Pointer to [**EmailTemplateContentVariablesSettingsInnerRecipientProperty**](EmailTemplateContentVariablesSettingsInnerRecipientProperty.md) |  | [optional] 

## Methods

### NewEmailTemplateContentVariablesSettingsInner

`func NewEmailTemplateContentVariablesSettingsInner() *EmailTemplateContentVariablesSettingsInner`

NewEmailTemplateContentVariablesSettingsInner instantiates a new EmailTemplateContentVariablesSettingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateContentVariablesSettingsInnerWithDefaults

`func NewEmailTemplateContentVariablesSettingsInnerWithDefaults() *EmailTemplateContentVariablesSettingsInner`

NewEmailTemplateContentVariablesSettingsInnerWithDefaults instantiates a new EmailTemplateContentVariablesSettingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *EmailTemplateContentVariablesSettingsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EmailTemplateContentVariablesSettingsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EmailTemplateContentVariablesSettingsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EmailTemplateContentVariablesSettingsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDefaultValue

`func (o *EmailTemplateContentVariablesSettingsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *EmailTemplateContentVariablesSettingsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *EmailTemplateContentVariablesSettingsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *EmailTemplateContentVariablesSettingsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRecipientProperty

`func (o *EmailTemplateContentVariablesSettingsInner) GetRecipientProperty() EmailTemplateContentVariablesSettingsInnerRecipientProperty`

GetRecipientProperty returns the RecipientProperty field if non-nil, zero value otherwise.

### GetRecipientPropertyOk

`func (o *EmailTemplateContentVariablesSettingsInner) GetRecipientPropertyOk() (*EmailTemplateContentVariablesSettingsInnerRecipientProperty, bool)`

GetRecipientPropertyOk returns a tuple with the RecipientProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientProperty

`func (o *EmailTemplateContentVariablesSettingsInner) SetRecipientProperty(v EmailTemplateContentVariablesSettingsInnerRecipientProperty)`

SetRecipientProperty sets RecipientProperty field to given value.

### HasRecipientProperty

`func (o *EmailTemplateContentVariablesSettingsInner) HasRecipientProperty() bool`

HasRecipientProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


