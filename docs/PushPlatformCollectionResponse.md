# PushPlatformCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionLinks**](CollectionLinks.md) |  | [optional] 
**Meta** | Pointer to [**CollectionMeta**](CollectionMeta.md) |  | [optional] 
**Data** | Pointer to [**[]PushPlatform**](PushPlatform.md) |  | [optional] 

## Methods

### NewPushPlatformCollectionResponse

`func NewPushPlatformCollectionResponse() *PushPlatformCollectionResponse`

NewPushPlatformCollectionResponse instantiates a new PushPlatformCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushPlatformCollectionResponseWithDefaults

`func NewPushPlatformCollectionResponseWithDefaults() *PushPlatformCollectionResponse`

NewPushPlatformCollectionResponseWithDefaults instantiates a new PushPlatformCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PushPlatformCollectionResponse) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PushPlatformCollectionResponse) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PushPlatformCollectionResponse) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PushPlatformCollectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *PushPlatformCollectionResponse) GetMeta() CollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PushPlatformCollectionResponse) GetMetaOk() (*CollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PushPlatformCollectionResponse) SetMeta(v CollectionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PushPlatformCollectionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *PushPlatformCollectionResponse) GetData() []PushPlatform`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PushPlatformCollectionResponse) GetDataOk() (*[]PushPlatform, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PushPlatformCollectionResponse) SetData(v []PushPlatform)`

SetData sets Data field to given value.

### HasData

`func (o *PushPlatformCollectionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


