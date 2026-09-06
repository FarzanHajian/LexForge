// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/FarzanHajian/lexforge/backend/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupDatabase(dsn string, migrate bool) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if migrate {
		err := db.AutoMigrate(&domain.User{}, &domain.Notebook{})
		if err != nil {
			return nil, fmt.Errorf("Failed to migrate database: %w", err)
		}
		log.Println("Database migrated successfully")
	}

	return db, nil
}

func teardownDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("failed to get underlying sql.DB during teardown: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("failed to close database connection: %v", err)
	}
}
