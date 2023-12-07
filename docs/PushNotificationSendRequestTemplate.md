# PushNotificationSendRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the template to use for the notification. The id or name parameter must be specified. | [optional] 
**Name** | Pointer to **string** | The name of the template to use for the notification. The id or name parameter must be specified. | [optional] 
**Locale** | Pointer to **string** | The locale of the template to use for the notification. | [optional] 
**Variables** | Pointer to **map[string]interface{}** | The parameters used to replace placeholders of the Template Content Where key - is a placeholder name, and value - is a placeholder value. If the value is not provided, the placeholder will use the default value from the Template Content. | [optional] 
**Data** | Pointer to **map[string]string** | The additional data used for send push notification if user wants to add some custom data in request | [optional] 

## Methods

### NewPushNotificationSendRequestTemplate

`func NewPushNotificationSendRequestTemplate() *PushNotificationSendRequestTemplate`

NewPushNotificationSendRequestTemplate instantiates a new PushNotificationSendRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushNotificationSendRequestTemplateWithDefaults

`func NewPushNotificationSendRequestTemplateWithDefaults() *PushNotificationSendRequestTemplate`

NewPushNotificationSendRequestTemplateWithDefaults instantiates a new PushNotificationSendRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PushNotificationSendRequestTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PushNotificationSendRequestTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PushNotificationSendRequestTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PushNotificationSendRequestTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PushNotificationSendRequestTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PushNotificationSendRequestTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PushNotificationSendRequestTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PushNotificationSendRequestTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocale

`func (o *PushNotificationSendRequestTemplate) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PushNotificationSendRequestTemplate) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PushNotificationSendRequestTemplate) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PushNotificationSendRequestTemplate) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetVariables

`func (o *PushNotificationSendRequestTemplate) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PushNotificationSendRequestTemplate) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PushNotificationSendRequestTemplate) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *PushNotificationSendRequestTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetData

`func (o *PushNotificationSendRequestTemplate) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PushNotificationSendRequestTemplate) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PushNotificationSendRequestTemplate) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *PushNotificationSendRequestTemplate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


