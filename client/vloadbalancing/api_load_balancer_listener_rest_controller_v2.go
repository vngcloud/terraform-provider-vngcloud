/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancing

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type LoadBalancerListenerRestControllerV2ApiService service

/*
LoadBalancerListenerRestControllerV2ApiService Create a listener
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param createListenerRequestV2 createListenerRequestV2
 * @param loadBalancerId The Load Balancer id
 * @param projectId The project id

@return map[string]string
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) CreateListenerUsingPOST(ctx context.Context, createListenerRequestV2 CreateListenerRequestV2, loadBalancerId string, projectId string) (map[string]string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue map[string]string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners"
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &createListenerRequestV2
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Create a l7 policy
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param createL7PolicyRequestV2 createL7PolicyRequestV2
 * @param loadBalancerId The Load Balancer id
 * @param listenerId The listener id
 * @param projectId The project id

@return map[string]string
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) CreatePolicyUsingPOST(ctx context.Context, createL7PolicyRequestV2 CreateL7PolicyRequestV2, loadBalancerId string, listenerId string, projectId string) (map[string]string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue map[string]string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies"
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listenerId"+"}", fmt.Sprintf("%v", listenerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &createL7PolicyRequestV2
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Delete a listener
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param listenerId listener id
 * @param projectId project id
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) DeleteListenerUsingDELETE(ctx context.Context, loadbalancerId string, listenerId string, projectId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadbalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listenerId"+"}", fmt.Sprintf("%v", listenerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Delete a policy
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loadBalancerId load balancer id
 * @param listenerId listener id
 * @param policyId policy id
 * @param projectId project id
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) DeletePolicyUsingDELETE(ctx context.Context, loadBalancerId string, listenerId string, policyId string, projectId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies/{l7PolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"listenerId"+"}", fmt.Sprintf("%v", listenerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"l7PolicyId"+"}", fmt.Sprintf("%v", policyId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Get a specific policy of a listener by l7PolicyId
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param l7PolicyId The l7Policy id
 * @param listenerId The listener id
 * @param loadBalancerId The Load Balancer id
 * @param projectId The project id
@return DataResponseL7Policy
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) GetL7PolicyUsingGET(ctx context.Context, l7PolicyId string, listenerId string, loadBalancerId string, projectId string) (DataResponseL7Policy, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue DataResponseL7Policy
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies/{l7PolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"l7PolicyId"+"}", fmt.Sprintf("%v", l7PolicyId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listenerId"+"}", fmt.Sprintf("%v", listenerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v DataResponseL7Policy
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Get list policies of a listener
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param listenerId The listener id
 * @param loadBalancerId The Load Balancer id
 * @param projectId The project id
@return DataResponseListL7Policy
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) GetListL7PoliciesUsingGET(ctx context.Context, listenerId string, loadBalancerId string, projectId string) (DataResponseListL7Policy, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue DataResponseListL7Policy
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies"
	localVarPath = strings.Replace(localVarPath, "{"+"listenerId"+"}", fmt.Sprintf("%v", listenerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v DataResponseListL7Policy
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Get list of listeners of a Load Balancer
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loadBalancerId The Load Balancer id
 * @param projectId The project id
@return DataResponseListListener
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) GetListListenersByLoadBalancerUsingGET(ctx context.Context, loadBalancerId string, projectId string) (DataResponseListListener, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue DataResponseListListener
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners"
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v DataResponseListListener
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Get a specific listener
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param listenerId The listener id
 * @param loadBalancerId The Load Balancer id
 * @param projectId The project id
@return DataResponseListener
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) GetListenersUsingGET(ctx context.Context, listenerId string, loadBalancerId string, projectId string) (DataResponseListener, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue DataResponseListener
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"listenerId"+"}", fmt.Sprintf("%v", listenerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v DataResponseListener
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Reorder the L7 policies
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param listenerId listener id
 * @param projectId project id
 * @param reorderPoliciesRequestV2 reorderPoliciesRequestV2
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) ReorderPoliciesUsingPUT(ctx context.Context, listenerId string, projectId string, reorderPoliciesRequestV2 ReorderPoliciesRequestV2) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/reorderL7Policies"
	localVarPath = strings.Replace(localVarPath, "{"+"listenerId"+"}", fmt.Sprintf("%v", listenerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &reorderPoliciesRequestV2
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Update a policy
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loadBalancerId The load balancer id
 * @param listenerId The listener id
 * @param policyId policy id
 * @param projectId project id
 * @param updateL7PolicyRequestV2 updateL7PolicyRequestV2
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) UpdateL7PolicyUsingPUT(ctx context.Context, loadBalancerId string, listenerId string, policyId string, projectId string, updateL7PolicyRequestV2 UpdateL7PolicyRequestV2) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}/l7policies/{policyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listenerId"+"}", fmt.Sprintf("%v", listenerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", fmt.Sprintf("%v", policyId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &updateL7PolicyRequestV2
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
LoadBalancerListenerRestControllerV2ApiService Update a listener
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param listenerId The listener id
 * @param loadBalancerId The Load Balancer id
 * @param projectId The project id
 * @param updateListenerRequestV2 updateListenerRequestV2
*/
func (a *LoadBalancerListenerRestControllerV2ApiService) UpdateListenerUsingPUT(ctx context.Context, listenerId string, loadBalancerId string, projectId string, updateListenerRequestV2 UpdateListenerRequestV2) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/listeners/{listenerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"listenerId"+"}", fmt.Sprintf("%v", listenerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &updateListenerRequestV2
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
LoadBalancerRestControllerV2ApiService Resize a load balancer
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param loadBalancerId The Load Balancer id
 * @param projectId The project id
 * @param request request


*/
func (a *LoadBalancerRestControllerV2ApiService) ResizeLoadBalancerUsingPUT(ctx context.Context, loadBalancerId string, projectId string, request ResizeLoadBalancerRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/{projectId}/loadBalancers/{loadBalancerId}/resize"
	localVarPath = strings.Replace(localVarPath, "{"+"loadBalancerId"+"}", fmt.Sprintf("%v", loadBalancerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	// body params
	localVarPostBody = &request
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
