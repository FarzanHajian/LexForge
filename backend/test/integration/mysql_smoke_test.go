// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

//go:build integration

package integration

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

// TestMySQLContainer_Smoke verifies that a MySQL Testcontainer can be
// started and connected to. It is a template for future
// repository/integration tests, not a test of application behavior.
func TestMySQLContainer_Smoke(t *testing.T) {
	ctx := context.Background()

	container, err := mysql.Run(ctx, "mysql:8.0",
		mysql.WithDatabase("lexforge_test"),
		mysql.WithUsername("lexforge"),
		mysql.WithPassword("lexforge"),
	)
	require.NoError(t, err)
	defer func() {
		require.NoError(t, container.Terminate(ctx))
	}()

	dsn, err := container.ConnectionString(ctx)
	require.NoError(t, err)

	db, err := sql.Open("mysql", dsn)
	require.NoError(t, err)
	defer db.Close()

	require.NoError(t, db.PingContext(ctx))
}
