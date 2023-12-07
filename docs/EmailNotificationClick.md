# EmailNotificationClick

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Click. | [optional] 
**Ip** | Pointer to **string** | The IP address from which the Click was made. | [optional] 
**UserAgent** | Pointer to **string** | The user agent from which the Click was made. | [optional] 
**Url** | Pointer to **string** | The URL that was clicked. | [optional] 
**Time** | Pointer to **time.Time** | The datetime when the Click was made. | [optional] 

## Methods

### NewEmailNotificationClick

`func NewEmailNotificationClick() *EmailNotificationClick`

NewEmailNotificationClick instantiates a new EmailNotificationClick object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationClickWithDefaults

`func NewEmailNotificationClickWithDefaults() *EmailNotificationClick`

NewEmailNotificationClickWithDefaults instantiates a new EmailNotificationClick object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailNotificationClick) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailNotificationClick) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailNotificationClick) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailNotificationClick) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *EmailNotificationClick) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *EmailNotificationClick) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *EmailNotificationClick) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *EmailNotificationClick) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *EmailNotificationClick) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *EmailNotificationClick) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *EmailNotificationClick) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *EmailNotificationClick) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUrl

`func (o *EmailNotificationClick) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmailNotificationClick) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmailNotificationClick) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EmailNotificationClick) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTime

`func (o *EmailNotificationClick) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EmailNotificationClick) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EmailNotificationClick) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *EmailNotificationClick) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


