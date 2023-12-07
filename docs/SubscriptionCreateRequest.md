# SubscriptionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A short explanation, what about this webhook subscription | 
**Events** | Pointer to **[]string** | List of events which webhook subscription subscribed | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **int32** | The status of the webhook subscription. Every digit matters:   1 - active,   2 - inactive. | [optional] 
**Url** | **string** | The url address to where need make post request | 

## Methods

### NewSubscriptionCreateRequest

`func NewSubscriptionCreateRequest(description string, url string, ) *SubscriptionCreateRequest`

NewSubscriptionCreateRequest instantiates a new SubscriptionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateRequestWithDefaults

`func NewSubscriptionCreateRequestWithDefaults() *SubscriptionCreateRequest`

NewSubscriptionCreateRequestWithDefaults instantiates a new SubscriptionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SubscriptionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEvents

`func (o *SubscriptionCreateRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SubscriptionCreateRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SubscriptionCreateRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SubscriptionCreateRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionCreateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionCreateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionCreateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionCreateRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionCreateRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionCreateRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *SubscriptionCreateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SubscriptionCreateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SubscriptionCreateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


