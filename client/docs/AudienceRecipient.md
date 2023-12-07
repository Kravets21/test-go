# AudienceRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Recipient. | [optional] 
**ExternalId** | Pointer to **NullableString** | An external identifier of the Recipient. This is optional and can be used to link the Recipient to an external system. | [optional] 
**Timezone** | Pointer to **NullableString** | The timezone of the Recipient. This is optional and can be used to set the timezone of the Recipient. | [optional] 
**Language** | Pointer to [**LocaleCode**](LocaleCode.md) |  | [optional] 
**LastActivityAt** | Pointer to **NullableTime** | Datetime of last activity of the Recipient. | [optional] 
**LastActivityType** | Pointer to **NullableString** | The last activity type of the Recipient. | [optional] 
**LastActivityEvent** | Pointer to **NullableString** | The last activity event of the Recipient. | [optional] 
**LastLoginAt** | Pointer to **NullableTime** | Datetime of last login of the Recipient. | [optional] 
**LoginsCount** | Pointer to **NullableInt32** | The number of logins of the Recipient. | [optional] 
**CustomFields** | Pointer to **map[string]string** | A json object of the custom data of the Recipient. | [optional] 
**Contacts** | Pointer to [**[]AudienceRecipientContact**](AudienceRecipientContact.md) | The contacts of the Recipient. | [optional] 
**Tags** | Pointer to [**[]AudienceRecipientTag**](AudienceRecipientTag.md) | The tags of the Recipient. | [optional] 
**Status** | Pointer to **int32** | The status of the Recipient. Every digit matters: 1 - active, 2 - inactive. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime when the Recipient was created. | [optional] 

## Methods

### NewAudienceRecipient

`func NewAudienceRecipient() *AudienceRecipient`

NewAudienceRecipient instantiates a new AudienceRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceRecipientWithDefaults

`func NewAudienceRecipientWithDefaults() *AudienceRecipient`

NewAudienceRecipientWithDefaults instantiates a new AudienceRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudienceRecipient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceRecipient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudienceRecipient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AudienceRecipient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *AudienceRecipient) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AudienceRecipient) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AudienceRecipient) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AudienceRecipient) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AudienceRecipient) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AudienceRecipient) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetTimezone

`func (o *AudienceRecipient) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AudienceRecipient) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AudienceRecipient) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AudienceRecipient) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *AudienceRecipient) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *AudienceRecipient) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetLanguage

`func (o *AudienceRecipient) GetLanguage() LocaleCode`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AudienceRecipient) GetLanguageOk() (*LocaleCode, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AudienceRecipient) SetLanguage(v LocaleCode)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AudienceRecipient) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *AudienceRecipient) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *AudienceRecipient) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *AudienceRecipient) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *AudienceRecipient) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### SetLastActivityAtNil

`func (o *AudienceRecipient) SetLastActivityAtNil(b bool)`

 SetLastActivityAtNil sets the value for LastActivityAt to be an explicit nil

### UnsetLastActivityAt
`func (o *AudienceRecipient) UnsetLastActivityAt()`

UnsetLastActivityAt ensures that no value is present for LastActivityAt, not even an explicit nil
### GetLastActivityType

`func (o *AudienceRecipient) GetLastActivityType() string`

GetLastActivityType returns the LastActivityType field if non-nil, zero value otherwise.

### GetLastActivityTypeOk

`func (o *AudienceRecipient) GetLastActivityTypeOk() (*string, bool)`

GetLastActivityTypeOk returns a tuple with the LastActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityType

`func (o *AudienceRecipient) SetLastActivityType(v string)`

SetLastActivityType sets LastActivityType field to given value.

### HasLastActivityType

`func (o *AudienceRecipient) HasLastActivityType() bool`

HasLastActivityType returns a boolean if a field has been set.

### SetLastActivityTypeNil

`func (o *AudienceRecipient) SetLastActivityTypeNil(b bool)`

 SetLastActivityTypeNil sets the value for LastActivityType to be an explicit nil

### UnsetLastActivityType
`func (o *AudienceRecipient) UnsetLastActivityType()`

UnsetLastActivityType ensures that no value is present for LastActivityType, not even an explicit nil
### GetLastActivityEvent

`func (o *AudienceRecipient) GetLastActivityEvent() string`

GetLastActivityEvent returns the LastActivityEvent field if non-nil, zero value otherwise.

### GetLastActivityEventOk

`func (o *AudienceRecipient) GetLastActivityEventOk() (*string, bool)`

GetLastActivityEventOk returns a tuple with the LastActivityEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityEvent

`func (o *AudienceRecipient) SetLastActivityEvent(v string)`

SetLastActivityEvent sets LastActivityEvent field to given value.

### HasLastActivityEvent

`func (o *AudienceRecipient) HasLastActivityEvent() bool`

HasLastActivityEvent returns a boolean if a field has been set.

### SetLastActivityEventNil

`func (o *AudienceRecipient) SetLastActivityEventNil(b bool)`

 SetLastActivityEventNil sets the value for LastActivityEvent to be an explicit nil

### UnsetLastActivityEvent
`func (o *AudienceRecipient) UnsetLastActivityEvent()`

UnsetLastActivityEvent ensures that no value is present for LastActivityEvent, not even an explicit nil
### GetLastLoginAt

`func (o *AudienceRecipient) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *AudienceRecipient) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *AudienceRecipient) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *AudienceRecipient) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### SetLastLoginAtNil

`func (o *AudienceRecipient) SetLastLoginAtNil(b bool)`

 SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil

### UnsetLastLoginAt
`func (o *AudienceRecipient) UnsetLastLoginAt()`

UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
### GetLoginsCount

`func (o *AudienceRecipient) GetLoginsCount() int32`

GetLoginsCount returns the LoginsCount field if non-nil, zero value otherwise.

### GetLoginsCountOk

`func (o *AudienceRecipient) GetLoginsCountOk() (*int32, bool)`

GetLoginsCountOk returns a tuple with the LoginsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginsCount

`func (o *AudienceRecipient) SetLoginsCount(v int32)`

SetLoginsCount sets LoginsCount field to given value.

### HasLoginsCount

`func (o *AudienceRecipient) HasLoginsCount() bool`

HasLoginsCount returns a boolean if a field has been set.

### SetLoginsCountNil

`func (o *AudienceRecipient) SetLoginsCountNil(b bool)`

 SetLoginsCountNil sets the value for LoginsCount to be an explicit nil

### UnsetLoginsCount
`func (o *AudienceRecipient) UnsetLoginsCount()`

UnsetLoginsCount ensures that no value is present for LoginsCount, not even an explicit nil
### GetCustomFields

`func (o *AudienceRecipient) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AudienceRecipient) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AudienceRecipient) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AudienceRecipient) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetContacts

`func (o *AudienceRecipient) GetContacts() []AudienceRecipientContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *AudienceRecipient) GetContactsOk() (*[]AudienceRecipientContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *AudienceRecipient) SetContacts(v []AudienceRecipientContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *AudienceRecipient) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetTags

`func (o *AudienceRecipient) GetTags() []AudienceRecipientTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AudienceRecipient) GetTagsOk() (*[]AudienceRecipientTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AudienceRecipient) SetTags(v []AudienceRecipientTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AudienceRecipient) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatus

`func (o *AudienceRecipient) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AudienceRecipient) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AudienceRecipient) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AudienceRecipient) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AudienceRecipient) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AudienceRecipient) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AudienceRecipient) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AudienceRecipient) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


