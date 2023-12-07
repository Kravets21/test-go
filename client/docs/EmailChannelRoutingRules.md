# EmailChannelRoutingRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the Routing-rule | [optional] 
**TransportId** | Pointer to **string** | The ID of the Transport | [optional] 
**SenderId** | Pointer to **string** | The ID of the Sender | [optional] 
**IsReserved** | Pointer to **bool** | Whether the Routing-rule is reserved or not | [optional] 
**Status** | Pointer to **int32** | The status of the Routing-rule. Every digit matters: 1 - active, 2 - inactive. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Template Content was created. | [optional] 

## Methods

### NewEmailChannelRoutingRules

`func NewEmailChannelRoutingRules() *EmailChannelRoutingRules`

NewEmailChannelRoutingRules instantiates a new EmailChannelRoutingRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailChannelRoutingRulesWithDefaults

`func NewEmailChannelRoutingRulesWithDefaults() *EmailChannelRoutingRules`

NewEmailChannelRoutingRulesWithDefaults instantiates a new EmailChannelRoutingRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailChannelRoutingRules) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailChannelRoutingRules) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailChannelRoutingRules) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailChannelRoutingRules) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTransportId

`func (o *EmailChannelRoutingRules) GetTransportId() string`

GetTransportId returns the TransportId field if non-nil, zero value otherwise.

### GetTransportIdOk

`func (o *EmailChannelRoutingRules) GetTransportIdOk() (*string, bool)`

GetTransportIdOk returns a tuple with the TransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportId

`func (o *EmailChannelRoutingRules) SetTransportId(v string)`

SetTransportId sets TransportId field to given value.

### HasTransportId

`func (o *EmailChannelRoutingRules) HasTransportId() bool`

HasTransportId returns a boolean if a field has been set.

### GetSenderId

`func (o *EmailChannelRoutingRules) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *EmailChannelRoutingRules) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *EmailChannelRoutingRules) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *EmailChannelRoutingRules) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetIsReserved

`func (o *EmailChannelRoutingRules) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *EmailChannelRoutingRules) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *EmailChannelRoutingRules) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.

### HasIsReserved

`func (o *EmailChannelRoutingRules) HasIsReserved() bool`

HasIsReserved returns a boolean if a field has been set.

### GetStatus

`func (o *EmailChannelRoutingRules) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailChannelRoutingRules) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailChannelRoutingRules) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailChannelRoutingRules) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EmailChannelRoutingRules) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailChannelRoutingRules) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailChannelRoutingRules) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EmailChannelRoutingRules) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


