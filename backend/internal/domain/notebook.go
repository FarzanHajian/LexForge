// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package domain

import "github.com/google/uuid"

type Notebook struct {
	Id       string `json:"id" gorm:"type:char(36);primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(255);uniqueIndex:idx_notebook_user_name;not null"`
	UserId   string `json:"userId" gorm:"type:char(36);uniqueIndex:idx_notebook_user_name;not null"`
	User     User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Template string `json:"template" gorm:"type:varchar(2048)"`
	Color    string `json:"color" gorm:"type:char(7);not null"` // Color is a "#RRGGBB" value,
}

func NewNotebook(name, userId, template, color string) (Notebook, error) {
	if err := ValidateNotebookColor(color); err != nil {
		return Notebook{}, err
	}

	return Notebook{
		Id:       uuid.NewString(),
		Name:     name,
		UserId:   userId,
		Template: template,
		Color:    color,
	}, nil
}
