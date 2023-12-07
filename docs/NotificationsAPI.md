# \NotificationsAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailNotification**](NotificationsAPI.md#GetEmailNotification) | **Get** /api/v1/email/notifications/{notification_id} | Retrieve an Existing Email Notification
[**GetEmailNotificationStates**](NotificationsAPI.md#GetEmailNotificationStates) | **Get** /api/v1/email/notifications/{notification_id}/states | Retrieve an Existing Email Notification States
[**GetPushNotification**](NotificationsAPI.md#GetPushNotification) | **Get** /api/v1/push/notifications/{notification_id} | Retrieve an Existing Push Notification
[**GetPushNotificationStates**](NotificationsAPI.md#GetPushNotificationStates) | **Get** /api/v1/push/notifications/{notification_id}/states | Retrieve an Existing Push Notification States
[**GetSmsNotification**](NotificationsAPI.md#GetSmsNotification) | **Get** /api/v1/sms/notifications/{notification_id} | Retrieve an Existing SMS Notification
[**GetSmsNotificationStates**](NotificationsAPI.md#GetSmsNotificationStates) | **Get** /api/v1/sms/notifications/{notification_id}/states | Retrieve an Existing Sms Notification States
[**ListEmailNotifications**](NotificationsAPI.md#ListEmailNotifications) | **Get** /api/v1/email/notifications | List All Email Notifications
[**ListPushNotifications**](NotificationsAPI.md#ListPushNotifications) | **Get** /api/v1/push/notifications | List All Push Notifications
[**ListSmsNotifications**](NotificationsAPI.md#ListSmsNotifications) | **Get** /api/v1/sms/notifications | List All SMS Notifications



## GetEmailNotification

> EmailNotificationResponse GetEmailNotification(ctx, notificationId).Execute()

Retrieve an Existing Email Notification



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
    notificationId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Notification

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.GetEmailNotification(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetEmailNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailNotification`: EmailNotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetEmailNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | The ID of the Notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailNotificationResponse**](EmailNotificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailNotificationStates

> EmailNotificationStateResponse GetEmailNotificationStates(ctx, notificationId).Execute()

Retrieve an Existing Email Notification States



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
    notificationId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Notification

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.GetEmailNotificationStates(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetEmailNotificationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailNotificationStates`: EmailNotificationStateResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetEmailNotificationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | The ID of the Notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailNotificationStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailNotificationStateResponse**](EmailNotificationStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushNotification

> PushNotificationResponse GetPushNotification(ctx, notificationId).Execute()

Retrieve an Existing Push Notification



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
    notificationId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Notification

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.GetPushNotification(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetPushNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushNotification`: PushNotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetPushNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | The ID of the Notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PushNotificationResponse**](PushNotificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushNotificationStates

> PushNotificationStateResponse GetPushNotificationStates(ctx, notificationId).Execute()

Retrieve an Existing Push Notification States



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
    notificationId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Notification

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.GetPushNotificationStates(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetPushNotificationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushNotificationStates`: PushNotificationStateResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetPushNotificationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | The ID of the Notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushNotificationStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PushNotificationStateResponse**](PushNotificationStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmsNotification

> SmsNotificationResponse GetSmsNotification(ctx, notificationId).Execute()

Retrieve an Existing SMS Notification



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
    notificationId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Notification

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.GetSmsNotification(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSmsNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmsNotification`: SmsNotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSmsNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | The ID of the Notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmsNotificationResponse**](SmsNotificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmsNotificationStates

> SmsNotificationStateResponse GetSmsNotificationStates(ctx, notificationId).Execute()

Retrieve an Existing Sms Notification States



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
    notificationId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Notification

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.GetSmsNotificationStates(context.Background(), notificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetSmsNotificationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmsNotificationStates`: SmsNotificationStateResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetSmsNotificationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | The ID of the Notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsNotificationStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmsNotificationStateResponse**](SmsNotificationStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailNotifications

> EmailNotificationCollectionResponse ListEmailNotifications(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).ToEmailAddress(toEmailAddress).TemplateId(templateId).TemplateName(templateName).RecipientId(recipientId).Execute()

List All Email Notifications



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
    toEmailAddress := "example@axample.com" // string | The email address of the Notification (optional)
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The template id of the Notification (optional)
    templateName := "welcome_to_documentation" // string | The template name of the Notification (optional)
    recipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The recipient id of the Notification (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.ListEmailNotifications(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).ToEmailAddress(toEmailAddress).TemplateId(templateId).TemplateName(templateName).RecipientId(recipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListEmailNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailNotifications`: EmailNotificationCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListEmailNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **toEmailAddress** | **string** | The email address of the Notification | 
 **templateId** | **string** | The template id of the Notification | 
 **templateName** | **string** | The template name of the Notification | 
 **recipientId** | **string** | The recipient id of the Notification | 

### Return type

[**EmailNotificationCollectionResponse**](EmailNotificationCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPushNotifications

> PushNotificationCollectionResponse ListPushNotifications(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).RecipientId(recipientId).Token(token).TemplateId(templateId).TemplateName(templateName).Execute()

List All Push Notifications



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
    recipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The recipient id of the Notification (optional)
    token := "01234567890:ABCDEFGpERNZ7em...aB06VJ-ABCD_XB34UgMJ6_abc_mJqnne8KxuHYxmemu21othozA3hqdNpKcmMr" // string | The device token of the Notification (optional)
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The template id of the Notification (optional)
    templateName := "welcome_to_documentation" // string | The template name of the Notification (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.ListPushNotifications(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).RecipientId(recipientId).Token(token).TemplateId(templateId).TemplateName(templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListPushNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPushNotifications`: PushNotificationCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListPushNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPushNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **recipientId** | **string** | The recipient id of the Notification | 
 **token** | **string** | The device token of the Notification | 
 **templateId** | **string** | The template id of the Notification | 
 **templateName** | **string** | The template name of the Notification | 

### Return type

[**PushNotificationCollectionResponse**](PushNotificationCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmsNotifications

> SmsNotificationCollectionResponse ListSmsNotifications(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).TemplateName(templateName).TemplateId(templateId).FromNumber(fromNumber).ToNumber(toNumber).RecipientId(recipientId).Execute()

List All SMS Notifications



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
    templateName := "welcome_to_documentation" // string | To filter records by the template name of the Notification. (optional)
    templateId := "00112233-4455-6677-8899-aabbccddeeff" // string | The template id of the Notification (optional)
    fromNumber := "+17779091327" // string | The filter value for getting records with provided from number. (optional)
    toNumber := "+19379091324" // string | The filter value for show the records of the notifications that have a similar to number. (optional)
    recipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The recipient id of the Notification (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.ListSmsNotifications(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).TemplateName(templateName).TemplateId(templateId).FromNumber(fromNumber).ToNumber(toNumber).RecipientId(recipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListSmsNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmsNotifications`: SmsNotificationCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListSmsNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSmsNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **templateName** | **string** | To filter records by the template name of the Notification. | 
 **templateId** | **string** | The template id of the Notification | 
 **fromNumber** | **string** | The filter value for getting records with provided from number. | 
 **toNumber** | **string** | The filter value for show the records of the notifications that have a similar to number. | 
 **recipientId** | **string** | The recipient id of the Notification | 

### Return type

[**SmsNotificationCollectionResponse**](SmsNotificationCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

