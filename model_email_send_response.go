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

// checks if the EmailSendResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSendResponse{}

// EmailSendResponse struct for EmailSendResponse
type EmailSendResponse struct {
	Data *EmailSendResponseData `json:"data,omitempty"`
}

// NewEmailSendResponse instantiates a new EmailSendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSendResponse() *EmailSendResponse {
	this := EmailSendResponse{}
	return &this
}

// NewEmailSendResponseWithDefaults instantiates a new EmailSendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSendResponseWithDefaults() *EmailSendResponse {
	this := EmailSendResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailSendResponse) GetData() EmailSendResponseData {
	if o == nil || IsNil(o.Data) {
		var ret EmailSendResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendResponse) GetDataOk() (*EmailSendResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailSendResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EmailSendResponseData and assigns it to the Data field.
func (o *EmailSendResponse) SetData(v EmailSendResponseData) {
	o.Data = &v
}

func (o EmailSendResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSendResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEmailSendResponse struct {
	value *EmailSendResponse
	isSet bool
}

func (v NullableEmailSendResponse) Get() *EmailSendResponse {
	return v.value
}

func (v *NullableEmailSendResponse) Set(val *EmailSendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSendResponse(val *EmailSendResponse) *NullableEmailSendResponse {
	return &NullableEmailSendResponse{value: val, isSet: true}
}

func (v NullableEmailSendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


