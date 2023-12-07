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

// checks if the EmailChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailChannel{}

// EmailChannel struct for EmailChannel
type EmailChannel struct {
	// The unique identifier of the Channel.
	Id *string `json:"id,omitempty"`
	// The name of the Channel.
	Name *string `json:"name,omitempty"`
	// The datetime when the Sender was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewEmailChannel instantiates a new EmailChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailChannel() *EmailChannel {
	this := EmailChannel{}
	return &this
}

// NewEmailChannelWithDefaults instantiates a new EmailChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailChannelWithDefaults() *EmailChannel {
	this := EmailChannel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailChannel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChannel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailChannel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailChannel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailChannel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChannel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailChannel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmailChannel) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EmailChannel) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChannel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EmailChannel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EmailChannel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o EmailChannel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableEmailChannel struct {
	value *EmailChannel
	isSet bool
}

func (v NullableEmailChannel) Get() *EmailChannel {
	return v.value
}

func (v *NullableEmailChannel) Set(val *EmailChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailChannel(val *EmailChannel) *NullableEmailChannel {
	return &NullableEmailChannel{value: val, isSet: true}
}

func (v NullableEmailChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


