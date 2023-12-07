# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sms** | Pointer to **[]string** |  | [optional] 
**Email** | Pointer to **[]string** |  | [optional] 
**Push** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSms

`func (o *Event) GetSms() []string`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *Event) GetSmsOk() (*[]string, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *Event) SetSms(v []string)`

SetSms sets Sms field to given value.

### HasSms

`func (o *Event) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetEmail

`func (o *Event) GetEmail() []string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Event) GetEmailOk() (*[]string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Event) SetEmail(v []string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Event) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPush

`func (o *Event) GetPush() []string`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *Event) GetPushOk() (*[]string, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *Event) SetPush(v []string)`

SetPush sets Push field to given value.

### HasPush

`func (o *Event) HasPush() bool`

HasPush returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


