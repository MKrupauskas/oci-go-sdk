// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v39/common"
	"net/http"
)

// ListErrataRequest wrapper for the ListErrata operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListErrata.go.html to see an example of how to use ListErrataRequest.
type ListErrataRequest struct {

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID of the erratum.
	ErratumId *string `mandatory:"false" contributesTo:"query" name:"erratumId"`

	// The assigned erratum name. It's unique and not changeable.
	// Example: `ELSA-2020-5804`
	AdvisoryName *string `mandatory:"false" contributesTo:"query" name:"advisoryName"`

	// The issue date after which to list all errata, in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	TimeIssueDateStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIssueDateStart"`

	// The issue date before which to list all errata, in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	TimeIssueDateEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIssueDateEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListErrataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort errata by. Only one sort order may be provided. Default order for ISSUEDATE is descending. Default order for ADVISORYNAME is ascending. If no value is specified ISSUEDATE is default.
	SortBy ListErrataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListErrataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListErrataRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListErrataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListErrataResponse wrapper for the ListErrata operation
type ListErrataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ErratumSummary instances
	Items []ErratumSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListErrataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListErrataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListErrataSortOrderEnum Enum with underlying type: string
type ListErrataSortOrderEnum string

// Set of constants representing the allowable values for ListErrataSortOrderEnum
const (
	ListErrataSortOrderAsc  ListErrataSortOrderEnum = "ASC"
	ListErrataSortOrderDesc ListErrataSortOrderEnum = "DESC"
)

var mappingListErrataSortOrder = map[string]ListErrataSortOrderEnum{
	"ASC":  ListErrataSortOrderAsc,
	"DESC": ListErrataSortOrderDesc,
}

// GetListErrataSortOrderEnumValues Enumerates the set of values for ListErrataSortOrderEnum
func GetListErrataSortOrderEnumValues() []ListErrataSortOrderEnum {
	values := make([]ListErrataSortOrderEnum, 0)
	for _, v := range mappingListErrataSortOrder {
		values = append(values, v)
	}
	return values
}

// ListErrataSortByEnum Enum with underlying type: string
type ListErrataSortByEnum string

// Set of constants representing the allowable values for ListErrataSortByEnum
const (
	ListErrataSortByIssuedate    ListErrataSortByEnum = "ISSUEDATE"
	ListErrataSortByAdvisoryname ListErrataSortByEnum = "ADVISORYNAME"
)

var mappingListErrataSortBy = map[string]ListErrataSortByEnum{
	"ISSUEDATE":    ListErrataSortByIssuedate,
	"ADVISORYNAME": ListErrataSortByAdvisoryname,
}

// GetListErrataSortByEnumValues Enumerates the set of values for ListErrataSortByEnum
func GetListErrataSortByEnumValues() []ListErrataSortByEnum {
	values := make([]ListErrataSortByEnum, 0)
	for _, v := range mappingListErrataSortBy {
		values = append(values, v)
	}
	return values
}
