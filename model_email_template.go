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

// checks if the EmailTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplate{}

// EmailTemplate struct for EmailTemplate
type EmailTemplate struct {
	// The unique identifier of the Template.
	Id *string `json:"id,omitempty"`
	// The unique identifier of the Channel.
	ChannelId *string `json:"channel_id,omitempty"`
	// The name of the Template.
	Name *string `json:"name,omitempty"`
	// The description of the Template.
	Description *string `json:"description,omitempty"`
	// List of locales Contents available for this Template.
	Locales []string `json:"locales,omitempty"`
	// The status of the Template. Every digit matters: 1 - active, 2 - inactive.
	Status *int32 `json:"status,omitempty"`
	// The datetime when the Template was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewEmailTemplate instantiates a new EmailTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplate() *EmailTemplate {
	this := EmailTemplate{}
	return &this
}

// NewEmailTemplateWithDefaults instantiates a new EmailTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateWithDefaults() *EmailTemplate {
	this := EmailTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailTemplate) SetId(v string) {
	o.Id = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *EmailTemplate) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplate) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *EmailTemplate) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *EmailTemplate) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmailTemplate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EmailTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EmailTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EmailTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetLocales returns the Locales field value if set, zero value otherwise.
func (o *EmailTemplate) GetLocales() []string {
	if o == nil || IsNil(o.Locales) {
		var ret []string
		return ret
	}
	return o.Locales
}

// GetLocalesOk returns a tuple with the Locales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplate) GetLocalesOk() ([]string, bool) {
	if o == nil || IsNil(o.Locales) {
		return nil, false
	}
	return o.Locales, true
}

// HasLocales returns a boolean if a field has been set.
func (o *EmailTemplate) HasLocales() bool {
	if o != nil && !IsNil(o.Locales) {
		return true
	}

	return false
}

// SetLocales gets a reference to the given []string and assigns it to the Locales field.
func (o *EmailTemplate) SetLocales(v []string) {
	o.Locales = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmailTemplate) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplate) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmailTemplate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *EmailTemplate) SetStatus(v int32) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EmailTemplate) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EmailTemplate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EmailTemplate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o EmailTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Locales) {
		toSerialize["locales"] = o.Locales
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableEmailTemplate struct {
	value *EmailTemplate
	isSet bool
}

func (v NullableEmailTemplate) Get() *EmailTemplate {
	return v.value
}

func (v *NullableEmailTemplate) Set(val *EmailTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplate(val *EmailTemplate) *NullableEmailTemplate {
	return &NullableEmailTemplate{value: val, isSet: true}
}

func (v NullableEmailTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


