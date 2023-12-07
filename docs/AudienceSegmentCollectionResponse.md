# AudienceSegmentCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionLinks**](CollectionLinks.md) |  | [optional] 
**Meta** | Pointer to [**CollectionMeta**](CollectionMeta.md) |  | [optional] 
**Data** | Pointer to [**[]AudienceSegment**](AudienceSegment.md) |  | [optional] 

## Methods

### NewAudienceSegmentCollectionResponse

`func NewAudienceSegmentCollectionResponse() *AudienceSegmentCollectionResponse`

NewAudienceSegmentCollectionResponse instantiates a new AudienceSegmentCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceSegmentCollectionResponseWithDefaults

`func NewAudienceSegmentCollectionResponseWithDefaults() *AudienceSegmentCollectionResponse`

NewAudienceSegmentCollectionResponseWithDefaults instantiates a new AudienceSegmentCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AudienceSegmentCollectionResponse) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AudienceSegmentCollectionResponse) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AudienceSegmentCollectionResponse) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AudienceSegmentCollectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *AudienceSegmentCollectionResponse) GetMeta() CollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AudienceSegmentCollectionResponse) GetMetaOk() (*CollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AudienceSegmentCollectionResponse) SetMeta(v CollectionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AudienceSegmentCollectionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *AudienceSegmentCollectionResponse) GetData() []AudienceSegment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AudienceSegmentCollectionResponse) GetDataOk() (*[]AudienceSegment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AudienceSegmentCollectionResponse) SetData(v []AudienceSegment)`

SetData sets Data field to given value.

### HasData

`func (o *AudienceSegmentCollectionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


