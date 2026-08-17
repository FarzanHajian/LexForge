// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package service

import (
	"github.com/FarzanHajian/lexforge/backend/internal/domain"
	"github.com/FarzanHajian/lexforge/backend/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetById(id string) (domain.User, error) {
	return s.repo.FindById(id)
}

func (s *UserService) GetAllAsLookup() ([]domain.UserLookup, error) {
	return s.repo.GetAllAsLookup()
}

func (s *UserService) UpdateSettings(id string, studyItemsPerSession int) (domain.User, error) {
	user, err := s.GetById(id)
	if err != nil {
		return domain.User{}, err
	}

	user.StudyItemsPerSession = studyItemsPerSession
	result, err := s.repo.Update(&user)
	if err != nil {
		return domain.User{}, err
	}

	return *result, nil
}
