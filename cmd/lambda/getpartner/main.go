package main

import (
	"fmt"
	"github.com/raulinoneto/partner-location-api/internal/adapters/secondary/dynamodbadapter"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/raulinoneto/partner-location-api/internal/adapters/primary/lambdaadapter"
	"github.com/raulinoneto/partner-location-api/internal/apierror"
	"github.com/raulinoneto/partner-location-api/pkg/domains/partners"
)

func handler(request lambdaadapter.Request) (lambdaadapter.Response, error) {
	pId, ok := request.PathParameters["partnerId"]
	if !ok {
		err := fmt.Errorf("the partnerId can't be empty")
		err = apierror.NewWarning(http.StatusBadRequest, err.Error(), err)
		return lambdaadapter.BuildBadRequestResponse(err), nil
	}
	service := partners.NewService(
		dynamodbadapter.NewAWSDocDBPartnerAdapter(os.Getenv("PARTNERS_TABLE_NAME"), nil),
	)
	return lambdaadapter.BuildCreatedResponse(service.GetPartner(pId)), nil
}

func main() {
	lambda.Start(handler)
}
