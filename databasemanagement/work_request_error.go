// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// WorkRequestError An error encountered while executing a work request.
type WorkRequestError struct {

	// The identifier of the work request erorr.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the work request.
	WorkRequestId *string `mandatory:"true" json:"workRequestId"`

	// A machine-usable code for the error that occured. Error codes are listed on
	// (https://docs.cloud.oracle.com/Content/API/References/apierrors.htm)
	Code *string `mandatory:"true" json:"code"`

	// A human readable description of the issue encountered.
	Message *string `mandatory:"true" json:"message"`

	// The time the error occured.
	// An RFC3339 formatted datetime string. The precision for the time object is milliseconds.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// To determine the work request error is retryable or not
	IsRetryable *bool `mandatory:"false" json:"isRetryable"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}
