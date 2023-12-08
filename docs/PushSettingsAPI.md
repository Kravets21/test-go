# \PushSettingsAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePushPlatform**](PushSettingsAPI.md#CreatePushPlatform) | **Post** /api/v1/push/platforms | Create a New Push Notification Platform
[**ListPushPlatforms**](PushSettingsAPI.md#ListPushPlatforms) | **Get** /api/v1/push/platforms | List All Push Notification Platforms



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

