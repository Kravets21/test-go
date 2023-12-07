# SmsTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Template. | [optional] 
**ChannelId** | Pointer to **string** | The unique identifier of the Channel. | [optional] 
**Name** | Pointer to **string** | The name of the Template. | [optional] 
**Description** | Pointer to **string** | The description of the Template. | [optional] 
**Locales** | Pointer to **[]string** | List of locales available for this Template. | [optional] 
**Status** | Pointer to **int32** | The status of the Template. Every digit matters: 1 - active, 2 - inactive. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Template was created. | [optional] 

## Methods

### NewSmsTemplate

`func NewSmsTemplate() *SmsTemplate`

NewSmsTemplate instantiates a new SmsTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsTemplateWithDefaults

`func NewSmsTemplateWithDefaults() *SmsTemplate`

NewSmsTemplateWithDefaults instantiates a new SmsTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelId

`func (o *SmsTemplate) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SmsTemplate) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SmsTemplate) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SmsTemplate) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetName

`func (o *SmsTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmsTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmsTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmsTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SmsTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SmsTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SmsTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SmsTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocales

`func (o *SmsTemplate) GetLocales() []string`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *SmsTemplate) GetLocalesOk() (*[]string, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *SmsTemplate) SetLocales(v []string)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *SmsTemplate) HasLocales() bool`

HasLocales returns a boolean if a field has been set.

### GetStatus

`func (o *SmsTemplate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsTemplate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsTemplate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SmsTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmsTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmsTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SmsTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


