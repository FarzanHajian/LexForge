// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id                   string `json:"id" gorm:"type:char(36);primaryKey"`
	Name                 string `json:"name" gorm:"type:varchar(255);uniqueIndex;not null"`
	ExternalId           string `json:"externalId" gorm:"type:char(36);uniqueIndex;not null"`
	StudyItemsPerSession int    `json:"studyItemsPerSession" gorm:"type:int;not null"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func NewUser(name, externalId string, studyItemsPerSession int) User {
	return User{
		Id:                   uuid.NewString(),
		Name:                 name,
		ExternalId:           externalId,
		StudyItemsPerSession: studyItemsPerSession,
	}
}

type UserLookup struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
