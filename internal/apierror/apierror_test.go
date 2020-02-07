package apierror

import (
	"errors"
	"net/http"
	"testing"
)

type errorTestCase struct {
	apiError *ApiError
	expected string
}

var testCases = map[string]errorTestCase{
	"Error Critical": {
		NewCritical("Test", errors.New("Test")),
		`{"status_code":500,"message":"Test","severity":2}`,
	},
	"Error Warning": {
		NewWarning(http.StatusBadRequest, "Test", errors.New("Test")),
		`{"status_code":400,"message":"Test","severity":1}`,
	},
	"Error Debug": {
		NewDebug(http.StatusOK, "Test", errors.New("Test")),
		`{"status_code":200,"message":"Test","severity":0}`,
	},
	"Error on Marshal": {
		nil,
		"null",
	},
}

func TestApiError_Error(t *testing.T) {
	for caseName, tCase := range testCases {
		result := tCase.apiError.Error()
		if result != tCase.expected {
			t.Errorf("case: %s\n expected: %s\n got: %s\n", caseName, tCase.expected, result)
		}
	}
}
