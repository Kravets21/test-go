# SmsRoutingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the routing rules. | [optional] 
**SenderId** | Pointer to **string** | The unique identifier of the Sender. | [optional] 
**IsReserved** | Pointer to **bool** | You can mark your number as Spare true and use it when main number can&#39;t send sms. | [optional] 
**AlphaNumber** | Pointer to **string** | You can use short alias, text Alphanumeric Sender ID instead digital number | [optional] 
**Priority** | Pointer to **int32** | Priority number will try to send sms firstly. | [optional] 
**Status** | Pointer to **int32** | The status of the Routing-rule. Every digit matters: 1 - active, 2 - inactive. | [optional] 
**CountryCode** | Pointer to [**LocaleCode**](LocaleCode.md) |  | [optional] 
**ExcludedCountries** | Pointer to [**[]LocaleCode**](LocaleCode.md) | List of excluded to send countries. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Channel Rule was created. | [optional] 

## Methods

### NewSmsRoutingRule

`func NewSmsRoutingRule() *SmsRoutingRule`

NewSmsRoutingRule instantiates a new SmsRoutingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsRoutingRuleWithDefaults

`func NewSmsRoutingRuleWithDefaults() *SmsRoutingRule`

NewSmsRoutingRuleWithDefaults instantiates a new SmsRoutingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsRoutingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsRoutingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsRoutingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsRoutingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSenderId

`func (o *SmsRoutingRule) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *SmsRoutingRule) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *SmsRoutingRule) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *SmsRoutingRule) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetIsReserved

`func (o *SmsRoutingRule) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *SmsRoutingRule) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *SmsRoutingRule) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.

### HasIsReserved

`func (o *SmsRoutingRule) HasIsReserved() bool`

HasIsReserved returns a boolean if a field has been set.

### GetAlphaNumber

`func (o *SmsRoutingRule) GetAlphaNumber() string`

GetAlphaNumber returns the AlphaNumber field if non-nil, zero value otherwise.

### GetAlphaNumberOk

`func (o *SmsRoutingRule) GetAlphaNumberOk() (*string, bool)`

GetAlphaNumberOk returns a tuple with the AlphaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaNumber

`func (o *SmsRoutingRule) SetAlphaNumber(v string)`

SetAlphaNumber sets AlphaNumber field to given value.

### HasAlphaNumber

`func (o *SmsRoutingRule) HasAlphaNumber() bool`

HasAlphaNumber returns a boolean if a field has been set.

### GetPriority

`func (o *SmsRoutingRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SmsRoutingRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SmsRoutingRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SmsRoutingRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *SmsRoutingRule) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsRoutingRule) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsRoutingRule) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsRoutingRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCountryCode

`func (o *SmsRoutingRule) GetCountryCode() LocaleCode`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SmsRoutingRule) GetCountryCodeOk() (*LocaleCode, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SmsRoutingRule) SetCountryCode(v LocaleCode)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SmsRoutingRule) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetExcludedCountries

`func (o *SmsRoutingRule) GetExcludedCountries() []LocaleCode`

GetExcludedCountries returns the ExcludedCountries field if non-nil, zero value otherwise.

### GetExcludedCountriesOk

`func (o *SmsRoutingRule) GetExcludedCountriesOk() (*[]LocaleCode, bool)`

GetExcludedCountriesOk returns a tuple with the ExcludedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCountries

`func (o *SmsRoutingRule) SetExcludedCountries(v []LocaleCode)`

SetExcludedCountries sets ExcludedCountries field to given value.

### HasExcludedCountries

`func (o *SmsRoutingRule) HasExcludedCountries() bool`

HasExcludedCountries returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SmsRoutingRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmsRoutingRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmsRoutingRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SmsRoutingRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


