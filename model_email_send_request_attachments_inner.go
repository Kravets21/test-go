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

// checks if the EmailSendRequestAttachmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSendRequestAttachmentsInner{}

// EmailSendRequestAttachmentsInner struct for EmailSendRequestAttachmentsInner
type EmailSendRequestAttachmentsInner struct {
	// The MIME type of the attachment file.
	ContentType *string `json:"content_type,omitempty"`
	// The filename of the attachment file.
	FileName *string `json:"file_name,omitempty"`
	// The content of the attachment file in base64 encoded characters.
	ContentBase64 *string `json:"content_base64,omitempty"`
}

// NewEmailSendRequestAttachmentsInner instantiates a new EmailSendRequestAttachmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSendRequestAttachmentsInner() *EmailSendRequestAttachmentsInner {
	this := EmailSendRequestAttachmentsInner{}
	return &this
}

// NewEmailSendRequestAttachmentsInnerWithDefaults instantiates a new EmailSendRequestAttachmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSendRequestAttachmentsInnerWithDefaults() *EmailSendRequestAttachmentsInner {
	this := EmailSendRequestAttachmentsInner{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *EmailSendRequestAttachmentsInner) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendRequestAttachmentsInner) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *EmailSendRequestAttachmentsInner) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *EmailSendRequestAttachmentsInner) SetContentType(v string) {
	o.ContentType = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *EmailSendRequestAttachmentsInner) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendRequestAttachmentsInner) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *EmailSendRequestAttachmentsInner) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *EmailSendRequestAttachmentsInner) SetFileName(v string) {
	o.FileName = &v
}

// GetContentBase64 returns the ContentBase64 field value if set, zero value otherwise.
func (o *EmailSendRequestAttachmentsInner) GetContentBase64() string {
	if o == nil || IsNil(o.ContentBase64) {
		var ret string
		return ret
	}
	return *o.ContentBase64
}

// GetContentBase64Ok returns a tuple with the ContentBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendRequestAttachmentsInner) GetContentBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.ContentBase64) {
		return nil, false
	}
	return o.ContentBase64, true
}

// HasContentBase64 returns a boolean if a field has been set.
func (o *EmailSendRequestAttachmentsInner) HasContentBase64() bool {
	if o != nil && !IsNil(o.ContentBase64) {
		return true
	}

	return false
}

// SetContentBase64 gets a reference to the given string and assigns it to the ContentBase64 field.
func (o *EmailSendRequestAttachmentsInner) SetContentBase64(v string) {
	o.ContentBase64 = &v
}

func (o EmailSendRequestAttachmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSendRequestAttachmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.ContentBase64) {
		toSerialize["content_base64"] = o.ContentBase64
	}
	return toSerialize, nil
}

type NullableEmailSendRequestAttachmentsInner struct {
	value *EmailSendRequestAttachmentsInner
	isSet bool
}

func (v NullableEmailSendRequestAttachmentsInner) Get() *EmailSendRequestAttachmentsInner {
	return v.value
}

func (v *NullableEmailSendRequestAttachmentsInner) Set(val *EmailSendRequestAttachmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSendRequestAttachmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSendRequestAttachmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSendRequestAttachmentsInner(val *EmailSendRequestAttachmentsInner) *NullableEmailSendRequestAttachmentsInner {
	return &NullableEmailSendRequestAttachmentsInner{value: val, isSet: true}
}

func (v NullableEmailSendRequestAttachmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSendRequestAttachmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

