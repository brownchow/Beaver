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
	e := &EtcdMock{}

	e.On("DoSomething", 1).Return(true, nil)

	g.Describe("#EtcdMock.DoSomething", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				arg        int
				wantResult bool
				wantError  bool
			}{
				{1, true, false},
			}

			for _, tt := range tests {
				result, err := e.DoSomething(tt.arg)
				g.Assert(result).Equal(tt.wantResult)
				g.Assert(err != nil).Equal(tt.wantError)
			}
		})
	})
}
