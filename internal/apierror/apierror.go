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
	StatusCode int    `json:"status_code"`
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

func NewWarning(status_code int, message string, err error) *ApiError {
	return &ApiError{status_code, message, Warning, err}
}

func NewDebug(status_code int, message string, err error) *ApiError {
	return &ApiError{status_code, message, Debug, err}
}
