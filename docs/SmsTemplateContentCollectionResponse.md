# SmsTemplateContentCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionLinks**](CollectionLinks.md) |  | [optional] 
**Meta** | Pointer to [**CollectionMeta**](CollectionMeta.md) |  | [optional] 
**Data** | Pointer to [**[]SmsTemplateContent**](SmsTemplateContent.md) |  | [optional] 

## Methods

### NewSmsTemplateContentCollectionResponse

`func NewSmsTemplateContentCollectionResponse() *SmsTemplateContentCollectionResponse`

NewSmsTemplateContentCollectionResponse instantiates a new SmsTemplateContentCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsTemplateContentCollectionResponseWithDefaults

`func NewSmsTemplateContentCollectionResponseWithDefaults() *SmsTemplateContentCollectionResponse`

NewSmsTemplateContentCollectionResponseWithDefaults instantiates a new SmsTemplateContentCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SmsTemplateContentCollectionResponse) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SmsTemplateContentCollectionResponse) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SmsTemplateContentCollectionResponse) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SmsTemplateContentCollectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *SmsTemplateContentCollectionResponse) GetMeta() CollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SmsTemplateContentCollectionResponse) GetMetaOk() (*CollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SmsTemplateContentCollectionResponse) SetMeta(v CollectionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SmsTemplateContentCollectionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *SmsTemplateContentCollectionResponse) GetData() []SmsTemplateContent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SmsTemplateContentCollectionResponse) GetDataOk() (*[]SmsTemplateContent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SmsTemplateContentCollectionResponse) SetData(v []SmsTemplateContent)`

SetData sets Data field to given value.

### HasData

`func (o *SmsTemplateContentCollectionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


