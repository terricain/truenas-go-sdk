# \IscsiGlobalAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetISCSIGlobalConfiguration**](IscsiGlobalAPI.md#GetISCSIGlobalConfiguration) | **Get** /iscsi/global | 



## GetISCSIGlobalConfiguration

> ISCSIGlobalConfiguration GetISCSIGlobalConfiguration(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/terricain/truenas-go-sdk/pkg/truenas"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IscsiGlobalAPI.GetISCSIGlobalConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiGlobalAPI.GetISCSIGlobalConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetISCSIGlobalConfiguration`: ISCSIGlobalConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IscsiGlobalAPI.GetISCSIGlobalConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetISCSIGlobalConfigurationRequest struct via the builder pattern


### Return type

[**ISCSIGlobalConfiguration**](ISCSIGlobalConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

