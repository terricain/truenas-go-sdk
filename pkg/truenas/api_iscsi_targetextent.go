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
	"io"
	"net/http"
	"net/url"
	"strings"
)

// IscsiTargetextentAPIService IscsiTargetextentAPI service
type IscsiTargetextentAPIService service

type ApiCreateISCSITargetExtentRequest struct {
	ctx                           context.Context
	ApiService                    *IscsiTargetextentAPIService
	createISCSITargetExtentParams *CreateISCSITargetExtentParams
}

func (r ApiCreateISCSITargetExtentRequest) CreateISCSITargetExtentParams(createISCSITargetExtentParams CreateISCSITargetExtentParams) ApiCreateISCSITargetExtentRequest {
	r.createISCSITargetExtentParams = &createISCSITargetExtentParams
	return r
}

func (r ApiCreateISCSITargetExtentRequest) Execute() (*ISCSITargetExtent, *http.Response, error) {
	return r.ApiService.CreateISCSITargetExtentExecute(r)
}

/*
CreateISCSITargetExtent Method for CreateISCSITargetExtent

Create an Associated Target.

`lunid` will be automatically assigned if it is not provided based on the `target`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateISCSITargetExtentRequest
*/
func (a *IscsiTargetextentAPIService) CreateISCSITargetExtent(ctx context.Context) ApiCreateISCSITargetExtentRequest {
	return ApiCreateISCSITargetExtentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ISCSITargetExtent
func (a *IscsiTargetextentAPIService) CreateISCSITargetExtentExecute(r ApiCreateISCSITargetExtentRequest) (*ISCSITargetExtent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ISCSITargetExtent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetextentAPIService.CreateISCSITargetExtent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/targetextent"

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
	localVarPostBody = r.createISCSITargetExtentParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeleteISCSITargetExtentRequest struct {
	ctx        context.Context
	ApiService *IscsiTargetextentAPIService
	id         int32
	body       *bool
}

func (r ApiDeleteISCSITargetExtentRequest) Body(body bool) ApiDeleteISCSITargetExtentRequest {
	r.body = &body
	return r
}

func (r ApiDeleteISCSITargetExtentRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteISCSITargetExtentExecute(r)
}

/*
DeleteISCSITargetExtent Method for DeleteISCSITargetExtent

Delete Associated Target of `id`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiDeleteISCSITargetExtentRequest
*/
func (a *IscsiTargetextentAPIService) DeleteISCSITargetExtent(ctx context.Context, id int32) ApiDeleteISCSITargetExtentRequest {
	return ApiDeleteISCSITargetExtentRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IscsiTargetextentAPIService) DeleteISCSITargetExtentExecute(r ApiDeleteISCSITargetExtentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetextentAPIService.DeleteISCSITargetExtent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/targetextent/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetISCSITargetExtentRequest struct {
	ctx        context.Context
	ApiService *IscsiTargetextentAPIService
	id         int32
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r ApiGetISCSITargetExtentRequest) Limit(limit int32) ApiGetISCSITargetExtentRequest {
	r.limit = &limit
	return r
}

func (r ApiGetISCSITargetExtentRequest) Offset(offset int32) ApiGetISCSITargetExtentRequest {
	r.offset = &offset
	return r
}

func (r ApiGetISCSITargetExtentRequest) Count(count bool) ApiGetISCSITargetExtentRequest {
	r.count = &count
	return r
}

func (r ApiGetISCSITargetExtentRequest) Sort(sort string) ApiGetISCSITargetExtentRequest {
	r.sort = &sort
	return r
}

func (r ApiGetISCSITargetExtentRequest) Execute() (*ISCSITargetExtent, *http.Response, error) {
	return r.ApiService.GetISCSITargetExtentExecute(r)
}

/*
GetISCSITargetExtent Method for GetISCSITargetExtent

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiGetISCSITargetExtentRequest
*/
func (a *IscsiTargetextentAPIService) GetISCSITargetExtent(ctx context.Context, id int32) ApiGetISCSITargetExtentRequest {
	return ApiGetISCSITargetExtentRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ISCSITargetExtent
func (a *IscsiTargetextentAPIService) GetISCSITargetExtentExecute(r ApiGetISCSITargetExtentRequest) (*ISCSITargetExtent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ISCSITargetExtent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetextentAPIService.GetISCSITargetExtent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/targetextent/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListISCSITargetExtentRequest struct {
	ctx        context.Context
	ApiService *IscsiTargetextentAPIService
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r ApiListISCSITargetExtentRequest) Limit(limit int32) ApiListISCSITargetExtentRequest {
	r.limit = &limit
	return r
}

func (r ApiListISCSITargetExtentRequest) Offset(offset int32) ApiListISCSITargetExtentRequest {
	r.offset = &offset
	return r
}

func (r ApiListISCSITargetExtentRequest) Count(count bool) ApiListISCSITargetExtentRequest {
	r.count = &count
	return r
}

func (r ApiListISCSITargetExtentRequest) Sort(sort string) ApiListISCSITargetExtentRequest {
	r.sort = &sort
	return r
}

func (r ApiListISCSITargetExtentRequest) Execute() ([]ISCSITargetExtent, *http.Response, error) {
	return r.ApiService.ListISCSITargetExtentExecute(r)
}

/*
ListISCSITargetExtent Method for ListISCSITargetExtent

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListISCSITargetExtentRequest
*/
func (a *IscsiTargetextentAPIService) ListISCSITargetExtent(ctx context.Context) ApiListISCSITargetExtentRequest {
	return ApiListISCSITargetExtentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ISCSITargetExtent
func (a *IscsiTargetextentAPIService) ListISCSITargetExtentExecute(r ApiListISCSITargetExtentRequest) ([]ISCSITargetExtent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ISCSITargetExtent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetextentAPIService.ListISCSITargetExtent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/targetextent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
