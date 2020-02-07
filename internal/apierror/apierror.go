package apierror

import (
	"encoding/json"
	"net/http"
)

const (
	Debug = iota
	Warning
	Critical
)

type ApiError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Severity   int    `json:"severity"`
	err        error
}

func (e *ApiError) Error() string {
	errJson, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(errJson)
}

func NewCritical(message string, err error) *ApiError {
	return &ApiError{http.StatusInternalServerError, message, Critical, err}
}

func NewWarning(statusCode int, message string, err error) *ApiError {
	return &ApiError{statusCode, message, Warning, err}
}

func NewDebug(statusCode int, message string, err error) *ApiError {
	return &ApiError{statusCode, message, Debug, err}
}
