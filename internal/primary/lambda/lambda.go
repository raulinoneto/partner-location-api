package lambda

import (
	"github.com/aws/aws-lambda-go/events"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func BuildResponse(response interface{}, err error) Response {
	return Response{}
}
