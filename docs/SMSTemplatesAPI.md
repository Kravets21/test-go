# \SMSTemplatesAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmsTemplate**](SMSTemplatesAPI.md#CreateSmsTemplate) | **Post** /api/v1/sms/templates | Create a New SMS Template
[**DeleteSmsTemplate**](SMSTemplatesAPI.md#DeleteSmsTemplate) | **Delete** /api/v1/sms/templates/{template_id} | Delete a SMS Template
[**DeleteSmsTemplateContent**](SMSTemplatesAPI.md#DeleteSmsTemplateContent) | **Delete** /api/v1/sms/templates/{template_id}/contents/{template_content_locale} | Delete a Content from the SMS Template
[**GetSmsTemplate**](SMSTemplatesAPI.md#GetSmsTemplate) | **Get** /api/v1/sms/templates/{template_id} | Retrieve an Existing SMS Template
[**GetSmsTemplateContent**](SMSTemplatesAPI.md#GetSmsTemplateContent) | **Get** /api/v1/sms/templates/{template_id}/contents/{template_content_locale} | Retrieve an Existing Content from the SMS Template
[**ListSmsTemplateContents**](SMSTemplatesAPI.md#ListSmsTemplateContents) | **Get** /api/v1/sms/templates/{template_id}/contents | List All Contents from the SMS Template
[**ListSmsTemplates**](SMSTemplatesAPI.md#ListSmsTemplates) | **Get** /api/v1/sms/templates | List All SMS Templates
[**SaveSmsTemplateContent**](SMSTemplatesAPI.md#SaveSmsTemplateContent) | **Post** /api/v1/sms/templates/{template_id}/contents/{template_content_locale} | Save a Content in the SMS Template
[**UpdateSmsTemplate**](SMSTemplatesAPI.md#UpdateSmsTemplate) | **Put** /api/v1/sms/templates/{template_id} | Update a SMS Template
[**UpdateSmsTemplateContent**](SMSTemplatesAPI.md#UpdateSmsTemplateContent) | **Put** /api/v1/sms/templates/{template_id}/contents/{template_content_locale} | Update a Content in the SMS Notification Template



## CreateSmsTemplate

> SmsTemplateResponse CreateSmsTemplate(ctx).SmsTemplateRequest(smsTemplateRequest).Execute()

Create a New SMS Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    smsTemplateRequest := *AirEml.NewSmsTemplateRequest() // SmsTemplateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSTemplatesAPI.CreateSmsTemplate(context.Background()).SmsTemplateRequest(smsTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.CreateSmsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSmsTemplate`: SmsTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `SMSTemplatesAPI.CreateSmsTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSmsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsTemplateRequest** | [**SmsTemplateRequest**](SmsTemplateRequest.md) |  | 

### Return type

[**SmsTemplateResponse**](SmsTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSmsTemplate

> DeleteSmsTemplate(ctx, templateId).Execute()

Delete a SMS Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Template

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.SMSTemplatesAPI.DeleteSmsTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.DeleteSmsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSmsTemplateContent

> DeleteSmsTemplateContent(ctx, templateId, templateContentLocale).Execute()

Delete a Content from the SMS Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Template
    templateContentLocale := AirEml.LocaleCode("en") // LocaleCode | The locale of the Template Content

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.SMSTemplatesAPI.DeleteSmsTemplateContent(context.Background(), templateId, templateContentLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.DeleteSmsTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmsTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmsTemplate

> SmsTemplateResponse GetSmsTemplate(ctx, templateId).Execute()

Retrieve an Existing SMS Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Template

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSTemplatesAPI.GetSmsTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.GetSmsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmsTemplate`: SmsTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `SMSTemplatesAPI.GetSmsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmsTemplateResponse**](SmsTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmsTemplateContent

> SmsTemplateContentResponse GetSmsTemplateContent(ctx, templateId, templateContentLocale).Execute()

Retrieve an Existing Content from the SMS Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Template
    templateContentLocale := AirEml.LocaleCode("en") // LocaleCode | The locale of the Template Content

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSTemplatesAPI.GetSmsTemplateContent(context.Background(), templateId, templateContentLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.GetSmsTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmsTemplateContent`: SmsTemplateContentResponse
    fmt.Fprintf(os.Stdout, "Response from `SMSTemplatesAPI.GetSmsTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SmsTemplateContentResponse**](SmsTemplateContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmsTemplateContents

> SmsTemplateContentCollectionResponse ListSmsTemplateContents(ctx, templateId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()

List All Contents from the SMS Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Template
    cursor := "eyJjcmVhdGVkX2F0IjoiMjAyMy0wMS0xOCAxMjo1ODoyMyIsImlkIjoiYWRhM2NkMGItODE2YS0zMDc2LWEyOGUtNzYyMjNkNTRlNDMyIiwiX3BvaW50c1RvTmV4dEl0ZW1zIjpmYWxzZX0" // string | Which 'cursor' of paginated results to return. (optional)
    perPage := int32(2) // int32 | Number of items returned per page (optional) (default to 20)
    orderBy := "latest" // string | The field to order the results by (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSTemplatesAPI.ListSmsTemplateContents(context.Background(), templateId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.ListSmsTemplateContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmsTemplateContents`: SmsTemplateContentCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `SMSTemplatesAPI.ListSmsTemplateContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSmsTemplateContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 

### Return type

[**SmsTemplateContentCollectionResponse**](SmsTemplateContentCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmsTemplates

> SmsTemplateCollectionResponse ListSmsTemplates(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Status(status).Execute()

List All SMS Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    cursor := "eyJjcmVhdGVkX2F0IjoiMjAyMy0wMS0xOCAxMjo1ODoyMyIsImlkIjoiYWRhM2NkMGItODE2YS0zMDc2LWEyOGUtNzYyMjNkNTRlNDMyIiwiX3BvaW50c1RvTmV4dEl0ZW1zIjpmYWxzZX0" // string | Which 'cursor' of paginated results to return. (optional)
    perPage := int32(2) // int32 | Number of items returned per page (optional) (default to 20)
    orderBy := "latest" // string | The field to order the results by (optional)
    name := "welcome_to_documentation" // string | The name of the Template (optional)
    status := int32(1) // int32 | The status of the Template (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSTemplatesAPI.ListSmsTemplates(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.ListSmsTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmsTemplates`: SmsTemplateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `SMSTemplatesAPI.ListSmsTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSmsTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **name** | **string** | The name of the Template | 
 **status** | **int32** | The status of the Template | 

### Return type

[**SmsTemplateCollectionResponse**](SmsTemplateCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveSmsTemplateContent

> SmsTemplateContentResponse SaveSmsTemplateContent(ctx, templateId, templateContentLocale).SmsTemplateContentRequest(smsTemplateContentRequest).Execute()

Save a Content in the SMS Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Template
    templateContentLocale := AirEml.LocaleCode("en") // LocaleCode | The locale of the Template Content
    smsTemplateContentRequest := *AirEml.NewSmsTemplateContentRequest() // SmsTemplateContentRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSTemplatesAPI.SaveSmsTemplateContent(context.Background(), templateId, templateContentLocale).SmsTemplateContentRequest(smsTemplateContentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.SaveSmsTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveSmsTemplateContent`: SmsTemplateContentResponse
    fmt.Fprintf(os.Stdout, "Response from `SMSTemplatesAPI.SaveSmsTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveSmsTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smsTemplateContentRequest** | [**SmsTemplateContentRequest**](SmsTemplateContentRequest.md) |  | 

### Return type

[**SmsTemplateContentResponse**](SmsTemplateContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmsTemplate

> SmsTemplateResponse UpdateSmsTemplate(ctx, templateId).SmsTemplateRequest(smsTemplateRequest).Execute()

Update a SMS Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Template
    smsTemplateRequest := *AirEml.NewSmsTemplateRequest() // SmsTemplateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSTemplatesAPI.UpdateSmsTemplate(context.Background(), templateId).SmsTemplateRequest(smsTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.UpdateSmsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmsTemplate`: SmsTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `SMSTemplatesAPI.UpdateSmsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsTemplateRequest** | [**SmsTemplateRequest**](SmsTemplateRequest.md) |  | 

### Return type

[**SmsTemplateResponse**](SmsTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmsTemplateContent

> SmsTemplateContentResponse UpdateSmsTemplateContent(ctx, templateId, templateContentLocale).SmsTemplateContentRequest(smsTemplateContentRequest).Execute()

Update a Content in the SMS Notification Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    AirEml "github.com/pdffiller/eml-go-api"
)

func main() {
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Template
    templateContentLocale := AirEml.LocaleCode("en") // LocaleCode | The locale of the Template Content
    smsTemplateContentRequest := *AirEml.NewSmsTemplateContentRequest() // SmsTemplateContentRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSTemplatesAPI.UpdateSmsTemplateContent(context.Background(), templateId, templateContentLocale).SmsTemplateContentRequest(smsTemplateContentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSTemplatesAPI.UpdateSmsTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmsTemplateContent`: SmsTemplateContentResponse
    fmt.Fprintf(os.Stdout, "Response from `SMSTemplatesAPI.UpdateSmsTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmsTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smsTemplateContentRequest** | [**SmsTemplateContentRequest**](SmsTemplateContentRequest.md) |  | 

### Return type

[**SmsTemplateContentResponse**](SmsTemplateContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

