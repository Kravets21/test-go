# \AudienceAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachAudienceRecipientContact**](AudienceAPI.md#AttachAudienceRecipientContact) | **Post** /api/v1/audience/recipients/{audience_recipient_id}/contacts/{audience_contact_type} | Attach a Contact to the Audience Recipient
[**AttachAudienceRecipientTag**](AudienceAPI.md#AttachAudienceRecipientTag) | **Post** /api/v1/audience/recipients/{audience_recipient_id}/tags/{audience_tag_name} | Attach a Tag to the Audience Recipient
[**CreateAudienceRecipient**](AudienceAPI.md#CreateAudienceRecipient) | **Post** /api/v1/audience/recipients | Create a New Recipient
[**CreateAudienceSegment**](AudienceAPI.md#CreateAudienceSegment) | **Post** /api/v1/audience/segments | Create a New Segment
[**CreateAudienceSegmentFilter**](AudienceAPI.md#CreateAudienceSegmentFilter) | **Post** /api/v1/audience/segments/{audience_segment_id}/filters | Create a New Segment Filter
[**CreateAudienceTag**](AudienceAPI.md#CreateAudienceTag) | **Post** /api/v1/audience/tags | Create a New Tag
[**DeleteAudienceRecipient**](AudienceAPI.md#DeleteAudienceRecipient) | **Delete** /api/v1/audience/recipients/{audience_recipient_id} | Delete a Recipient
[**DeleteAudienceRecipientContact**](AudienceAPI.md#DeleteAudienceRecipientContact) | **Delete** /api/v1/audience/recipients/{audience_recipient_id}/contacts/{audience_contact_type} | Delete a Contact from the Audience Recipient
[**DeleteAudienceRecipientTag**](AudienceAPI.md#DeleteAudienceRecipientTag) | **Delete** /api/v1/audience/recipients/{audience_recipient_id}/tags/{audience_tag_name} | Delete a Tag from the Audience Recipient
[**DeleteAudienceSegment**](AudienceAPI.md#DeleteAudienceSegment) | **Delete** /api/v1/audience/segments/{audience_segment_id} | Delete a Segment
[**DeleteAudienceSegmentFilter**](AudienceAPI.md#DeleteAudienceSegmentFilter) | **Delete** /api/v1/audience/segments/{audience_segment_id}/filters/{audience_filter_id} | Delete a Filter from the Audience Segment
[**DeleteAudienceTag**](AudienceAPI.md#DeleteAudienceTag) | **Delete** /api/v1/audience/tags/{audience_tag_id} | Delete a Tag
[**GetAudienceRecipient**](AudienceAPI.md#GetAudienceRecipient) | **Get** /api/v1/audience/recipients/{audience_recipient_id} | Retrieve an Existing Recipient
[**GetAudienceSegment**](AudienceAPI.md#GetAudienceSegment) | **Get** /api/v1/audience/segments/{audience_segment_id} | Retrieve an Existing Segment
[**GetAudienceSegmentFilter**](AudienceAPI.md#GetAudienceSegmentFilter) | **Get** /api/v1/audience/segments/{audience_segment_id}/filters/{audience_filter_id} | Retrieve an Existing Filter from the Audience Segment
[**GetAudienceTag**](AudienceAPI.md#GetAudienceTag) | **Get** /api/v1/audience/tags/{audience_tag_id} | Retrieve an Existing Tag
[**ListAudienceRecipientSegments**](AudienceAPI.md#ListAudienceRecipientSegments) | **Get** /api/v1/audience/recipients/{audience_recipient_id}/segments | List All Segments from the Recipient
[**ListAudienceRecipients**](AudienceAPI.md#ListAudienceRecipients) | **Get** /api/v1/audience/recipients | List All Recipients
[**ListAudienceSegmentFilters**](AudienceAPI.md#ListAudienceSegmentFilters) | **Get** /api/v1/audience/segments/{audience_segment_id}/filters | List All Filters from the Segment
[**ListAudienceSegments**](AudienceAPI.md#ListAudienceSegments) | **Get** /api/v1/audience/segments | List All Segments
[**ListAudienceTags**](AudienceAPI.md#ListAudienceTags) | **Get** /api/v1/audience/tags | List All Tags
[**UpdateAssignActiveStatusToAudienceSegment**](AudienceAPI.md#UpdateAssignActiveStatusToAudienceSegment) | **Put** /api/v1/audience/segments/{audience_segment_id}/activate | Activate a Segment
[**UpdateAssignInactiveStatusToAudienceSegment**](AudienceAPI.md#UpdateAssignInactiveStatusToAudienceSegment) | **Put** /api/v1/audience/segments/{audience_segment_id}/deactivate | Deactivate a Segment
[**UpdateAudienceRecipient**](AudienceAPI.md#UpdateAudienceRecipient) | **Put** /api/v1/audience/recipients/{audience_recipient_id} | Update a Recipient
[**UpdateAudienceSegment**](AudienceAPI.md#UpdateAudienceSegment) | **Put** /api/v1/audience/segments/{audience_segment_id} | Update a Segment
[**UpdateAudienceSegmentFilter**](AudienceAPI.md#UpdateAudienceSegmentFilter) | **Put** /api/v1/audience/segments/{audience_segment_id}/filters/{audience_filter_id} | Update a Filter in the Audience Segment
[**UpdateAudienceTag**](AudienceAPI.md#UpdateAudienceTag) | **Put** /api/v1/audience/tags/{audience_tag_id} | Update a Tag



## AttachAudienceRecipientContact

> AttachAudienceRecipientContact(ctx, audienceRecipientId, audienceContactType).AudienceRecipientContactRequest(audienceRecipientContactRequest).Execute()

Attach a Contact to the Audience Recipient



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
    audienceRecipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Recipient
    audienceContactType := "email" // string | The type of the Contact
    audienceRecipientContactRequest := *AirEml.NewAudienceRecipientContactRequest() // AudienceRecipientContactRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.AttachAudienceRecipientContact(context.Background(), audienceRecipientId, audienceContactType).AudienceRecipientContactRequest(audienceRecipientContactRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.AttachAudienceRecipientContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceRecipientId** | **string** | The ID of the Recipient | 
**audienceContactType** | **string** | The type of the Contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachAudienceRecipientContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **audienceRecipientContactRequest** | [**AudienceRecipientContactRequest**](AudienceRecipientContactRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachAudienceRecipientTag

> AttachAudienceRecipientTag(ctx, audienceRecipientId, audienceTagName).Execute()

Attach a Tag to the Audience Recipient



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
    audienceRecipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Recipient
    audienceTagName := "paid_recipients" // string | The name of the Tag

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.AttachAudienceRecipientTag(context.Background(), audienceRecipientId, audienceTagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.AttachAudienceRecipientTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceRecipientId** | **string** | The ID of the Recipient | 
**audienceTagName** | **string** | The name of the Tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachAudienceRecipientTagRequest struct via the builder pattern


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


## CreateAudienceRecipient

> AudienceRecipientResponse CreateAudienceRecipient(ctx).AudienceRecipientRequest(audienceRecipientRequest).Execute()

Create a New Recipient



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
    audienceRecipientRequest := *AirEml.NewAudienceRecipientRequest() // AudienceRecipientRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.CreateAudienceRecipient(context.Background()).AudienceRecipientRequest(audienceRecipientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.CreateAudienceRecipient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAudienceRecipient`: AudienceRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.CreateAudienceRecipient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudienceRecipientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceRecipientRequest** | [**AudienceRecipientRequest**](AudienceRecipientRequest.md) |  | 

### Return type

[**AudienceRecipientResponse**](AudienceRecipientResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAudienceSegment

> AudienceSegmentResponse CreateAudienceSegment(ctx).AudienceSegmentRequest(audienceSegmentRequest).Execute()

Create a New Segment



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
    audienceSegmentRequest := *AirEml.NewAudienceSegmentRequest() // AudienceSegmentRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.CreateAudienceSegment(context.Background()).AudienceSegmentRequest(audienceSegmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.CreateAudienceSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAudienceSegment`: AudienceSegmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.CreateAudienceSegment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudienceSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceSegmentRequest** | [**AudienceSegmentRequest**](AudienceSegmentRequest.md) |  | 

### Return type

[**AudienceSegmentResponse**](AudienceSegmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAudienceSegmentFilter

> AudienceFilterResponse CreateAudienceSegmentFilter(ctx, audienceSegmentId).AudienceFilterRequest(audienceFilterRequest).Execute()

Create a New Segment Filter



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment
    audienceFilterRequest := *AirEml.NewAudienceFilterRequest() // AudienceFilterRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.CreateAudienceSegmentFilter(context.Background(), audienceSegmentId).AudienceFilterRequest(audienceFilterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.CreateAudienceSegmentFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAudienceSegmentFilter`: AudienceFilterResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.CreateAudienceSegmentFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudienceSegmentFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **audienceFilterRequest** | [**AudienceFilterRequest**](AudienceFilterRequest.md) |  | 

### Return type

[**AudienceFilterResponse**](AudienceFilterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAudienceTag

> AudienceTagResponse CreateAudienceTag(ctx).AudienceTagRequest(audienceTagRequest).Execute()

Create a New Tag



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
    audienceTagRequest := *AirEml.NewAudienceTagRequest() // AudienceTagRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.CreateAudienceTag(context.Background()).AudienceTagRequest(audienceTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.CreateAudienceTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAudienceTag`: AudienceTagResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.CreateAudienceTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudienceTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceTagRequest** | [**AudienceTagRequest**](AudienceTagRequest.md) |  | 

### Return type

[**AudienceTagResponse**](AudienceTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAudienceRecipient

> DeleteAudienceRecipient(ctx, audienceRecipientId).Execute()

Delete a Recipient



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
    audienceRecipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Recipient

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.DeleteAudienceRecipient(context.Background(), audienceRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.DeleteAudienceRecipient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceRecipientId** | **string** | The ID of the Recipient | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceRecipientRequest struct via the builder pattern


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


## DeleteAudienceRecipientContact

> DeleteAudienceRecipientContact(ctx, audienceRecipientId, audienceContactType).Execute()

Delete a Contact from the Audience Recipient



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
    audienceRecipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Recipient
    audienceContactType := "email" // string | The type of the Contact

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.DeleteAudienceRecipientContact(context.Background(), audienceRecipientId, audienceContactType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.DeleteAudienceRecipientContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceRecipientId** | **string** | The ID of the Recipient | 
**audienceContactType** | **string** | The type of the Contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceRecipientContactRequest struct via the builder pattern


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


## DeleteAudienceRecipientTag

> DeleteAudienceRecipientTag(ctx, audienceRecipientId, audienceTagName).Execute()

Delete a Tag from the Audience Recipient



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
    audienceRecipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Recipient
    audienceTagName := "paid_recipients" // string | The name of the Tag

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.DeleteAudienceRecipientTag(context.Background(), audienceRecipientId, audienceTagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.DeleteAudienceRecipientTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceRecipientId** | **string** | The ID of the Recipient | 
**audienceTagName** | **string** | The name of the Tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceRecipientTagRequest struct via the builder pattern


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


## DeleteAudienceSegment

> DeleteAudienceSegment(ctx, audienceSegmentId).Execute()

Delete a Segment



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.DeleteAudienceSegment(context.Background(), audienceSegmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.DeleteAudienceSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceSegmentRequest struct via the builder pattern


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


## DeleteAudienceSegmentFilter

> DeleteAudienceSegmentFilter(ctx, audienceSegmentId, audienceFilterId).Execute()

Delete a Filter from the Audience Segment



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment
    audienceFilterId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Filter

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.DeleteAudienceSegmentFilter(context.Background(), audienceSegmentId, audienceFilterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.DeleteAudienceSegmentFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 
**audienceFilterId** | **string** | The ID of the Filter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceSegmentFilterRequest struct via the builder pattern


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


## DeleteAudienceTag

> DeleteAudienceTag(ctx, audienceTagId).Execute()

Delete a Tag



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
    audienceTagId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Tag

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.DeleteAudienceTag(context.Background(), audienceTagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.DeleteAudienceTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceTagId** | **string** | The ID of the Tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceTagRequest struct via the builder pattern


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


## GetAudienceRecipient

> AudienceRecipientResponse GetAudienceRecipient(ctx, audienceRecipientId).Execute()

Retrieve an Existing Recipient



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
    audienceRecipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Recipient

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.GetAudienceRecipient(context.Background(), audienceRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.GetAudienceRecipient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudienceRecipient`: AudienceRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.GetAudienceRecipient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceRecipientId** | **string** | The ID of the Recipient | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceRecipientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AudienceRecipientResponse**](AudienceRecipientResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudienceSegment

> AudienceSegmentResponse GetAudienceSegment(ctx, audienceSegmentId).Execute()

Retrieve an Existing Segment



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.GetAudienceSegment(context.Background(), audienceSegmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.GetAudienceSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudienceSegment`: AudienceSegmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.GetAudienceSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AudienceSegmentResponse**](AudienceSegmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudienceSegmentFilter

> AudienceFilterResponse GetAudienceSegmentFilter(ctx, audienceSegmentId, audienceFilterId).Execute()

Retrieve an Existing Filter from the Audience Segment



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment
    audienceFilterId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Filter

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.GetAudienceSegmentFilter(context.Background(), audienceSegmentId, audienceFilterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.GetAudienceSegmentFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudienceSegmentFilter`: AudienceFilterResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.GetAudienceSegmentFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 
**audienceFilterId** | **string** | The ID of the Filter | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceSegmentFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AudienceFilterResponse**](AudienceFilterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudienceTag

> AudienceTagResponse GetAudienceTag(ctx, audienceTagId).Execute()

Retrieve an Existing Tag



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
    audienceTagId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Tag

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.GetAudienceTag(context.Background(), audienceTagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.GetAudienceTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudienceTag`: AudienceTagResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.GetAudienceTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceTagId** | **string** | The ID of the Tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AudienceTagResponse**](AudienceTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAudienceRecipientSegments

> AudienceRecipientSegmentCollectionResponse ListAudienceRecipientSegments(ctx, audienceRecipientId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()

List All Segments from the Recipient



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
    audienceRecipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Recipient
    cursor := "eyJjcmVhdGVkX2F0IjoiMjAyMy0wMS0xOCAxMjo1ODoyMyIsImlkIjoiYWRhM2NkMGItODE2YS0zMDc2LWEyOGUtNzYyMjNkNTRlNDMyIiwiX3BvaW50c1RvTmV4dEl0ZW1zIjpmYWxzZX0" // string | Which 'cursor' of paginated results to return. (optional)
    perPage := int32(2) // int32 | Number of items returned per page (optional) (default to 20)
    orderBy := "latest" // string | The field to order the results by (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.ListAudienceRecipientSegments(context.Background(), audienceRecipientId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.ListAudienceRecipientSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAudienceRecipientSegments`: AudienceRecipientSegmentCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.ListAudienceRecipientSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceRecipientId** | **string** | The ID of the Recipient | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceRecipientSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 

### Return type

[**AudienceRecipientSegmentCollectionResponse**](AudienceRecipientSegmentCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAudienceRecipients

> AudienceRecipientCollectionResponse ListAudienceRecipients(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).ExternalId(externalId).RecipientId(recipientId).ContactAddress(contactAddress).Execute()

List All Recipients



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
    externalId := "123456789" // string | The external id of the Recipient (optional)
    recipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Recipient (optional)
    contactAddress := "john.doe@example.com" // string | The Contact Address of the Recipient (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.ListAudienceRecipients(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).ExternalId(externalId).RecipientId(recipientId).ContactAddress(contactAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.ListAudienceRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAudienceRecipients`: AudienceRecipientCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.ListAudienceRecipients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **externalId** | **string** | The external id of the Recipient | 
 **recipientId** | **string** | The ID of the Recipient | 
 **contactAddress** | **string** | The Contact Address of the Recipient | 

### Return type

[**AudienceRecipientCollectionResponse**](AudienceRecipientCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAudienceSegmentFilters

> AudienceFilterCollectionResponse ListAudienceSegmentFilters(ctx, audienceSegmentId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()

List All Filters from the Segment



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment
    cursor := "eyJjcmVhdGVkX2F0IjoiMjAyMy0wMS0xOCAxMjo1ODoyMyIsImlkIjoiYWRhM2NkMGItODE2YS0zMDc2LWEyOGUtNzYyMjNkNTRlNDMyIiwiX3BvaW50c1RvTmV4dEl0ZW1zIjpmYWxzZX0" // string | Which 'cursor' of paginated results to return. (optional)
    perPage := int32(2) // int32 | Number of items returned per page (optional) (default to 20)
    orderBy := "latest" // string | The field to order the results by (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.ListAudienceSegmentFilters(context.Background(), audienceSegmentId).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.ListAudienceSegmentFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAudienceSegmentFilters`: AudienceFilterCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.ListAudienceSegmentFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceSegmentFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 

### Return type

[**AudienceFilterCollectionResponse**](AudienceFilterCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAudienceSegments

> AudienceSegmentCollectionResponse ListAudienceSegments(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Execute()

List All Segments



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
    name := "paid" // string | The name of the Segment (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.ListAudienceSegments(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.ListAudienceSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAudienceSegments`: AudienceSegmentCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.ListAudienceSegments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **name** | **string** | The name of the Segment | 

### Return type

[**AudienceSegmentCollectionResponse**](AudienceSegmentCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAudienceTags

> AudienceTagCollectionResponse ListAudienceTags(ctx).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Execute()

List All Tags



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
    name := "paid" // string | The name of the Tag (optional)

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.ListAudienceTags(context.Background()).Cursor(cursor).PerPage(perPage).OrderBy(orderBy).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.ListAudienceTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAudienceTags`: AudienceTagCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.ListAudienceTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAudienceTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Which &#39;cursor&#39; of paginated results to return. | 
 **perPage** | **int32** | Number of items returned per page | [default to 20]
 **orderBy** | **string** | The field to order the results by | 
 **name** | **string** | The name of the Tag | 

### Return type

[**AudienceTagCollectionResponse**](AudienceTagCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssignActiveStatusToAudienceSegment

> UpdateAssignActiveStatusToAudienceSegment(ctx, audienceSegmentId).Execute()

Activate a Segment



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.UpdateAssignActiveStatusToAudienceSegment(context.Background(), audienceSegmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.UpdateAssignActiveStatusToAudienceSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssignActiveStatusToAudienceSegmentRequest struct via the builder pattern


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


## UpdateAssignInactiveStatusToAudienceSegment

> UpdateAssignInactiveStatusToAudienceSegment(ctx, audienceSegmentId).Execute()

Deactivate a Segment



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    r, err := apiClient.AudienceAPI.UpdateAssignInactiveStatusToAudienceSegment(context.Background(), audienceSegmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.UpdateAssignInactiveStatusToAudienceSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssignInactiveStatusToAudienceSegmentRequest struct via the builder pattern


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


## UpdateAudienceRecipient

> AudienceRecipientResponse UpdateAudienceRecipient(ctx, audienceRecipientId).AudienceRecipientRequest(audienceRecipientRequest).Execute()

Update a Recipient



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
    audienceRecipientId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Recipient
    audienceRecipientRequest := *AirEml.NewAudienceRecipientRequest() // AudienceRecipientRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.UpdateAudienceRecipient(context.Background(), audienceRecipientId).AudienceRecipientRequest(audienceRecipientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.UpdateAudienceRecipient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAudienceRecipient`: AudienceRecipientResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.UpdateAudienceRecipient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceRecipientId** | **string** | The ID of the Recipient | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceRecipientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **audienceRecipientRequest** | [**AudienceRecipientRequest**](AudienceRecipientRequest.md) |  | 

### Return type

[**AudienceRecipientResponse**](AudienceRecipientResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudienceSegment

> AudienceSegmentResponse UpdateAudienceSegment(ctx, audienceSegmentId).AudienceSegmentRequest(audienceSegmentRequest).Execute()

Update a Segment



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment
    audienceSegmentRequest := *AirEml.NewAudienceSegmentRequest() // AudienceSegmentRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.UpdateAudienceSegment(context.Background(), audienceSegmentId).AudienceSegmentRequest(audienceSegmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.UpdateAudienceSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAudienceSegment`: AudienceSegmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.UpdateAudienceSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **audienceSegmentRequest** | [**AudienceSegmentRequest**](AudienceSegmentRequest.md) |  | 

### Return type

[**AudienceSegmentResponse**](AudienceSegmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudienceSegmentFilter

> AudienceFilterResponse UpdateAudienceSegmentFilter(ctx, audienceSegmentId, audienceFilterId).AudienceFilterRequest(audienceFilterRequest).Execute()

Update a Filter in the Audience Segment



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
    audienceSegmentId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Segment
    audienceFilterId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Filter
    audienceFilterRequest := *AirEml.NewAudienceFilterRequest() // AudienceFilterRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.UpdateAudienceSegmentFilter(context.Background(), audienceSegmentId, audienceFilterId).AudienceFilterRequest(audienceFilterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.UpdateAudienceSegmentFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAudienceSegmentFilter`: AudienceFilterResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.UpdateAudienceSegmentFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceSegmentId** | **string** | The ID of the Segment | 
**audienceFilterId** | **string** | The ID of the Filter | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceSegmentFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **audienceFilterRequest** | [**AudienceFilterRequest**](AudienceFilterRequest.md) |  | 

### Return type

[**AudienceFilterResponse**](AudienceFilterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudienceTag

> AudienceTagResponse UpdateAudienceTag(ctx, audienceTagId).AudienceTagRequest(audienceTagRequest).Execute()

Update a Tag



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
    audienceTagId := "00112233-4455-6677-8899-aabbccddeeff" // string | The ID of the Tag
    audienceTagRequest := *AirEml.NewAudienceTagRequest() // AudienceTagRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AudienceAPI.UpdateAudienceTag(context.Background(), audienceTagId).AudienceTagRequest(audienceTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudienceAPI.UpdateAudienceTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAudienceTag`: AudienceTagResponse
    fmt.Fprintf(os.Stdout, "Response from `AudienceAPI.UpdateAudienceTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceTagId** | **string** | The ID of the Tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **audienceTagRequest** | [**AudienceTagRequest**](AudienceTagRequest.md) |  | 

### Return type

[**AudienceTagResponse**](AudienceTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

