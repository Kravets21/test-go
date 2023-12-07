/*
airEml - API

airEml - API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AirEml

import (
	"encoding/json"
	"time"
)

// checks if the AudienceRecipientContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceRecipientContact{}

// AudienceRecipientContact struct for AudienceRecipientContact
type AudienceRecipientContact struct {
	// The type of the Contact.
	Type *string `json:"type,omitempty"`
	// The address of the Contact.
	Address NullableString `json:"address,omitempty"`
	// The datetime when the Contact was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewAudienceRecipientContact instantiates a new AudienceRecipientContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceRecipientContact() *AudienceRecipientContact {
	this := AudienceRecipientContact{}
	return &this
}

// NewAudienceRecipientContactWithDefaults instantiates a new AudienceRecipientContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceRecipientContactWithDefaults() *AudienceRecipientContact {
	this := AudienceRecipientContact{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AudienceRecipientContact) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceRecipientContact) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AudienceRecipientContact) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AudienceRecipientContact) SetType(v string) {
	o.Type = &v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AudienceRecipientContact) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AudienceRecipientContact) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *AudienceRecipientContact) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *AudienceRecipientContact) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *AudienceRecipientContact) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *AudienceRecipientContact) UnsetAddress() {
	o.Address.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AudienceRecipientContact) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceRecipientContact) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AudienceRecipientContact) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AudienceRecipientContact) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o AudienceRecipientContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceRecipientContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableAudienceRecipientContact struct {
	value *AudienceRecipientContact
	isSet bool
}

func (v NullableAudienceRecipientContact) Get() *AudienceRecipientContact {
	return v.value
}

func (v *NullableAudienceRecipientContact) Set(val *AudienceRecipientContact) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceRecipientContact) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceRecipientContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceRecipientContact(val *AudienceRecipientContact) *NullableAudienceRecipientContact {
	return &NullableAudienceRecipientContact{value: val, isSet: true}
}

func (v NullableAudienceRecipientContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceRecipientContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

