// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package service

/*
import (
	"testing"

	"github.com/FarzanHajian/lexforge/backend/internal/domain"
	"github.com/FarzanHajian/lexforge/backend/internal/repository"
	"github.com/FarzanHajian/lexforge/backend/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_GetOrCreate_ExistingUser(t *testing.T) {
	repo := new(testutil.UserRepositoryMock)
	existing := domain.User{Id: "u-1", Name: "Ada", ExternalId: "auth0|123", StudyItemsPerSession: 10}
	repo.On("FindByExternalId", "auth0|123").Return(existing, nil)

	svc := NewUserService(repo)
	user, err := svc.GetOrCreate("Ada", "auth0|123", 10)

	assert.NoError(t, err)
	assert.Equal(t, existing, user)
	repo.AssertNotCalled(t, "Create", mock.Anything)
	repo.AssertExpectations(t)
}

func TestUserService_GetOrCreate_NewUser(t *testing.T) {
	repo := new(testutil.UserRepositoryMock)
	repo.On("FindByExternalId", "auth0|456").Return(domain.User{}, repository.ErrNotFound)
	repo.On("Create", mock.MatchedBy(func(u domain.User) bool {
		return u.Name == "Grace" && u.ExternalId == "auth0|456" && u.StudyItemsPerSession == 10
	})).Return(nil)

	svc := NewUserService(repo)
	user, err := svc.GetOrCreate("Grace", "auth0|456", 10)

	assert.NoError(t, err)
	assert.Equal(t, "Grace", user.Name)
	assert.Equal(t, "auth0|456", user.ExternalId)
	assert.NotEmpty(t, user.Id)
	repo.AssertExpectations(t)
}

func TestUserService_GetById(t *testing.T) {
	repo := new(testutil.UserRepositoryMock)
	want := domain.User{Id: "u-1", Name: "Ada", ExternalId: "auth0|123"}
	repo.On("FindById", "u-1").Return(want, nil)

	svc := NewUserService(repo)
	user, err := svc.GetById("u-1")

	assert.NoError(t, err)
	assert.Equal(t, want, user)
}
*/
