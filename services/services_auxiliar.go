package services

import (
	"errors"
	"net/http"
)

// Error Structure
type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

// Possibles service error
var (
	ErrNotFound   = errors.New("resource not found")
	ErrInternal   = errors.New("internal server error")
	ErrJson       = errors.New("JSON format is not valid")
	ErrBadRequest = errors.New("request is not valid")
)

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
		Err:  "request is no valid",
	}
)
