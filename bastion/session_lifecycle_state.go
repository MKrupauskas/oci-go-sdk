// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Oracle Cloud Infrastructure Bastion provides restricted and time-limited access to target resources that don't have public endpoints. Through the configuration of a bastion, you can let authorized users connect from specific IP addresses to target resources by way of Secure Shell (SSH) sessions hosted on the bastion.
//

package bastion

// SessionLifecycleStateEnum Enum with underlying type: string
type SessionLifecycleStateEnum string

// Set of constants representing the allowable values for SessionLifecycleStateEnum
const (
	SessionLifecycleStateCreating SessionLifecycleStateEnum = "CREATING"
	SessionLifecycleStateActive   SessionLifecycleStateEnum = "ACTIVE"
	SessionLifecycleStateDeleting SessionLifecycleStateEnum = "DELETING"
	SessionLifecycleStateDeleted  SessionLifecycleStateEnum = "DELETED"
	SessionLifecycleStateFailed   SessionLifecycleStateEnum = "FAILED"
)

var mappingSessionLifecycleState = map[string]SessionLifecycleStateEnum{
	"CREATING": SessionLifecycleStateCreating,
	"ACTIVE":   SessionLifecycleStateActive,
	"DELETING": SessionLifecycleStateDeleting,
	"DELETED":  SessionLifecycleStateDeleted,
	"FAILED":   SessionLifecycleStateFailed,
}

// GetSessionLifecycleStateEnumValues Enumerates the set of values for SessionLifecycleStateEnum
func GetSessionLifecycleStateEnumValues() []SessionLifecycleStateEnum {
	values := make([]SessionLifecycleStateEnum, 0)
	for _, v := range mappingSessionLifecycleState {
		values = append(values, v)
	}
	return values
}
