// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package repository

import (
	"context"
	"errors"

	"github.com/FarzanHajian/lexforge/backend/internal/domain"
	"github.com/FarzanHajian/lexforge/backend/internal/errorhandler"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r *GormUserRepository) FindById(id string) (domain.User, error) {
	ctx := context.Background()

	result, err := gorm.G[domain.User](r.db).Where(&domain.User{Id: id}).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, errorhandler.NotFound("User not found")
		}
		return domain.User{}, err
	}
	return result, nil
}

func (r *GormUserRepository) GetAllAsLookup() ([]domain.UserLookup, error) {
	var lookups []domain.UserLookup
	if err := r.db.Model(&domain.User{}).Find(&lookups).Error; err != nil {
		return nil, err
	}
	return lookups, nil
}

func (r *GormUserRepository) Update(user domain.User) (domain.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}
