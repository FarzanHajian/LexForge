// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

// Package integration contains repository/integration tests that run
// against a real MySQL instance provisioned via Testcontainers, as
// opposed to the unit tests colocated with each internal package
// (which use mocks/fakes and no database).
//
// These tests are gated behind the "integration" build tag so that a
// plain `go test ./...` never requires Docker. Run them explicitly:
//
//	go test -tags=integration ./test/integration/...
//
// Docker (or a compatible container runtime) must be running locally.
package integration
