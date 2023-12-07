# EmailRoutingRuleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportId** | **string** | The ID of the Transport | 
**SenderId** | **string** | The ID of the Sender | 
**IsReserved** | Pointer to **bool** | Whether the Routing-rule is reserved or not | [optional] 

## Methods

### NewEmailRoutingRuleCreateRequest

`func NewEmailRoutingRuleCreateRequest(transportId string, senderId string, ) *EmailRoutingRuleCreateRequest`

NewEmailRoutingRuleCreateRequest instantiates a new EmailRoutingRuleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRoutingRuleCreateRequestWithDefaults

`func NewEmailRoutingRuleCreateRequestWithDefaults() *EmailRoutingRuleCreateRequest`

NewEmailRoutingRuleCreateRequestWithDefaults instantiates a new EmailRoutingRuleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportId

`func (o *EmailRoutingRuleCreateRequest) GetTransportId() string`

GetTransportId returns the TransportId field if non-nil, zero value otherwise.

### GetTransportIdOk

`func (o *EmailRoutingRuleCreateRequest) GetTransportIdOk() (*string, bool)`

GetTransportIdOk returns a tuple with the TransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportId

`func (o *EmailRoutingRuleCreateRequest) SetTransportId(v string)`

SetTransportId sets TransportId field to given value.


### GetSenderId

`func (o *EmailRoutingRuleCreateRequest) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *EmailRoutingRuleCreateRequest) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *EmailRoutingRuleCreateRequest) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.


### GetIsReserved

`func (o *EmailRoutingRuleCreateRequest) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *EmailRoutingRuleCreateRequest) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *EmailRoutingRuleCreateRequest) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.

### HasIsReserved

`func (o *EmailRoutingRuleCreateRequest) HasIsReserved() bool`

HasIsReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


