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

// checks if the AudienceSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceSegment{}

// AudienceSegment struct for AudienceSegment
type AudienceSegment struct {
	// The unique identifier of the Segment.
	Id *string `json:"id,omitempty"`
	// The name of the Segment.
	Name *string `json:"name,omitempty"`
	// The status of the Segment.
	Status *int32 `json:"status,omitempty"`
	// The datetime when the Platform was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewAudienceSegment instantiates a new AudienceSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceSegment() *AudienceSegment {
	this := AudienceSegment{}
	return &this
}

// NewAudienceSegmentWithDefaults instantiates a new AudienceSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceSegmentWithDefaults() *AudienceSegment {
	this := AudienceSegment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AudienceSegment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSegment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AudienceSegment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AudienceSegment) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AudienceSegment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSegment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AudienceSegment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AudienceSegment) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AudienceSegment) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSegment) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AudienceSegment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *AudienceSegment) SetStatus(v int32) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AudienceSegment) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSegment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AudienceSegment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AudienceSegment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o AudienceSegment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableAudienceSegment struct {
	value *AudienceSegment
	isSet bool
}

func (v NullableAudienceSegment) Get() *AudienceSegment {
	return v.value
}

func (v *NullableAudienceSegment) Set(val *AudienceSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceSegment(val *AudienceSegment) *NullableAudienceSegment {
	return &NullableAudienceSegment{value: val, isSet: true}
}

func (v NullableAudienceSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


