// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

import (
	"github.com/stretchr/testify/mock"
)

type EtcdMock struct {
	mock.Mock
}

func (e *EtcdMock) DoSomething(number int) (bool, error) {
	args := e.Called(number)
	return args.Bool(0), args.Error(1)
}
