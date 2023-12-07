# \WebhooksAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookSubscription**](WebhooksAPI.md#CreateWebhookSubscription) | **Post** /api/v1/webhook/subscriptions | Create a New Subscription
[**DeleteWebhookSubscription**](WebhooksAPI.md#DeleteWebhookSubscription) | **Delete** /api/v1/webhook/subscriptions/{webhook_subscription_id} | Delete a webhook subscription
[**GetWebhookSubscription**](WebhooksAPI.md#GetWebhookSubscription) | **Get** /api/v1/webhook/subscriptions/{webhook_subscription_id} | Retrieve an Existing webhook subscription
[**ListWebhookEvents**](WebhooksAPI.md#ListWebhookEvents) | **Get** /api/v1/webhook/events | List All Webhook events
[**ListWebhookSubscriptions**](WebhooksAPI.md#ListWebhookSubscriptions) | **Get** /api/v1/webhook/subscriptions | List All webhook Subscriptions
[**UpdateWebhookSubscription**](WebhooksAPI.md#UpdateWebhookSubscription) | **Put** /api/v1/webhook/subscriptions/{webhook_subscription_id} | Update a webhook subscription



## CreateWebhookSubscription

> SubscriptionResponse CreateWebhookSubscription(ctx).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()

Create a New Subscription



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
    subscriptionCreateRequest := *AirEml.NewSubscriptionCreateRequest("Using for check sms status", "https://aireml.com/api/callback") // SubscriptionCreateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.CreateWebhookSubscription(context.Background()).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateWebhookSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhookSubscription`: SubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CreateWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionCreateRequest** | [**SubscriptionCreateRequest**](SubscriptionCreateRequest.md) |  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookSubscription

> DeleteWebhookSubscription(ctx, webhookSubscriptionId).Execute()

Delete a webhook subscription



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
    webhookSubscriptionId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Webhook subscription

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.WebhooksAPI.DeleteWebhookSubscription(context.Background(), webhookSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhookSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookSubscriptionId** | **string** | The ID of the Webhook subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookSubscriptionRequest struct via the builder pattern


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


## GetWebhookSubscription

> SubscriptionResponse GetWebhookSubscription(ctx, webhookSubscriptionId).Execute()

Retrieve an Existing webhook subscription



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
    webhookSubscriptionId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Webhook subscription

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.GetWebhookSubscription(context.Background(), webhookSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhookSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookSubscription`: SubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookSubscriptionId** | **string** | The ID of the Webhook subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookEvents

> EventResponse ListWebhookEvents(ctx).Execute()

List All Webhook events



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

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.ListWebhookEvents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhookEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookEvents`: EventResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhookEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookEventsRequest struct via the builder pattern


### Return type

[**EventResponse**](EventResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookSubscriptions

> SubscriptionCollectionResponse ListWebhookSubscriptions(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Search(search).Status(status).Execute()

List All webhook Subscriptions



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
    search := "example_subscription" // string | The name of the Subscription. (optional)
    status := int32(1) // int32 | The status of the Subscription. (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.ListWebhookSubscriptions(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Search(search).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhookSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookSubscriptions`: SubscriptionCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhookSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **search** | **string** | The name of the Subscription. | 
 **status** | **int32** | The status of the Subscription. | 

### Return type

[**SubscriptionCollectionResponse**](SubscriptionCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookSubscription

> SubscriptionResponse UpdateWebhookSubscription(ctx, webhookSubscriptionId).SubscriptionUpdateRequest(subscriptionUpdateRequest).Execute()

Update a webhook subscription



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
    webhookSubscriptionId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the webhook subscription
    subscriptionUpdateRequest := *AirEml.NewSubscriptionUpdateRequest("Using for check sms status", []string{"Events_example"}, map[string]interface{}({"isTest":true,"externalId":1,"source":"gclick_123"}), int32(2), "https://aireml.com/api/callback") // SubscriptionUpdateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.UpdateWebhookSubscription(context.Background(), webhookSubscriptionId).SubscriptionUpdateRequest(subscriptionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.UpdateWebhookSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookSubscription`: SubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.UpdateWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookSubscriptionId** | **string** | The ID of the webhook subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionUpdateRequest** | [**SubscriptionUpdateRequest**](SubscriptionUpdateRequest.md) |  | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

