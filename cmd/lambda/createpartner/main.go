package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/raulinoneto/partner-location-api/internal/apierror"
	"github.com/raulinoneto/partner-location-api/internal/primary/lambdaadapter"
	"github.com/raulinoneto/partner-location-api/pkg/domains/partners"
	"net/http"
)

func handler(request lambdaadapter.Request) (lambdaadapter.Response, error) {
	p := new(partners.Partner)
	if err := json.Unmarshal([]byte(request.Body), p); err != nil {
		err = apierror.NewWarning(http.StatusBadRequest, err.Error(), err)
		return lambdaadapter.BuildResponse(0, nil, err), nil
	}
	service := partners.NewService(nil)
	return lambdaadapter.BuildCreatedResponse(service.CreatePartner(p)), nil
}

func main() {
	lambda.Start(handler)
}
