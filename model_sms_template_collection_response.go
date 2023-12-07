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

// checks if the SmsTemplateCollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsTemplateCollectionResponse{}

// SmsTemplateCollectionResponse struct for SmsTemplateCollectionResponse
type SmsTemplateCollectionResponse struct {
	Links *CollectionLinks `json:"links,omitempty"`
	Meta *CollectionMeta `json:"meta,omitempty"`
	Data []SmsTemplate `json:"data,omitempty"`
}

// NewSmsTemplateCollectionResponse instantiates a new SmsTemplateCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsTemplateCollectionResponse() *SmsTemplateCollectionResponse {
	this := SmsTemplateCollectionResponse{}
	return &this
}

// NewSmsTemplateCollectionResponseWithDefaults instantiates a new SmsTemplateCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsTemplateCollectionResponseWithDefaults() *SmsTemplateCollectionResponse {
	this := SmsTemplateCollectionResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SmsTemplateCollectionResponse) GetLinks() CollectionLinks {
	if o == nil || IsNil(o.Links) {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplateCollectionResponse) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SmsTemplateCollectionResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *SmsTemplateCollectionResponse) SetLinks(v CollectionLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SmsTemplateCollectionResponse) GetMeta() CollectionMeta {
	if o == nil || IsNil(o.Meta) {
		var ret CollectionMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplateCollectionResponse) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SmsTemplateCollectionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMeta and assigns it to the Meta field.
func (o *SmsTemplateCollectionResponse) SetMeta(v CollectionMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SmsTemplateCollectionResponse) GetData() []SmsTemplate {
	if o == nil || IsNil(o.Data) {
		var ret []SmsTemplate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplateCollectionResponse) GetDataOk() ([]SmsTemplate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SmsTemplateCollectionResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SmsTemplate and assigns it to the Data field.
func (o *SmsTemplateCollectionResponse) SetData(v []SmsTemplate) {
	o.Data = v
}

func (o SmsTemplateCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsTemplateCollectionResponse) ToMap() (map[string]interface{}, error) {
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

type NullableSmsTemplateCollectionResponse struct {
	value *SmsTemplateCollectionResponse
	isSet bool
}

func (v NullableSmsTemplateCollectionResponse) Get() *SmsTemplateCollectionResponse {
	return v.value
}

func (v *NullableSmsTemplateCollectionResponse) Set(val *SmsTemplateCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsTemplateCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsTemplateCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsTemplateCollectionResponse(val *SmsTemplateCollectionResponse) *NullableSmsTemplateCollectionResponse {
	return &NullableSmsTemplateCollectionResponse{value: val, isSet: true}
}

func (v NullableSmsTemplateCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsTemplateCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

