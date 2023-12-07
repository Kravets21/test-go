# \EmailSettingsAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailChannel**](EmailSettingsAPI.md#CreateEmailChannel) | **Post** /api/v1/email/channels | Create a New Email Notification Channel
[**CreateEmailChannelRoutingRule**](EmailSettingsAPI.md#CreateEmailChannelRoutingRule) | **Post** /api/v1/email/channels/{channel_id}/routing-rules | Create a Routing-rule in the Email Notification Channel
[**CreateEmailSender**](EmailSettingsAPI.md#CreateEmailSender) | **Post** /api/v1/email/senders | Create a New Email Notification Sender
[**CreateEmailTransport**](EmailSettingsAPI.md#CreateEmailTransport) | **Post** /api/v1/email/transports | Create a New Email Notification Transport
[**DeleteEmailChannel**](EmailSettingsAPI.md#DeleteEmailChannel) | **Delete** /api/v1/email/channels/{channel_id} | Delete a Email Notification Channel
[**DeleteEmailChannelRoutingRule**](EmailSettingsAPI.md#DeleteEmailChannelRoutingRule) | **Delete** /api/v1/email/channels/{channel_id}/routing-rules/{routing_rule_id} | Delete a Routing-rule from the Email Notification Channel
[**DeleteEmailSender**](EmailSettingsAPI.md#DeleteEmailSender) | **Delete** /api/v1/email/senders/{sender_id} | Delete a Email Notification Sender
[**DeleteEmailTransport**](EmailSettingsAPI.md#DeleteEmailTransport) | **Delete** /api/v1/email/transports/{transport_id} | Delete a Email Notification Transport
[**GetEmailChannel**](EmailSettingsAPI.md#GetEmailChannel) | **Get** /api/v1/email/channels/{channel_id} | Retrieve an Existing Email Notification Channel
[**GetEmailChannelRoutingRule**](EmailSettingsAPI.md#GetEmailChannelRoutingRule) | **Get** /api/v1/email/channels/{channel_id}/routing-rules/{routing_rule_id} | Retrieve an Existing Routing-rule from the Email Notification Channel
[**GetEmailSender**](EmailSettingsAPI.md#GetEmailSender) | **Get** /api/v1/email/senders/{sender_id} | Retrieve an Existing Email Notification Sender
[**GetEmailTransport**](EmailSettingsAPI.md#GetEmailTransport) | **Get** /api/v1/email/transports/{transport_id} | Retrieve an Existing Email Notification Transport
[**ListEmailChannelRoutingRules**](EmailSettingsAPI.md#ListEmailChannelRoutingRules) | **Get** /api/v1/email/channels/{channel_id}/routing-rules | List All Routing-rules from the Email Notification Channel
[**ListEmailChannels**](EmailSettingsAPI.md#ListEmailChannels) | **Get** /api/v1/email/channels | List All Email Notification Channels
[**ListEmailSenders**](EmailSettingsAPI.md#ListEmailSenders) | **Get** /api/v1/email/senders | List All Email Notification Senders
[**ListEmailTransports**](EmailSettingsAPI.md#ListEmailTransports) | **Get** /api/v1/email/transports | List All Email Notification Transports
[**UpdateEmailChannel**](EmailSettingsAPI.md#UpdateEmailChannel) | **Put** /api/v1/email/channels/{channel_id} | Update a Email Notification Channel
[**UpdateEmailChannelRoutingRule**](EmailSettingsAPI.md#UpdateEmailChannelRoutingRule) | **Put** /api/v1/email/channels/{channel_id}/routing-rules/{routing_rule_id} | Update a Routing-rule in the Email Notification Channel
[**UpdateEmailSender**](EmailSettingsAPI.md#UpdateEmailSender) | **Put** /api/v1/email/senders/{sender_id} | Update a Email Notification Sender
[**UpdateEmailTransport**](EmailSettingsAPI.md#UpdateEmailTransport) | **Put** /api/v1/email/transports/{transport_id} | Update a Email Notification Transport



## CreateEmailChannel

> EmailChannelResponse CreateEmailChannel(ctx).EmailChannelRequest(emailChannelRequest).Execute()

Create a New Email Notification Channel



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
    emailChannelRequest := *AirEml.NewEmailChannelRequest("example channel") // EmailChannelRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.CreateEmailChannel(context.Background()).EmailChannelRequest(emailChannelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.CreateEmailChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailChannel`: EmailChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.CreateEmailChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailChannelRequest** | [**EmailChannelRequest**](EmailChannelRequest.md) |  | 

### Return type

[**EmailChannelResponse**](EmailChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmailChannelRoutingRule

> EmailChannelRoutingRuleResponse CreateEmailChannelRoutingRule(ctx, channelId).EmailRoutingRuleCreateRequest(emailRoutingRuleCreateRequest).Execute()

Create a Routing-rule in the Email Notification Channel



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
    channelId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Channel
    emailRoutingRuleCreateRequest := *AirEml.NewEmailRoutingRuleCreateRequest("00112233-4455-6677-8899-aabbccddeeff", "00112233-4455-6677-8899-aabbccddeeff") // EmailRoutingRuleCreateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.CreateEmailChannelRoutingRule(context.Background(), channelId).EmailRoutingRuleCreateRequest(emailRoutingRuleCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.CreateEmailChannelRoutingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailChannelRoutingRule`: EmailChannelRoutingRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.CreateEmailChannelRoutingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | The ID of the Channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailChannelRoutingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailRoutingRuleCreateRequest** | [**EmailRoutingRuleCreateRequest**](EmailRoutingRuleCreateRequest.md) |  | 

### Return type

[**EmailChannelRoutingRuleResponse**](EmailChannelRoutingRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmailSender

> EmailSenderResponse CreateEmailSender(ctx).EmailSenderCreateRequest(emailSenderCreateRequest).Execute()

Create a New Email Notification Sender



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
    emailSenderCreateRequest := *AirEml.NewEmailSenderCreateRequest("John Doe", "john.doe@example.com") // EmailSenderCreateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.CreateEmailSender(context.Background()).EmailSenderCreateRequest(emailSenderCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.CreateEmailSender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailSender`: EmailSenderResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.CreateEmailSender`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailSenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailSenderCreateRequest** | [**EmailSenderCreateRequest**](EmailSenderCreateRequest.md) |  | 

### Return type

[**EmailSenderResponse**](EmailSenderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmailTransport

> EmailTransportResponse CreateEmailTransport(ctx).EmailTransportCreateRequest(emailTransportCreateRequest).Execute()

Create a New Email Notification Transport



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
    emailTransportCreateRequest := *AirEml.NewEmailTransportCreateRequest("main", "smtp", map[string]interface{}({"host":"smtp.example.com","port":5877,"username":"username","password":"password"})) // EmailTransportCreateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.CreateEmailTransport(context.Background()).EmailTransportCreateRequest(emailTransportCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.CreateEmailTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailTransport`: EmailTransportResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.CreateEmailTransport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailTransportCreateRequest** | [**EmailTransportCreateRequest**](EmailTransportCreateRequest.md) |  | 

### Return type

[**EmailTransportResponse**](EmailTransportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailChannel

> DeleteEmailChannel(ctx, channelId).Execute()

Delete a Email Notification Channel



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
    channelId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Channel

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.EmailSettingsAPI.DeleteEmailChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.DeleteEmailChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | The ID of the Channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailChannelRequest struct via the builder pattern


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


## DeleteEmailChannelRoutingRule

> DeleteEmailChannelRoutingRule(ctx, channelId, routingRuleId).Execute()

Delete a Routing-rule from the Email Notification Channel



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
    channelId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Channel
    routingRuleId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Routing-rule

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.EmailSettingsAPI.DeleteEmailChannelRoutingRule(context.Background(), channelId, routingRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.DeleteEmailChannelRoutingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | The ID of the Channel | 
**routingRuleId** | **string** | The ID of the Routing-rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailChannelRoutingRuleRequest struct via the builder pattern


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


## DeleteEmailSender

> DeleteEmailSender(ctx, senderId).Execute()

Delete a Email Notification Sender



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
    senderId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Sender

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.EmailSettingsAPI.DeleteEmailSender(context.Background(), senderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.DeleteEmailSender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**senderId** | **string** | The ID of the Sender | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailSenderRequest struct via the builder pattern


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


## DeleteEmailTransport

> DeleteEmailTransport(ctx, transportId).Execute()

Delete a Email Notification Transport



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
    transportId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Transport

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.EmailSettingsAPI.DeleteEmailTransport(context.Background(), transportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.DeleteEmailTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | The ID of the Transport | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailTransportRequest struct via the builder pattern


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


## GetEmailChannel

> EmailChannelResponse GetEmailChannel(ctx, channelId).Execute()

Retrieve an Existing Email Notification Channel



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
    channelId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Channel

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.GetEmailChannel(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.GetEmailChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailChannel`: EmailChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.GetEmailChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | The ID of the Channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailChannelResponse**](EmailChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailChannelRoutingRule

> EmailChannelRoutingRuleResponse GetEmailChannelRoutingRule(ctx, channelId, routingRuleId).Execute()

Retrieve an Existing Routing-rule from the Email Notification Channel



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
    channelId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Channel
    routingRuleId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Routing-rule

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.GetEmailChannelRoutingRule(context.Background(), channelId, routingRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.GetEmailChannelRoutingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailChannelRoutingRule`: EmailChannelRoutingRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.GetEmailChannelRoutingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | The ID of the Channel | 
**routingRuleId** | **string** | The ID of the Routing-rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailChannelRoutingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailChannelRoutingRuleResponse**](EmailChannelRoutingRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSender

> EmailSenderResponse GetEmailSender(ctx, senderId).Execute()

Retrieve an Existing Email Notification Sender



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
    senderId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Sender

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.GetEmailSender(context.Background(), senderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.GetEmailSender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailSender`: EmailSenderResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.GetEmailSender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**senderId** | **string** | The ID of the Sender | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailSenderResponse**](EmailSenderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTransport

> EmailTransportResponse GetEmailTransport(ctx, transportId).Execute()

Retrieve an Existing Email Notification Transport



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
    transportId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Transport

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.GetEmailTransport(context.Background(), transportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.GetEmailTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailTransport`: EmailTransportResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.GetEmailTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | The ID of the Transport | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailTransportResponse**](EmailTransportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailChannelRoutingRules

> EmailChannelRoutingRulesCollectionResponse ListEmailChannelRoutingRules(ctx, channelId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()

List All Routing-rules from the Email Notification Channel



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
    channelId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Channel
    cursor := "eyJjcmVhdGVkX2F0IjoiMjAyMy0wMS0xOCAxMjo1ODoyMyIsImlkIjoiYWRhM2NkMGItODE2YS0zMDc2LWEyOGUtNzYyMjNkNTRlNDMyIiwiX3BvaW50c1RvTmV4dEl0ZW1zIjpmYWxzZX0" // string | Which 'cursor' of paginated results to return. (optional)
    perPage := int32(2) // int32 | Number of items returned per page (optional) (default to 20)
    orderBy := "latest" // string | The field to order the results by (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.ListEmailChannelRoutingRules(context.Background(), channelId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.ListEmailChannelRoutingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailChannelRoutingRules`: EmailChannelRoutingRulesCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.ListEmailChannelRoutingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | The ID of the Channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailChannelRoutingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 

### Return type

[**EmailChannelRoutingRulesCollectionResponse**](EmailChannelRoutingRulesCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailChannels

> EmailChannelCollectionResponse ListEmailChannels(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Execute()

List All Email Notification Channels



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
    name := "Example Channel" // string | The filter value for show the records of the channels that have a specific name. (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.ListEmailChannels(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.ListEmailChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailChannels`: EmailChannelCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.ListEmailChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **name** | **string** | The filter value for show the records of the channels that have a specific name. | 

### Return type

[**EmailChannelCollectionResponse**](EmailChannelCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailSenders

> EmailSenderCollectionResponse ListEmailSenders(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Execute()

List All Email Notification Senders



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
    name := "John Doe" // string | The filter value for show the records of the senders that have a specific name. (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.ListEmailSenders(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.ListEmailSenders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailSenders`: EmailSenderCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.ListEmailSenders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailSendersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **name** | **string** | The filter value for show the records of the senders that have a specific name. | 

### Return type

[**EmailSenderCollectionResponse**](EmailSenderCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailTransports

> EmailTransportCollectionResponse ListEmailTransports(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Status(status).Execute()

List All Email Notification Transports



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
    name := "SMTP" // string | The filter value for show the records of the transports that have a specific name. (optional)
    status := int32(1) // int32 | The filter value for show the records of the transports that have a specific status. (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.ListEmailTransports(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.ListEmailTransports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailTransports`: EmailTransportCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.ListEmailTransports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailTransportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **name** | **string** | The filter value for show the records of the transports that have a specific name. | 
 **status** | **int32** | The filter value for show the records of the transports that have a specific status. | 

### Return type

[**EmailTransportCollectionResponse**](EmailTransportCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailChannel

> EmailChannelResponse UpdateEmailChannel(ctx, channelId).EmailChannelRequest(emailChannelRequest).Execute()

Update a Email Notification Channel



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
    channelId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Channel
    emailChannelRequest := *AirEml.NewEmailChannelRequest("example channel") // EmailChannelRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.UpdateEmailChannel(context.Background(), channelId).EmailChannelRequest(emailChannelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.UpdateEmailChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailChannel`: EmailChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.UpdateEmailChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | The ID of the Channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailChannelRequest** | [**EmailChannelRequest**](EmailChannelRequest.md) |  | 

### Return type

[**EmailChannelResponse**](EmailChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailChannelRoutingRule

> EmailChannelRoutingRuleResponse UpdateEmailChannelRoutingRule(ctx, channelId, routingRuleId).EmailRoutingRuleUpdateRequest(emailRoutingRuleUpdateRequest).Execute()

Update a Routing-rule in the Email Notification Channel



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
    channelId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Channel
    routingRuleId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Routing-rule
    emailRoutingRuleUpdateRequest := *AirEml.NewEmailRoutingRuleUpdateRequest() // EmailRoutingRuleUpdateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.UpdateEmailChannelRoutingRule(context.Background(), channelId, routingRuleId).EmailRoutingRuleUpdateRequest(emailRoutingRuleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.UpdateEmailChannelRoutingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailChannelRoutingRule`: EmailChannelRoutingRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.UpdateEmailChannelRoutingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | The ID of the Channel | 
**routingRuleId** | **string** | The ID of the Routing-rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailChannelRoutingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailRoutingRuleUpdateRequest** | [**EmailRoutingRuleUpdateRequest**](EmailRoutingRuleUpdateRequest.md) |  | 

### Return type

[**EmailChannelRoutingRuleResponse**](EmailChannelRoutingRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailSender

> EmailSenderResponse UpdateEmailSender(ctx, senderId).EmailSenderUpdateRequest(emailSenderUpdateRequest).Execute()

Update a Email Notification Sender



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
    senderId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Sender
    emailSenderUpdateRequest := *AirEml.NewEmailSenderUpdateRequest() // EmailSenderUpdateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.UpdateEmailSender(context.Background(), senderId).EmailSenderUpdateRequest(emailSenderUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.UpdateEmailSender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailSender`: EmailSenderResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.UpdateEmailSender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**senderId** | **string** | The ID of the Sender | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailSenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailSenderUpdateRequest** | [**EmailSenderUpdateRequest**](EmailSenderUpdateRequest.md) |  | 

### Return type

[**EmailSenderResponse**](EmailSenderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailTransport

> EmailTransportResponse UpdateEmailTransport(ctx, transportId).EmailTransportUpdateRequest(emailTransportUpdateRequest).Execute()

Update a Email Notification Transport



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
    transportId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Transport
    emailTransportUpdateRequest := *AirEml.NewEmailTransportUpdateRequest() // EmailTransportUpdateRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailSettingsAPI.UpdateEmailTransport(context.Background(), transportId).EmailTransportUpdateRequest(emailTransportUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailSettingsAPI.UpdateEmailTransport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailTransport`: EmailTransportResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailSettingsAPI.UpdateEmailTransport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportId** | **string** | The ID of the Transport | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailTransportUpdateRequest** | [**EmailTransportUpdateRequest**](EmailTransportUpdateRequest.md) |  | 

### Return type

[**EmailTransportResponse**](EmailTransportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

