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

// checks if the AudienceSegmentCollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceSegmentCollectionResponse{}

// AudienceSegmentCollectionResponse struct for AudienceSegmentCollectionResponse
type AudienceSegmentCollectionResponse struct {
	Links *CollectionLinks `json:"links,omitempty"`
	Meta *CollectionMeta `json:"meta,omitempty"`
	Data []AudienceSegment `json:"data,omitempty"`
}

// NewAudienceSegmentCollectionResponse instantiates a new AudienceSegmentCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceSegmentCollectionResponse() *AudienceSegmentCollectionResponse {
	this := AudienceSegmentCollectionResponse{}
	return &this
}

// NewAudienceSegmentCollectionResponseWithDefaults instantiates a new AudienceSegmentCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceSegmentCollectionResponseWithDefaults() *AudienceSegmentCollectionResponse {
	this := AudienceSegmentCollectionResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AudienceSegmentCollectionResponse) GetLinks() CollectionLinks {
	if o == nil || IsNil(o.Links) {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSegmentCollectionResponse) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AudienceSegmentCollectionResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *AudienceSegmentCollectionResponse) SetLinks(v CollectionLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AudienceSegmentCollectionResponse) GetMeta() CollectionMeta {
	if o == nil || IsNil(o.Meta) {
		var ret CollectionMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSegmentCollectionResponse) GetMetaOk() (*CollectionMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AudienceSegmentCollectionResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMeta and assigns it to the Meta field.
func (o *AudienceSegmentCollectionResponse) SetMeta(v CollectionMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AudienceSegmentCollectionResponse) GetData() []AudienceSegment {
	if o == nil || IsNil(o.Data) {
		var ret []AudienceSegment
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSegmentCollectionResponse) GetDataOk() ([]AudienceSegment, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AudienceSegmentCollectionResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AudienceSegment and assigns it to the Data field.
func (o *AudienceSegmentCollectionResponse) SetData(v []AudienceSegment) {
	o.Data = v
}

func (o AudienceSegmentCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudienceSegmentCollectionResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAudienceSegmentCollectionResponse struct {
	value *AudienceSegmentCollectionResponse
	isSet bool
}

func (v NullableAudienceSegmentCollectionResponse) Get() *AudienceSegmentCollectionResponse {
	return v.value
}

func (v *NullableAudienceSegmentCollectionResponse) Set(val *AudienceSegmentCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceSegmentCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceSegmentCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceSegmentCollectionResponse(val *AudienceSegmentCollectionResponse) *NullableAudienceSegmentCollectionResponse {
	return &NullableAudienceSegmentCollectionResponse{value: val, isSet: true}
}

func (v NullableAudienceSegmentCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceSegmentCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


