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

// checks if the AuthTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenResponse{}

// AuthTokenResponse struct for AuthTokenResponse
type AuthTokenResponse struct {
	Data *AuthToken `json:"data,omitempty"`
}

// NewAuthTokenResponse instantiates a new AuthTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenResponse() *AuthTokenResponse {
	this := AuthTokenResponse{}
	return &this
}

// NewAuthTokenResponseWithDefaults instantiates a new AuthTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenResponseWithDefaults() *AuthTokenResponse {
	this := AuthTokenResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AuthTokenResponse) GetData() AuthToken {
	if o == nil || IsNil(o.Data) {
		var ret AuthToken
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenResponse) GetDataOk() (*AuthToken, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AuthTokenResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AuthToken and assigns it to the Data field.
func (o *AuthTokenResponse) SetData(v AuthToken) {
	o.Data = &v
}

func (o AuthTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAuthTokenResponse struct {
	value *AuthTokenResponse
	isSet bool
}

func (v NullableAuthTokenResponse) Get() *AuthTokenResponse {
	return v.value
}

func (v *NullableAuthTokenResponse) Set(val *AuthTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenResponse(val *AuthTokenResponse) *NullableAuthTokenResponse {
	return &NullableAuthTokenResponse{value: val, isSet: true}
}

func (v NullableAuthTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

