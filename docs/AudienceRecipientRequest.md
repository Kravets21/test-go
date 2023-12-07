# AudienceRecipientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **NullableString** | An external identifier of the Recipient. This is optional and can be used to link the Recipient to an external system. | [optional] 
**Timezone** | Pointer to **NullableString** | The timezone of the Recipient. This is optional and can be used to set the timezone of the Recipient. | [optional] 
**Language** | Pointer to [**LocaleCode**](LocaleCode.md) |  | [optional] 
**LastLoginAt** | Pointer to **NullableTime** | Datetime of last login of the Recipient. | [optional] 
**LoginsCount** | Pointer to **NullableInt32** | The number of logins of the Recipient. | [optional] 
**CustomFields** | Pointer to **map[string]string** | A json object of the custom data of the Recipient. | [optional] 

## Methods

### NewAudienceRecipientRequest

`func NewAudienceRecipientRequest() *AudienceRecipientRequest`

NewAudienceRecipientRequest instantiates a new AudienceRecipientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceRecipientRequestWithDefaults

`func NewAudienceRecipientRequestWithDefaults() *AudienceRecipientRequest`

NewAudienceRecipientRequestWithDefaults instantiates a new AudienceRecipientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *AudienceRecipientRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AudienceRecipientRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AudienceRecipientRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AudienceRecipientRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AudienceRecipientRequest) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AudienceRecipientRequest) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetTimezone

`func (o *AudienceRecipientRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AudienceRecipientRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AudienceRecipientRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AudienceRecipientRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *AudienceRecipientRequest) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *AudienceRecipientRequest) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetLanguage

`func (o *AudienceRecipientRequest) GetLanguage() LocaleCode`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AudienceRecipientRequest) GetLanguageOk() (*LocaleCode, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AudienceRecipientRequest) SetLanguage(v LocaleCode)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AudienceRecipientRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *AudienceRecipientRequest) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *AudienceRecipientRequest) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *AudienceRecipientRequest) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *AudienceRecipientRequest) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### SetLastLoginAtNil

`func (o *AudienceRecipientRequest) SetLastLoginAtNil(b bool)`

 SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil

### UnsetLastLoginAt
`func (o *AudienceRecipientRequest) UnsetLastLoginAt()`

UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
### GetLoginsCount

`func (o *AudienceRecipientRequest) GetLoginsCount() int32`

GetLoginsCount returns the LoginsCount field if non-nil, zero value otherwise.

### GetLoginsCountOk

`func (o *AudienceRecipientRequest) GetLoginsCountOk() (*int32, bool)`

GetLoginsCountOk returns a tuple with the LoginsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginsCount

`func (o *AudienceRecipientRequest) SetLoginsCount(v int32)`

SetLoginsCount sets LoginsCount field to given value.

### HasLoginsCount

`func (o *AudienceRecipientRequest) HasLoginsCount() bool`

HasLoginsCount returns a boolean if a field has been set.

### SetLoginsCountNil

`func (o *AudienceRecipientRequest) SetLoginsCountNil(b bool)`

 SetLoginsCountNil sets the value for LoginsCount to be an explicit nil

### UnsetLoginsCount
`func (o *AudienceRecipientRequest) UnsetLoginsCount()`

UnsetLoginsCount ensures that no value is present for LoginsCount, not even an explicit nil
### GetCustomFields

`func (o *AudienceRecipientRequest) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AudienceRecipientRequest) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AudienceRecipientRequest) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AudienceRecipientRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


