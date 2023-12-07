# EmailTemplateCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionLinks**](CollectionLinks.md) |  | [optional] 
**Meta** | Pointer to [**CollectionMeta**](CollectionMeta.md) |  | [optional] 
**Data** | Pointer to [**[]EmailTemplate**](EmailTemplate.md) |  | [optional] 

## Methods

### NewEmailTemplateCollectionResponse

`func NewEmailTemplateCollectionResponse() *EmailTemplateCollectionResponse`

NewEmailTemplateCollectionResponse instantiates a new EmailTemplateCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateCollectionResponseWithDefaults

`func NewEmailTemplateCollectionResponseWithDefaults() *EmailTemplateCollectionResponse`

NewEmailTemplateCollectionResponseWithDefaults instantiates a new EmailTemplateCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EmailTemplateCollectionResponse) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmailTemplateCollectionResponse) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmailTemplateCollectionResponse) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EmailTemplateCollectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *EmailTemplateCollectionResponse) GetMeta() CollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EmailTemplateCollectionResponse) GetMetaOk() (*CollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EmailTemplateCollectionResponse) SetMeta(v CollectionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EmailTemplateCollectionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *EmailTemplateCollectionResponse) GetData() []EmailTemplate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmailTemplateCollectionResponse) GetDataOk() (*[]EmailTemplate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmailTemplateCollectionResponse) SetData(v []EmailTemplate)`

SetData sets Data field to given value.

### HasData

`func (o *EmailTemplateCollectionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


