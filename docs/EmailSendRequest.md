# EmailSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to [**EmailSendRequestTemplate**](EmailSendRequestTemplate.md) |  | [optional] 
**To** | Pointer to [**EmailSendRequestTo**](EmailSendRequestTo.md) |  | [optional] 
**From** | Pointer to [**EmailSendRequestFrom**](EmailSendRequestFrom.md) |  | [optional] 
**Attachments** | Pointer to [**[]EmailSendRequestAttachmentsInner**](EmailSendRequestAttachmentsInner.md) |  | [optional] 
**CustomParams** | Pointer to **map[string]interface{}** | The custom parameters of the notification. | [optional] 

## Methods

### NewEmailSendRequest

`func NewEmailSendRequest() *EmailSendRequest`

NewEmailSendRequest instantiates a new EmailSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendRequestWithDefaults

`func NewEmailSendRequestWithDefaults() *EmailSendRequest`

NewEmailSendRequestWithDefaults instantiates a new EmailSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *EmailSendRequest) GetTemplate() EmailSendRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EmailSendRequest) GetTemplateOk() (*EmailSendRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EmailSendRequest) SetTemplate(v EmailSendRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EmailSendRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTo

`func (o *EmailSendRequest) GetTo() EmailSendRequestTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailSendRequest) GetToOk() (*EmailSendRequestTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailSendRequest) SetTo(v EmailSendRequestTo)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailSendRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *EmailSendRequest) GetFrom() EmailSendRequestFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailSendRequest) GetFromOk() (*EmailSendRequestFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailSendRequest) SetFrom(v EmailSendRequestFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailSendRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetAttachments

`func (o *EmailSendRequest) GetAttachments() []EmailSendRequestAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *EmailSendRequest) GetAttachmentsOk() (*[]EmailSendRequestAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *EmailSendRequest) SetAttachments(v []EmailSendRequestAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *EmailSendRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCustomParams

`func (o *EmailSendRequest) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *EmailSendRequest) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *EmailSendRequest) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *EmailSendRequest) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### SetCustomParamsNil

`func (o *EmailSendRequest) SetCustomParamsNil(b bool)`

 SetCustomParamsNil sets the value for CustomParams to be an explicit nil

### UnsetCustomParams
`func (o *EmailSendRequest) UnsetCustomParams()`

UnsetCustomParams ensures that no value is present for CustomParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


