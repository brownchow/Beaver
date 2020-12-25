// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"

	"github.com/franela/goblin"
)

// TestCluster test cases
func TestCluster(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#Func", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				input      int
				wantOutput int
			}{
				{1, 1},
			}

			for _, tt := range tests {
				g.Assert(Sum(tt.input)).Equal(tt.wantOutput)
			}
		})
	})
}
