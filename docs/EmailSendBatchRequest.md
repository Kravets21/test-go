# EmailSendBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | [**[]EmailSendRequest**](EmailSendRequest.md) | Array of notifications which need to send | 

## Methods

### NewEmailSendBatchRequest

`func NewEmailSendBatchRequest(notifications []EmailSendRequest, ) *EmailSendBatchRequest`

NewEmailSendBatchRequest instantiates a new EmailSendBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendBatchRequestWithDefaults

`func NewEmailSendBatchRequestWithDefaults() *EmailSendBatchRequest`

NewEmailSendBatchRequestWithDefaults instantiates a new EmailSendBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *EmailSendBatchRequest) GetNotifications() []EmailSendRequest`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *EmailSendBatchRequest) GetNotificationsOk() (*[]EmailSendRequest, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *EmailSendBatchRequest) SetNotifications(v []EmailSendRequest)`

SetNotifications sets Notifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


