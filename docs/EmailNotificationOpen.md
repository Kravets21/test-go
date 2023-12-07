# EmailNotificationOpen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Open. | [optional] 
**Ip** | Pointer to **string** | The IP address from which the Open was made. | [optional] 
**UserAgent** | Pointer to **string** | The user agent from which the Open was made. | [optional] 
**Time** | Pointer to **time.Time** | The datetime when the Open was made. | [optional] 

## Methods

### NewEmailNotificationOpen

`func NewEmailNotificationOpen() *EmailNotificationOpen`

NewEmailNotificationOpen instantiates a new EmailNotificationOpen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationOpenWithDefaults

`func NewEmailNotificationOpenWithDefaults() *EmailNotificationOpen`

NewEmailNotificationOpenWithDefaults instantiates a new EmailNotificationOpen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailNotificationOpen) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailNotificationOpen) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailNotificationOpen) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailNotificationOpen) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *EmailNotificationOpen) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *EmailNotificationOpen) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *EmailNotificationOpen) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *EmailNotificationOpen) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *EmailNotificationOpen) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *EmailNotificationOpen) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *EmailNotificationOpen) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *EmailNotificationOpen) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetTime

`func (o *EmailNotificationOpen) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EmailNotificationOpen) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EmailNotificationOpen) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *EmailNotificationOpen) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


