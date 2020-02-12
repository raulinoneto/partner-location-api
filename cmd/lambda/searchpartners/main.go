package main

import (
	"fmt"
	"github.com/raulinoneto/partner-location-api/internal/adapters/secondary/mongodbadapter"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/raulinoneto/partner-location-api/internal/adapters/primary/lambdaadapter"
	"github.com/raulinoneto/partner-location-api/internal/apierror"
	"github.com/raulinoneto/partner-location-api/pkg/domains/partners"
)

var (
	errLat = fmt.Errorf("the lat is invalid or empty")
	errLng = fmt.Errorf("the lng is invalid or empty")
)

func handler(request lambdaadapter.Request) (lambdaadapter.Response, error) {
	point, err := getPointByQueryParams(request.PathParameters)
	if err != nil {
		return lambdaadapter.BuildBadRequestResponse(err), nil
	}
	service := partners.NewService(
		mongodbadapter.NewMongoDBPartnerAdapter(
			os.Getenv("DB_NAME"),
			os.Getenv("PARTNERS_TABLE_NAME"),
			nil,
		),
	)
	return lambdaadapter.BuildCreatedResponse(service.SearchPartners(point)), nil
}

func getPointByQueryParams(query map[string]string) (*partners.Point, error) {
	lat, ok := query["lat"]
	if !ok {
		return nil, apierror.NewWarning(http.StatusBadRequest, errLat.Error(), errLat)
	}
	latFloat, err := strconv.ParseFloat(lat, 0)
	if err != nil {
		return nil, apierror.NewWarning(http.StatusBadRequest, errLat.Error(), errLat)
	}
	lng, ok := query["lng"]
	if !ok {
		return nil, apierror.NewWarning(http.StatusBadRequest, errLng.Error(), errLng)
	}
	lngFloat, err := strconv.ParseFloat(lng, 0)
	if err != nil {
		return nil, apierror.NewWarning(http.StatusBadRequest, errLng.Error(), errLng)
	}
	return &partners.Point{latFloat, lngFloat}, nil

}

func main() {
	lambda.Start(handler)
}
