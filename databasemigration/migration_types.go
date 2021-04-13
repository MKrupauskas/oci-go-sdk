// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// REST API for Zero Downtime Migration (Oracle Database Migration Service --ODMS-- as customer-facing service name)
//
// Provides users the ability to perform Zero Downtime migration operations
//

package databasemigration

// MigrationTypesEnum Enum with underlying type: string
type MigrationTypesEnum string

// Set of constants representing the allowable values for MigrationTypesEnum
const (
	MigrationTypesOnline  MigrationTypesEnum = "ONLINE"
	MigrationTypesOffline MigrationTypesEnum = "OFFLINE"
)

var mappingMigrationTypes = map[string]MigrationTypesEnum{
	"ONLINE":  MigrationTypesOnline,
	"OFFLINE": MigrationTypesOffline,
}

// GetMigrationTypesEnumValues Enumerates the set of values for MigrationTypesEnum
func GetMigrationTypesEnumValues() []MigrationTypesEnum {
	values := make([]MigrationTypesEnum, 0)
	for _, v := range mappingMigrationTypes {
		values = append(values, v)
	}
	return values
}
