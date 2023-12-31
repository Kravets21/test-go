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

// checks if the EmailTemplateContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateContent{}

// EmailTemplateContent struct for EmailTemplateContent
type EmailTemplateContent struct {
	Locale *LocaleCode `json:"locale,omitempty"`
	// The is a subject of the email notification.
	Subject *string `json:"subject,omitempty"`
	// The is a text body of the email notification.
	TextPlain *string `json:"text_plain,omitempty"`
	// The is a html body of the email notification. Must be valid HTML.
	TextHtml *string `json:"text_html,omitempty"`
	// The is a amp body of the email notification.
	TextAmp *string `json:"text_amp,omitempty"`
	// This is settings for UI Template Builder. Must be valid JSON.
	BuilderHtml *string `json:"builder_html,omitempty"`
	VariablesSettings []EmailTemplateContentVariablesSettingsInner `json:"variables_settings,omitempty"`
	// The datetime when the Template Content was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewEmailTemplateContent instantiates a new EmailTemplateContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplateContent() *EmailTemplateContent {
	this := EmailTemplateContent{}
	return &this
}

// NewEmailTemplateContentWithDefaults instantiates a new EmailTemplateContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateContentWithDefaults() *EmailTemplateContent {
	this := EmailTemplateContent{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *EmailTemplateContent) GetLocale() LocaleCode {
	if o == nil || IsNil(o.Locale) {
		var ret LocaleCode
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContent) GetLocaleOk() (*LocaleCode, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *EmailTemplateContent) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given LocaleCode and assigns it to the Locale field.
func (o *EmailTemplateContent) SetLocale(v LocaleCode) {
	o.Locale = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailTemplateContent) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContent) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailTemplateContent) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *EmailTemplateContent) SetSubject(v string) {
	o.Subject = &v
}

// GetTextPlain returns the TextPlain field value if set, zero value otherwise.
func (o *EmailTemplateContent) GetTextPlain() string {
	if o == nil || IsNil(o.TextPlain) {
		var ret string
		return ret
	}
	return *o.TextPlain
}

// GetTextPlainOk returns a tuple with the TextPlain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContent) GetTextPlainOk() (*string, bool) {
	if o == nil || IsNil(o.TextPlain) {
		return nil, false
	}
	return o.TextPlain, true
}

// HasTextPlain returns a boolean if a field has been set.
func (o *EmailTemplateContent) HasTextPlain() bool {
	if o != nil && !IsNil(o.TextPlain) {
		return true
	}

	return false
}

// SetTextPlain gets a reference to the given string and assigns it to the TextPlain field.
func (o *EmailTemplateContent) SetTextPlain(v string) {
	o.TextPlain = &v
}

// GetTextHtml returns the TextHtml field value if set, zero value otherwise.
func (o *EmailTemplateContent) GetTextHtml() string {
	if o == nil || IsNil(o.TextHtml) {
		var ret string
		return ret
	}
	return *o.TextHtml
}

// GetTextHtmlOk returns a tuple with the TextHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContent) GetTextHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.TextHtml) {
		return nil, false
	}
	return o.TextHtml, true
}

// HasTextHtml returns a boolean if a field has been set.
func (o *EmailTemplateContent) HasTextHtml() bool {
	if o != nil && !IsNil(o.TextHtml) {
		return true
	}

	return false
}

// SetTextHtml gets a reference to the given string and assigns it to the TextHtml field.
func (o *EmailTemplateContent) SetTextHtml(v string) {
	o.TextHtml = &v
}

// GetTextAmp returns the TextAmp field value if set, zero value otherwise.
func (o *EmailTemplateContent) GetTextAmp() string {
	if o == nil || IsNil(o.TextAmp) {
		var ret string
		return ret
	}
	return *o.TextAmp
}

// GetTextAmpOk returns a tuple with the TextAmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContent) GetTextAmpOk() (*string, bool) {
	if o == nil || IsNil(o.TextAmp) {
		return nil, false
	}
	return o.TextAmp, true
}

// HasTextAmp returns a boolean if a field has been set.
func (o *EmailTemplateContent) HasTextAmp() bool {
	if o != nil && !IsNil(o.TextAmp) {
		return true
	}

	return false
}

// SetTextAmp gets a reference to the given string and assigns it to the TextAmp field.
func (o *EmailTemplateContent) SetTextAmp(v string) {
	o.TextAmp = &v
}

// GetBuilderHtml returns the BuilderHtml field value if set, zero value otherwise.
func (o *EmailTemplateContent) GetBuilderHtml() string {
	if o == nil || IsNil(o.BuilderHtml) {
		var ret string
		return ret
	}
	return *o.BuilderHtml
}

// GetBuilderHtmlOk returns a tuple with the BuilderHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContent) GetBuilderHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.BuilderHtml) {
		return nil, false
	}
	return o.BuilderHtml, true
}

// HasBuilderHtml returns a boolean if a field has been set.
func (o *EmailTemplateContent) HasBuilderHtml() bool {
	if o != nil && !IsNil(o.BuilderHtml) {
		return true
	}

	return false
}

// SetBuilderHtml gets a reference to the given string and assigns it to the BuilderHtml field.
func (o *EmailTemplateContent) SetBuilderHtml(v string) {
	o.BuilderHtml = &v
}

// GetVariablesSettings returns the VariablesSettings field value if set, zero value otherwise.
func (o *EmailTemplateContent) GetVariablesSettings() []EmailTemplateContentVariablesSettingsInner {
	if o == nil || IsNil(o.VariablesSettings) {
		var ret []EmailTemplateContentVariablesSettingsInner
		return ret
	}
	return o.VariablesSettings
}

// GetVariablesSettingsOk returns a tuple with the VariablesSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContent) GetVariablesSettingsOk() ([]EmailTemplateContentVariablesSettingsInner, bool) {
	if o == nil || IsNil(o.VariablesSettings) {
		return nil, false
	}
	return o.VariablesSettings, true
}

// HasVariablesSettings returns a boolean if a field has been set.
func (o *EmailTemplateContent) HasVariablesSettings() bool {
	if o != nil && !IsNil(o.VariablesSettings) {
		return true
	}

	return false
}

// SetVariablesSettings gets a reference to the given []EmailTemplateContentVariablesSettingsInner and assigns it to the VariablesSettings field.
func (o *EmailTemplateContent) SetVariablesSettings(v []EmailTemplateContentVariablesSettingsInner) {
	o.VariablesSettings = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EmailTemplateContent) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateContent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EmailTemplateContent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EmailTemplateContent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o EmailTemplateContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.TextPlain) {
		toSerialize["text_plain"] = o.TextPlain
	}
	if !IsNil(o.TextHtml) {
		toSerialize["text_html"] = o.TextHtml
	}
	if !IsNil(o.TextAmp) {
		toSerialize["text_amp"] = o.TextAmp
	}
	if !IsNil(o.BuilderHtml) {
		toSerialize["builder_html"] = o.BuilderHtml
	}
	if !IsNil(o.VariablesSettings) {
		toSerialize["variables_settings"] = o.VariablesSettings
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableEmailTemplateContent struct {
	value *EmailTemplateContent
	isSet bool
}

func (v NullableEmailTemplateContent) Get() *EmailTemplateContent {
	return v.value
}

func (v *NullableEmailTemplateContent) Set(val *EmailTemplateContent) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateContent) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateContent(val *EmailTemplateContent) *NullableEmailTemplateContent {
	return &NullableEmailTemplateContent{value: val, isSet: true}
}

func (v NullableEmailTemplateContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


