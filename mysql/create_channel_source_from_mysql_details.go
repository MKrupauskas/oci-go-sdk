// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CreateChannelSourceFromMysqlDetails Parameters detailing how to provision the source endpoint that is a MySQL Server.
// Typically a MySQL Server that is not managed by the MySQL Database Service.
type CreateChannelSourceFromMysqlDetails struct {

	// The network address of the MySQL instance.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The name of the replication user on the source MySQL instance.
	// The username has a maximum length of 96 characters. For more information,
	// please see the MySQL documentation (https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html)
	Username *string `mandatory:"true" json:"username"`

	// The password for the replication user. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character,
	// and 1 special (nonalphanumeric) character.
	Password *string `mandatory:"true" json:"password"`

	// The port the source MySQL instance listens on.
	Port *int `mandatory:"false" json:"port"`

	SslCaCertificate CaCertificate `mandatory:"false" json:"sslCaCertificate"`

	// The SSL mode of the Channel.
	SslMode ChannelSourceMysqlSslModeEnum `mandatory:"true" json:"sslMode"`
}

func (m CreateChannelSourceFromMysqlDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateChannelSourceFromMysqlDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateChannelSourceFromMysqlDetails CreateChannelSourceFromMysqlDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateChannelSourceFromMysqlDetails
	}{
		"MYSQL",
		(MarshalTypeCreateChannelSourceFromMysqlDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateChannelSourceFromMysqlDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Port             *int                          `json:"port"`
		SslCaCertificate cacertificate                 `json:"sslCaCertificate"`
		Hostname         *string                       `json:"hostname"`
		Username         *string                       `json:"username"`
		Password         *string                       `json:"password"`
		SslMode          ChannelSourceMysqlSslModeEnum `json:"sslMode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Port = model.Port

	nn, e = model.SslCaCertificate.UnmarshalPolymorphicJSON(model.SslCaCertificate.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SslCaCertificate = nn.(CaCertificate)
	} else {
		m.SslCaCertificate = nil
	}

	m.Hostname = model.Hostname

	m.Username = model.Username

	m.Password = model.Password

	m.SslMode = model.SslMode

	return
}
