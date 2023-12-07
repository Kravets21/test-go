# AudienceRecipientSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Segment. | [optional] 
**Name** | Pointer to **string** | The name of the Segment. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Recipient was added to the Segment. | [optional] 

## Methods

### NewAudienceRecipientSegment

`func NewAudienceRecipientSegment() *AudienceRecipientSegment`

NewAudienceRecipientSegment instantiates a new AudienceRecipientSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceRecipientSegmentWithDefaults

`func NewAudienceRecipientSegmentWithDefaults() *AudienceRecipientSegment`

NewAudienceRecipientSegmentWithDefaults instantiates a new AudienceRecipientSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudienceRecipientSegment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceRecipientSegment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudienceRecipientSegment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AudienceRecipientSegment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AudienceRecipientSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudienceRecipientSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudienceRecipientSegment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AudienceRecipientSegment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AudienceRecipientSegment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AudienceRecipientSegment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AudienceRecipientSegment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AudienceRecipientSegment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


