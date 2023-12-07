# PushTemplateContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The is a title of the push notification. | 
**Body** | **string** | The is a body of the push notification. | 
**Icon** | Pointer to **NullableString** | The notification&#39;s icon. Sets the notification icon to myicon for drawable resource myicon. If you don&#39;t send this key in the request, FCM displays the launcher icon specified in your app manifest. | [optional] 
**Image** | Pointer to **NullableString** | Contains the URL of an image that is going to be downloaded on the device and displayed in a notification. JPEG, PNG, BMP have full support across platforms. Animated GIF and video only work on iOS. WebP and HEIF have varying levels of support across platforms and platform versions. Android has 1MB image size limit.  | [optional] 
**Link** | Pointer to **NullableString** | The link to open when the user clicks on the notification. For all URL values, HTTPS is required. | [optional] 
**Data** | Pointer to **map[string]string** | The additional data used for send push notification if user wants to add some custom data in request | [optional] 

## Methods

### NewPushTemplateContentRequest

`func NewPushTemplateContentRequest(title string, body string, ) *PushTemplateContentRequest`

NewPushTemplateContentRequest instantiates a new PushTemplateContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushTemplateContentRequestWithDefaults

`func NewPushTemplateContentRequestWithDefaults() *PushTemplateContentRequest`

NewPushTemplateContentRequestWithDefaults instantiates a new PushTemplateContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PushTemplateContentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PushTemplateContentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PushTemplateContentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *PushTemplateContentRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PushTemplateContentRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PushTemplateContentRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetIcon

`func (o *PushTemplateContentRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PushTemplateContentRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PushTemplateContentRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PushTemplateContentRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *PushTemplateContentRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *PushTemplateContentRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetImage

`func (o *PushTemplateContentRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PushTemplateContentRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PushTemplateContentRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PushTemplateContentRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *PushTemplateContentRequest) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *PushTemplateContentRequest) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetLink

`func (o *PushTemplateContentRequest) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PushTemplateContentRequest) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PushTemplateContentRequest) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PushTemplateContentRequest) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *PushTemplateContentRequest) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *PushTemplateContentRequest) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetData

`func (o *PushTemplateContentRequest) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PushTemplateContentRequest) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PushTemplateContentRequest) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *PushTemplateContentRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


