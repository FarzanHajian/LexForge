// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"errors"
	"net/http"

	"github.com/FarzanHajian/lexforge/backend/internal/errorhandler"
	"github.com/labstack/echo/v5"
)

func httpErrorHandler(fallback echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(c *echo.Context, err error) {
		if r, _ := echo.UnwrapResponse(c.Response()); r != nil && r.Committed {
			return
		}

		if appErr, ok := errors.AsType[errorhandler.AppError](err); ok {
			status := statusFor(appErr.Code())
			if status >= http.StatusInternalServerError {
				c.Logger().Error("Request failed.", "error", err)
			}
			if cErr := c.JSON(status, map[string]string{"code": string(appErr.Code()), "message": appErr.Message()}); cErr != nil {
				c.Logger().Error("Failed to send error response.", "error", cErr)
			}
		} else {
			if echo.StatusCode(err) == 0 {
				c.Logger().Error("Unhandled error.", "error", err)
			}
			fallback(c, err)
		}
	}
}

func statusFor(code errorhandler.Code) int {
	switch code {
	case errorhandler.ErrorCodeNotFound:
		return http.StatusNotFound
	case errorhandler.ErrorCodeBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
