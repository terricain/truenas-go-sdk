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

// IscsiInitiatorAPIService IscsiInitiatorAPI service
type IscsiInitiatorAPIService service

type ApiCreateISCSIInitiatorRequest struct {
	ctx                        context.Context
	ApiService                 *IscsiInitiatorAPIService
	createISCSIInitiatorParams *CreateISCSIInitiatorParams
}

// Created iSCSI Initiators
func (r ApiCreateISCSIInitiatorRequest) CreateISCSIInitiatorParams(createISCSIInitiatorParams CreateISCSIInitiatorParams) ApiCreateISCSIInitiatorRequest {
	r.createISCSIInitiatorParams = &createISCSIInitiatorParams
	return r
}

func (r ApiCreateISCSIInitiatorRequest) Execute() (*ISCSIInitiator, *http.Response, error) {
	return r.ApiService.CreateISCSIInitiatorExecute(r)
}

/*
CreateISCSIInitiator Method for CreateISCSIInitiator

Create an iSCSI Initiator.

`initiators` is a list of initiator hostnames which are authorized to access an iSCSI Target. To allow all
possible initiators, `initiators` can be left empty.

`auth_network` is a list of IP/CIDR addresses which are allowed to use this initiator. If all networks are
to be allowed, this field should be left empty.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateISCSIInitiatorRequest
*/
func (a *IscsiInitiatorAPIService) CreateISCSIInitiator(ctx context.Context) ApiCreateISCSIInitiatorRequest {
	return ApiCreateISCSIInitiatorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ISCSIInitiator
func (a *IscsiInitiatorAPIService) CreateISCSIInitiatorExecute(r ApiCreateISCSIInitiatorRequest) (*ISCSIInitiator, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ISCSIInitiator
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiInitiatorAPIService.CreateISCSIInitiator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/initiator"

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
	localVarPostBody = r.createISCSIInitiatorParams
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

type ApiDeleteISCSIInitiatorRequest struct {
	ctx        context.Context
	ApiService *IscsiInitiatorAPIService
	id         int32
}

func (r ApiDeleteISCSIInitiatorRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteISCSIInitiatorExecute(r)
}

/*
DeleteISCSIInitiator Method for DeleteISCSIInitiator

Delete iSCSI initiator of `id`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiDeleteISCSIInitiatorRequest
*/
func (a *IscsiInitiatorAPIService) DeleteISCSIInitiator(ctx context.Context, id int32) ApiDeleteISCSIInitiatorRequest {
	return ApiDeleteISCSIInitiatorRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IscsiInitiatorAPIService) DeleteISCSIInitiatorExecute(r ApiDeleteISCSIInitiatorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiInitiatorAPIService.DeleteISCSIInitiator")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/initiator/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetISCSIInitiatorRequest struct {
	ctx        context.Context
	ApiService *IscsiInitiatorAPIService
	id         int32
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r ApiGetISCSIInitiatorRequest) Limit(limit int32) ApiGetISCSIInitiatorRequest {
	r.limit = &limit
	return r
}

func (r ApiGetISCSIInitiatorRequest) Offset(offset int32) ApiGetISCSIInitiatorRequest {
	r.offset = &offset
	return r
}

func (r ApiGetISCSIInitiatorRequest) Count(count bool) ApiGetISCSIInitiatorRequest {
	r.count = &count
	return r
}

func (r ApiGetISCSIInitiatorRequest) Sort(sort string) ApiGetISCSIInitiatorRequest {
	r.sort = &sort
	return r
}

func (r ApiGetISCSIInitiatorRequest) Execute() (*ISCSIInitiator, *http.Response, error) {
	return r.ApiService.GetISCSIInitiatorExecute(r)
}

/*
GetISCSIInitiator Method for GetISCSIInitiator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiGetISCSIInitiatorRequest
*/
func (a *IscsiInitiatorAPIService) GetISCSIInitiator(ctx context.Context, id int32) ApiGetISCSIInitiatorRequest {
	return ApiGetISCSIInitiatorRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ISCSIInitiator
func (a *IscsiInitiatorAPIService) GetISCSIInitiatorExecute(r ApiGetISCSIInitiatorRequest) (*ISCSIInitiator, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ISCSIInitiator
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiInitiatorAPIService.GetISCSIInitiator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/initiator/id/{id}"
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

type ApiListISCSIInitiatorRequest struct {
	ctx        context.Context
	ApiService *IscsiInitiatorAPIService
	limit      *int32
	offset     *int32
	count      *bool
	sort       *string
}

func (r ApiListISCSIInitiatorRequest) Limit(limit int32) ApiListISCSIInitiatorRequest {
	r.limit = &limit
	return r
}

func (r ApiListISCSIInitiatorRequest) Offset(offset int32) ApiListISCSIInitiatorRequest {
	r.offset = &offset
	return r
}

func (r ApiListISCSIInitiatorRequest) Count(count bool) ApiListISCSIInitiatorRequest {
	r.count = &count
	return r
}

func (r ApiListISCSIInitiatorRequest) Sort(sort string) ApiListISCSIInitiatorRequest {
	r.sort = &sort
	return r
}

func (r ApiListISCSIInitiatorRequest) Execute() ([]ISCSIInitiator, *http.Response, error) {
	return r.ApiService.ListISCSIInitiatorExecute(r)
}

/*
ListISCSIInitiator Method for ListISCSIInitiator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListISCSIInitiatorRequest
*/
func (a *IscsiInitiatorAPIService) ListISCSIInitiator(ctx context.Context) ApiListISCSIInitiatorRequest {
	return ApiListISCSIInitiatorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ISCSIInitiator
func (a *IscsiInitiatorAPIService) ListISCSIInitiatorExecute(r ApiListISCSIInitiatorRequest) ([]ISCSIInitiator, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ISCSIInitiator
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiInitiatorAPIService.ListISCSIInitiator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/initiator"

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
