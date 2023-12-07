# \PushTemplatesAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePushTemplate**](PushTemplatesAPI.md#CreatePushTemplate) | **Post** /api/v1/push/templates | Create a New Push Notification Template
[**CreatePushTemplateContent**](PushTemplatesAPI.md#CreatePushTemplateContent) | **Post** /api/v1/push/templates/{template_id}/contents/{template_content_locale} | Create a Content in the Push Notification Template
[**DeletePushTemplate**](PushTemplatesAPI.md#DeletePushTemplate) | **Delete** /api/v1/push/templates/{template_id} | Delete a Push Notification Template
[**DeletePushTemplateContent**](PushTemplatesAPI.md#DeletePushTemplateContent) | **Delete** /api/v1/push/templates/{template_id}/contents/{template_content_locale} | Delete a Content from the Push Notification Template
[**GetPushTemplate**](PushTemplatesAPI.md#GetPushTemplate) | **Get** /api/v1/push/templates/{template_id} | Retrieve an Existing Push Notification Template
[**GetPushTemplateContent**](PushTemplatesAPI.md#GetPushTemplateContent) | **Get** /api/v1/push/templates/{template_id}/contents/{template_content_locale} | Retrieve an Existing Content from the Push Notification Template
[**ListPushTemplateContents**](PushTemplatesAPI.md#ListPushTemplateContents) | **Get** /api/v1/push/templates/{template_id}/contents | List All Contents from the Push Notification Template
[**ListPushTemplates**](PushTemplatesAPI.md#ListPushTemplates) | **Get** /api/v1/push/templates | List All Push Notification Templates
[**UpdatePushTemplate**](PushTemplatesAPI.md#UpdatePushTemplate) | **Put** /api/v1/push/templates/{template_id} | Update a Push Notification Template
[**UpdatePushTemplateContent**](PushTemplatesAPI.md#UpdatePushTemplateContent) | **Put** /api/v1/push/templates/{template_id}/contents/{template_content_locale} | Update a Content in the Push Notification Template



## CreatePushTemplate

> PushTemplateResponse CreatePushTemplate(ctx).PushTemplateRequest(pushTemplateRequest).Execute()

Create a New Push Notification Template



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
    pushTemplateRequest := *AirEml.NewPushTemplateRequest("welcome_to_documentation") // PushTemplateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.PushTemplatesAPI.CreatePushTemplate(context.Background()).PushTemplateRequest(pushTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.CreatePushTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePushTemplate`: PushTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `PushTemplatesAPI.CreatePushTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushTemplateRequest** | [**PushTemplateRequest**](PushTemplateRequest.md) |  | 

### Return type

[**PushTemplateResponse**](PushTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePushTemplateContent

> PushTemplateContentResponse CreatePushTemplateContent(ctx, templateId, templateContentLocale).PushTemplateContentRequest(pushTemplateContentRequest).Execute()

Create a Content in the Push Notification Template



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
    pushTemplateContentRequest := *AirEml.NewPushTemplateContentRequest("Welcome to documentation", "This is a body of the push notification. It can be a long text with multiple lines.") // PushTemplateContentRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.PushTemplatesAPI.CreatePushTemplateContent(context.Background(), templateId, templateContentLocale).PushTemplateContentRequest(pushTemplateContentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.CreatePushTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePushTemplateContent`: PushTemplateContentResponse
    fmt.Fprintf(os.Stdout, "Response from `PushTemplatesAPI.CreatePushTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pushTemplateContentRequest** | [**PushTemplateContentRequest**](PushTemplateContentRequest.md) |  | 

### Return type

[**PushTemplateContentResponse**](PushTemplateContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePushTemplate

> DeletePushTemplate(ctx, templateId).Execute()

Delete a Push Notification Template



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
    r, err := apiClient.PushTemplatesAPI.DeletePushTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.DeletePushTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePushTemplateRequest struct via the builder pattern


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


## DeletePushTemplateContent

> DeletePushTemplateContent(ctx, templateId, templateContentLocale).Execute()

Delete a Content from the Push Notification Template



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
    r, err := apiClient.PushTemplatesAPI.DeletePushTemplateContent(context.Background(), templateId, templateContentLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.DeletePushTemplateContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePushTemplateContentRequest struct via the builder pattern


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


## GetPushTemplate

> PushTemplateResponse GetPushTemplate(ctx, templateId).Execute()

Retrieve an Existing Push Notification Template



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
    resp, r, err := apiClient.PushTemplatesAPI.GetPushTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.GetPushTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushTemplate`: PushTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `PushTemplatesAPI.GetPushTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PushTemplateResponse**](PushTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushTemplateContent

> PushTemplateContentResponse GetPushTemplateContent(ctx, templateId, templateContentLocale).Execute()

Retrieve an Existing Content from the Push Notification Template



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
    resp, r, err := apiClient.PushTemplatesAPI.GetPushTemplateContent(context.Background(), templateId, templateContentLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.GetPushTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushTemplateContent`: PushTemplateContentResponse
    fmt.Fprintf(os.Stdout, "Response from `PushTemplatesAPI.GetPushTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PushTemplateContentResponse**](PushTemplateContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPushTemplateContents

> PushTemplateContentCollectionResponse ListPushTemplateContents(ctx, templateId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()

List All Contents from the Push Notification Template



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
    resp, r, err := apiClient.PushTemplatesAPI.ListPushTemplateContents(context.Background(), templateId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.ListPushTemplateContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPushTemplateContents`: PushTemplateContentCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `PushTemplatesAPI.ListPushTemplateContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPushTemplateContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 

### Return type

[**PushTemplateContentCollectionResponse**](PushTemplateContentCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPushTemplates

> PushTemplateCollectionResponse ListPushTemplates(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Status(status).Execute()

List All Push Notification Templates



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
    status := int32(1) // int32 | The status of the Template. (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.PushTemplatesAPI.ListPushTemplates(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.ListPushTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPushTemplates`: PushTemplateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `PushTemplatesAPI.ListPushTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPushTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **name** | **string** | The name of the Template. | 
 **status** | **int32** | The status of the Template. | 

### Return type

[**PushTemplateCollectionResponse**](PushTemplateCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePushTemplate

> PushTemplateResponse UpdatePushTemplate(ctx, templateId).PushTemplateUpdateRequest(pushTemplateUpdateRequest).Execute()

Update a Push Notification Template



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
    pushTemplateUpdateRequest := *AirEml.NewPushTemplateUpdateRequest() // PushTemplateUpdateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.PushTemplatesAPI.UpdatePushTemplate(context.Background(), templateId).PushTemplateUpdateRequest(pushTemplateUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.UpdatePushTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePushTemplate`: PushTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `PushTemplatesAPI.UpdatePushTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePushTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pushTemplateUpdateRequest** | [**PushTemplateUpdateRequest**](PushTemplateUpdateRequest.md) |  | 

### Return type

[**PushTemplateResponse**](PushTemplateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePushTemplateContent

> PushTemplateContentResponse UpdatePushTemplateContent(ctx, templateId, templateContentLocale).PushTemplateContentRequest(pushTemplateContentRequest).Execute()

Update a Content in the Push Notification Template



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
    pushTemplateContentRequest := *AirEml.NewPushTemplateContentRequest("Welcome to documentation", "This is a body of the push notification. It can be a long text with multiple lines.") // PushTemplateContentRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.PushTemplatesAPI.UpdatePushTemplateContent(context.Background(), templateId, templateContentLocale).PushTemplateContentRequest(pushTemplateContentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushTemplatesAPI.UpdatePushTemplateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePushTemplateContent`: PushTemplateContentResponse
    fmt.Fprintf(os.Stdout, "Response from `PushTemplatesAPI.UpdatePushTemplateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | The ID of the Template | 
**templateContentLocale** | [**LocaleCode**](.md) | The locale of the Template Content | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePushTemplateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pushTemplateContentRequest** | [**PushTemplateContentRequest**](PushTemplateContentRequest.md) |  | 

### Return type

[**PushTemplateContentResponse**](PushTemplateContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

