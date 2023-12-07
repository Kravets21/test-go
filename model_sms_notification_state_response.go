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

// checks if the SmsNotificationStateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsNotificationStateResponse{}

// SmsNotificationStateResponse struct for SmsNotificationStateResponse
type SmsNotificationStateResponse struct {
	// Notification States
	Data []EmailNotificationStateInner `json:"data,omitempty"`
}

// NewSmsNotificationStateResponse instantiates a new SmsNotificationStateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsNotificationStateResponse() *SmsNotificationStateResponse {
	this := SmsNotificationStateResponse{}
	return &this
}

// NewSmsNotificationStateResponseWithDefaults instantiates a new SmsNotificationStateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsNotificationStateResponseWithDefaults() *SmsNotificationStateResponse {
	this := SmsNotificationStateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SmsNotificationStateResponse) GetData() []EmailNotificationStateInner {
	if o == nil || IsNil(o.Data) {
		var ret []EmailNotificationStateInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsNotificationStateResponse) GetDataOk() ([]EmailNotificationStateInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SmsNotificationStateResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EmailNotificationStateInner and assigns it to the Data field.
func (o *SmsNotificationStateResponse) SetData(v []EmailNotificationStateInner) {
	o.Data = v
}

func (o SmsNotificationStateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsNotificationStateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSmsNotificationStateResponse struct {
	value *SmsNotificationStateResponse
	isSet bool
}

func (v NullableSmsNotificationStateResponse) Get() *SmsNotificationStateResponse {
	return v.value
}

func (v *NullableSmsNotificationStateResponse) Set(val *SmsNotificationStateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsNotificationStateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsNotificationStateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsNotificationStateResponse(val *SmsNotificationStateResponse) *NullableSmsNotificationStateResponse {
	return &NullableSmsNotificationStateResponse{value: val, isSet: true}
}

func (v NullableSmsNotificationStateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsNotificationStateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

