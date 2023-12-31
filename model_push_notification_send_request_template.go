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

// checks if the PushNotificationSendRequestTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushNotificationSendRequestTemplate{}

// PushNotificationSendRequestTemplate struct for PushNotificationSendRequestTemplate
type PushNotificationSendRequestTemplate struct {
	// The ID of the template to use for the notification. The id or name parameter must be specified.
	Id *string `json:"id,omitempty"`
	// The name of the template to use for the notification. The id or name parameter must be specified.
	Name *string `json:"name,omitempty"`
	// The locale of the template to use for the notification.
	Locale *string `json:"locale,omitempty"`
	// The parameters used to replace placeholders of the Template Content Where key - is a placeholder name, and value - is a placeholder value. If the value is not provided, the placeholder will use the default value from the Template Content.
	Variables map[string]interface{} `json:"variables,omitempty"`
	// The additional data used for send push notification if user wants to add some custom data in request
	Data *map[string]string `json:"data,omitempty"`
}

// NewPushNotificationSendRequestTemplate instantiates a new PushNotificationSendRequestTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushNotificationSendRequestTemplate() *PushNotificationSendRequestTemplate {
	this := PushNotificationSendRequestTemplate{}
	return &this
}

// NewPushNotificationSendRequestTemplateWithDefaults instantiates a new PushNotificationSendRequestTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushNotificationSendRequestTemplateWithDefaults() *PushNotificationSendRequestTemplate {
	this := PushNotificationSendRequestTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PushNotificationSendRequestTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushNotificationSendRequestTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PushNotificationSendRequestTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PushNotificationSendRequestTemplate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PushNotificationSendRequestTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushNotificationSendRequestTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PushNotificationSendRequestTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PushNotificationSendRequestTemplate) SetName(v string) {
	o.Name = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *PushNotificationSendRequestTemplate) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushNotificationSendRequestTemplate) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *PushNotificationSendRequestTemplate) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *PushNotificationSendRequestTemplate) SetLocale(v string) {
	o.Locale = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *PushNotificationSendRequestTemplate) GetVariables() map[string]interface{} {
	if o == nil || IsNil(o.Variables) {
		var ret map[string]interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushNotificationSendRequestTemplate) GetVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Variables) {
		return map[string]interface{}{}, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *PushNotificationSendRequestTemplate) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *PushNotificationSendRequestTemplate) SetVariables(v map[string]interface{}) {
	o.Variables = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PushNotificationSendRequestTemplate) GetData() map[string]string {
	if o == nil || IsNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushNotificationSendRequestTemplate) GetDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PushNotificationSendRequestTemplate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *PushNotificationSendRequestTemplate) SetData(v map[string]string) {
	o.Data = &v
}

func (o PushNotificationSendRequestTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushNotificationSendRequestTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePushNotificationSendRequestTemplate struct {
	value *PushNotificationSendRequestTemplate
	isSet bool
}

func (v NullablePushNotificationSendRequestTemplate) Get() *PushNotificationSendRequestTemplate {
	return v.value
}

func (v *NullablePushNotificationSendRequestTemplate) Set(val *PushNotificationSendRequestTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullablePushNotificationSendRequestTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullablePushNotificationSendRequestTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushNotificationSendRequestTemplate(val *PushNotificationSendRequestTemplate) *NullablePushNotificationSendRequestTemplate {
	return &NullablePushNotificationSendRequestTemplate{value: val, isSet: true}
}

func (v NullablePushNotificationSendRequestTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushNotificationSendRequestTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


