/*
airEml - API

airEml - API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AirEml

import (
	"encoding/json"
	"fmt"
)

// checks if the EmailTemplateContentCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateContentCreateRequest{}

// EmailTemplateContentCreateRequest struct for EmailTemplateContentCreateRequest
type EmailTemplateContentCreateRequest struct {
	// The is a subject of the email notification.
	Subject string `json:"subject"`
	// This is an HTML part of the body of the email notification.
	TextHtml string `json:"text_html"`
	// This is a TXT part of the body of the email notification.
	TextPlain *string `json:"text_plain,omitempty"`
	// This is an AMP part of the body of the email notification.
	TextAmp *string `json:"text_amp,omitempty"`
	// This is settings for UI Template Builder. Must be valid JSON.
	BuilderHtml *string `json:"builder_html,omitempty"`
}

type _EmailTemplateContentCreateRequest EmailTemplateContentCreateRequest

// NewEmailTemplateContentCreateRequest instantiates a new EmailTemplateContentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplateContentCreateRequest(subject string, textHtml string) *EmailTemplateContentCreateRequest {
	this := EmailTemplateContentCreateRequest{}
	this.Subject = subject
	this.TextHtml = textHtml
	return &this
}

// NewEmailTemplateContentCreateRequestWithDefaults instantiates a new EmailTemplateContentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateContentCreateRequestWithDefaults() *EmailTemplateContentCreateRequest {
	this := EmailTemplateContentCreateRequest{}
	return &this
}

// GetSubject returns the Subject field value
func (o *EmailTemplateContentCreateRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EmailTemplateContentCreateRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EmailTemplateContentCreateRequest) SetSubject(v string) {
	o.Subject = v
}

// GetTextHtml returns the TextHtml field value
func (o *EmailTemplateContentCreateRequest) GetTextHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TextHtml
}

// GetTextHtmlOk returns a tuple with the TextHtml field value
// and a boolean to check if the value has been set.
func (o *EmailTemplateContentCreateRequest) GetTextHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TextHtml, true
}

// SetTextHtml sets field value
func (o *EmailTemplateContentCreateRequest) SetTextHtml(v string) {
	o.TextHtml = v
}

// GetTextPlain returns the TextPlain field value if set, zero value otherwise.
func (o *EmailTemplateContentCreateRequest) GetTextPlain() string {
	if o == nil || IsNil(o.TextPlain) {
		var ret string
		return ret
	}
	return *o.TextPlain
}

// GetTextPlainOk returns a tuple with the TextPlain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContentCreateRequest) GetTextPlainOk() (*string, bool) {
	if o == nil || IsNil(o.TextPlain) {
		return nil, false
	}
	return o.TextPlain, true
}

// HasTextPlain returns a boolean if a field has been set.
func (o *EmailTemplateContentCreateRequest) HasTextPlain() bool {
	if o != nil && !IsNil(o.TextPlain) {
		return true
	}

	return false
}

// SetTextPlain gets a reference to the given string and assigns it to the TextPlain field.
func (o *EmailTemplateContentCreateRequest) SetTextPlain(v string) {
	o.TextPlain = &v
}

// GetTextAmp returns the TextAmp field value if set, zero value otherwise.
func (o *EmailTemplateContentCreateRequest) GetTextAmp() string {
	if o == nil || IsNil(o.TextAmp) {
		var ret string
		return ret
	}
	return *o.TextAmp
}

// GetTextAmpOk returns a tuple with the TextAmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContentCreateRequest) GetTextAmpOk() (*string, bool) {
	if o == nil || IsNil(o.TextAmp) {
		return nil, false
	}
	return o.TextAmp, true
}

// HasTextAmp returns a boolean if a field has been set.
func (o *EmailTemplateContentCreateRequest) HasTextAmp() bool {
	if o != nil && !IsNil(o.TextAmp) {
		return true
	}

	return false
}

// SetTextAmp gets a reference to the given string and assigns it to the TextAmp field.
func (o *EmailTemplateContentCreateRequest) SetTextAmp(v string) {
	o.TextAmp = &v
}

// GetBuilderHtml returns the BuilderHtml field value if set, zero value otherwise.
func (o *EmailTemplateContentCreateRequest) GetBuilderHtml() string {
	if o == nil || IsNil(o.BuilderHtml) {
		var ret string
		return ret
	}
	return *o.BuilderHtml
}

// GetBuilderHtmlOk returns a tuple with the BuilderHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContentCreateRequest) GetBuilderHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.BuilderHtml) {
		return nil, false
	}
	return o.BuilderHtml, true
}

// HasBuilderHtml returns a boolean if a field has been set.
func (o *EmailTemplateContentCreateRequest) HasBuilderHtml() bool {
	if o != nil && !IsNil(o.BuilderHtml) {
		return true
	}

	return false
}

// SetBuilderHtml gets a reference to the given string and assigns it to the BuilderHtml field.
func (o *EmailTemplateContentCreateRequest) SetBuilderHtml(v string) {
	o.BuilderHtml = &v
}

func (o EmailTemplateContentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateContentCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subject"] = o.Subject
	toSerialize["text_html"] = o.TextHtml
	if !IsNil(o.TextPlain) {
		toSerialize["text_plain"] = o.TextPlain
	}
	if !IsNil(o.TextAmp) {
		toSerialize["text_amp"] = o.TextAmp
	}
	if !IsNil(o.BuilderHtml) {
		toSerialize["builder_html"] = o.BuilderHtml
	}
	return toSerialize, nil
}

func (o *EmailTemplateContentCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subject",
		"text_html",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEmailTemplateContentCreateRequest := _EmailTemplateContentCreateRequest{}

	err = json.Unmarshal(bytes, &varEmailTemplateContentCreateRequest)

	if err != nil {
		return err
	}

	*o = EmailTemplateContentCreateRequest(varEmailTemplateContentCreateRequest)

	return err
}

type NullableEmailTemplateContentCreateRequest struct {
	value *EmailTemplateContentCreateRequest
	isSet bool
}

func (v NullableEmailTemplateContentCreateRequest) Get() *EmailTemplateContentCreateRequest {
	return v.value
}

func (v *NullableEmailTemplateContentCreateRequest) Set(val *EmailTemplateContentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateContentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateContentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateContentCreateRequest(val *EmailTemplateContentCreateRequest) *NullableEmailTemplateContentCreateRequest {
	return &NullableEmailTemplateContentCreateRequest{value: val, isSet: true}
}

func (v NullableEmailTemplateContentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateContentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


