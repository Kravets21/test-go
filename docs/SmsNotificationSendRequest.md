# SmsNotificationSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to [**SmsNotificationSendRequestTemplate**](SmsNotificationSendRequestTemplate.md) |  | [optional] 
**To** | Pointer to [**SmsNotificationSendRequestTo**](SmsNotificationSendRequestTo.md) |  | [optional] 
**CustomParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSmsNotificationSendRequest

`func NewSmsNotificationSendRequest() *SmsNotificationSendRequest`

NewSmsNotificationSendRequest instantiates a new SmsNotificationSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsNotificationSendRequestWithDefaults

`func NewSmsNotificationSendRequestWithDefaults() *SmsNotificationSendRequest`

NewSmsNotificationSendRequestWithDefaults instantiates a new SmsNotificationSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *SmsNotificationSendRequest) GetTemplate() SmsNotificationSendRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SmsNotificationSendRequest) GetTemplateOk() (*SmsNotificationSendRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SmsNotificationSendRequest) SetTemplate(v SmsNotificationSendRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SmsNotificationSendRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTo

`func (o *SmsNotificationSendRequest) GetTo() SmsNotificationSendRequestTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SmsNotificationSendRequest) GetToOk() (*SmsNotificationSendRequestTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SmsNotificationSendRequest) SetTo(v SmsNotificationSendRequestTo)`

SetTo sets To field to given value.

### HasTo

`func (o *SmsNotificationSendRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCustomParams

`func (o *SmsNotificationSendRequest) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *SmsNotificationSendRequest) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *SmsNotificationSendRequest) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *SmsNotificationSendRequest) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### SetCustomParamsNil

`func (o *SmsNotificationSendRequest) SetCustomParamsNil(b bool)`

 SetCustomParamsNil sets the value for CustomParams to be an explicit nil

### UnsetCustomParams
`func (o *SmsNotificationSendRequest) UnsetCustomParams()`

UnsetCustomParams ensures that no value is present for CustomParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


