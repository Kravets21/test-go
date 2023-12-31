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

// checks if the PushTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushTemplateRequest{}

// PushTemplateRequest struct for PushTemplateRequest
type PushTemplateRequest struct {
	// The name of the Template.
	Name string `json:"name"`
	// The badge number to display on the app icon.
	Badge *int32 `json:"badge,omitempty"`
	// The notification delivery priority. Pass 2 to send the notification immediately, or pass 1 to send the notification at a time that conserves the battery on the user's device.
	Priority *int32 `json:"priority,omitempty"`
	// The unique identifier of the Channel.
	ChannelId *string `json:"channel_id,omitempty"`
	// The time-to-live (TTL) value, in seconds, for the notification. The TTL value is the amount of time that FCM attempts to deliver the notification. If the notification is not delivered within the TTL value, FCM no longer attempts to deliver the notification.
	Ttl *int32 `json:"ttl,omitempty"`
}

type _PushTemplateRequest PushTemplateRequest

// NewPushTemplateRequest instantiates a new PushTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushTemplateRequest(name string) *PushTemplateRequest {
	this := PushTemplateRequest{}
	this.Name = name
	return &this
}

// NewPushTemplateRequestWithDefaults instantiates a new PushTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushTemplateRequestWithDefaults() *PushTemplateRequest {
	this := PushTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PushTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PushTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PushTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetBadge returns the Badge field value if set, zero value otherwise.
func (o *PushTemplateRequest) GetBadge() int32 {
	if o == nil || IsNil(o.Badge) {
		var ret int32
		return ret
	}
	return *o.Badge
}

// GetBadgeOk returns a tuple with the Badge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTemplateRequest) GetBadgeOk() (*int32, bool) {
	if o == nil || IsNil(o.Badge) {
		return nil, false
	}
	return o.Badge, true
}

// HasBadge returns a boolean if a field has been set.
func (o *PushTemplateRequest) HasBadge() bool {
	if o != nil && !IsNil(o.Badge) {
		return true
	}

	return false
}

// SetBadge gets a reference to the given int32 and assigns it to the Badge field.
func (o *PushTemplateRequest) SetBadge(v int32) {
	o.Badge = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PushTemplateRequest) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTemplateRequest) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PushTemplateRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *PushTemplateRequest) SetPriority(v int32) {
	o.Priority = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *PushTemplateRequest) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTemplateRequest) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *PushTemplateRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *PushTemplateRequest) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *PushTemplateRequest) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushTemplateRequest) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *PushTemplateRequest) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *PushTemplateRequest) SetTtl(v int32) {
	o.Ttl = &v
}

func (o PushTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Badge) {
		toSerialize["badge"] = o.Badge
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	return toSerialize, nil
}

func (o *PushTemplateRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varPushTemplateRequest := _PushTemplateRequest{}

	err = json.Unmarshal(bytes, &varPushTemplateRequest)

	if err != nil {
		return err
	}

	*o = PushTemplateRequest(varPushTemplateRequest)

	return err
}

type NullablePushTemplateRequest struct {
	value *PushTemplateRequest
	isSet bool
}

func (v NullablePushTemplateRequest) Get() *PushTemplateRequest {
	return v.value
}

func (v *NullablePushTemplateRequest) Set(val *PushTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePushTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePushTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushTemplateRequest(val *PushTemplateRequest) *NullablePushTemplateRequest {
	return &NullablePushTemplateRequest{value: val, isSet: true}
}

func (v NullablePushTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


