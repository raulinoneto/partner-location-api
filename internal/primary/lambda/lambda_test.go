package lambda

import (
	"errors"
	"github.com/raulinoneto/partner-location-api/internal/apierror"
	"net/http"
	"reflect"
	"testing"
)

type TestPayload struct{ Test string }
type testCase struct {
	status   int
	body     *TestPayload
	err      error
	expected Response
}

var testCases = map[string]testCase{
	"Status Ok": {
		http.StatusOK,
		&TestPayload{"Test"},
		nil,
		newResponse(http.StatusOK, `{"Test":"Test"}`),
	},
	"Status Bad Request": {
		http.StatusBadRequest,
		nil,
		apierror.NewWarning(http.StatusBadRequest, "Test", errors.New("test")),
		newResponse(http.StatusBadRequest, `{"status_code":400,"message":"Test","severity":1}`),
	},
	"Status Internal Server Error": {
		http.StatusInternalServerError,
		nil,
		errors.New("test"),
		newResponse(http.StatusInternalServerError, `{"status_code":500,"message":"test","severity":2}`),
	},
}

func TestBuildResponse(t *testing.T) {
	for caseName, tCase := range testCases {
		result := BuildResponse(tCase.status, tCase.body, tCase.err)
		if !reflect.DeepEqual(tCase.expected, result) {
			t.Errorf("case: %s\n expected: %v\n got: %v\n", caseName, tCase.expected, result)
		}
	}
}
