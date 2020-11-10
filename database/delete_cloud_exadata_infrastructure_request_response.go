// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"net/http"
)

// DeleteCloudExadataInfrastructureRequest wrapper for the DeleteCloudExadataInfrastructure operation
type DeleteCloudExadataInfrastructureRequest struct {

	// The cloud Exadata infrastructure OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CloudExadataInfrastructureId *string `mandatory:"true" contributesTo:"path" name:"cloudExadataInfrastructureId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// If `true`, forces the deletion the specified cloud Exadata infrastructure resource as well as all associated VM clusters. If `false`, the cloud Exadata infrastructure resource can be deleted only if it has no associated VM clusters. Default value is `false`.
	IsDeleteVmClusters *bool `mandatory:"false" contributesTo:"query" name:"isDeleteVmClusters"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteCloudExadataInfrastructureRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteCloudExadataInfrastructureRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteCloudExadataInfrastructureRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteCloudExadataInfrastructureResponse wrapper for the DeleteCloudExadataInfrastructure operation
type DeleteCloudExadataInfrastructureResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request. Multiple OCID values are returned in a comma-separated list. Use GetWorkRequest with a work request OCID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteCloudExadataInfrastructureResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteCloudExadataInfrastructureResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
