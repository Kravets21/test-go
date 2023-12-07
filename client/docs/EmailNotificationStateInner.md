# EmailNotificationStateInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | The state of the Notification status.  - 1: Notification Accepted  - 2: Notification Blocked  - 3: Notification Sent  - 4: Notification Not Sent  - 5: Notification Delivered  - 6: Notification Not Delivered | [optional] 
**Reason** | Pointer to **string** | The message of the Notification status. Contains the error message if the Notification was not sent or additional information about status state. | [optional] 
**Data** | Pointer to **map[string]interface{}** | Json string Data about this notification state step.  Helps to debug or understanding problem or just information. Notice: fields can be differently depends status and errors. | [optional] 
**Time** | Pointer to **time.Time** | The datetime when the Notification Status was updated. | [optional] 

## Methods

### NewEmailNotificationStateInner

`func NewEmailNotificationStateInner() *EmailNotificationStateInner`

NewEmailNotificationStateInner instantiates a new EmailNotificationStateInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationStateInnerWithDefaults

`func NewEmailNotificationStateInnerWithDefaults() *EmailNotificationStateInner`

NewEmailNotificationStateInnerWithDefaults instantiates a new EmailNotificationStateInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EmailNotificationStateInner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailNotificationStateInner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailNotificationStateInner) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailNotificationStateInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *EmailNotificationStateInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailNotificationStateInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailNotificationStateInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EmailNotificationStateInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetData

`func (o *EmailNotificationStateInner) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmailNotificationStateInner) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmailNotificationStateInner) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EmailNotificationStateInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTime

`func (o *EmailNotificationStateInner) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EmailNotificationStateInner) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EmailNotificationStateInner) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *EmailNotificationStateInner) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


