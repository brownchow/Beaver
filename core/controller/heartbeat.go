// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"time"

	"github.com/clivern/beaver/core/cluster"

	log "github.com/sirupsen/logrus"
)

// Heartbeat node heartbeat
func Heartbeat() {
	node := &cluster.Node{}
	err := node.Init()

	log.Info(`Start heartbeat daemon`)

	if err != nil {
		panic(fmt.Sprintf(
			"Error while connecting to etcd: %s",
			err.Error(),
		))
	}

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
