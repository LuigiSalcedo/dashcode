package services

import (
	"net/http"
)

// Error Structure
type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

// Errors to send as reponse
var (
	ErrorNotFound = &Error{
		Code: http.StatusNotFound,
		Err:  "resource not found",
	}

	ErrorInternal = &Error{
		Code: http.StatusInternalServerError,
		Err:  "internal server error",
	}

	ErrorJson = &Error{
		Code: http.StatusBadRequest,
		Err:  "json format is not valid",
	}

	ErrorBadRequest = &Error{
		Code: http.StatusBadRequest,
		Err:  "request is not valid",
	}

	ErrorUnauthorized = &Error{
		Code: http.StatusUnauthorized,
		Err:  "the access to this resource is unauthorized",
	}

	ErrorJWT = &Error{
		Code: http.StatusForbidden,
		Err:  "JWT expired or not valid",
	}

	ErrorPathParam = &Error{
		Code: http.StatusBadRequest,
		Err:  "path parameter is not valid",
	}
)
