/*
airEml - API

airEml - API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AirEml

import (
	"encoding/json"
)

// checks if the SmsTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsTemplateResponse{}

// SmsTemplateResponse struct for SmsTemplateResponse
type SmsTemplateResponse struct {
	Data *SmsTemplate `json:"data,omitempty"`
}

// NewSmsTemplateResponse instantiates a new SmsTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsTemplateResponse() *SmsTemplateResponse {
	this := SmsTemplateResponse{}
	return &this
}

// NewSmsTemplateResponseWithDefaults instantiates a new SmsTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsTemplateResponseWithDefaults() *SmsTemplateResponse {
	this := SmsTemplateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SmsTemplateResponse) GetData() SmsTemplate {
	if o == nil || IsNil(o.Data) {
		var ret SmsTemplate
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplateResponse) GetDataOk() (*SmsTemplate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SmsTemplateResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SmsTemplate and assigns it to the Data field.
func (o *SmsTemplateResponse) SetData(v SmsTemplate) {
	o.Data = &v
}

func (o SmsTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSmsTemplateResponse struct {
	value *SmsTemplateResponse
	isSet bool
}

func (v NullableSmsTemplateResponse) Get() *SmsTemplateResponse {
	return v.value
}

func (v *NullableSmsTemplateResponse) Set(val *SmsTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsTemplateResponse(val *SmsTemplateResponse) *NullableSmsTemplateResponse {
	return &NullableSmsTemplateResponse{value: val, isSet: true}
}

func (v NullableSmsTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


