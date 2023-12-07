# \SendingAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendEmail**](SendingAPI.md#SendEmail) | **Post** /api/v1/email/send | Send Email
[**SendEmailBatch**](SendingAPI.md#SendEmailBatch) | **Post** /api/v1/email/send/batch | Send Email Batch
[**SendPush**](SendingAPI.md#SendPush) | **Post** /api/v1/push/send | Send Push Notification
[**SendSmsBatchNotification**](SendingAPI.md#SendSmsBatchNotification) | **Post** /api/v1/sms/send/batch | Send SMS Batch Notification
[**SendSmsNotification**](SendingAPI.md#SendSmsNotification) | **Post** /api/v1/sms/send | Send SMS Notification



## SendEmail

> EmailSendResponse SendEmail(ctx).EmailSendRequest(emailSendRequest).Execute()

Send Email



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
    emailSendRequest := *AirEml.NewEmailSendRequest() // EmailSendRequest |  (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SendingAPI.SendEmail(context.Background()).EmailSendRequest(emailSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SendingAPI.SendEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendEmail`: EmailSendResponse
    fmt.Fprintf(os.Stdout, "Response from `SendingAPI.SendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailSendRequest** | [**EmailSendRequest**](EmailSendRequest.md) |  | 

### Return type

[**EmailSendResponse**](EmailSendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmailBatch

> EmailSendBatchResponse SendEmailBatch(ctx).EmailSendBatchRequest(emailSendBatchRequest).Execute()

Send Email Batch



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
    emailSendBatchRequest := *AirEml.NewEmailSendBatchRequest([]AirEml.EmailSendRequest{*AirEml.NewEmailSendRequest()}) // EmailSendBatchRequest |  (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SendingAPI.SendEmailBatch(context.Background()).EmailSendBatchRequest(emailSendBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SendingAPI.SendEmailBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendEmailBatch`: EmailSendBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `SendingAPI.SendEmailBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailSendBatchRequest** | [**EmailSendBatchRequest**](EmailSendBatchRequest.md) |  | 

### Return type

[**EmailSendBatchResponse**](EmailSendBatchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPush

> PushNotificationSentResponse SendPush(ctx).PushNotificationSendRequest(pushNotificationSendRequest).Execute()

Send Push Notification



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
    pushNotificationSendRequest := *AirEml.NewPushNotificationSendRequest() // PushNotificationSendRequest |  (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SendingAPI.SendPush(context.Background()).PushNotificationSendRequest(pushNotificationSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SendingAPI.SendPush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendPush`: PushNotificationSentResponse
    fmt.Fprintf(os.Stdout, "Response from `SendingAPI.SendPush`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushNotificationSendRequest** | [**PushNotificationSendRequest**](PushNotificationSendRequest.md) |  | 

### Return type

[**PushNotificationSentResponse**](PushNotificationSentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSmsBatchNotification

> SmsMultipleNotificationSentResponse SendSmsBatchNotification(ctx).SmsNotificationSendMultipleRequest(smsNotificationSendMultipleRequest).Execute()

Send SMS Batch Notification



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
    smsNotificationSendMultipleRequest := *AirEml.NewSmsNotificationSendMultipleRequest() // SmsNotificationSendMultipleRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SendingAPI.SendSmsBatchNotification(context.Background()).SmsNotificationSendMultipleRequest(smsNotificationSendMultipleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SendingAPI.SendSmsBatchNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSmsBatchNotification`: SmsMultipleNotificationSentResponse
    fmt.Fprintf(os.Stdout, "Response from `SendingAPI.SendSmsBatchNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendSmsBatchNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsNotificationSendMultipleRequest** | [**SmsNotificationSendMultipleRequest**](SmsNotificationSendMultipleRequest.md) |  | 

### Return type

[**SmsMultipleNotificationSentResponse**](SmsMultipleNotificationSentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSmsNotification

> SmsNotificationSentResponse SendSmsNotification(ctx).SmsNotificationSendRequest(smsNotificationSendRequest).Execute()

Send SMS Notification



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
    smsNotificationSendRequest := *AirEml.NewSmsNotificationSendRequest() // SmsNotificationSendRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.SendingAPI.SendSmsNotification(context.Background()).SmsNotificationSendRequest(smsNotificationSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SendingAPI.SendSmsNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSmsNotification`: SmsNotificationSentResponse
    fmt.Fprintf(os.Stdout, "Response from `SendingAPI.SendSmsNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendSmsNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsNotificationSendRequest** | [**SmsNotificationSendRequest**](SmsNotificationSendRequest.md) |  | 

### Return type

[**SmsNotificationSentResponse**](SmsNotificationSentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

