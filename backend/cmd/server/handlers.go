// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"net/http"

	"github.com/FarzanHajian/lexforge/backend/internal/auth"
	"github.com/FarzanHajian/lexforge/backend/internal/errorhandler"
	"github.com/FarzanHajian/lexforge/backend/internal/httpapi"
	"github.com/labstack/echo/v5"
)

func registerRoutes(e *echo.Echo, api *echo.Group, services *ServiceContainer) {

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World! LexForge API is running."})
	})

	api.GET("/users/me", func(c *echo.Context) error {
		sub, err := auth.Subject(c.Request())
		if err != nil {
			return errorhandler.BadRequest("The provided access token is invalid.", err)
		}

		user, err := services.userService.GetByExternalId(sub)
		if err != nil {
			if err, ok := err.(errorhandler.AppError); ok && err.Code() == errorhandler.ErrorCodeNotFound {
				return err
			}
			return errorhandler.Internal(err)
		}
		return c.JSON(http.StatusOK, user)
	})

	api.PUT("/users/me", func(c *echo.Context) error {
		sub, err := auth.Subject(c.Request())
		if err != nil {
			return errorhandler.BadRequest("The provided access token is invalid.", err)
		}

		var req httpapi.UserUpdateRequest
		if err := c.Bind(&req); err != nil {
			return errorhandler.BadRequest("The provided request body is invalid.", err)
		}

		user, err := services.userService.UpdateSettings(sub, req.StudyItemsPerSession)
		if err != nil {
			return errorhandler.Internal(err)
		}
		return c.JSON(http.StatusOK, user)
	})
}
