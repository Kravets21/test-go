# \EmailTemplatesAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailTemplate**](EmailTemplatesAPI.md#CreateEmailTemplate) | **Post** /api/v1/email/templates | Create a New Email Notification Template
[**CreateEmailTemplateContent**](EmailTemplatesAPI.md#CreateEmailTemplateContent) | **Post** /api/v1/email/templates/{template_id}/contents/{template_content_locale} | Create a Content in the Email Notification Template
[**DeleteEmailTemplate**](EmailTemplatesAPI.md#DeleteEmailTemplate) | **Delete** /api/v1/email/templates/{template_id} | Delete a Email Notification Template
[**DeleteEmailTemplateContent**](EmailTemplatesAPI.md#DeleteEmailTemplateContent) | **Delete** /api/v1/email/templates/{template_id}/contents/{template_content_locale} | Delete a Content from the Email Notification Template
[**GetEmailTemplate**](EmailTemplatesAPI.md#GetEmailTemplate) | **Get** /api/v1/email/templates/{template_id} | Retrieve an Existing Email Notification Template
[**GetEmailTemplateContent**](EmailTemplatesAPI.md#GetEmailTemplateContent) | **Get** /api/v1/email/templates/{template_id}/contents/{template_content_locale} | Retrieve an Existing Content from the Email Notification Template
[**ListEmailTemplateContents**](EmailTemplatesAPI.md#ListEmailTemplateContents) | **Get** /api/v1/email/templates/{template_id}/contents | List All Contents from the Email Notification Template
[**ListEmailTemplates**](EmailTemplatesAPI.md#ListEmailTemplates) | **Get** /api/v1/email/templates | List All Email Notification Templates
[**UpdateEmailTemplate**](EmailTemplatesAPI.md#UpdateEmailTemplate) | **Put** /api/v1/email/templates/{template_id} | Update a Email Notification Template
[**UpdateEmailTemplateContent**](EmailTemplatesAPI.md#UpdateEmailTemplateContent) | **Put** /api/v1/email/templates/{template_id}/contents/{template_content_locale} | Update a Content in the Email Notification Template



## CreateEmailTemplate

> EmailTemplateResponse CreateEmailTemplate(ctx).EmailTemplateCreateRequest(emailTemplateCreateRequest).Execute()

Create a New Email Notification Template



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
    emailTemplateCreateRequest := *AirEml.NewEmailTemplateCreateRequest("00112233-4455-6677-8899-aabbccddeeff", "welcome_to_documentation") // EmailTemplateCreateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailTemplatesAPI.CreateEmailTemplate(context.Background()).EmailTemplateCreateRequest(emailTemplateCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.CreateEmailTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailTemplate`: EmailTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.CreateEmailTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailTemplateCreateRequest** | [**EmailTemplateCreateRequest**](EmailTemplateCreateRequest.md) |  | 

### Return type

[**EmailTemplateResponse**](EmailTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmailTemplateContent

> EmailTemplateContentResponse CreateEmailTemplateContent(ctx, templateId, templateContentLocale).EmailTemplateContentCreateRequest(emailTemplateContentCreateRequest).Execute()

Create a Content in the Email Notification Template



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
    emailTemplateContentCreateRequest := *AirEml.NewEmailTemplateContentCreateRequest("Welcome to documentation {{ name }}.", "<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
  <html lang="en">
  <head>
    <meta http-equiv="Content-Type" content="text/html" charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes, maximum-scale=5">
    <title>Welcome</title>
  </head>
  <body>
  	<p>Welcome {{ name }}. This is welcome email. It is sent to new users.</p>
      <p><img src={{ link }}></p>
  </body>
  </html>") // EmailTemplateContentCreateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailTemplatesAPI.CreateEmailTemplateContent(context.Background(), templateId, templateContentLocale).EmailTemplateContentCreateRequest(emailTemplateContentCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.CreateEmailTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailTemplateContent`: EmailTemplateContentResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.CreateEmailTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailTemplateContentCreateRequest** | [**EmailTemplateContentCreateRequest**](EmailTemplateContentCreateRequest.md) |  | 

### Return type

[**EmailTemplateContentResponse**](EmailTemplateContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailTemplate

> DeleteEmailTemplate(ctx, templateId).Execute()

Delete a Email Notification Template



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
    r, err := apiClient.EmailTemplatesAPI.DeleteEmailTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.DeleteEmailTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEmailTemplateRequest struct via the builder pattern


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


## DeleteEmailTemplateContent

> DeleteEmailTemplateContent(ctx, templateId, templateContentLocale).Execute()

Delete a Content from the Email Notification Template



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
    r, err := apiClient.EmailTemplatesAPI.DeleteEmailTemplateContent(context.Background(), templateId, templateContentLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.DeleteEmailTemplateContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEmailTemplateContentRequest struct via the builder pattern


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


## GetEmailTemplate

> EmailTemplateResponse GetEmailTemplate(ctx, templateId).Execute()

Retrieve an Existing Email Notification Template



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
    resp, r, err := apiClient.EmailTemplatesAPI.GetEmailTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.GetEmailTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailTemplate`: EmailTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.GetEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailTemplateResponse**](EmailTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTemplateContent

> EmailTemplateContentResponse GetEmailTemplateContent(ctx, templateId, templateContentLocale).Execute()

Retrieve an Existing Content from the Email Notification Template



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
    resp, r, err := apiClient.EmailTemplatesAPI.GetEmailTemplateContent(context.Background(), templateId, templateContentLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.GetEmailTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailTemplateContent`: EmailTemplateContentResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.GetEmailTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailTemplateContentResponse**](EmailTemplateContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailTemplateContents

> EmailTemplateContentCollectionResponse ListEmailTemplateContents(ctx, templateId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()

List All Contents from the Email Notification Template



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
    resp, r, err := apiClient.EmailTemplatesAPI.ListEmailTemplateContents(context.Background(), templateId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.ListEmailTemplateContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailTemplateContents`: EmailTemplateContentCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.ListEmailTemplateContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailTemplateContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 

### Return type

[**EmailTemplateContentCollectionResponse**](EmailTemplateContentCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailTemplates

> EmailTemplateCollectionResponse ListEmailTemplates(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Status(status).Execute()

List All Email Notification Templates



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
    name := "welcome_to_documentation" // string | The name of the Template. (optional)
    status := "active" // string | The status of the Template. (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailTemplatesAPI.ListEmailTemplates(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.ListEmailTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailTemplates`: EmailTemplateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.ListEmailTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **name** | **string** | The name of the Template. | 
 **status** | **string** | The status of the Template. | 

### Return type

[**EmailTemplateCollectionResponse**](EmailTemplateCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailTemplate

> EmailTemplateResponse UpdateEmailTemplate(ctx, templateId).EmailTemplateUpdateRequest(emailTemplateUpdateRequest).Execute()

Update a Email Notification Template



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
    emailTemplateUpdateRequest := *AirEml.NewEmailTemplateUpdateRequest() // EmailTemplateUpdateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailTemplatesAPI.UpdateEmailTemplate(context.Background(), templateId).EmailTemplateUpdateRequest(emailTemplateUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.UpdateEmailTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailTemplate`: EmailTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.UpdateEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailTemplateUpdateRequest** | [**EmailTemplateUpdateRequest**](EmailTemplateUpdateRequest.md) |  | 

### Return type

[**EmailTemplateResponse**](EmailTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailTemplateContent

> EmailTemplateContentResponse UpdateEmailTemplateContent(ctx, templateId, templateContentLocale).EmailTemplateContentUpdateRequest(emailTemplateContentUpdateRequest).Execute()

Update a Content in the Email Notification Template



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
    emailTemplateContentUpdateRequest := *AirEml.NewEmailTemplateContentUpdateRequest() // EmailTemplateContentUpdateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailTemplatesAPI.UpdateEmailTemplateContent(context.Background(), templateId, templateContentLocale).EmailTemplateContentUpdateRequest(emailTemplateContentUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.UpdateEmailTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailTemplateContent`: EmailTemplateContentResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.UpdateEmailTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailTemplateContentUpdateRequest** | [**EmailTemplateContentUpdateRequest**](EmailTemplateContentUpdateRequest.md) |  | 

### Return type

[**EmailTemplateContentResponse**](EmailTemplateContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

