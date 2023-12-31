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

// checks if the SmsTransportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsTransportResponse{}

// SmsTransportResponse struct for SmsTransportResponse
type SmsTransportResponse struct {
	Data *SmsTransport `json:"data,omitempty"`
}

// NewSmsTransportResponse instantiates a new SmsTransportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsTransportResponse() *SmsTransportResponse {
	this := SmsTransportResponse{}
	return &this
}

// NewSmsTransportResponseWithDefaults instantiates a new SmsTransportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsTransportResponseWithDefaults() *SmsTransportResponse {
	this := SmsTransportResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SmsTransportResponse) GetData() SmsTransport {
	if o == nil || IsNil(o.Data) {
		var ret SmsTransport
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTransportResponse) GetDataOk() (*SmsTransport, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SmsTransportResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SmsTransport and assigns it to the Data field.
func (o *SmsTransportResponse) SetData(v SmsTransport) {
	o.Data = &v
}

func (o SmsTransportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsTransportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSmsTransportResponse struct {
	value *SmsTransportResponse
	isSet bool
}

func (v NullableSmsTransportResponse) Get() *SmsTransportResponse {
	return v.value
}

func (v *NullableSmsTransportResponse) Set(val *SmsTransportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsTransportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsTransportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsTransportResponse(val *SmsTransportResponse) *NullableSmsTransportResponse {
	return &NullableSmsTransportResponse{value: val, isSet: true}
}

func (v NullableSmsTransportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsTransportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


