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

// checks if the PushTemplateCollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushTemplateCollectionResponse{}

// PushTemplateCollectionResponse struct for PushTemplateCollectionResponse
type PushTemplateCollectionResponse struct {
	Links *CollectionLinks `json:"links,omitempty"`
	Meta *CollectionMeta `json:"meta,omitempty"`
	Data []PushTemplate `json:"data,omitempty"`
}

// NewPushTemplateCollectionResponse instantiates a new PushTemplateCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushTemplateCollectionResponse() *PushTemplateCollectionResponse {
	this := PushTemplateCollectionResponse{}
	return &this
}

// NewPushTemplateCollectionResponseWithDefaults instantiates a new PushTemplateCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushTemplateCollectionResponseWithDefaults() *PushTemplateCollectionResponse {
	this := PushTemplateCollectionResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PushTemplateCollectionResponse) GetLinks() CollectionLinks {
	if o == nil || IsNil(o.Links) {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTemplateCollectionResponse) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PushTemplateCollectionResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *PushTemplateCollectionResponse) SetLinks(v CollectionLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PushTemplateCollectionResponse) GetMeta() CollectionMeta {
	if o == nil || IsNil(o.Meta) {
		var ret CollectionMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTemplateCollectionResponse) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PushTemplateCollectionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMeta and assigns it to the Meta field.
func (o *PushTemplateCollectionResponse) SetMeta(v CollectionMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PushTemplateCollectionResponse) GetData() []PushTemplate {
	if o == nil || IsNil(o.Data) {
		var ret []PushTemplate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTemplateCollectionResponse) GetDataOk() ([]PushTemplate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PushTemplateCollectionResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PushTemplate and assigns it to the Data field.
func (o *PushTemplateCollectionResponse) SetData(v []PushTemplate) {
	o.Data = v
}

func (o PushTemplateCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushTemplateCollectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePushTemplateCollectionResponse struct {
	value *PushTemplateCollectionResponse
	isSet bool
}

func (v NullablePushTemplateCollectionResponse) Get() *PushTemplateCollectionResponse {
	return v.value
}

func (v *NullablePushTemplateCollectionResponse) Set(val *PushTemplateCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePushTemplateCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePushTemplateCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushTemplateCollectionResponse(val *PushTemplateCollectionResponse) *NullablePushTemplateCollectionResponse {
	return &NullablePushTemplateCollectionResponse{value: val, isSet: true}
}

func (v NullablePushTemplateCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushTemplateCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


