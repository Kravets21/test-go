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

// checks if the PushNotificationStateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushNotificationStateResponse{}

// PushNotificationStateResponse struct for PushNotificationStateResponse
type PushNotificationStateResponse struct {
	// Notification States
	Data []EmailNotificationStateInner `json:"data,omitempty"`
}

// NewPushNotificationStateResponse instantiates a new PushNotificationStateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushNotificationStateResponse() *PushNotificationStateResponse {
	this := PushNotificationStateResponse{}
	return &this
}

// NewPushNotificationStateResponseWithDefaults instantiates a new PushNotificationStateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushNotificationStateResponseWithDefaults() *PushNotificationStateResponse {
	this := PushNotificationStateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PushNotificationStateResponse) GetData() []EmailNotificationStateInner {
	if o == nil || IsNil(o.Data) {
		var ret []EmailNotificationStateInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushNotificationStateResponse) GetDataOk() ([]EmailNotificationStateInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PushNotificationStateResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EmailNotificationStateInner and assigns it to the Data field.
func (o *PushNotificationStateResponse) SetData(v []EmailNotificationStateInner) {
	o.Data = v
}

func (o PushNotificationStateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushNotificationStateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePushNotificationStateResponse struct {
	value *PushNotificationStateResponse
	isSet bool
}

func (v NullablePushNotificationStateResponse) Get() *PushNotificationStateResponse {
	return v.value
}

func (v *NullablePushNotificationStateResponse) Set(val *PushNotificationStateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePushNotificationStateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePushNotificationStateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushNotificationStateResponse(val *PushNotificationStateResponse) *NullablePushNotificationStateResponse {
	return &NullablePushNotificationStateResponse{value: val, isSet: true}
}

func (v NullablePushNotificationStateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushNotificationStateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


