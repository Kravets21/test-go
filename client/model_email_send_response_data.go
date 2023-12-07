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

// checks if the EmailSendResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSendResponseData{}

// EmailSendResponseData struct for EmailSendResponseData
type EmailSendResponseData struct {
	// The unique identifier of the Notification.
	Id *string `json:"id,omitempty"`
}

// NewEmailSendResponseData instantiates a new EmailSendResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSendResponseData() *EmailSendResponseData {
	this := EmailSendResponseData{}
	return &this
}

// NewEmailSendResponseDataWithDefaults instantiates a new EmailSendResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSendResponseDataWithDefaults() *EmailSendResponseData {
	this := EmailSendResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailSendResponseData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendResponseData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailSendResponseData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailSendResponseData) SetId(v string) {
	o.Id = &v
}

func (o EmailSendResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSendResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableEmailSendResponseData struct {
	value *EmailSendResponseData
	isSet bool
}

func (v NullableEmailSendResponseData) Get() *EmailSendResponseData {
	return v.value
}

func (v *NullableEmailSendResponseData) Set(val *EmailSendResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSendResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSendResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSendResponseData(val *EmailSendResponseData) *NullableEmailSendResponseData {
	return &NullableEmailSendResponseData{value: val, isSet: true}
}

func (v NullableEmailSendResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSendResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


