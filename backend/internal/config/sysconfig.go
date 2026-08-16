// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package config

import (
	"fmt"
	"regexp"

	"github.com/caarlos0/env/v11"
)

type SysConfig struct {
	// MaxSessionsPerDay is the maximum number of sessions per topic a user can have per day.
	MaxSessionsPerDay int `env:"LEXFORGE_MAX_SESSIONS_PER_DAY" envDefault:"1"`

	MigrateDatabase bool   `env:"LEXFORGE_MIGRATE_DATABASE" envDefault:"false"`
	DatabaseDsn     string `env:"LEXFORGE_DATABASE_DSN" envDefault:"lexforge:lexforge@tcp(localhost:3306)/lexforge?charset=utf8mb4&parseTime=True&loc=Local"`
}

func LoadSystemConfig(cfg *SysConfig) error {
	err := env.Parse(&cfg)
	if err != nil {
		return fmt.Errorf("Failed to load environment variables: %w", err)
	}
	return nil
}

func (cfg *SysConfig) String() string {
	var credentialsPattern = regexp.MustCompile(`[^:/@]+:[^@]*@`)

	return fmt.Sprintf(
		"\tMaxSessionsPerDay: %d\tMigrateDatabase: %t\tDatabaseDsn: %s",
		cfg.MaxSessionsPerDay,
		cfg.MigrateDatabase,
		credentialsPattern.ReplaceAllString(cfg.DatabaseDsn, "***:***@"))
}
