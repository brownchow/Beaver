// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"fmt"
	"os"
	"strings"

	"github.com/clivern/beaver/core/driver"

	"github.com/spf13/viper"
)

// Node type
type Node struct {
	Driver *driver.Etcd
}

// Init init node object
func (n *Node) Init() error {
	var err error

	n.Driver, err = driver.NewEtcdDriver()

	if err != nil {
		return err
	}

	return nil
}

// Alive report the node as live to etcd
func (n *Node) Alive(seconds int64) error {

	hostname, err := n.GetHostname()

	if err != nil {
		return err
	}

	key := fmt.Sprintf(
		"%s/node/%s__%s",
		viper.GetString("app.database.etcd.databaseName"),
		hostname,
		viper.GetString("app.name"),
	)

	leaseID, err := n.Driver.CreateLease(seconds)

	if err != nil {
		return err
	}

	err = n.Driver.PutWithLease(fmt.Sprintf("%s/state", key), "alive", leaseID)

	if err != nil {
		return err
	}

	err = n.Driver.PutWithLease(fmt.Sprintf("%s/url", key), viper.GetString("app.url"), leaseID)

	if err != nil {
		return err
	}

	err = n.Driver.RenewLease(leaseID)

	if err != nil {
		return err
	}

	return nil
}

// GetHostname gets the hostname
func (n *Node) GetHostname() (string, error) {
	hostname, err := os.Hostname()

	if err != nil {
		return "", err
	}

	return strings.ToLower(hostname), nil
}
