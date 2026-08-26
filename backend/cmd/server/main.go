// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"log"

	"github.com/FarzanHajian/lexforge/backend/internal/auth"
	"github.com/FarzanHajian/lexforge/backend/internal/config"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

const version = "0.1.0"

func main() {
	log.Printf("LexForge server %s started...\n", version)

	cfg := config.SysConfig{}
	err := config.LoadSystemConfig(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("System Configuration: %s\n", cfg.String())

	db, err := setupDatabase(cfg.DatabaseDsn, cfg.MigrateDatabase)
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer teardownDatabase(db)

	jwtMiddleware, err := auth.NewJWTMiddleware(cfg.OAuthAuthority, cfg.OAuthAudience)
	if err != nil {
		log.Fatalf("Failed to set up JWT middleware: %v", err)
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	api := e.Group("/api", echo.WrapMiddleware(jwtMiddleware.CheckJWT))
	registerRoutes(e, api)

	if err := e.Start(cfg.ServerAddress); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
