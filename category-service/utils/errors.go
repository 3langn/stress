package utils

import (
	"errors"
	"net/http"
)

type AppError error

func GetStatusCode(err error) (int, bool) {
	if v, ok := MapErrorStatusCode[err.Error()]; !ok {
		return http.StatusInternalServerError, false
	} else {
		return v, true
	}
}

var (
	ErrInternalServerError AppError = errors.New("Internal Server Error")
	ErrorBadParamInput              = "Bad param input %v"
)

const (
	ErrCatExists        = "Category already exists"
	ErrCatNotFound      = "Category not found"
	ErrCatAlreadyExists = "Category already exists"
)

var MapErrorStatusCode = map[string]int{
	ErrCatExists: http.StatusConflict,
}
