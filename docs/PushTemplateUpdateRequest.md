# PushTemplateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Template. | [optional] 
**Badge** | Pointer to **int32** | The badge number to display on the app icon. | [optional] 
**Priority** | Pointer to **int32** | The notification delivery priority. Pass 2 to send the notification immediately, or pass 1 to send the notification at a time that conserves the battery on the user&#39;s device. | [optional] 
**Ttl** | Pointer to **int32** | The time-to-live (TTL) value, in seconds, for the notification. The TTL value is the amount of time that FCM attempts to deliver the notification. If the notification is not delivered within the TTL value, FCM no longer attempts to deliver the notification. | [optional] 
**Status** | Pointer to **int32** | The status of the Template. | [optional] 

## Methods

### NewPushTemplateUpdateRequest

`func NewPushTemplateUpdateRequest() *PushTemplateUpdateRequest`

NewPushTemplateUpdateRequest instantiates a new PushTemplateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushTemplateUpdateRequestWithDefaults

`func NewPushTemplateUpdateRequestWithDefaults() *PushTemplateUpdateRequest`

NewPushTemplateUpdateRequestWithDefaults instantiates a new PushTemplateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PushTemplateUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PushTemplateUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PushTemplateUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PushTemplateUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBadge

`func (o *PushTemplateUpdateRequest) GetBadge() int32`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *PushTemplateUpdateRequest) GetBadgeOk() (*int32, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *PushTemplateUpdateRequest) SetBadge(v int32)`

SetBadge sets Badge field to given value.

### HasBadge

`func (o *PushTemplateUpdateRequest) HasBadge() bool`

HasBadge returns a boolean if a field has been set.

### GetPriority

`func (o *PushTemplateUpdateRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PushTemplateUpdateRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PushTemplateUpdateRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PushTemplateUpdateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTtl

`func (o *PushTemplateUpdateRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PushTemplateUpdateRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PushTemplateUpdateRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PushTemplateUpdateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetStatus

`func (o *PushTemplateUpdateRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PushTemplateUpdateRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PushTemplateUpdateRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PushTemplateUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


