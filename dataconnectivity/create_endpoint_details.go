// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateEndpointDetails The information about new Endpoint.
type CreateEndpointDetails struct {

	// Data Connectivity Management Registry display name, registries can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// VCN Identifier where the subnet resides.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// Subnet Identifier for customer connected databases
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// List of DNS zones to be used by the data assets to be harvested.
	// Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com
	DnsZones []string `mandatory:"false" json:"dnsZones"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Data Connectivity Management Registry description
	Description *string `mandatory:"false" json:"description"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Endpoint size for reverse connection capacity.
	EndpointSize *int `mandatory:"false" json:"endpointSize"`

	// List of NSGs to which the Private Endpoint VNIC must be added.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m CreateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
