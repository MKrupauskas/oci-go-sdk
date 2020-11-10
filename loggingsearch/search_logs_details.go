// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Search API
//
// Search for logs in your compartments, log groups, and log objects.
//

package loggingsearch

import (
	"github.com/oracle/oci-go-sdk/v28/common"
)

// SearchLogsDetails Search request object.
type SearchLogsDetails struct {

	// Start filter log's date and time, in the format defined by RFC3339.
	TimeStart *common.SDKTime `mandatory:"true" json:"timeStart"`

	// End filter log's date and time, in the format defined by RFC3339.
	TimeEnd *common.SDKTime `mandatory:"true" json:"timeEnd"`

	// Query corresponding to the search operation. This query is parsed and validated before execution and
	// should follow the specification. For more information on the query language specification, see
	// Logging Query Language Specification (https://docs.cloud.oracle.com/Content/Logging/Reference/query_language_specification.htm).
	SearchQuery *string `mandatory:"true" json:"searchQuery"`

	// Whether to return field schema information for the log stream specified in searchQuery.
	IsReturnFieldInfo *bool `mandatory:"false" json:"isReturnFieldInfo"`
}

func (m SearchLogsDetails) String() string {
	return common.PointerString(m)
}
