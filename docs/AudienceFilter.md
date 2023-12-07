# AudienceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Segment. | [optional] 
**Rules** | Pointer to [**[]AudienceRule**](AudienceRule.md) | The rules of the Filter. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Platform was created. | [optional] 

## Methods

### NewAudienceFilter

`func NewAudienceFilter() *AudienceFilter`

NewAudienceFilter instantiates a new AudienceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceFilterWithDefaults

`func NewAudienceFilterWithDefaults() *AudienceFilter`

NewAudienceFilterWithDefaults instantiates a new AudienceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudienceFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudienceFilter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AudienceFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRules

`func (o *AudienceFilter) GetRules() []AudienceRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AudienceFilter) GetRulesOk() (*[]AudienceRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AudienceFilter) SetRules(v []AudienceRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AudienceFilter) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AudienceFilter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AudienceFilter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AudienceFilter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AudienceFilter) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


