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

// checks if the EmailNotificationStateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailNotificationStateResponse{}

// EmailNotificationStateResponse struct for EmailNotificationStateResponse
type EmailNotificationStateResponse struct {
	// Notification States
	Data []EmailNotificationStateInner `json:"data,omitempty"`
}

// NewEmailNotificationStateResponse instantiates a new EmailNotificationStateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationStateResponse() *EmailNotificationStateResponse {
	this := EmailNotificationStateResponse{}
	return &this
}

// NewEmailNotificationStateResponseWithDefaults instantiates a new EmailNotificationStateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationStateResponseWithDefaults() *EmailNotificationStateResponse {
	this := EmailNotificationStateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailNotificationStateResponse) GetData() []EmailNotificationStateInner {
	if o == nil || IsNil(o.Data) {
		var ret []EmailNotificationStateInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationStateResponse) GetDataOk() ([]EmailNotificationStateInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailNotificationStateResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EmailNotificationStateInner and assigns it to the Data field.
func (o *EmailNotificationStateResponse) SetData(v []EmailNotificationStateInner) {
	o.Data = v
}

func (o EmailNotificationStateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailNotificationStateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEmailNotificationStateResponse struct {
	value *EmailNotificationStateResponse
	isSet bool
}

func (v NullableEmailNotificationStateResponse) Get() *EmailNotificationStateResponse {
	return v.value
}

func (v *NullableEmailNotificationStateResponse) Set(val *EmailNotificationStateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationStateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationStateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationStateResponse(val *EmailNotificationStateResponse) *NullableEmailNotificationStateResponse {
	return &NullableEmailNotificationStateResponse{value: val, isSet: true}
}

func (v NullableEmailNotificationStateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationStateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


