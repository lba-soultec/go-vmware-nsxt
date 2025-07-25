/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nsxt

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"github.com/vmware/go-vmware-nsxt/clustercontrolplane"
)

// Linger please
var (
	_ context.Context
)

type SystemAdministrationPolicyClusterControlPlaneApiService service

/*
SystemAdministrationPolicyClusterControlPlaneApiService Create or Update Cluster Control Plane to NSX-T
Joins a Cluster Control Plane to NSX-T.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param siteId
  - @param enforcementpointId
  - @param clusterControlPlaneId

@return ClusterControlPlane
*/
func (a *SystemAdministrationPolicyClusterControlPlaneApiService) CreateOrUpdateClusterControlPlane(ctx context.Context, body clustercontrolplane.ClusterControlPlane, siteId string, enforcementpointId string, clusterControlPlaneId string) (clustercontrolplane.ClusterControlPlane, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue clustercontrolplane.ClusterControlPlane
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/cluster-control-planes/{cluster-control-plane-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", fmt.Sprintf("%v", siteId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"enforcementpoint-id"+"}", fmt.Sprintf("%v", enforcementpointId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cluster-control-plane-id"+"}", fmt.Sprintf("%v", clusterControlPlaneId), -1)

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
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
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

		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SystemAdministrationPolicyClusterControlPlaneApiService Delete a specified cluster control plane.
Removes a specified cluster control plane.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param siteId
 * @param enforcementpointId
 * @param clusterControlPlaneId
 * @param optional nil or *DeleteClusterControlPlaneOpts - Optional Parameters:
     * @param "Cascade" (optional.Bool) -  Flag to indicate if force delete cluster references from the firewall security policies

*/

type DeleteClusterControlPlaneOpts struct {
	Cascade optional.Bool
}

func (a *SystemAdministrationPolicyClusterControlPlaneApiService) DeleteClusterControlPlane(ctx context.Context, siteId string, enforcementpointId string, clusterControlPlaneId string, localVarOptionals *DeleteClusterControlPlaneOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/cluster-control-planes/{cluster-control-plane-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", fmt.Sprintf("%v", siteId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"enforcementpoint-id"+"}", fmt.Sprintf("%v", enforcementpointId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cluster-control-plane-id"+"}", fmt.Sprintf("%v", clusterControlPlaneId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cascade.IsSet() {
		localVarQueryParams.Add("cascade", parameterToString(localVarOptionals.Cascade.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

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
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHttpResponse, newErr
		}
		newErr.model = v
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
SystemAdministrationPolicyClusterControlPlaneApiService Return a specified cluster control plane.
Returns a specified cluster control plane.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param siteId
  - @param enforcementpointId
  - @param clusterControlPlaneId

@return ClusterControlPlane
*/
func (a *SystemAdministrationPolicyClusterControlPlaneApiService) GetClusterControlPlane(ctx context.Context, siteId string, enforcementpointId string, clusterControlPlaneId string) (clustercontrolplane.ClusterControlPlane, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue clustercontrolplane.ClusterControlPlane
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/cluster-control-planes/{cluster-control-plane-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", fmt.Sprintf("%v", siteId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"enforcementpoint-id"+"}", fmt.Sprintf("%v", enforcementpointId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cluster-control-plane-id"+"}", fmt.Sprintf("%v", clusterControlPlaneId), -1)

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
	localVarHttpHeaderAccepts := []string{"application/json"}

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
	localVarHttpResponse.Body.Close()
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

		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SystemAdministrationPolicyClusterControlPlaneApiService Return the List of cluster control planes.
Returns information about all cluster control planes.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param siteId
 * @param enforcementpointId
 * @param optional nil or *ListClusterControlPlaneOpts - Optional Parameters:
     * @param "Cursor" (optional.String) -  Opaque cursor to be used for getting next page of records (supplied by current result page)
     * @param "IncludeMarkForDeleteObjects" (optional.Bool) -  Include objects that are marked for deletion in results
     * @param "IncludedFields" (optional.String) -  Comma separated list of fields that should be included in query result
     * @param "PageSize" (optional.Int64) -  Maximum number of results to return in this page (server may return fewer)
     * @param "SortAscending" (optional.Bool) -
     * @param "SortBy" (optional.String) -  Field by which records are sorted
@return ClusterControlPlaneListResult
*/

type ListClusterControlPlaneOpts struct {
	Cursor                      optional.String
	IncludeMarkForDeleteObjects optional.Bool
	IncludedFields              optional.String
	PageSize                    optional.Int64
	SortAscending               optional.Bool
	SortBy                      optional.String
}

func (a *SystemAdministrationPolicyClusterControlPlaneApiService) ListClusterControlPlane(ctx context.Context, siteId string, enforcementpointId string, localVarOptionals *ListClusterControlPlaneOpts) (clustercontrolplane.ClusterControlPlaneListResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue clustercontrolplane.ClusterControlPlaneListResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/cluster-control-planes"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", fmt.Sprintf("%v", siteId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"enforcementpoint-id"+"}", fmt.Sprintf("%v", enforcementpointId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeMarkForDeleteObjects.IsSet() {
		localVarQueryParams.Add("include_mark_for_delete_objects", parameterToString(localVarOptionals.IncludeMarkForDeleteObjects.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludedFields.IsSet() {
		localVarQueryParams.Add("included_fields", parameterToString(localVarOptionals.IncludedFields.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortAscending.IsSet() {
		localVarQueryParams.Add("sort_ascending", parameterToString(localVarOptionals.SortAscending.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortBy.IsSet() {
		localVarQueryParams.Add("sort_by", parameterToString(localVarOptionals.SortBy.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

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
	localVarHttpResponse.Body.Close()
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

		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SystemAdministrationPolicyClusterControlPlaneApiService Patch a Cluster Control Plane.
Patch a Cluster Control Plane.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param siteId
  - @param enforcementpointId
  - @param clusterControlPlaneId

@return ClusterControlPlane
*/
func (a *SystemAdministrationPolicyClusterControlPlaneApiService) PatchClusterControlPlane(ctx context.Context, body clustercontrolplane.ClusterControlPlane, siteId string, enforcementpointId string, clusterControlPlaneId string) (clustercontrolplane.ClusterControlPlane, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Patch")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue clustercontrolplane.ClusterControlPlane
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/cluster-control-planes/{cluster-control-plane-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", fmt.Sprintf("%v", siteId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"enforcementpoint-id"+"}", fmt.Sprintf("%v", enforcementpointId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cluster-control-plane-id"+"}", fmt.Sprintf("%v", clusterControlPlaneId), -1)

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
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
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

		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SystemAdministrationPolicyClusterControlPlaneApiService Get status of a specified cluster control plane.
Returns the status of a specified cluster control plane including controller, adapter, and agent information.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param siteId
  - @param enforcementpointId
  - @param clusterControlPlaneId

@return ClusterStatus
*/
func (a *SystemAdministrationPolicyClusterControlPlaneApiService) GetClusterControlPlaneStatus(ctx context.Context, siteId string, enforcementpointId string, clusterControlPlaneId string) (clustercontrolplane.ClusterStatus, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue clustercontrolplane.ClusterStatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/infra/sites/{site-id}/enforcement-points/{enforcementpoint-id}/cluster-control-planes/{cluster-control-plane-id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", fmt.Sprintf("%v", siteId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"enforcementpoint-id"+"}", fmt.Sprintf("%v", enforcementpointId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cluster-control-plane-id"+"}", fmt.Sprintf("%v", clusterControlPlaneId), -1)

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
	localVarHttpHeaderAccepts := []string{"application/json"}

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
	localVarHttpResponse.Body.Close()
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

		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
