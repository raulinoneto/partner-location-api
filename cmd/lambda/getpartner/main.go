package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/raulinoneto/partner-location-api/internal/apierror"
	"github.com/raulinoneto/partner-location-api/internal/primary/lambdaadapter"
	"github.com/raulinoneto/partner-location-api/pkg/domains/partners"
)

func handler(request lambdaadapter.Request) (lambdaadapter.Response, error) {
	pId, ok := request.PathParameters["partnerId"]
	if !ok {
		err := fmt.Errorf("the partnerId can't be empty")
		err = apierror.NewWarning(http.StatusBadRequest, err.Error(), err)
		return lambdaadapter.BuildBadRequestResponse(err), nil
	}
	service := partners.NewService(nil)
	return lambdaadapter.BuildCreatedResponse(service.GetPartner(pId)), nil
}

func main() {
	lambda.Start(handler)
}
