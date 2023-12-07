# PushNotificationSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to [**PushNotificationSendRequestTemplate**](PushNotificationSendRequestTemplate.md) |  | [optional] 
**To** | Pointer to [**PushNotificationSendRequestTo**](PushNotificationSendRequestTo.md) |  | [optional] 
**CustomParams** | Pointer to **map[string]interface{}** | The custom parameters of the notification. | [optional] 

## Methods

### NewPushNotificationSendRequest

`func NewPushNotificationSendRequest() *PushNotificationSendRequest`

NewPushNotificationSendRequest instantiates a new PushNotificationSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationSendRequestWithDefaults

`func NewPushNotificationSendRequestWithDefaults() *PushNotificationSendRequest`

NewPushNotificationSendRequestWithDefaults instantiates a new PushNotificationSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *PushNotificationSendRequest) GetTemplate() PushNotificationSendRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PushNotificationSendRequest) GetTemplateOk() (*PushNotificationSendRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PushNotificationSendRequest) SetTemplate(v PushNotificationSendRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PushNotificationSendRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTo

`func (o *PushNotificationSendRequest) GetTo() PushNotificationSendRequestTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PushNotificationSendRequest) GetToOk() (*PushNotificationSendRequestTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PushNotificationSendRequest) SetTo(v PushNotificationSendRequestTo)`

SetTo sets To field to given value.

### HasTo

`func (o *PushNotificationSendRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCustomParams

`func (o *PushNotificationSendRequest) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *PushNotificationSendRequest) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *PushNotificationSendRequest) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *PushNotificationSendRequest) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### SetCustomParamsNil

`func (o *PushNotificationSendRequest) SetCustomParamsNil(b bool)`

 SetCustomParamsNil sets the value for CustomParams to be an explicit nil

### UnsetCustomParams
`func (o *PushNotificationSendRequest) UnsetCustomParams()`

UnsetCustomParams ensures that no value is present for CustomParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


