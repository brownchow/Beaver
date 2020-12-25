// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/core/util"

	"github.com/spf13/viper"
)

// Node type
type Node struct {
	Driver *driver.Etcd
}

// Init init node object
func (n *Node) Init() error {
	n.Driver, err = driver.NewEtcdDriver()

	if err != nil {
		return err
	}

	return nil
}

// Alive report the node as live to etcd
func (n *Node) Alive(seconds int) error {

	hostname, err := n.GetHostname()

	if err != nil {
		return err
	}

	key := fmt.Sprintf(
		"%s/node/%s__%s",
		viper.GetString("database.etcd.databaseName"),
		hostname,
		util.GenerateUUID4(),
	)

	leaseId, err := n.Driver.CreateLease(seconds)

	if err != nil {
		return err
	}

	err = n.Driver.PutWithLease(key, "alive", leaseId)

	if err != nil {
		return err
	}

	err = n.Driver.RenewLease(leaseId)

	if err != nil {
		return err
	}
}

// GetHostname gets the hostname
func (n *Node) GetHostname() (string, error) {
	hostname, err := os.Hostname()

	if err != nil {
		return "", err
	}

	return strings.ToLower(hostname), nil
}
