# ErrorMultiple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A message providing information about the error. | 
**Errors** | **map[string][]string** |  | 

## Methods

### NewErrorMultiple

`func NewErrorMultiple(message string, errors map[string][]string, ) *ErrorMultiple`

NewErrorMultiple instantiates a new ErrorMultiple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorMultipleWithDefaults

`func NewErrorMultipleWithDefaults() *ErrorMultiple`

NewErrorMultipleWithDefaults instantiates a new ErrorMultiple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorMultiple) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorMultiple) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorMultiple) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *ErrorMultiple) GetErrors() map[string][]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorMultiple) GetErrorsOk() (*map[string][]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorMultiple) SetErrors(v map[string][]string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


