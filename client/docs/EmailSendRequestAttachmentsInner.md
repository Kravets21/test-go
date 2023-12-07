# EmailSendRequestAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | The MIME type of the attachment file. | [optional] 
**FileName** | Pointer to **string** | The filename of the attachment file. | [optional] 
**ContentBase64** | Pointer to **string** | The content of the attachment file in base64 encoded characters. | [optional] 

## Methods

### NewEmailSendRequestAttachmentsInner

`func NewEmailSendRequestAttachmentsInner() *EmailSendRequestAttachmentsInner`

NewEmailSendRequestAttachmentsInner instantiates a new EmailSendRequestAttachmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendRequestAttachmentsInnerWithDefaults

`func NewEmailSendRequestAttachmentsInnerWithDefaults() *EmailSendRequestAttachmentsInner`

NewEmailSendRequestAttachmentsInnerWithDefaults instantiates a new EmailSendRequestAttachmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *EmailSendRequestAttachmentsInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *EmailSendRequestAttachmentsInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *EmailSendRequestAttachmentsInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *EmailSendRequestAttachmentsInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetFileName

`func (o *EmailSendRequestAttachmentsInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *EmailSendRequestAttachmentsInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *EmailSendRequestAttachmentsInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *EmailSendRequestAttachmentsInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetContentBase64

`func (o *EmailSendRequestAttachmentsInner) GetContentBase64() string`

GetContentBase64 returns the ContentBase64 field if non-nil, zero value otherwise.

### GetContentBase64Ok

`func (o *EmailSendRequestAttachmentsInner) GetContentBase64Ok() (*string, bool)`

GetContentBase64Ok returns a tuple with the ContentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBase64

`func (o *EmailSendRequestAttachmentsInner) SetContentBase64(v string)`

SetContentBase64 sets ContentBase64 field to given value.

### HasContentBase64

`func (o *EmailSendRequestAttachmentsInner) HasContentBase64() bool`

HasContentBase64 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


