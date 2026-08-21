// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package service

import (
	"errors"
	"testing"

	"github.com/FarzanHajian/lexforge/backend/internal/domain"
	"github.com/FarzanHajian/lexforge/backend/internal/errorhandler"
	"github.com/FarzanHajian/lexforge/backend/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_WhenUserNotFound_ReturnsNotFoundError(t *testing.T) {
	repo := new(testutil.UserRepositoryMock)
	repo.On("FindById", "u-1").Return(domain.User{}, errorhandler.NotFound("User not found"))
	svc := NewUserService(repo)

	user, err := svc.GetById("u-1")

	assert.Equal(t, domain.User{}, user)
	var appErr errorhandler.AppError
	assert.ErrorAs(t, err, &appErr)
	assert.Equal(t, errorhandler.ErrorCodeNotFound, appErr.Code())
	assert.Equal(t, "User not found", appErr.Message())
}

func TestUserService_UpdateSettings_Success(t *testing.T) {
	repo := new(testutil.UserRepositoryMock)
	existing := domain.User{Id: "u-1", Name: "Ada", ExternalId: "auth0|123", StudyItemsPerSession: 10}
	repo.On("FindById", "u-1").Return(existing, nil)

	updated := existing
	updated.StudyItemsPerSession = 20
	repo.On("Update", &updated).Return(&updated, nil)

	svc := NewUserService(repo)
	user, err := svc.UpdateSettings("u-1", 20)

	assert.NoError(t, err)
	assert.Equal(t, updated, user)
	repo.AssertExpectations(t)
}

func TestUserService_UpdateSettings_GetByIdError(t *testing.T) {
	repo := new(testutil.UserRepositoryMock)
	repo.On("FindById", "missing").Return(domain.User{}, errorhandler.NotFound("User not found"))

	svc := NewUserService(repo)
	_, err := svc.UpdateSettings("missing", 20)

	assert.Error(t, err)
	repo.AssertNotCalled(t, "Update", mock.Anything)
	repo.AssertExpectations(t)
}

func TestUserService_UpdateSettings_UpdateError(t *testing.T) {
	repo := new(testutil.UserRepositoryMock)
	existing := domain.User{Id: "u-1", Name: "Ada", ExternalId: "auth0|123", StudyItemsPerSession: 10}
	repo.On("FindById", "u-1").Return(existing, nil)

	updated := existing
	updated.StudyItemsPerSession = 20
	repo.On("Update", &updated).Return((*domain.User)(nil), errors.New("db error"))

	svc := NewUserService(repo)
	_, err := svc.UpdateSettings("u-1", 20)

	assert.Error(t, err)
	repo.AssertExpectations(t)
}
