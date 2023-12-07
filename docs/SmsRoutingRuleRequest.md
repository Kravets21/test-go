# SmsRoutingRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderId** | **string** | The unique identifier of the Sender. | 
**CountryCode** | Pointer to [**LocaleCode**](LocaleCode.md) |  | [optional] 
**ExcludedCountries** | Pointer to [**[]LocaleCode**](LocaleCode.md) |  | [optional] 
**IsReserved** | Pointer to **bool** | You can mark your number true and use it when main number can&#39;t send sms | [optional] 
**AlphaNumber** | Pointer to **string** | You can use short alias, text Alphanumeric Sender ID instead digital number | [optional] 
**Priority** | Pointer to **int32** | Priority number will try to send sms firstly | [optional] 
**Status** | Pointer to **int32** | The status of the Routing-rule. Every digit matters: 1 - active, 2 - inactive. | [optional] 

## Methods

### NewSmsRoutingRuleRequest

`func NewSmsRoutingRuleRequest(senderId string, ) *SmsRoutingRuleRequest`

NewSmsRoutingRuleRequest instantiates a new SmsRoutingRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsRoutingRuleRequestWithDefaults

`func NewSmsRoutingRuleRequestWithDefaults() *SmsRoutingRuleRequest`

NewSmsRoutingRuleRequestWithDefaults instantiates a new SmsRoutingRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenderId

`func (o *SmsRoutingRuleRequest) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *SmsRoutingRuleRequest) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *SmsRoutingRuleRequest) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.


### GetCountryCode

`func (o *SmsRoutingRuleRequest) GetCountryCode() LocaleCode`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SmsRoutingRuleRequest) GetCountryCodeOk() (*LocaleCode, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SmsRoutingRuleRequest) SetCountryCode(v LocaleCode)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SmsRoutingRuleRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetExcludedCountries

`func (o *SmsRoutingRuleRequest) GetExcludedCountries() []LocaleCode`

GetExcludedCountries returns the ExcludedCountries field if non-nil, zero value otherwise.

### GetExcludedCountriesOk

`func (o *SmsRoutingRuleRequest) GetExcludedCountriesOk() (*[]LocaleCode, bool)`

GetExcludedCountriesOk returns a tuple with the ExcludedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCountries

`func (o *SmsRoutingRuleRequest) SetExcludedCountries(v []LocaleCode)`

SetExcludedCountries sets ExcludedCountries field to given value.

### HasExcludedCountries

`func (o *SmsRoutingRuleRequest) HasExcludedCountries() bool`

HasExcludedCountries returns a boolean if a field has been set.

### GetIsReserved

`func (o *SmsRoutingRuleRequest) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *SmsRoutingRuleRequest) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *SmsRoutingRuleRequest) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.

### HasIsReserved

`func (o *SmsRoutingRuleRequest) HasIsReserved() bool`

HasIsReserved returns a boolean if a field has been set.

### GetAlphaNumber

`func (o *SmsRoutingRuleRequest) GetAlphaNumber() string`

GetAlphaNumber returns the AlphaNumber field if non-nil, zero value otherwise.

### GetAlphaNumberOk

`func (o *SmsRoutingRuleRequest) GetAlphaNumberOk() (*string, bool)`

GetAlphaNumberOk returns a tuple with the AlphaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaNumber

`func (o *SmsRoutingRuleRequest) SetAlphaNumber(v string)`

SetAlphaNumber sets AlphaNumber field to given value.

### HasAlphaNumber

`func (o *SmsRoutingRuleRequest) HasAlphaNumber() bool`

HasAlphaNumber returns a boolean if a field has been set.

### GetPriority

`func (o *SmsRoutingRuleRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SmsRoutingRuleRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SmsRoutingRuleRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SmsRoutingRuleRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *SmsRoutingRuleRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsRoutingRuleRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsRoutingRuleRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsRoutingRuleRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


