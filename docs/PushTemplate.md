# PushTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Template. | [optional] 
**ChannelId** | Pointer to **string** | The unique identifier of the Channel. | [optional] 
**Name** | Pointer to **string** | The name of the Template. | [optional] 
**Badge** | Pointer to **int32** | The badge number to display on the app icon. | [optional] 
**Priority** | Pointer to **int32** | The notification delivery priority. Pass 2 to send the notification immediately, or pass 1 to send the notification at a time that conserves the battery on the user&#39;s device. | [optional] 
**Ttl** | Pointer to **int32** | The time-to-live (TTL) value, in seconds, for the notification. The TTL value is the amount of time that FCM attempts to deliver the notification. If the notification is not delivered within the TTL value, FCM no longer attempts to deliver the notification. | [optional] 
**Status** | Pointer to **int32** | The status of the Template. Every digit matters: 1 - active, 2 - inactive. | [optional] 
**Locales** | Pointer to **[]string** | List of locales available for this Template. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Template was created. | [optional] 

## Methods

### NewPushTemplate

`func NewPushTemplate() *PushTemplate`

NewPushTemplate instantiates a new PushTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushTemplateWithDefaults

`func NewPushTemplateWithDefaults() *PushTemplate`

NewPushTemplateWithDefaults instantiates a new PushTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PushTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PushTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PushTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PushTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelId

`func (o *PushTemplate) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PushTemplate) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PushTemplate) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *PushTemplate) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetName

`func (o *PushTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PushTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PushTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PushTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBadge

`func (o *PushTemplate) GetBadge() int32`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *PushTemplate) GetBadgeOk() (*int32, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *PushTemplate) SetBadge(v int32)`

SetBadge sets Badge field to given value.

### HasBadge

`func (o *PushTemplate) HasBadge() bool`

HasBadge returns a boolean if a field has been set.

### GetPriority

`func (o *PushTemplate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PushTemplate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PushTemplate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PushTemplate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTtl

`func (o *PushTemplate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PushTemplate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PushTemplate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PushTemplate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetStatus

`func (o *PushTemplate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PushTemplate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PushTemplate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PushTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLocales

`func (o *PushTemplate) GetLocales() []string`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *PushTemplate) GetLocalesOk() (*[]string, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *PushTemplate) SetLocales(v []string)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *PushTemplate) HasLocales() bool`

HasLocales returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PushTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PushTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PushTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PushTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


