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

// checks if the EmailChannelRoutingRulesCollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailChannelRoutingRulesCollectionResponse{}

// EmailChannelRoutingRulesCollectionResponse struct for EmailChannelRoutingRulesCollectionResponse
type EmailChannelRoutingRulesCollectionResponse struct {
	Links *CollectionLinks `json:"links,omitempty"`
	Meta *CollectionMeta `json:"meta,omitempty"`
	Data []EmailChannelRoutingRules `json:"data,omitempty"`
}

// NewEmailChannelRoutingRulesCollectionResponse instantiates a new EmailChannelRoutingRulesCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailChannelRoutingRulesCollectionResponse() *EmailChannelRoutingRulesCollectionResponse {
	this := EmailChannelRoutingRulesCollectionResponse{}
	return &this
}

// NewEmailChannelRoutingRulesCollectionResponseWithDefaults instantiates a new EmailChannelRoutingRulesCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailChannelRoutingRulesCollectionResponseWithDefaults() *EmailChannelRoutingRulesCollectionResponse {
	this := EmailChannelRoutingRulesCollectionResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailChannelRoutingRulesCollectionResponse) GetLinks() CollectionLinks {
	if o == nil || IsNil(o.Links) {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChannelRoutingRulesCollectionResponse) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailChannelRoutingRulesCollectionResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *EmailChannelRoutingRulesCollectionResponse) SetLinks(v CollectionLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *EmailChannelRoutingRulesCollectionResponse) GetMeta() CollectionMeta {
	if o == nil || IsNil(o.Meta) {
		var ret CollectionMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChannelRoutingRulesCollectionResponse) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *EmailChannelRoutingRulesCollectionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMeta and assigns it to the Meta field.
func (o *EmailChannelRoutingRulesCollectionResponse) SetMeta(v CollectionMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailChannelRoutingRulesCollectionResponse) GetData() []EmailChannelRoutingRules {
	if o == nil || IsNil(o.Data) {
		var ret []EmailChannelRoutingRules
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChannelRoutingRulesCollectionResponse) GetDataOk() ([]EmailChannelRoutingRules, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailChannelRoutingRulesCollectionResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EmailChannelRoutingRules and assigns it to the Data field.
func (o *EmailChannelRoutingRulesCollectionResponse) SetData(v []EmailChannelRoutingRules) {
	o.Data = v
}

func (o EmailChannelRoutingRulesCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailChannelRoutingRulesCollectionResponse) ToMap() (map[string]interface{}, error) {
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

type NullableEmailChannelRoutingRulesCollectionResponse struct {
	value *EmailChannelRoutingRulesCollectionResponse
	isSet bool
}

func (v NullableEmailChannelRoutingRulesCollectionResponse) Get() *EmailChannelRoutingRulesCollectionResponse {
	return v.value
}

func (v *NullableEmailChannelRoutingRulesCollectionResponse) Set(val *EmailChannelRoutingRulesCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailChannelRoutingRulesCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailChannelRoutingRulesCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailChannelRoutingRulesCollectionResponse(val *EmailChannelRoutingRulesCollectionResponse) *NullableEmailChannelRoutingRulesCollectionResponse {
	return &NullableEmailChannelRoutingRulesCollectionResponse{value: val, isSet: true}
}

func (v NullableEmailChannelRoutingRulesCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailChannelRoutingRulesCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


