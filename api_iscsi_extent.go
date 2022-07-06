/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// IscsiExtentApiService IscsiExtentApi service
type IscsiExtentApiService service

type ApiCreateISCSIExtentRequest struct {
	ctx context.Context
	ApiService *IscsiExtentApiService
	iSCSIExtent *ISCSIExtent
}

func (r ApiCreateISCSIExtentRequest) ISCSIExtent(iSCSIExtent ISCSIExtent) ApiCreateISCSIExtentRequest {
	r.iSCSIExtent = &iSCSIExtent
	return r
}

func (r ApiCreateISCSIExtentRequest) Execute() (*CreateISCSIExtentParams, *http.Response, error) {
	return r.ApiService.CreateISCSIExtentExecute(r)
}

/*
CreateISCSIExtent Method for CreateISCSIExtent

Create an iSCSI Extent.

When `type` is set to FILE, attribute `filesize` is used and it represents number of bytes. `filesize` if
not zero should be a multiple of `blocksize`. `path` is a required attribute with `type` set as FILE and it
should be ensured that it does not come under a jail root.

With `type` being set to DISK, a valid ZVOL or DISK should be provided.

`insecure_tpc` when enabled allows an initiator to bypass normal access control and access any scannable
target. This allows xcopy operations otherwise blocked by access control.

`xen` is a boolean value which is set to true if Xen is being used as the iSCSI initiator.

`ro` when set to true prevents the initiator from writing to this LUN.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateISCSIExtentRequest
*/
func (a *IscsiExtentApiService) CreateISCSIExtent(ctx context.Context) ApiCreateISCSIExtentRequest {
	return ApiCreateISCSIExtentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateISCSIExtentParams
func (a *IscsiExtentApiService) CreateISCSIExtentExecute(r ApiCreateISCSIExtentRequest) (*CreateISCSIExtentParams, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateISCSIExtentParams
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiExtentApiService.CreateISCSIExtent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/extent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.iSCSIExtent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteISCSIExtentRequest struct {
	ctx context.Context
	ApiService *IscsiExtentApiService
	id int32
	deleteISCSIExtentParams *DeleteISCSIExtentParams
}

func (r ApiDeleteISCSIExtentRequest) DeleteISCSIExtentParams(deleteISCSIExtentParams DeleteISCSIExtentParams) ApiDeleteISCSIExtentRequest {
	r.deleteISCSIExtentParams = &deleteISCSIExtentParams
	return r
}

func (r ApiDeleteISCSIExtentRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteISCSIExtentExecute(r)
}

/*
DeleteISCSIExtent Method for DeleteISCSIExtent

Delete iSCSI Extent of `id`.

If `id` iSCSI Extent's `type` was configured to FILE, `remove` can be set to remove the configured file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiDeleteISCSIExtentRequest
*/
func (a *IscsiExtentApiService) DeleteISCSIExtent(ctx context.Context, id int32) ApiDeleteISCSIExtentRequest {
	return ApiDeleteISCSIExtentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *IscsiExtentApiService) DeleteISCSIExtentExecute(r ApiDeleteISCSIExtentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiExtentApiService.DeleteISCSIExtent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/extent/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.deleteISCSIExtentParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetISCSIExtentRequest struct {
	ctx context.Context
	ApiService *IscsiExtentApiService
	id int32
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiGetISCSIExtentRequest) Limit(limit int32) ApiGetISCSIExtentRequest {
	r.limit = &limit
	return r
}

func (r ApiGetISCSIExtentRequest) Offset(offset int32) ApiGetISCSIExtentRequest {
	r.offset = &offset
	return r
}

func (r ApiGetISCSIExtentRequest) Count(count bool) ApiGetISCSIExtentRequest {
	r.count = &count
	return r
}

func (r ApiGetISCSIExtentRequest) Sort(sort string) ApiGetISCSIExtentRequest {
	r.sort = &sort
	return r
}

func (r ApiGetISCSIExtentRequest) Execute() (*ISCSIExtent, *http.Response, error) {
	return r.ApiService.GetISCSIExtentExecute(r)
}

/*
GetISCSIExtent Method for GetISCSIExtent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetISCSIExtentRequest
*/
func (a *IscsiExtentApiService) GetISCSIExtent(ctx context.Context, id int32) ApiGetISCSIExtentRequest {
	return ApiGetISCSIExtentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ISCSIExtent
func (a *IscsiExtentApiService) GetISCSIExtentExecute(r ApiGetISCSIExtentRequest) (*ISCSIExtent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ISCSIExtent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiExtentApiService.GetISCSIExtent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/extent/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetISCSIExtentsRequest struct {
	ctx context.Context
	ApiService *IscsiExtentApiService
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiGetISCSIExtentsRequest) Limit(limit int32) ApiGetISCSIExtentsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetISCSIExtentsRequest) Offset(offset int32) ApiGetISCSIExtentsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetISCSIExtentsRequest) Count(count bool) ApiGetISCSIExtentsRequest {
	r.count = &count
	return r
}

func (r ApiGetISCSIExtentsRequest) Sort(sort string) ApiGetISCSIExtentsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetISCSIExtentsRequest) Execute() ([]ISCSIExtent, *http.Response, error) {
	return r.ApiService.GetISCSIExtentsExecute(r)
}

/*
GetISCSIExtents Method for GetISCSIExtents

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetISCSIExtentsRequest
*/
func (a *IscsiExtentApiService) GetISCSIExtents(ctx context.Context) ApiGetISCSIExtentsRequest {
	return ApiGetISCSIExtentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ISCSIExtent
func (a *IscsiExtentApiService) GetISCSIExtentsExecute(r ApiGetISCSIExtentsRequest) ([]ISCSIExtent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ISCSIExtent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiExtentApiService.GetISCSIExtents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/extent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
