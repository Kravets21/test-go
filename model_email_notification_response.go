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

// checks if the EmailNotificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailNotificationResponse{}

// EmailNotificationResponse struct for EmailNotificationResponse
type EmailNotificationResponse struct {
	Data *EmailNotificationDetails `json:"data,omitempty"`
}

// NewEmailNotificationResponse instantiates a new EmailNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationResponse() *EmailNotificationResponse {
	this := EmailNotificationResponse{}
	return &this
}

// NewEmailNotificationResponseWithDefaults instantiates a new EmailNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationResponseWithDefaults() *EmailNotificationResponse {
	this := EmailNotificationResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailNotificationResponse) GetData() EmailNotificationDetails {
	if o == nil || IsNil(o.Data) {
		var ret EmailNotificationDetails
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationResponse) GetDataOk() (*EmailNotificationDetails, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EmailNotificationDetails and assigns it to the Data field.
func (o *EmailNotificationResponse) SetData(v EmailNotificationDetails) {
	o.Data = &v
}

func (o EmailNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEmailNotificationResponse struct {
	value *EmailNotificationResponse
	isSet bool
}

func (v NullableEmailNotificationResponse) Get() *EmailNotificationResponse {
	return v.value
}

func (v *NullableEmailNotificationResponse) Set(val *EmailNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationResponse(val *EmailNotificationResponse) *NullableEmailNotificationResponse {
	return &NullableEmailNotificationResponse{value: val, isSet: true}
}

func (v NullableEmailNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


