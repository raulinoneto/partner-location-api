package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/raulinoneto/partner-location-api/internal/adapters/primary/lambdaadapter"
	"github.com/raulinoneto/partner-location-api/internal/adapters/secondary/mongodbadapter"
	"github.com/raulinoneto/partner-location-api/internal/apierror"
	"github.com/raulinoneto/partner-location-api/pkg/domains/partners"
	"net/http"
	"os"
)

func handler(request lambdaadapter.Request) (lambdaadapter.Response, error) {
	p := new(partners.Partner)
	if err := json.Unmarshal([]byte(request.Body), p); err != nil {
		err = apierror.NewWarning(http.StatusBadRequest, err.Error(), err)
		return lambdaadapter.BuildBadRequestResponse(err), nil
	}
	service := partners.NewService(
		mongodbadapter.NewMongoDBPartnerAdapter(
			os.Getenv("DB_NAME"),
			os.Getenv("PARTNERS_TABLE_NAME"),
			nil,
		),
	)
	return lambdaadapter.BuildCreatedResponse(service.CreatePartner(p)), nil
}

func main() {
	lambda.Start(handler)
}
