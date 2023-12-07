# SubscriptionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A short explanation, what about this webhook subscription | 
**Events** | **[]string** | List of events which webhook subscription subscribed | 
**Metadata** | **map[string]interface{}** |  | 
**Status** | **int32** | The status of the webhook subscription. Every digit matters:   1 - active,   2 - inactive. | 
**Url** | **string** | The url address to where need make post request | 

## Methods

### NewSubscriptionUpdateRequest

`func NewSubscriptionUpdateRequest(description string, events []string, metadata map[string]interface{}, status int32, url string, ) *SubscriptionUpdateRequest`

NewSubscriptionUpdateRequest instantiates a new SubscriptionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUpdateRequestWithDefaults

`func NewSubscriptionUpdateRequestWithDefaults() *SubscriptionUpdateRequest`

NewSubscriptionUpdateRequestWithDefaults instantiates a new SubscriptionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SubscriptionUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEvents

`func (o *SubscriptionUpdateRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SubscriptionUpdateRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SubscriptionUpdateRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetMetadata

`func (o *SubscriptionUpdateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionUpdateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionUpdateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *SubscriptionUpdateRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionUpdateRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionUpdateRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetUrl

`func (o *SubscriptionUpdateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SubscriptionUpdateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SubscriptionUpdateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


