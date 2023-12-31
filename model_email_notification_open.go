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

// checks if the EmailNotificationOpen type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailNotificationOpen{}

// EmailNotificationOpen struct for EmailNotificationOpen
type EmailNotificationOpen struct {
	// The unique identifier of the Open.
	Id *string `json:"id,omitempty"`
	// The IP address from which the Open was made.
	Ip *string `json:"ip,omitempty"`
	// The user agent from which the Open was made.
	UserAgent *string `json:"user_agent,omitempty"`
	// The datetime when the Open was made.
	Time *time.Time `json:"time,omitempty"`
}

// NewEmailNotificationOpen instantiates a new EmailNotificationOpen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationOpen() *EmailNotificationOpen {
	this := EmailNotificationOpen{}
	return &this
}

// NewEmailNotificationOpenWithDefaults instantiates a new EmailNotificationOpen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationOpenWithDefaults() *EmailNotificationOpen {
	this := EmailNotificationOpen{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailNotificationOpen) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationOpen) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailNotificationOpen) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailNotificationOpen) SetId(v string) {
	o.Id = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *EmailNotificationOpen) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationOpen) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *EmailNotificationOpen) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *EmailNotificationOpen) SetIp(v string) {
	o.Ip = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *EmailNotificationOpen) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationOpen) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *EmailNotificationOpen) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *EmailNotificationOpen) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *EmailNotificationOpen) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationOpen) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *EmailNotificationOpen) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *EmailNotificationOpen) SetTime(v time.Time) {
	o.Time = &v
}

func (o EmailNotificationOpen) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailNotificationOpen) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.UserAgent) {
		toSerialize["user_agent"] = o.UserAgent
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableEmailNotificationOpen struct {
	value *EmailNotificationOpen
	isSet bool
}

func (v NullableEmailNotificationOpen) Get() *EmailNotificationOpen {
	return v.value
}

func (v *NullableEmailNotificationOpen) Set(val *EmailNotificationOpen) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationOpen) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationOpen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationOpen(val *EmailNotificationOpen) *NullableEmailNotificationOpen {
	return &NullableEmailNotificationOpen{value: val, isSet: true}
}

func (v NullableEmailNotificationOpen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationOpen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


