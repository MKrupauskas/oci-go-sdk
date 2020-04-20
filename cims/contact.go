// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Contact Contact Details of the Customer
type Contact struct {

	// Contact person name
	ContactName *string `mandatory:"false" json:"contactName"`

	// Contact person email
	ContactEmail *string `mandatory:"false" json:"contactEmail"`

	// Contact person phone number
	ContactPhone *string `mandatory:"false" json:"contactPhone"`

	// ContactType enum. eg: MANAGER, PRIMARY
	ContactType ContactContactTypeEnum `mandatory:"false" json:"contactType,omitempty"`
}

func (m Contact) String() string {
	return common.PointerString(m)
}

// ContactContactTypeEnum Enum with underlying type: string
type ContactContactTypeEnum string

// Set of constants representing the allowable values for ContactContactTypeEnum
const (
	ContactContactTypePrimary   ContactContactTypeEnum = "PRIMARY"
	ContactContactTypeAlternate ContactContactTypeEnum = "ALTERNATE"
	ContactContactTypeSecondary ContactContactTypeEnum = "SECONDARY"
	ContactContactTypeAdmin     ContactContactTypeEnum = "ADMIN"
	ContactContactTypeManager   ContactContactTypeEnum = "MANAGER"
)

var mappingContactContactType = map[string]ContactContactTypeEnum{
	"PRIMARY":   ContactContactTypePrimary,
	"ALTERNATE": ContactContactTypeAlternate,
	"SECONDARY": ContactContactTypeSecondary,
	"ADMIN":     ContactContactTypeAdmin,
	"MANAGER":   ContactContactTypeManager,
}

// GetContactContactTypeEnumValues Enumerates the set of values for ContactContactTypeEnum
func GetContactContactTypeEnumValues() []ContactContactTypeEnum {
	values := make([]ContactContactTypeEnum, 0)
	for _, v := range mappingContactContactType {
		values = append(values, v)
	}
	return values
}
