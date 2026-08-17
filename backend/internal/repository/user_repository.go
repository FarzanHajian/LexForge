// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package repository

import "github.com/FarzanHajian/lexforge/backend/internal/domain"

type UserRepository interface {
	FindById(id string) (domain.User, error)
	GetAllAsLookup() ([]domain.UserLookup, error)
	Update(user *domain.User) (*domain.User, error)
}
