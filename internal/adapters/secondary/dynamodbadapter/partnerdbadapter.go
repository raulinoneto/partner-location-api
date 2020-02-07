package dynamodbadapter

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/raulinoneto/partner-location-api/pkg/domains/partners"
)

type AWSDocDBPartnerAdapter struct {
	conn dynamodbiface.DynamoDBAPI
}

func NewAWSDocDBPartnerAdapter(conn dynamodbiface.DynamoDBAPI) *AWSDocDBPartnerAdapter {
	return &AWSDocDBPartnerAdapter{
		conn,
	}
}

func (a *AWSDocDBPartnerAdapter) SavePartner(partner *partners.Partner) error {
	return nil
}

func (a *AWSDocDBPartnerAdapter) GetPartner(id string) (*partners.Partner, error) {
	return nil, nil
}

func (a *AWSDocDBPartnerAdapter) SearchPartners(point *partners.Point) ([]partners.Partner, error) {
	return nil, nil
}
