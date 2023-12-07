# AudienceTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Tag. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Tag was created. | [optional] 

## Methods

### NewAudienceTag

`func NewAudienceTag() *AudienceTag`

NewAudienceTag instantiates a new AudienceTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceTagWithDefaults

`func NewAudienceTagWithDefaults() *AudienceTag`

NewAudienceTagWithDefaults instantiates a new AudienceTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AudienceTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudienceTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudienceTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AudienceTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AudienceTag) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AudienceTag) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AudienceTag) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AudienceTag) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


