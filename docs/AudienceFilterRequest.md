# AudienceFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]AudienceRule**](AudienceRule.md) | The multiple filter rules. | [optional] 

## Methods

### NewAudienceFilterRequest

`func NewAudienceFilterRequest() *AudienceFilterRequest`

NewAudienceFilterRequest instantiates a new AudienceFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceFilterRequestWithDefaults

`func NewAudienceFilterRequestWithDefaults() *AudienceFilterRequest`

NewAudienceFilterRequestWithDefaults instantiates a new AudienceFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *AudienceFilterRequest) GetRules() []AudienceRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AudienceFilterRequest) GetRulesOk() (*[]AudienceRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AudienceFilterRequest) SetRules(v []AudienceRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AudienceFilterRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


