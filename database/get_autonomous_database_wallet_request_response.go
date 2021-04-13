// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v39/common"
	"net/http"
)

// GetAutonomousDatabaseWalletRequest wrapper for the GetAutonomousDatabaseWallet operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetAutonomousDatabaseWallet.go.html to see an example of how to use GetAutonomousDatabaseWalletRequest.
type GetAutonomousDatabaseWalletRequest struct {

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousDatabaseId *string `mandatory:"true" contributesTo:"path" name:"autonomousDatabaseId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAutonomousDatabaseWalletRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAutonomousDatabaseWalletRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAutonomousDatabaseWalletRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAutonomousDatabaseWalletResponse wrapper for the GetAutonomousDatabaseWallet operation
type GetAutonomousDatabaseWalletResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AutonomousDatabaseWallet instance
	AutonomousDatabaseWallet `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAutonomousDatabaseWalletResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAutonomousDatabaseWalletResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
