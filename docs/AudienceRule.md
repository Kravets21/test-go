# AudienceRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to [**AudienceRuleProperty**](AudienceRuleProperty.md) |  | [optional] 
**Operator** | Pointer to **string** | The operator of the filter. | [optional] 
**Value** | Pointer to **string** | The value of the filter. | [optional] 

## Methods

### NewAudienceRule

`func NewAudienceRule() *AudienceRule`

NewAudienceRule instantiates a new AudienceRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceRuleWithDefaults

`func NewAudienceRuleWithDefaults() *AudienceRule`

NewAudienceRuleWithDefaults instantiates a new AudienceRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *AudienceRule) GetProperty() AudienceRuleProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AudienceRule) GetPropertyOk() (*AudienceRuleProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AudienceRule) SetProperty(v AudienceRuleProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AudienceRule) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *AudienceRule) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AudienceRule) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AudienceRule) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AudienceRule) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *AudienceRule) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AudienceRule) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AudienceRule) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AudienceRule) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


