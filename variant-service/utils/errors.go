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
	ErrVariantExists = "Variant already exists"
	VariantNotFound  = "Variant not found"
)

var MapErrorStatusCode = map[string]int{
	ErrVariantExists: http.StatusConflict,
	VariantNotFound:  http.StatusNotFound,
}
