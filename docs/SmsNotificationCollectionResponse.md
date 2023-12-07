# SmsNotificationCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionLinks**](CollectionLinks.md) |  | [optional] 
**Meta** | Pointer to [**CollectionMeta**](CollectionMeta.md) |  | [optional] 
**Data** | Pointer to [**[]SmsNotification**](SmsNotification.md) |  | [optional] 

## Methods

### NewSmsNotificationCollectionResponse

`func NewSmsNotificationCollectionResponse() *SmsNotificationCollectionResponse`

NewSmsNotificationCollectionResponse instantiates a new SmsNotificationCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsNotificationCollectionResponseWithDefaults

`func NewSmsNotificationCollectionResponseWithDefaults() *SmsNotificationCollectionResponse`

NewSmsNotificationCollectionResponseWithDefaults instantiates a new SmsNotificationCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SmsNotificationCollectionResponse) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SmsNotificationCollectionResponse) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SmsNotificationCollectionResponse) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SmsNotificationCollectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *SmsNotificationCollectionResponse) GetMeta() CollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SmsNotificationCollectionResponse) GetMetaOk() (*CollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SmsNotificationCollectionResponse) SetMeta(v CollectionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SmsNotificationCollectionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *SmsNotificationCollectionResponse) GetData() []SmsNotification`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SmsNotificationCollectionResponse) GetDataOk() (*[]SmsNotification, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SmsNotificationCollectionResponse) SetData(v []SmsNotification)`

SetData sets Data field to given value.

### HasData

`func (o *SmsNotificationCollectionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


