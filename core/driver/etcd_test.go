// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

import (
	"testing"

	"github.com/franela/goblin"
)

// TestEtcd
func TestEtcd(t *testing.T) {
	g := goblin.Goblin(t)
	e := new(EtcdMock)

	e.On("Get", "parent/child").Return(
		map[string]string{"foo": "foo", "bar": "bar"},
		nil,
	)

	g.Describe("#EtcdMock.Get", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				arg        string
				wantResult map[string]string
				wantError  bool
			}{
				{"parent/child", map[string]string{"foo": "foo", "bar": "bar"}, false},
			}

			for _, tt := range tests {
				result, err := e.Get(tt.arg)
				g.Assert(result).Equal(tt.wantResult)
				g.Assert(err != nil).Equal(tt.wantError)
			}
		})
	})
}
