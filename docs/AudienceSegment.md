# AudienceSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Segment. | [optional] 
**Name** | Pointer to **string** | The name of the Segment. | [optional] 
**Status** | Pointer to **int32** | The status of the Segment. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Platform was created. | [optional] 

## Methods

### NewAudienceSegment

`func NewAudienceSegment() *AudienceSegment`

NewAudienceSegment instantiates a new AudienceSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceSegmentWithDefaults

`func NewAudienceSegmentWithDefaults() *AudienceSegment`

NewAudienceSegmentWithDefaults instantiates a new AudienceSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudienceSegment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceSegment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudienceSegment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AudienceSegment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AudienceSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudienceSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudienceSegment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AudienceSegment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AudienceSegment) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AudienceSegment) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AudienceSegment) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AudienceSegment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AudienceSegment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AudienceSegment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AudienceSegment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AudienceSegment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


