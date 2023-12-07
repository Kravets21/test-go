# EmailRoutingRuleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportId** | Pointer to **NullableString** | The ID of the Transport | [optional] 
**SenderId** | Pointer to **string** | The ID of the Sender | [optional] 
**IsReserved** | Pointer to **NullableBool** | Whether the Routing-rule is reserved or not | [optional] 
**Status** | Pointer to **NullableInt32** | The status of the Routing-rule. Every digit matters: 1 - active, 2 - inactive. | [optional] 

## Methods

### NewEmailRoutingRuleUpdateRequest

`func NewEmailRoutingRuleUpdateRequest() *EmailRoutingRuleUpdateRequest`

NewEmailRoutingRuleUpdateRequest instantiates a new EmailRoutingRuleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRoutingRuleUpdateRequestWithDefaults

`func NewEmailRoutingRuleUpdateRequestWithDefaults() *EmailRoutingRuleUpdateRequest`

NewEmailRoutingRuleUpdateRequestWithDefaults instantiates a new EmailRoutingRuleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportId

`func (o *EmailRoutingRuleUpdateRequest) GetTransportId() string`

GetTransportId returns the TransportId field if non-nil, zero value otherwise.

### GetTransportIdOk

`func (o *EmailRoutingRuleUpdateRequest) GetTransportIdOk() (*string, bool)`

GetTransportIdOk returns a tuple with the TransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportId

`func (o *EmailRoutingRuleUpdateRequest) SetTransportId(v string)`

SetTransportId sets TransportId field to given value.

### HasTransportId

`func (o *EmailRoutingRuleUpdateRequest) HasTransportId() bool`

HasTransportId returns a boolean if a field has been set.

### SetTransportIdNil

`func (o *EmailRoutingRuleUpdateRequest) SetTransportIdNil(b bool)`

 SetTransportIdNil sets the value for TransportId to be an explicit nil

### UnsetTransportId
`func (o *EmailRoutingRuleUpdateRequest) UnsetTransportId()`

UnsetTransportId ensures that no value is present for TransportId, not even an explicit nil
### GetSenderId

`func (o *EmailRoutingRuleUpdateRequest) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *EmailRoutingRuleUpdateRequest) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *EmailRoutingRuleUpdateRequest) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *EmailRoutingRuleUpdateRequest) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetIsReserved

`func (o *EmailRoutingRuleUpdateRequest) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *EmailRoutingRuleUpdateRequest) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *EmailRoutingRuleUpdateRequest) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.

### HasIsReserved

`func (o *EmailRoutingRuleUpdateRequest) HasIsReserved() bool`

HasIsReserved returns a boolean if a field has been set.

### SetIsReservedNil

`func (o *EmailRoutingRuleUpdateRequest) SetIsReservedNil(b bool)`

 SetIsReservedNil sets the value for IsReserved to be an explicit nil

### UnsetIsReserved
`func (o *EmailRoutingRuleUpdateRequest) UnsetIsReserved()`

UnsetIsReserved ensures that no value is present for IsReserved, not even an explicit nil
### GetStatus

`func (o *EmailRoutingRuleUpdateRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailRoutingRuleUpdateRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailRoutingRuleUpdateRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailRoutingRuleUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EmailRoutingRuleUpdateRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EmailRoutingRuleUpdateRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


