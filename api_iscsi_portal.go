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


// IscsiPortalApiService IscsiPortalApi service
type IscsiPortalApiService service

type ApiGetISCSIPortalRequest struct {
	ctx context.Context
	ApiService *IscsiPortalApiService
	id int32
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiGetISCSIPortalRequest) Limit(limit int32) ApiGetISCSIPortalRequest {
	r.limit = &limit
	return r
}

func (r ApiGetISCSIPortalRequest) Offset(offset int32) ApiGetISCSIPortalRequest {
	r.offset = &offset
	return r
}

func (r ApiGetISCSIPortalRequest) Count(count bool) ApiGetISCSIPortalRequest {
	r.count = &count
	return r
}

func (r ApiGetISCSIPortalRequest) Sort(sort string) ApiGetISCSIPortalRequest {
	r.sort = &sort
	return r
}

func (r ApiGetISCSIPortalRequest) Execute() (*ISCSIPortal, *http.Response, error) {
	return r.ApiService.GetISCSIPortalExecute(r)
}

/*
GetISCSIPortal Method for GetISCSIPortal

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetISCSIPortalRequest
*/
func (a *IscsiPortalApiService) GetISCSIPortal(ctx context.Context, id int32) ApiGetISCSIPortalRequest {
	return ApiGetISCSIPortalRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ISCSIPortal
func (a *IscsiPortalApiService) GetISCSIPortalExecute(r ApiGetISCSIPortalRequest) (*ISCSIPortal, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ISCSIPortal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiPortalApiService.GetISCSIPortal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/portal/id/{id}"
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

type ApiListISCSIPortalRequest struct {
	ctx context.Context
	ApiService *IscsiPortalApiService
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiListISCSIPortalRequest) Limit(limit int32) ApiListISCSIPortalRequest {
	r.limit = &limit
	return r
}

func (r ApiListISCSIPortalRequest) Offset(offset int32) ApiListISCSIPortalRequest {
	r.offset = &offset
	return r
}

func (r ApiListISCSIPortalRequest) Count(count bool) ApiListISCSIPortalRequest {
	r.count = &count
	return r
}

func (r ApiListISCSIPortalRequest) Sort(sort string) ApiListISCSIPortalRequest {
	r.sort = &sort
	return r
}

func (r ApiListISCSIPortalRequest) Execute() ([]ISCSIPortal, *http.Response, error) {
	return r.ApiService.ListISCSIPortalExecute(r)
}

/*
ListISCSIPortal Method for ListISCSIPortal

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListISCSIPortalRequest
*/
func (a *IscsiPortalApiService) ListISCSIPortal(ctx context.Context) ApiListISCSIPortalRequest {
	return ApiListISCSIPortalRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ISCSIPortal
func (a *IscsiPortalApiService) ListISCSIPortalExecute(r ApiListISCSIPortalRequest) ([]ISCSIPortal, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ISCSIPortal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiPortalApiService.ListISCSIPortal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iscsi/portal"

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