package errors

import (
	"net/http"
)

// ErrorCodeInvalidToken means that token is invalid
const ErrorCodeInvalidToken = "invalid-token"

func NewInvalidTokenError() AppError {
	return AppError{
		StatusCode: http.StatusUnauthorized,
		Code:       ErrorCodeInvalidToken,
	}
}
