# \PushSettingsAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePushPlatform**](PushSettingsAPI.md#CreatePushPlatform) | **Post** /api/v1/push/platforms | Create a New Push Notification Platform
[**DeletePushPlatforms**](PushSettingsAPI.md#DeletePushPlatforms) | **Delete** /api/v1/push/platforms/{platform_id} | Delete a Push Notification Platform
[**GetPushPlatform**](PushSettingsAPI.md#GetPushPlatform) | **Get** /api/v1/push/platforms/{platform_id} | Retrieve an Existing Push Notification Platform
[**ListPushPlatforms**](PushSettingsAPI.md#ListPushPlatforms) | **Get** /api/v1/push/platforms | List All Push Notification Platforms
[**UpdatePushPlatform**](PushSettingsAPI.md#UpdatePushPlatform) | **Put** /api/v1/push/platforms/{platform_id} | Update a Push Notification Platform



## CreatePushPlatform

> PushPlatformResponse CreatePushPlatform(ctx).PushPlatformRequest(pushPlatformRequest).Execute()

Create a New Push Notification Platform



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
    pushPlatformRequest := *AirEml.NewPushPlatformRequest() // PushPlatformRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.PushSettingsAPI.CreatePushPlatform(context.Background()).PushPlatformRequest(pushPlatformRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushSettingsAPI.CreatePushPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePushPlatform`: PushPlatformResponse
    fmt.Fprintf(os.Stdout, "Response from `PushSettingsAPI.CreatePushPlatform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushPlatformRequest** | [**PushPlatformRequest**](PushPlatformRequest.md) |  | 

### Return type

[**PushPlatformResponse**](PushPlatformResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePushPlatforms

> DeletePushPlatforms(ctx, platformId).Execute()

Delete a Push Notification Platform



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
    platformId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Platform

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.PushSettingsAPI.DeletePushPlatforms(context.Background(), platformId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushSettingsAPI.DeletePushPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformId** | **string** | The ID of the Platform | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePushPlatformsRequest struct via the builder pattern


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


## GetPushPlatform

> PushPlatformResponse GetPushPlatform(ctx, platformId).Execute()

Retrieve an Existing Push Notification Platform



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
    platformId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Platform

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.PushSettingsAPI.GetPushPlatform(context.Background(), platformId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushSettingsAPI.GetPushPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushPlatform`: PushPlatformResponse
    fmt.Fprintf(os.Stdout, "Response from `PushSettingsAPI.GetPushPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformId** | **string** | The ID of the Platform | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PushPlatformResponse**](PushPlatformResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPushPlatforms

> PushPlatformCollectionResponse ListPushPlatforms(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Status(status).Execute()

List All Push Notification Platforms



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
    status := int32(1) // int32 | The filter value for show the records of the platforms that have a specific status. (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.PushSettingsAPI.ListPushPlatforms(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushSettingsAPI.ListPushPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPushPlatforms`: PushPlatformCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `PushSettingsAPI.ListPushPlatforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPushPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **status** | **int32** | The filter value for show the records of the platforms that have a specific status. | 

### Return type

[**PushPlatformCollectionResponse**](PushPlatformCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePushPlatform

> PushPlatformResponse UpdatePushPlatform(ctx, platformId).PushPlatformRequest(pushPlatformRequest).Execute()

Update a Push Notification Platform



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
    platformId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Platform
    pushPlatformRequest := *AirEml.NewPushPlatformRequest() // PushPlatformRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.PushSettingsAPI.UpdatePushPlatform(context.Background(), platformId).PushPlatformRequest(pushPlatformRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushSettingsAPI.UpdatePushPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePushPlatform`: PushPlatformResponse
    fmt.Fprintf(os.Stdout, "Response from `PushSettingsAPI.UpdatePushPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformId** | **string** | The ID of the Platform | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePushPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pushPlatformRequest** | [**PushPlatformRequest**](PushPlatformRequest.md) |  | 

### Return type

[**PushPlatformResponse**](PushPlatformResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

