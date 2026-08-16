// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package errorhandler

import "fmt"

type Code string

const (
	ErrorCodeNotFound Code = "NOT_FOUND"
	ErrorCodeInternal Code = "INTERNAL"
)

type AppError interface {
	error
	Code() Code
	Message() string
}

type appError struct {
	code    Code
	message string
	cause   error
}

func newError(code Code, message string, cause error) *appError {
	return &appError{code: code, message: message, cause: cause}
}

func NotFound(message string) AppError {
	return newError(ErrorCodeNotFound, message, nil)
}

func Internal(cause error) AppError {
	return newError(ErrorCodeInternal, "Something went wrong. Please try again later.", cause)
}

func (e *appError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %s: %v", e.code, e.message, e.cause)
	}
	return fmt.Sprintf("%s: %s", e.code, e.message)
}

func (e *appError) Code() Code {
	return e.code
}

func (e *appError) Message() string {
	return e.message
}

func (e *appError) Unwrap() error {
	return e.cause
}
