// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// CreateAnnouncementSubscriptionDetails The details for creating a new announcement subscription.
type CreateAnnouncementSubscriptionDetails struct {

	// A user-friendly name for the announcement subscription. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the announcement subscription.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Notifications service topic that is the target for publishing announcements that match the configured announcement subscription. The caller of the operation needs the ONS_TOPIC_PUBLISH permission for the targeted Notifications service topic. For more information about Notifications permissions, see Details for Notifications (https://docs.cloud.oracle.com/Content/Identity/policyreference/notificationpolicyreference.htm).
	OnsTopicId *string `mandatory:"true" json:"onsTopicId"`

	// A description of the announcement subscription. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A list of filter groups for the announcement subscription. A filter group combines one or more filters that the Announcements service applies to announcements for matching purposes.
	FilterGroups map[string]FilterGroupDetails `mandatory:"false" json:"filterGroups"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAnnouncementSubscriptionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAnnouncementSubscriptionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
