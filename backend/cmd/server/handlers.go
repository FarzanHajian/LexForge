// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"net/http"

	"github.com/FarzanHajian/lexforge/backend/internal/auth"
	"github.com/labstack/echo/v5"
)

func registerRoutes(e *echo.Echo, api *echo.Group) {

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	api.GET("/me", func(c *echo.Context) error {
		sub, err := auth.Subject(c.Request())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to read claims"})
		}
		return c.JSON(http.StatusOK, map[string]string{"sub": sub})
	})
}
