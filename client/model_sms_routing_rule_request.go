/*
airEml - API

airEml - API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AirEml

import (
	"encoding/json"
	"fmt"
)

// checks if the SmsRoutingRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsRoutingRuleRequest{}

// SmsRoutingRuleRequest struct for SmsRoutingRuleRequest
type SmsRoutingRuleRequest struct {
	// The unique identifier of the Sender.
	SenderId string `json:"sender_id"`
	CountryCode *LocaleCode `json:"country_code,omitempty"`
	ExcludedCountries []LocaleCode `json:"excluded_countries,omitempty"`
	// You can mark your number true and use it when main number can't send sms
	IsReserved *bool `json:"is_reserved,omitempty"`
	// You can use short alias, text Alphanumeric Sender ID instead digital number
	AlphaNumber *string `json:"alpha_number,omitempty"`
	// Priority number will try to send sms firstly
	Priority *int32 `json:"priority,omitempty"`
	// The status of the Routing-rule. Every digit matters: 1 - active, 2 - inactive.
	Status *int32 `json:"status,omitempty"`
}

type _SmsRoutingRuleRequest SmsRoutingRuleRequest

// NewSmsRoutingRuleRequest instantiates a new SmsRoutingRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsRoutingRuleRequest(senderId string) *SmsRoutingRuleRequest {
	this := SmsRoutingRuleRequest{}
	this.SenderId = senderId
	return &this
}

// NewSmsRoutingRuleRequestWithDefaults instantiates a new SmsRoutingRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsRoutingRuleRequestWithDefaults() *SmsRoutingRuleRequest {
	this := SmsRoutingRuleRequest{}
	return &this
}

// GetSenderId returns the SenderId field value
func (o *SmsRoutingRuleRequest) GetSenderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value
// and a boolean to check if the value has been set.
func (o *SmsRoutingRuleRequest) GetSenderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderId, true
}

// SetSenderId sets field value
func (o *SmsRoutingRuleRequest) SetSenderId(v string) {
	o.SenderId = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *SmsRoutingRuleRequest) GetCountryCode() LocaleCode {
	if o == nil || IsNil(o.CountryCode) {
		var ret LocaleCode
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRoutingRuleRequest) GetCountryCodeOk() (*LocaleCode, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *SmsRoutingRuleRequest) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given LocaleCode and assigns it to the CountryCode field.
func (o *SmsRoutingRuleRequest) SetCountryCode(v LocaleCode) {
	o.CountryCode = &v
}

// GetExcludedCountries returns the ExcludedCountries field value if set, zero value otherwise.
func (o *SmsRoutingRuleRequest) GetExcludedCountries() []LocaleCode {
	if o == nil || IsNil(o.ExcludedCountries) {
		var ret []LocaleCode
		return ret
	}
	return o.ExcludedCountries
}

// GetExcludedCountriesOk returns a tuple with the ExcludedCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRoutingRuleRequest) GetExcludedCountriesOk() ([]LocaleCode, bool) {
	if o == nil || IsNil(o.ExcludedCountries) {
		return nil, false
	}
	return o.ExcludedCountries, true
}

// HasExcludedCountries returns a boolean if a field has been set.
func (o *SmsRoutingRuleRequest) HasExcludedCountries() bool {
	if o != nil && !IsNil(o.ExcludedCountries) {
		return true
	}

	return false
}

// SetExcludedCountries gets a reference to the given []LocaleCode and assigns it to the ExcludedCountries field.
func (o *SmsRoutingRuleRequest) SetExcludedCountries(v []LocaleCode) {
	o.ExcludedCountries = v
}

// GetIsReserved returns the IsReserved field value if set, zero value otherwise.
func (o *SmsRoutingRuleRequest) GetIsReserved() bool {
	if o == nil || IsNil(o.IsReserved) {
		var ret bool
		return ret
	}
	return *o.IsReserved
}

// GetIsReservedOk returns a tuple with the IsReserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRoutingRuleRequest) GetIsReservedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReserved) {
		return nil, false
	}
	return o.IsReserved, true
}

// HasIsReserved returns a boolean if a field has been set.
func (o *SmsRoutingRuleRequest) HasIsReserved() bool {
	if o != nil && !IsNil(o.IsReserved) {
		return true
	}

	return false
}

// SetIsReserved gets a reference to the given bool and assigns it to the IsReserved field.
func (o *SmsRoutingRuleRequest) SetIsReserved(v bool) {
	o.IsReserved = &v
}

// GetAlphaNumber returns the AlphaNumber field value if set, zero value otherwise.
func (o *SmsRoutingRuleRequest) GetAlphaNumber() string {
	if o == nil || IsNil(o.AlphaNumber) {
		var ret string
		return ret
	}
	return *o.AlphaNumber
}

// GetAlphaNumberOk returns a tuple with the AlphaNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRoutingRuleRequest) GetAlphaNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AlphaNumber) {
		return nil, false
	}
	return o.AlphaNumber, true
}

// HasAlphaNumber returns a boolean if a field has been set.
func (o *SmsRoutingRuleRequest) HasAlphaNumber() bool {
	if o != nil && !IsNil(o.AlphaNumber) {
		return true
	}

	return false
}

// SetAlphaNumber gets a reference to the given string and assigns it to the AlphaNumber field.
func (o *SmsRoutingRuleRequest) SetAlphaNumber(v string) {
	o.AlphaNumber = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SmsRoutingRuleRequest) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRoutingRuleRequest) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SmsRoutingRuleRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *SmsRoutingRuleRequest) SetPriority(v int32) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SmsRoutingRuleRequest) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRoutingRuleRequest) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SmsRoutingRuleRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *SmsRoutingRuleRequest) SetStatus(v int32) {
	o.Status = &v
}

func (o SmsRoutingRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsRoutingRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sender_id"] = o.SenderId
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.ExcludedCountries) {
		toSerialize["excluded_countries"] = o.ExcludedCountries
	}
	if !IsNil(o.IsReserved) {
		toSerialize["is_reserved"] = o.IsReserved
	}
	if !IsNil(o.AlphaNumber) {
		toSerialize["alpha_number"] = o.AlphaNumber
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *SmsRoutingRuleRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sender_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSmsRoutingRuleRequest := _SmsRoutingRuleRequest{}

	err = json.Unmarshal(bytes, &varSmsRoutingRuleRequest)

	if err != nil {
		return err
	}

	*o = SmsRoutingRuleRequest(varSmsRoutingRuleRequest)

	return err
}

type NullableSmsRoutingRuleRequest struct {
	value *SmsRoutingRuleRequest
	isSet bool
}

func (v NullableSmsRoutingRuleRequest) Get() *SmsRoutingRuleRequest {
	return v.value
}

func (v *NullableSmsRoutingRuleRequest) Set(val *SmsRoutingRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsRoutingRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsRoutingRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsRoutingRuleRequest(val *SmsRoutingRuleRequest) *NullableSmsRoutingRuleRequest {
	return &NullableSmsRoutingRuleRequest{value: val, isSet: true}
}

func (v NullableSmsRoutingRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsRoutingRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


