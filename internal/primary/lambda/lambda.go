package lambda

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/raulinoneto/partner-location-api/internal/apierror"
	"net/http"
	"reflect"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func newResponse(status int, body string) Response {
	return Response{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
		Body: body,
	}
}

func BuildResponse(status int, body interface{}, err error) Response {
	if err != nil {
		return asError(err)
	}
	respJson, err := json.Marshal(body)
	if err != nil {
		return newResponse(http.StatusInternalServerError, err.Error())
	}
	return newResponse(status, string(respJson))
}

func asError(errResponse error) Response {
	if reflect.TypeOf(errResponse) != reflect.TypeOf(&apierror.ApiError{}) {
		errResponse = apierror.NewCritical(errResponse.Error(), errResponse)
	}
	errJson, err := json.Marshal(errResponse)
	if err != nil {
		return newResponse(http.StatusInternalServerError, err.Error())
	}
	return newResponse(errResponse.(*apierror.ApiError).StatusCode, string(errJson))
}
