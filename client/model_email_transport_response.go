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

// checks if the EmailTransportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTransportResponse{}

// EmailTransportResponse struct for EmailTransportResponse
type EmailTransportResponse struct {
	Data *EmailTransport `json:"data,omitempty"`
}

// NewEmailTransportResponse instantiates a new EmailTransportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTransportResponse() *EmailTransportResponse {
	this := EmailTransportResponse{}
	return &this
}

// NewEmailTransportResponseWithDefaults instantiates a new EmailTransportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTransportResponseWithDefaults() *EmailTransportResponse {
	this := EmailTransportResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailTransportResponse) GetData() EmailTransport {
	if o == nil || IsNil(o.Data) {
		var ret EmailTransport
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTransportResponse) GetDataOk() (*EmailTransport, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailTransportResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EmailTransport and assigns it to the Data field.
func (o *EmailTransportResponse) SetData(v EmailTransport) {
	o.Data = &v
}

func (o EmailTransportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTransportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEmailTransportResponse struct {
	value *EmailTransportResponse
	isSet bool
}

func (v NullableEmailTransportResponse) Get() *EmailTransportResponse {
	return v.value
}

func (v *NullableEmailTransportResponse) Set(val *EmailTransportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTransportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTransportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTransportResponse(val *EmailTransportResponse) *NullableEmailTransportResponse {
	return &NullableEmailTransportResponse{value: val, isSet: true}
}

func (v NullableEmailTransportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTransportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

