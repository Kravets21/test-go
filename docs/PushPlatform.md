# PushPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Platform. | [optional] 
**Name** | Pointer to **string** | The name of the Platform. | [optional] 
**Status** | Pointer to **int32** | The status of the Platform. Every digit matters: 1 - active, 2 - inactive. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Platform was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The datetime when the Platform was updated. | [optional] 

## Methods

### NewPushPlatform

`func NewPushPlatform() *PushPlatform`

NewPushPlatform instantiates a new PushPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushPlatformWithDefaults

`func NewPushPlatformWithDefaults() *PushPlatform`

NewPushPlatformWithDefaults instantiates a new PushPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PushPlatform) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PushPlatform) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PushPlatform) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PushPlatform) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PushPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PushPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PushPlatform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PushPlatform) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PushPlatform) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PushPlatform) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PushPlatform) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PushPlatform) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PushPlatform) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PushPlatform) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PushPlatform) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PushPlatform) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PushPlatform) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PushPlatform) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PushPlatform) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PushPlatform) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


