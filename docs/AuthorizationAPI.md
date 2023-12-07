# \AuthorizationAPI

All URIs are relative to *https://aireml.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthTokenForServiceAccount**](AuthorizationAPI.md#CreateAuthTokenForServiceAccount) | **Post** /api/v1/auth/tokens/service-account | Create Access Token for Service Account



## CreateAuthTokenForServiceAccount

> AuthTokenResponse CreateAuthTokenForServiceAccount(ctx).AuthTokenForServiceAccountRequest(authTokenForServiceAccountRequest).Execute()

Create Access Token for Service Account



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
    authTokenForServiceAccountRequest := *AirEml.NewAuthTokenForServiceAccountRequest() // AuthTokenForServiceAccountRequest | 

    configuration := AirEml.NewConfiguration()
    apiClient := AirEml.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationAPI.CreateAuthTokenForServiceAccount(context.Background()).AuthTokenForServiceAccountRequest(authTokenForServiceAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.CreateAuthTokenForServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthTokenForServiceAccount`: AuthTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.CreateAuthTokenForServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthTokenForServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authTokenForServiceAccountRequest** | [**AuthTokenForServiceAccountRequest**](AuthTokenForServiceAccountRequest.md) |  | 

### Return type

[**AuthTokenResponse**](AuthTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

