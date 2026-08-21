// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package testutil

import (
	"github.com/FarzanHajian/lexforge/backend/internal/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) FindById(id string) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepositoryMock) LoadAll(buffer any, where ...any) error {
	callArgs := append([]any{buffer}, where...)
	args := m.Called(callArgs...)
	return args.Error(0)
}

func (m *UserRepositoryMock) Update(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}
