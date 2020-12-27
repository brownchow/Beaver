// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"fmt"
	"testing"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
	"go.etcd.io/etcd/clientv3"
)

// TestNode test cases
func TestNode(t *testing.T) {
	// Load Configs
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("#Node.Alive", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				input      int64
				mockReturn error
				wantError  bool
			}{
				{5, nil, false},
				{6, nil, false},
				{7, nil, false},
				{5, fmt.Errorf("error1"), true},
				{6, fmt.Errorf("error2"), true},
				{7, fmt.Errorf("error3"), true},
			}

			for _, tt := range tests {
				e := new(driver.EtcdMock)
				n := NewNode(e)

				e.On("CreateLease", int64(tt.input)).Return(clientv3.LeaseID(1234567), tt.mockReturn)
				e.On("RenewLease", clientv3.LeaseID(1234567)).Return(tt.mockReturn)
				e.On("PutWithLease", "beaver_v2/node/clivern-2.local__x-x-x-x/state", "alive", clientv3.LeaseID(1234567)).Return(tt.mockReturn)
				e.On("PutWithLease", "beaver_v2/node/clivern-2.local__x-x-x-x/url", "http://127.0.0.1:8080", clientv3.LeaseID(1234567)).Return(tt.mockReturn)
				e.On("PutWithLease", "beaver_v2/node/clivern-2.local__x-x-x-x/load", "0", clientv3.LeaseID(1234567)).Return(tt.mockReturn)

				g.Assert(n.Alive(tt.input) != nil).Equal(tt.wantError)
			}
		})
	})

	g.Describe("#Node.GetHostname", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := new(driver.EtcdMock)
			n := NewNode(e)

			result, err := n.GetHostname()
			g.Assert(err).Equal(nil)
			g.Assert(result != "").Equal(true)
		})
	})
}
