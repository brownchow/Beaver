// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"github.com/clivern/beaver/core/driver"
)

// Stats type
type Stats struct {
	Driver *driver.Etcd
}

// Init init node object
func (n *Stats) Init() error {
	var err error

	n.Driver, err = driver.NewEtcdDriver()

	if err != nil {
		return err
	}

	return nil
}

// GetTotalNodes
func (n *Stats) GetTotalNodes() (int, error) {
	return 0, nil
}

// GetTotalChannels
func (n *Stats) GetTotalChannels() (int, error) {
	return 0, nil
}

// GetTotalClients
func (n *Stats) GetTotalClients() (int, error) {
	return 0, nil
}

// GetTotalConnectedClients
func (n *Stats) GetTotalConnectedClients() (int, error) {
	return 0, nil
}
