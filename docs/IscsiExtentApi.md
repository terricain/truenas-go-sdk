# \IscsiExtentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateISCSIExtent**](IscsiExtentApi.md#CreateISCSIExtent) | **Post** /iscsi/extent | 
[**DeleteISCSIExtent**](IscsiExtentApi.md#DeleteISCSIExtent) | **Delete** /iscsi/extent/id/{id} | 
[**GetISCSIExtent**](IscsiExtentApi.md#GetISCSIExtent) | **Get** /iscsi/extent/id/{id} | 
[**GetISCSIExtents**](IscsiExtentApi.md#GetISCSIExtents) | **Get** /iscsi/extent | 



## CreateISCSIExtent

> CreateISCSIExtentParams CreateISCSIExtent(ctx).ISCSIExtent(iSCSIExtent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    iSCSIExtent := *openapiclient.NewISCSIExtent() // ISCSIExtent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IscsiExtentApi.CreateISCSIExtent(context.Background()).ISCSIExtent(iSCSIExtent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiExtentApi.CreateISCSIExtent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateISCSIExtent`: CreateISCSIExtentParams
    fmt.Fprintf(os.Stdout, "Response from `IscsiExtentApi.CreateISCSIExtent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateISCSIExtentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iSCSIExtent** | [**ISCSIExtent**](ISCSIExtent.md) |  | 

### Return type

[**CreateISCSIExtentParams**](CreateISCSIExtentParams.md)

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
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    deleteISCSIExtentParams := *openapiclient.NewDeleteISCSIExtentParams(false, false) // DeleteISCSIExtentParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IscsiExtentApi.DeleteISCSIExtent(context.Background(), id).DeleteISCSIExtentParams(deleteISCSIExtentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiExtentApi.DeleteISCSIExtent``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IscsiExtentApi.GetISCSIExtent(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiExtentApi.GetISCSIExtent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetISCSIExtent`: ISCSIExtent
    fmt.Fprintf(os.Stdout, "Response from `IscsiExtentApi.GetISCSIExtent`: %v\n", resp)
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


## GetISCSIExtents

> []ISCSIExtent GetISCSIExtents(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IscsiExtentApi.GetISCSIExtents(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiExtentApi.GetISCSIExtents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetISCSIExtents`: []ISCSIExtent
    fmt.Fprintf(os.Stdout, "Response from `IscsiExtentApi.GetISCSIExtents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetISCSIExtentsRequest struct via the builder pattern


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

