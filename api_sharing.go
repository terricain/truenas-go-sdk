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


// SharingApiService SharingApi service
type SharingApiService service

type ApiCreateShareNFSRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	createShareNFSParams *CreateShareNFSParams
}

func (r ApiCreateShareNFSRequest) CreateShareNFSParams(createShareNFSParams CreateShareNFSParams) ApiCreateShareNFSRequest {
	r.createShareNFSParams = &createShareNFSParams
	return r
}

func (r ApiCreateShareNFSRequest) Execute() (*ShareNFS, *http.Response, error) {
	return r.ApiService.CreateShareNFSExecute(r)
}

/*
CreateShareNFS Method for CreateShareNFS

Create a new NFS Share.

`paths` is a list of valid paths which are configured to be shared on this share.

`networks` is a list of authorized networks that are allowed to access the share having format
"network/mask" CIDR notation. If empty, all networks are allowed.

`hosts` is a list of IP's/hostnames which are allowed to access the share. If empty, all IP's/hostnames are
allowed.

`alldirs` is a boolean value which when set indicates that the client can mount any subdirectories of the
selected pool or dataset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateShareNFSRequest
*/
func (a *SharingApiService) CreateShareNFS(ctx context.Context) ApiCreateShareNFSRequest {
	return ApiCreateShareNFSRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ShareNFS
func (a *SharingApiService) CreateShareNFSExecute(r ApiCreateShareNFSRequest) (*ShareNFS, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShareNFS
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.CreateShareNFS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/nfs"

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
	localVarPostBody = r.createShareNFSParams
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

type ApiCreateShareSMBRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	createShareSMBParams *CreateShareSMBParams
}

func (r ApiCreateShareSMBRequest) CreateShareSMBParams(createShareSMBParams CreateShareSMBParams) ApiCreateShareSMBRequest {
	r.createShareSMBParams = &createShareSMBParams
	return r
}

func (r ApiCreateShareSMBRequest) Execute() (*ShareSMB, *http.Response, error) {
	return r.ApiService.CreateShareSMBExecute(r)
}

/*
CreateShareSMB Method for CreateShareSMB

Create an SMB Share.

`purpose` applies common configuration presets depending on intended purpose.

`timemachine` when set, enables Time Machine backups for this share.

`ro` when enabled, prohibits write access to the share.

`guestok` when enabled, allows access to this share without a password.

`hostsallow` is a list of hostnames / IP addresses which have access to this share.

`hostsdeny` is a list of hostnames / IP addresses which are not allowed access to this share. If a handful
of hostnames are to be only allowed access, `hostsdeny` can be passed "ALL" which means that it will deny
access to ALL hostnames except for the ones which have been listed in `hostsallow`.

`acl` enables support for storing the SMB Security Descriptor as a Filesystem ACL.

`streams` enables support for storing alternate datastreams as filesystem extended attributes.

`fsrvp` enables support for the filesystem remote VSS protocol. This allows clients to create
ZFS snapshots through RPC.

`shadowcopy` enables support for the volume shadow copy service.

`auxsmbconf` is a string of additional smb4.conf parameters not covered by the system's API.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateShareSMBRequest
*/
func (a *SharingApiService) CreateShareSMB(ctx context.Context) ApiCreateShareSMBRequest {
	return ApiCreateShareSMBRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ShareSMB
func (a *SharingApiService) CreateShareSMBExecute(r ApiCreateShareSMBRequest) (*ShareSMB, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShareSMB
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.CreateShareSMB")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb"

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
	localVarPostBody = r.createShareSMBParams
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

type ApiGetShareNFSRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	id int32
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiGetShareNFSRequest) Limit(limit int32) ApiGetShareNFSRequest {
	r.limit = &limit
	return r
}

func (r ApiGetShareNFSRequest) Offset(offset int32) ApiGetShareNFSRequest {
	r.offset = &offset
	return r
}

func (r ApiGetShareNFSRequest) Count(count bool) ApiGetShareNFSRequest {
	r.count = &count
	return r
}

func (r ApiGetShareNFSRequest) Sort(sort string) ApiGetShareNFSRequest {
	r.sort = &sort
	return r
}

func (r ApiGetShareNFSRequest) Execute() (*ShareNFS, *http.Response, error) {
	return r.ApiService.GetShareNFSExecute(r)
}

/*
GetShareNFS Method for GetShareNFS

Get NFS Share of `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetShareNFSRequest
*/
func (a *SharingApiService) GetShareNFS(ctx context.Context, id int32) ApiGetShareNFSRequest {
	return ApiGetShareNFSRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ShareNFS
func (a *SharingApiService) GetShareNFSExecute(r ApiGetShareNFSRequest) (*ShareNFS, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShareNFS
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.GetShareNFS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/nfs/id/{id}"
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

type ApiGetShareSMBRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	id int32
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiGetShareSMBRequest) Limit(limit int32) ApiGetShareSMBRequest {
	r.limit = &limit
	return r
}

func (r ApiGetShareSMBRequest) Offset(offset int32) ApiGetShareSMBRequest {
	r.offset = &offset
	return r
}

func (r ApiGetShareSMBRequest) Count(count bool) ApiGetShareSMBRequest {
	r.count = &count
	return r
}

func (r ApiGetShareSMBRequest) Sort(sort string) ApiGetShareSMBRequest {
	r.sort = &sort
	return r
}

func (r ApiGetShareSMBRequest) Execute() (*ShareSMB, *http.Response, error) {
	return r.ApiService.GetShareSMBExecute(r)
}

/*
GetShareSMB Method for GetShareSMB

Get SMB Share of `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetShareSMBRequest
*/
func (a *SharingApiService) GetShareSMB(ctx context.Context, id int32) ApiGetShareSMBRequest {
	return ApiGetShareSMBRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ShareSMB
func (a *SharingApiService) GetShareSMBExecute(r ApiGetShareSMBRequest) (*ShareSMB, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShareSMB
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.GetShareSMB")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb/id/{id}"
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

type ApiListSharesNFSRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiListSharesNFSRequest) Limit(limit int32) ApiListSharesNFSRequest {
	r.limit = &limit
	return r
}

func (r ApiListSharesNFSRequest) Offset(offset int32) ApiListSharesNFSRequest {
	r.offset = &offset
	return r
}

func (r ApiListSharesNFSRequest) Count(count bool) ApiListSharesNFSRequest {
	r.count = &count
	return r
}

func (r ApiListSharesNFSRequest) Sort(sort string) ApiListSharesNFSRequest {
	r.sort = &sort
	return r
}

func (r ApiListSharesNFSRequest) Execute() ([]ShareNFS, *http.Response, error) {
	return r.ApiService.ListSharesNFSExecute(r)
}

/*
ListSharesNFS Method for ListSharesNFS

Get a list of NFS shares

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListSharesNFSRequest
*/
func (a *SharingApiService) ListSharesNFS(ctx context.Context) ApiListSharesNFSRequest {
	return ApiListSharesNFSRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ShareNFS
func (a *SharingApiService) ListSharesNFSExecute(r ApiListSharesNFSRequest) ([]ShareNFS, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ShareNFS
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.ListSharesNFS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/nfs"

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

type ApiListSharesSMBRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiListSharesSMBRequest) Limit(limit int32) ApiListSharesSMBRequest {
	r.limit = &limit
	return r
}

func (r ApiListSharesSMBRequest) Offset(offset int32) ApiListSharesSMBRequest {
	r.offset = &offset
	return r
}

func (r ApiListSharesSMBRequest) Count(count bool) ApiListSharesSMBRequest {
	r.count = &count
	return r
}

func (r ApiListSharesSMBRequest) Sort(sort string) ApiListSharesSMBRequest {
	r.sort = &sort
	return r
}

func (r ApiListSharesSMBRequest) Execute() ([]ShareSMB, *http.Response, error) {
	return r.ApiService.ListSharesSMBExecute(r)
}

/*
ListSharesSMB Method for ListSharesSMB

Get a list of SMB shares

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListSharesSMBRequest
*/
func (a *SharingApiService) ListSharesSMB(ctx context.Context) ApiListSharesSMBRequest {
	return ApiListSharesSMBRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ShareSMB
func (a *SharingApiService) ListSharesSMBExecute(r ApiListSharesSMBRequest) ([]ShareSMB, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ShareSMB
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.ListSharesSMB")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb"

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

type ApiRemoveShareNFSRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	id int32
}

func (r ApiRemoveShareNFSRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveShareNFSExecute(r)
}

/*
RemoveShareNFS Method for RemoveShareNFS

Delete NFS Share of `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiRemoveShareNFSRequest
*/
func (a *SharingApiService) RemoveShareNFS(ctx context.Context, id int32) ApiRemoveShareNFSRequest {
	return ApiRemoveShareNFSRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *SharingApiService) RemoveShareNFSExecute(r ApiRemoveShareNFSRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.RemoveShareNFS")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/nfs/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiRemoveShareSMBRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	id int32
}

func (r ApiRemoveShareSMBRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveShareSMBExecute(r)
}

/*
RemoveShareSMB Method for RemoveShareSMB

Delete SMB Share of `id`. This will forcibly disconnect SMB clients
that are accessing the share.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiRemoveShareSMBRequest
*/
func (a *SharingApiService) RemoveShareSMB(ctx context.Context, id int32) ApiRemoveShareSMBRequest {
	return ApiRemoveShareSMBRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *SharingApiService) RemoveShareSMBExecute(r ApiRemoveShareSMBRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.RemoveShareSMB")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiUpdateShareNFSRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	id int32
	createShareNFSParams *CreateShareNFSParams
}

func (r ApiUpdateShareNFSRequest) CreateShareNFSParams(createShareNFSParams CreateShareNFSParams) ApiUpdateShareNFSRequest {
	r.createShareNFSParams = &createShareNFSParams
	return r
}

func (r ApiUpdateShareNFSRequest) Execute() (*ShareNFS, *http.Response, error) {
	return r.ApiService.UpdateShareNFSExecute(r)
}

/*
UpdateShareNFS Method for UpdateShareNFS

Update NFS Share of `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateShareNFSRequest
*/
func (a *SharingApiService) UpdateShareNFS(ctx context.Context, id int32) ApiUpdateShareNFSRequest {
	return ApiUpdateShareNFSRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ShareNFS
func (a *SharingApiService) UpdateShareNFSExecute(r ApiUpdateShareNFSRequest) (*ShareNFS, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShareNFS
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.UpdateShareNFS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/nfs/id/{id}"
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createShareNFSParams
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

type ApiUpdateShareSMBRequest struct {
	ctx context.Context
	ApiService *SharingApiService
	id int32
	createShareSMBParams *CreateShareSMBParams
}

func (r ApiUpdateShareSMBRequest) CreateShareSMBParams(createShareSMBParams CreateShareSMBParams) ApiUpdateShareSMBRequest {
	r.createShareSMBParams = &createShareSMBParams
	return r
}

func (r ApiUpdateShareSMBRequest) Execute() (*ShareSMB, *http.Response, error) {
	return r.ApiService.UpdateShareSMBExecute(r)
}

/*
UpdateShareSMB Method for UpdateShareSMB

Update SMB Share of `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateShareSMBRequest
*/
func (a *SharingApiService) UpdateShareSMB(ctx context.Context, id int32) ApiUpdateShareSMBRequest {
	return ApiUpdateShareSMBRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ShareSMB
func (a *SharingApiService) UpdateShareSMBExecute(r ApiUpdateShareSMBRequest) (*ShareSMB, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShareSMB
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharingApiService.UpdateShareSMB")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sharing/smb/id/{id}"
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createShareSMBParams
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
