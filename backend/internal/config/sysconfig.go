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

const (
	EnvServerAddress     = "LEXFORGE_SERVER_ADDRESS"
	EnvMaxSessionsPerDay = "LEXFORGE_MAX_SESSIONS_PER_DAY"
	EnvMigrateDatabase   = "LEXFORGE_MIGRATE_DATABASE"
	EnvDatabaseDsn       = "LEXFORGE_DATABASE_DSN"
)

type SysConfig struct {
	ServerAddress string `env:"LEXFORGE_SERVER_ADDRESS" envDefault:"http://localhost:5656"`

	// MaxSessionsPerDay is the maximum number of sessions per topic a user can have per day.
	MaxSessionsPerDay int `env:"LEXFORGE_MAX_SESSIONS_PER_DAY" envDefault:"1"`

	MigrateDatabase bool   `env:"LEXFORGE_MIGRATE_DATABASE" envDefault:"false"`
	DatabaseDsn     string `env:"LEXFORGE_DATABASE_DSN,notEmpty,required"`

	OAuthAuthority string `env:"AUTH0_DOMAIN,notEmpty,required"`
	OAuthAudience  string `env:"AUTH0_AUDIENCE,notEmpty,required"`
}

func LoadSystemConfig(cfg *SysConfig) error {
	err := env.Parse(cfg)
	if err != nil {
		return fmt.Errorf("Failed to load environment variables: %w", err)
	}
	return nil
}

func (cfg *SysConfig) String() string {
	var credentialsPattern = regexp.MustCompile(`[^:/@]+:[^@]*@`)

	return fmt.Sprintf(
		"\n\tMaxSessionsPerDay: %d\n\tMigrateDatabase: %t\n\tDatabaseDsn: %s",
		cfg.MaxSessionsPerDay,
		cfg.MigrateDatabase,
		credentialsPattern.ReplaceAllString(cfg.DatabaseDsn, "***:***@"))
}
