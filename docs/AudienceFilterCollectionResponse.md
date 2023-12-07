# AudienceFilterCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionLinks**](CollectionLinks.md) |  | [optional] 
**Meta** | Pointer to [**CollectionMeta**](CollectionMeta.md) |  | [optional] 
**Data** | Pointer to [**[]AudienceFilter**](AudienceFilter.md) |  | [optional] 

## Methods

### NewAudienceFilterCollectionResponse

`func NewAudienceFilterCollectionResponse() *AudienceFilterCollectionResponse`

NewAudienceFilterCollectionResponse instantiates a new AudienceFilterCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceFilterCollectionResponseWithDefaults

`func NewAudienceFilterCollectionResponseWithDefaults() *AudienceFilterCollectionResponse`

NewAudienceFilterCollectionResponseWithDefaults instantiates a new AudienceFilterCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AudienceFilterCollectionResponse) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AudienceFilterCollectionResponse) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AudienceFilterCollectionResponse) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AudienceFilterCollectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AudienceFilterCollectionResponse) GetMeta() CollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AudienceFilterCollectionResponse) GetMetaOk() (*CollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AudienceFilterCollectionResponse) SetMeta(v CollectionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AudienceFilterCollectionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *AudienceFilterCollectionResponse) GetData() []AudienceFilter`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AudienceFilterCollectionResponse) GetDataOk() (*[]AudienceFilter, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AudienceFilterCollectionResponse) SetData(v []AudienceFilter)`

SetData sets Data field to given value.

### HasData

`func (o *AudienceFilterCollectionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


