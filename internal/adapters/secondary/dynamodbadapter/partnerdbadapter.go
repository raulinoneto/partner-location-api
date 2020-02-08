package dynamodbadapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/raulinoneto/partner-location-api/pkg/domains/partners"
	"github.com/raulinoneto/partner-location-api/pkg/helpers"
)

type AWSDocDBPartnerAdapter struct {
	tableName string
	conn      dynamodbiface.DynamoDBAPI
}

func NewAWSDocDBPartnerAdapter(tableName string, conn dynamodbiface.DynamoDBAPI) *AWSDocDBPartnerAdapter {
	if conn == nil {
		conn = CreateDynamoSess()
	}
	return &AWSDocDBPartnerAdapter{
		tableName,
		conn,
	}
}

func (a *AWSDocDBPartnerAdapter) SavePartner(partner *partners.Partner) (*partners.Partner, error) {
	item, err := dynamodbattribute.MarshalMap(&partner)
	if err != nil {
		return nil, err
	}
	partner.ID = helpers.GenerateUUID()
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(a.tableName),
	}
	result, err := a.conn.PutItem(input)
	if err != nil {
		return nil, err
	}
	err = dynamodbattribute.UnmarshalMap(result.Attributes, partner)
	return partner, err
}

func (a *AWSDocDBPartnerAdapter) GetPartner(id string) (*partners.Partner, error) {
	partner := new(partners.Partner)
	query := &dynamodb.GetItemInput{
		TableName: aws.String(a.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ConsistentRead: aws.Bool(true),
	}
	result, err := a.conn.GetItem(query)
	if err != nil {
		return nil, err
	}
	if len(result.Item) < 1 {
		return nil, err
	}
	err = dynamodbattribute.UnmarshalMap(result.Item, partner)
	return partner, nil
}

func (a *AWSDocDBPartnerAdapter) SearchPartners(point *partners.Point) ([]partners.Partner, error) {
	return nil, nil
}
