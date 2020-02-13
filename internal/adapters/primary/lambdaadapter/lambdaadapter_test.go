package lambdaadapter

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/raulinoneto/partner-location-api/internal/apierror"
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
		result := buildResponse(tCase.status, tCase.body, tCase.err)
		if !reflect.DeepEqual(tCase.expected, result) {
			t.Errorf("case: %s\n expected: %v\n got: %v\n", caseName, tCase.expected, result)
		}
	}
}

func TestBuildOKResponse(t *testing.T) {
	expected := newResponse(http.StatusOK, `{"Test":"Test"}`)
	result := BuildOKResponse(&TestPayload{"Test"}, nil)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v\n got: %v\n", expected, result)
	}
}

func TestBuildCreatedResponse(t *testing.T) {
	expected := newResponse(http.StatusCreated, `{"Test":"Test"}`)
	result := BuildCreatedResponse(&TestPayload{"Test"}, nil)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v\n got: %v\n", expected, result)
	}
}

func TestBuildBadRequestResponse(t *testing.T) {
	expected := newResponse(http.StatusBadRequest, `{"status_code":400,"message":"test","severity":1}`)
	err := apierror.NewWarning(http.StatusBadRequest, "test", errors.New("test"))
	result := BuildBadRequestResponse(err)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected: %v\n got: %v\n", expected, result)
	}
}
