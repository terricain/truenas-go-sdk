# \NetworkAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkConfiguration**](NetworkAPI.md#GetNetworkConfiguration) | **Get** /network/configuration | 
[**GetNetworkSummary**](NetworkAPI.md#GetNetworkSummary) | **Get** /network/general/summary | 



## GetNetworkConfiguration

> NetworkConfig GetNetworkConfiguration(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/terrycain/truenas-go-sdk/pkg/truenas"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkAPI.GetNetworkConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkConfiguration`: NetworkConfig
    fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkConfigurationRequest struct via the builder pattern


### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSummary

> NetworkSummary GetNetworkSummary(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/terrycain/truenas-go-sdk/pkg/truenas"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkAPI.GetNetworkSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSummary`: NetworkSummary
    fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSummaryRequest struct via the builder pattern


### Return type

[**NetworkSummary**](NetworkSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

