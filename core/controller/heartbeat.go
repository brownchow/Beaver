// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"time"

	"github.com/clivern/beaver/core/cluster"
	"github.com/clivern/beaver/core/driver"

	log "github.com/sirupsen/logrus"
)

// Heartbeat node heartbeat
func Heartbeat() {
	db := driver.NewEtcdDriver()

	err := db.Connect()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while connecting to etcd: %s",
			err.Error(),
		))
	}

	node := cluster.NewNode(db)

	log.Info(`Start heartbeat daemon`)

	for {
		err := node.Alive(5)

		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Error(`Error while connecting to etcd`)
		} else {
			log.Debug(`Node heartbeat done`)
		}

		time.Sleep(3 * time.Second)
	}
}
