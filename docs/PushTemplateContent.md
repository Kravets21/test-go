# PushTemplateContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to [**LocaleCode**](LocaleCode.md) |  | [optional] 
**Title** | Pointer to **string** | The is a title of the push notification. | [optional] 
**Body** | Pointer to **string** | The is a body of the push notification. | [optional] 
**Icon** | Pointer to **NullableString** | The notification&#39;s icon. Sets the notification icon to myicon for drawable resource myicon. If you don&#39;t send this key in the request, FCM displays the launcher icon specified in your app manifest. | [optional] 
**Image** | Pointer to **NullableString** | Contains the URL of an image that is going to be downloaded on the device and displayed in a notification. JPEG, PNG, BMP have full support across platforms. Animated GIF and video only work on iOS. WebP and HEIF have varying levels of support across platforms and platform versions. Android has 1MB image size limit.  | [optional] 
**Link** | Pointer to **NullableString** | The link to open when the user clicks on the notification. For all URL values, HTTPS is required. | [optional] 
**VariablesSettings** | Pointer to [**[]EmailTemplateContentVariablesSettingsInner**](EmailTemplateContentVariablesSettingsInner.md) |  | [optional] 
**Data** | Pointer to **map[string]string** | The additional data used for send push notification if user wants to add some custom data in request | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Template Content was created. | [optional] 

## Methods

### NewPushTemplateContent

`func NewPushTemplateContent() *PushTemplateContent`

NewPushTemplateContent instantiates a new PushTemplateContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushTemplateContentWithDefaults

`func NewPushTemplateContentWithDefaults() *PushTemplateContent`

NewPushTemplateContentWithDefaults instantiates a new PushTemplateContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *PushTemplateContent) GetLocale() LocaleCode`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PushTemplateContent) GetLocaleOk() (*LocaleCode, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PushTemplateContent) SetLocale(v LocaleCode)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PushTemplateContent) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetTitle

`func (o *PushTemplateContent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PushTemplateContent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PushTemplateContent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PushTemplateContent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *PushTemplateContent) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PushTemplateContent) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PushTemplateContent) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PushTemplateContent) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetIcon

`func (o *PushTemplateContent) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PushTemplateContent) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PushTemplateContent) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PushTemplateContent) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *PushTemplateContent) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *PushTemplateContent) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetImage

`func (o *PushTemplateContent) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PushTemplateContent) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PushTemplateContent) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PushTemplateContent) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *PushTemplateContent) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *PushTemplateContent) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetLink

`func (o *PushTemplateContent) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PushTemplateContent) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PushTemplateContent) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PushTemplateContent) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *PushTemplateContent) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *PushTemplateContent) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetVariablesSettings

`func (o *PushTemplateContent) GetVariablesSettings() []EmailTemplateContentVariablesSettingsInner`

GetVariablesSettings returns the VariablesSettings field if non-nil, zero value otherwise.

### GetVariablesSettingsOk

`func (o *PushTemplateContent) GetVariablesSettingsOk() (*[]EmailTemplateContentVariablesSettingsInner, bool)`

GetVariablesSettingsOk returns a tuple with the VariablesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesSettings

`func (o *PushTemplateContent) SetVariablesSettings(v []EmailTemplateContentVariablesSettingsInner)`

SetVariablesSettings sets VariablesSettings field to given value.

### HasVariablesSettings

`func (o *PushTemplateContent) HasVariablesSettings() bool`

HasVariablesSettings returns a boolean if a field has been set.

### GetData

`func (o *PushTemplateContent) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PushTemplateContent) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PushTemplateContent) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *PushTemplateContent) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PushTemplateContent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PushTemplateContent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PushTemplateContent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PushTemplateContent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


