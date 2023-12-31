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

// checks if the EmailTemplateContentVariablesSettingsInnerRecipientProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateContentVariablesSettingsInnerRecipientProperty{}

// EmailTemplateContentVariablesSettingsInnerRecipientProperty struct for EmailTemplateContentVariablesSettingsInnerRecipientProperty
type EmailTemplateContentVariablesSettingsInnerRecipientProperty struct {
	// Mapping to the recipient property name.
	Name *string `json:"name,omitempty"`
	// Mapping to the recipient custom field key (required if name=custom_fields).
	Key NullableString `json:"key,omitempty"`
}

// NewEmailTemplateContentVariablesSettingsInnerRecipientProperty instantiates a new EmailTemplateContentVariablesSettingsInnerRecipientProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplateContentVariablesSettingsInnerRecipientProperty() *EmailTemplateContentVariablesSettingsInnerRecipientProperty {
	this := EmailTemplateContentVariablesSettingsInnerRecipientProperty{}
	return &this
}

// NewEmailTemplateContentVariablesSettingsInnerRecipientPropertyWithDefaults instantiates a new EmailTemplateContentVariablesSettingsInnerRecipientProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateContentVariablesSettingsInnerRecipientPropertyWithDefaults() *EmailTemplateContentVariablesSettingsInnerRecipientProperty {
	this := EmailTemplateContentVariablesSettingsInnerRecipientProperty{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *EmailTemplateContentVariablesSettingsInnerRecipientProperty) UnsetKey() {
	o.Key.Unset()
}

func (o EmailTemplateContentVariablesSettingsInnerRecipientProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateContentVariablesSettingsInnerRecipientProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	return toSerialize, nil
}

type NullableEmailTemplateContentVariablesSettingsInnerRecipientProperty struct {
	value *EmailTemplateContentVariablesSettingsInnerRecipientProperty
	isSet bool
}

func (v NullableEmailTemplateContentVariablesSettingsInnerRecipientProperty) Get() *EmailTemplateContentVariablesSettingsInnerRecipientProperty {
	return v.value
}

func (v *NullableEmailTemplateContentVariablesSettingsInnerRecipientProperty) Set(val *EmailTemplateContentVariablesSettingsInnerRecipientProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateContentVariablesSettingsInnerRecipientProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateContentVariablesSettingsInnerRecipientProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateContentVariablesSettingsInnerRecipientProperty(val *EmailTemplateContentVariablesSettingsInnerRecipientProperty) *NullableEmailTemplateContentVariablesSettingsInnerRecipientProperty {
	return &NullableEmailTemplateContentVariablesSettingsInnerRecipientProperty{value: val, isSet: true}
}

func (v NullableEmailTemplateContentVariablesSettingsInnerRecipientProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateContentVariablesSettingsInnerRecipientProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


