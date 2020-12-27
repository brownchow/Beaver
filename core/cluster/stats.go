// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"fmt"

	"github.com/clivern/beaver/core/driver"

	"github.com/spf13/viper"
)

// Stats type
type Stats struct {
	db driver.Database
}

// NewStats creates a stats instance
func NewStats(db driver.Database) *Stats {
	result := new(Stats)
	result.db = db

	return result
}

// GetTotalNodes
func (n *Stats) GetTotalNodes() (int, error) {

	key := fmt.Sprintf(
		"%s/node",
		viper.GetString("app.database.etcd.databaseName"),
	)

	keys, err := n.db.GetKeys(key)

	if err != nil {
		return 0, err
	}

	return len(keys), nil
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
