# \IscsiExtentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateISCSIExtent**](IscsiExtentAPI.md#CreateISCSIExtent) | **Post** /iscsi/extent | 
[**DeleteISCSIExtent**](IscsiExtentAPI.md#DeleteISCSIExtent) | **Delete** /iscsi/extent/id/{id} | 
[**GetISCSIExtent**](IscsiExtentAPI.md#GetISCSIExtent) | **Get** /iscsi/extent/id/{id} | 
[**ListISCSIExtent**](IscsiExtentAPI.md#ListISCSIExtent) | **Get** /iscsi/extent | 



## CreateISCSIExtent

> ISCSIExtent CreateISCSIExtent(ctx).CreateISCSIExtentParams(createISCSIExtentParams).Execute()





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
    createISCSIExtentParams := *openapiclient.NewCreateISCSIExtentParams("Name_example", "Type_example") // CreateISCSIExtentParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IscsiExtentAPI.CreateISCSIExtent(context.Background()).CreateISCSIExtentParams(createISCSIExtentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiExtentAPI.CreateISCSIExtent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateISCSIExtent`: ISCSIExtent
    fmt.Fprintf(os.Stdout, "Response from `IscsiExtentAPI.CreateISCSIExtent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateISCSIExtentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createISCSIExtentParams** | [**CreateISCSIExtentParams**](CreateISCSIExtentParams.md) |  | 

### Return type

[**ISCSIExtent**](ISCSIExtent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteISCSIExtent

> DeleteISCSIExtent(ctx, id).DeleteISCSIExtentParams(deleteISCSIExtentParams).Execute()





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
    id := int32(56) // int32 | 
    deleteISCSIExtentParams := *openapiclient.NewDeleteISCSIExtentParams(false, false) // DeleteISCSIExtentParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IscsiExtentAPI.DeleteISCSIExtent(context.Background(), id).DeleteISCSIExtentParams(deleteISCSIExtentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiExtentAPI.DeleteISCSIExtent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteISCSIExtentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteISCSIExtentParams** | [**DeleteISCSIExtentParams**](DeleteISCSIExtentParams.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetISCSIExtent

> ISCSIExtent GetISCSIExtent(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



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
    id := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IscsiExtentAPI.GetISCSIExtent(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiExtentAPI.GetISCSIExtent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetISCSIExtent`: ISCSIExtent
    fmt.Fprintf(os.Stdout, "Response from `IscsiExtentAPI.GetISCSIExtent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetISCSIExtentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 

### Return type

[**ISCSIExtent**](ISCSIExtent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListISCSIExtent

> []ISCSIExtent ListISCSIExtent(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IscsiExtentAPI.ListISCSIExtent(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiExtentAPI.ListISCSIExtent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListISCSIExtent`: []ISCSIExtent
    fmt.Fprintf(os.Stdout, "Response from `IscsiExtentAPI.ListISCSIExtent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListISCSIExtentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 

### Return type

[**[]ISCSIExtent**](ISCSIExtent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

