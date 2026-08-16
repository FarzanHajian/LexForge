// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package domain

import "github.com/google/uuid"

type Topic struct {
	Id       string `json:"id" gorm:"type:char(36);primaryKey"`
	Name     string `json:"name" gorm:"uniqueIndex;not null"`
	UserId   string `json:"userId" gorm:"type:char(36);not null"`
	User     User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Template string `json:"template" gorm:"type:varchar(2048)"`
}

func NewTopic(name, userId, template string) Topic {
	return Topic{
		Id:       uuid.NewString(),
		Name:     name,
		UserId:   userId,
		Template: template,
	}
}
