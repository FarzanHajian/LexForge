// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"log"

	"github.com/FarzanHajian/lexforge/backend/internal/config"
)

const version = "0.1.0"

func main() {
	cfg := config.SysConfig{}
	err := config.LoadSystemConfig(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := setupDatabase(cfg.DatabaseDsn, cfg.MigrateDatabase)
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer teardownDatabase(db)

	log.Printf("lexforge server %s started...\n", version)
	log.Printf("System Configuration: %s\n", cfg.String())
}
