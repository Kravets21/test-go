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

// checks if the SmsTransportUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsTransportUpdateRequest{}

// SmsTransportUpdateRequest struct for SmsTransportUpdateRequest
type SmsTransportUpdateRequest struct {
	// The name of the Transport.
	Name NullableString `json:"name,omitempty"`
	// The status of the Transport. Every digit matters: 1 - active, 2 - inactive,
	Status NullableInt32 `json:"status,omitempty"`
	// The credentials of the Transport for type.
	Credentials map[string]interface{} `json:"credentials,omitempty"`
}

// NewSmsTransportUpdateRequest instantiates a new SmsTransportUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsTransportUpdateRequest() *SmsTransportUpdateRequest {
	this := SmsTransportUpdateRequest{}
	return &this
}

// NewSmsTransportUpdateRequestWithDefaults instantiates a new SmsTransportUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsTransportUpdateRequestWithDefaults() *SmsTransportUpdateRequest {
	this := SmsTransportUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmsTransportUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmsTransportUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SmsTransportUpdateRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SmsTransportUpdateRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SmsTransportUpdateRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SmsTransportUpdateRequest) UnsetName() {
	o.Name.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmsTransportUpdateRequest) GetStatus() int32 {
	if o == nil || IsNil(o.Status.Get()) {
		var ret int32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmsTransportUpdateRequest) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *SmsTransportUpdateRequest) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableInt32 and assigns it to the Status field.
func (o *SmsTransportUpdateRequest) SetStatus(v int32) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *SmsTransportUpdateRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *SmsTransportUpdateRequest) UnsetStatus() {
	o.Status.Unset()
}

// GetCredentials returns the Credentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmsTransportUpdateRequest) GetCredentials() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmsTransportUpdateRequest) GetCredentialsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Credentials) {
		return map[string]interface{}{}, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *SmsTransportUpdateRequest) HasCredentials() bool {
	if o != nil && IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given map[string]interface{} and assigns it to the Credentials field.
func (o *SmsTransportUpdateRequest) SetCredentials(v map[string]interface{}) {
	o.Credentials = v
}

func (o SmsTransportUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsTransportUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	return toSerialize, nil
}

type NullableSmsTransportUpdateRequest struct {
	value *SmsTransportUpdateRequest
	isSet bool
}

func (v NullableSmsTransportUpdateRequest) Get() *SmsTransportUpdateRequest {
	return v.value
}

func (v *NullableSmsTransportUpdateRequest) Set(val *SmsTransportUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsTransportUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsTransportUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsTransportUpdateRequest(val *SmsTransportUpdateRequest) *NullableSmsTransportUpdateRequest {
	return &NullableSmsTransportUpdateRequest{value: val, isSet: true}
}

func (v NullableSmsTransportUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsTransportUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


