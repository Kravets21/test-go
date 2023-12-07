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

// checks if the SubscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionResponse{}

// SubscriptionResponse struct for SubscriptionResponse
type SubscriptionResponse struct {
	Data *Subscription `json:"data,omitempty"`
}

// NewSubscriptionResponse instantiates a new SubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionResponse() *SubscriptionResponse {
	this := SubscriptionResponse{}
	return &this
}

// NewSubscriptionResponseWithDefaults instantiates a new SubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionResponseWithDefaults() *SubscriptionResponse {
	this := SubscriptionResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetData() Subscription {
	if o == nil || IsNil(o.Data) {
		var ret Subscription
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetDataOk() (*Subscription, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Subscription and assigns it to the Data field.
func (o *SubscriptionResponse) SetData(v Subscription) {
	o.Data = &v
}

func (o SubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSubscriptionResponse struct {
	value *SubscriptionResponse
	isSet bool
}

func (v NullableSubscriptionResponse) Get() *SubscriptionResponse {
	return v.value
}

func (v *NullableSubscriptionResponse) Set(val *SubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionResponse(val *SubscriptionResponse) *NullableSubscriptionResponse {
	return &NullableSubscriptionResponse{value: val, isSet: true}
}

func (v NullableSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


