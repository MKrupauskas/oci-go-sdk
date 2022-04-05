// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NotificationsTargetDetails The topic used for the Notifications target.
// For configuration instructions, see
// To create a service connector (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/managingconnectors.htm#create).
type NotificationsTargetDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
	TopicId *string `mandatory:"true" json:"topicId"`

	// Whether to apply a simplified, user-friendly format to the message. Applies only when friendly formatting is supported by the service connector source and the subscription protocol.
	// Example: `true`
	EnableFormattedMessaging *bool `mandatory:"false" json:"enableFormattedMessaging"`
}

func (m NotificationsTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotificationsTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NotificationsTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNotificationsTargetDetails NotificationsTargetDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeNotificationsTargetDetails
	}{
		"notifications",
		(MarshalTypeNotificationsTargetDetails)(m),
	}

	return json.Marshal(&s)
}
