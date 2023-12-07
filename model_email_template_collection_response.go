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

// checks if the EmailTemplateCollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateCollectionResponse{}

// EmailTemplateCollectionResponse struct for EmailTemplateCollectionResponse
type EmailTemplateCollectionResponse struct {
	Links *CollectionLinks `json:"links,omitempty"`
	Meta *CollectionMeta `json:"meta,omitempty"`
	Data []EmailTemplate `json:"data,omitempty"`
}

// NewEmailTemplateCollectionResponse instantiates a new EmailTemplateCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplateCollectionResponse() *EmailTemplateCollectionResponse {
	this := EmailTemplateCollectionResponse{}
	return &this
}

// NewEmailTemplateCollectionResponseWithDefaults instantiates a new EmailTemplateCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateCollectionResponseWithDefaults() *EmailTemplateCollectionResponse {
	this := EmailTemplateCollectionResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailTemplateCollectionResponse) GetLinks() CollectionLinks {
	if o == nil || IsNil(o.Links) {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateCollectionResponse) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailTemplateCollectionResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *EmailTemplateCollectionResponse) SetLinks(v CollectionLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *EmailTemplateCollectionResponse) GetMeta() CollectionMeta {
	if o == nil || IsNil(o.Meta) {
		var ret CollectionMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateCollectionResponse) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *EmailTemplateCollectionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMeta and assigns it to the Meta field.
func (o *EmailTemplateCollectionResponse) SetMeta(v CollectionMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailTemplateCollectionResponse) GetData() []EmailTemplate {
	if o == nil || IsNil(o.Data) {
		var ret []EmailTemplate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateCollectionResponse) GetDataOk() ([]EmailTemplate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailTemplateCollectionResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EmailTemplate and assigns it to the Data field.
func (o *EmailTemplateCollectionResponse) SetData(v []EmailTemplate) {
	o.Data = v
}

func (o EmailTemplateCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateCollectionResponse) ToMap() (map[string]interface{}, error) {
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

type NullableEmailTemplateCollectionResponse struct {
	value *EmailTemplateCollectionResponse
	isSet bool
}

func (v NullableEmailTemplateCollectionResponse) Get() *EmailTemplateCollectionResponse {
	return v.value
}

func (v *NullableEmailTemplateCollectionResponse) Set(val *EmailTemplateCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateCollectionResponse(val *EmailTemplateCollectionResponse) *NullableEmailTemplateCollectionResponse {
	return &NullableEmailTemplateCollectionResponse{value: val, isSet: true}
}

func (v NullableEmailTemplateCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

