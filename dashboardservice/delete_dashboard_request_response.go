// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dashboardservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// DeleteDashboardRequest wrapper for the DeleteDashboard operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/DeleteDashboard.go.html to see an example of how to use DeleteDashboardRequest.
type DeleteDashboardRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard.
	DashboardId *string `mandatory:"true" contributesTo:"path" name:"dashboardId"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// To identify if the call is cross-regional. In CRUD calls for a resource, to
	// identify that the call originates from different region, set the
	// `CrossRegionIdentifierHeader` parameter to a region name (ex - `US-ASHBURN-1`)
	// The call will be served from a Replicated bucket.
	// For same-region calls, the value is unassigned.
	OpcCrossRegion *string `mandatory:"false" contributesTo:"header" name:"opc-cross-region"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteDashboardRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteDashboardRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteDashboardRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteDashboardRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DeleteDashboardRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeleteDashboardResponse wrapper for the DeleteDashboard operation
type DeleteDashboardResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteDashboardResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteDashboardResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
