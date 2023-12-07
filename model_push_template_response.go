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

// checks if the PushTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushTemplateResponse{}

// PushTemplateResponse struct for PushTemplateResponse
type PushTemplateResponse struct {
	Data *PushTemplate `json:"data,omitempty"`
}

// NewPushTemplateResponse instantiates a new PushTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushTemplateResponse() *PushTemplateResponse {
	this := PushTemplateResponse{}
	return &this
}

// NewPushTemplateResponseWithDefaults instantiates a new PushTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushTemplateResponseWithDefaults() *PushTemplateResponse {
	this := PushTemplateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PushTemplateResponse) GetData() PushTemplate {
	if o == nil || IsNil(o.Data) {
		var ret PushTemplate
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTemplateResponse) GetDataOk() (*PushTemplate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PushTemplateResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PushTemplate and assigns it to the Data field.
func (o *PushTemplateResponse) SetData(v PushTemplate) {
	o.Data = &v
}

func (o PushTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePushTemplateResponse struct {
	value *PushTemplateResponse
	isSet bool
}

func (v NullablePushTemplateResponse) Get() *PushTemplateResponse {
	return v.value
}

func (v *NullablePushTemplateResponse) Set(val *PushTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePushTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePushTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushTemplateResponse(val *PushTemplateResponse) *NullablePushTemplateResponse {
	return &NullablePushTemplateResponse{value: val, isSet: true}
}

func (v NullablePushTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


