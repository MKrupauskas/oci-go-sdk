// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v39/common"
	"net/http"
)

// ListBackendsRequest wrapper for the ListBackends operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkloadbalancer/ListBackends.go.html to see an example of how to use ListBackendsRequest.
type ListBackendsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"networkLoadBalancerId"`

	// The name of the backend set associated with the backend servers.
	// Example: `example_backend_set`
	BackendSetName *string `mandatory:"true" contributesTo:"path" name:"backendSetName"`

	// The unique Oracle-assigned identifier for the request. If you must contact Oracle about a
	// particular request, then provide the request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The system returns the requested resource, with a 200 status, only if the resource has no etag
	// matching the one specified. If the condition fails for the GET and HEAD methods, then the system returns the
	// HTTP status code `304 (Not Modified)`.
	// Example: `example-etag`
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// For list pagination. The maximum number of results per page or items to return, in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from which to start retrieving results.
	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' (ascending) or 'desc' (descending).
	SortOrder ListBackendsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBackendsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBackendsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBackendsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBackendsResponse wrapper for the ListBackends operation
type ListBackendsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BackendCollection instances
	BackendCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBackendsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBackendsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBackendsSortOrderEnum Enum with underlying type: string
type ListBackendsSortOrderEnum string

// Set of constants representing the allowable values for ListBackendsSortOrderEnum
const (
	ListBackendsSortOrderAsc  ListBackendsSortOrderEnum = "ASC"
	ListBackendsSortOrderDesc ListBackendsSortOrderEnum = "DESC"
)

var mappingListBackendsSortOrder = map[string]ListBackendsSortOrderEnum{
	"ASC":  ListBackendsSortOrderAsc,
	"DESC": ListBackendsSortOrderDesc,
}

// GetListBackendsSortOrderEnumValues Enumerates the set of values for ListBackendsSortOrderEnum
func GetListBackendsSortOrderEnumValues() []ListBackendsSortOrderEnum {
	values := make([]ListBackendsSortOrderEnum, 0)
	for _, v := range mappingListBackendsSortOrder {
		values = append(values, v)
	}
	return values
}
