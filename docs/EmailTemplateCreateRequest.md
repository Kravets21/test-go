# EmailTemplateCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** | The unique identifier of the Channel. | 
**Name** | **string** | The name of the Template. | 
**Description** | Pointer to **string** | The description of the Template. | [optional] 

## Methods

### NewEmailTemplateCreateRequest

`func NewEmailTemplateCreateRequest(channelId string, name string, ) *EmailTemplateCreateRequest`

NewEmailTemplateCreateRequest instantiates a new EmailTemplateCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateCreateRequestWithDefaults

`func NewEmailTemplateCreateRequestWithDefaults() *EmailTemplateCreateRequest`

NewEmailTemplateCreateRequestWithDefaults instantiates a new EmailTemplateCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *EmailTemplateCreateRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *EmailTemplateCreateRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *EmailTemplateCreateRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetName

`func (o *EmailTemplateCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailTemplateCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailTemplateCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EmailTemplateCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailTemplateCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailTemplateCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailTemplateCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


