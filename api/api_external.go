/*
@ledge/external-api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgeapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ExternalAPIService ExternalAPI service
type ExternalAPIService service

type ApiRegisterUserRequest struct {
	ctx context.Context
	ApiService *ExternalAPIService
	apiKey *string
	externalUser *ExternalUser
}

func (r ApiRegisterUserRequest) ApiKey(apiKey string) ApiRegisterUserRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiRegisterUserRequest) ExternalUser(externalUser ExternalUser) ApiRegisterUserRequest {
	r.externalUser = &externalUser
	return r
}

func (r ApiRegisterUserRequest) Execute() (*RegisterUser200Response, *http.Response, error) {
	return r.ApiService.RegisterUserExecute(r)
}

/*
RegisterUser Method for RegisterUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRegisterUserRequest
*/
func (a *ExternalAPIService) RegisterUser(ctx context.Context) ApiRegisterUserRequest {
	return ApiRegisterUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RegisterUser200Response
func (a *ExternalAPIService) RegisterUserExecute(r ApiRegisterUserRequest) (*RegisterUser200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegisterUser200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAPIService.RegisterUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.externalUser == nil {
		return localVarReturnValue, nil, reportError("externalUser is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "ApiKey", r.apiKey, "")
	// body params
	localVarPostBody = r.externalUser
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiTrackActivityRequest struct {
	ctx context.Context
	ApiService *ExternalAPIService
	apiKey *string
	trackActivityInput *TrackActivityInput
}

func (r ApiTrackActivityRequest) ApiKey(apiKey string) ApiTrackActivityRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiTrackActivityRequest) TrackActivityInput(trackActivityInput TrackActivityInput) ApiTrackActivityRequest {
	r.trackActivityInput = &trackActivityInput
	return r
}

func (r ApiTrackActivityRequest) Execute() (*TrackActivity200Response, *http.Response, error) {
	return r.ApiService.TrackActivityExecute(r)
}

/*
TrackActivity Method for TrackActivity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTrackActivityRequest
*/
func (a *ExternalAPIService) TrackActivity(ctx context.Context) ApiTrackActivityRequest {
	return ApiTrackActivityRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TrackActivity200Response
func (a *ExternalAPIService) TrackActivityExecute(r ApiTrackActivityRequest) (*TrackActivity200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TrackActivity200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAPIService.TrackActivity")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/activities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.trackActivityInput == nil {
		return localVarReturnValue, nil, reportError("trackActivityInput is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "ApiKey", r.apiKey, "")
	// body params
	localVarPostBody = r.trackActivityInput
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
