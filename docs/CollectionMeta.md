# CollectionMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The base URL for the current request. | [optional] 
**PerPage** | Pointer to **int32** | The number of items to be shown on each page. | [optional] 
**NextCursor** | Pointer to **NullableString** | The cursor to the next page of the results. | [optional] 
**PrevCursor** | Pointer to **NullableString** | The cursor to the previous page of the results. | [optional] 

## Methods

### NewCollectionMeta

`func NewCollectionMeta() *CollectionMeta`

NewCollectionMeta instantiates a new CollectionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionMetaWithDefaults

`func NewCollectionMetaWithDefaults() *CollectionMeta`

NewCollectionMetaWithDefaults instantiates a new CollectionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CollectionMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CollectionMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CollectionMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CollectionMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPerPage

`func (o *CollectionMeta) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *CollectionMeta) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *CollectionMeta) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *CollectionMeta) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetNextCursor

`func (o *CollectionMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CollectionMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CollectionMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *CollectionMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *CollectionMeta) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *CollectionMeta) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPrevCursor

`func (o *CollectionMeta) GetPrevCursor() string`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *CollectionMeta) GetPrevCursorOk() (*string, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *CollectionMeta) SetPrevCursor(v string)`

SetPrevCursor sets PrevCursor field to given value.

### HasPrevCursor

`func (o *CollectionMeta) HasPrevCursor() bool`

HasPrevCursor returns a boolean if a field has been set.

### SetPrevCursorNil

`func (o *CollectionMeta) SetPrevCursorNil(b bool)`

 SetPrevCursorNil sets the value for PrevCursor to be an explicit nil

### UnsetPrevCursor
`func (o *CollectionMeta) UnsetPrevCursor()`

UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


