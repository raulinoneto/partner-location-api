package dynamodbadapter

import (
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/raulinoneto/partner-location-api/pkg/domains/partners"
)

type DynamoDBAPIMock struct {
	dynamodbiface.DynamoDBAPI
	PutItemOutput *dynamodb.PutItemOutput
	GetItemOutput *dynamodb.GetItemOutput
	ScanOutput    *dynamodb.ScanOutput
}

var (
	nilPutItemInput = errors.New("nil PutItemInput")
	nilGetItemInput = errors.New("nil GetItemInput")
	nilScanInput    = errors.New("nil ScanInput")
)

func (d *DynamoDBAPIMock) PutItem(pi *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if len(pi.Item) <= 0 {
		return nil, nilPutItemInput
	}
	return d.PutItemOutput, nil
}

func (d *DynamoDBAPIMock) GetItem(gi *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	if len(gi.Key) <= 0 {
		return nil, nilGetItemInput
	}
	return d.GetItemOutput, nil
}

func (d *DynamoDBAPIMock) Scan(si *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if len(si.ExpressionAttributeValues) <= 0 {
		return nil, nilScanInput
	}
	return d.ScanOutput, nil
}

type saveTestCase struct {
	payload *partners.Partner
	conn    *DynamoDBAPIMock
	err     error
}

var saveTestCases = map[string]saveTestCase{
	"Success": {
		payload: &partners.Partner{ID: "test"},
		conn: &DynamoDBAPIMock{
			PutItemOutput: &dynamodb.PutItemOutput{
				Attributes: map[string]*dynamodb.AttributeValue{
					"id": {
						S: aws.String("test"),
					},
				},
			},
		},
		err: nil,
	},
	"Error": {
		payload: nil,
		conn:    &DynamoDBAPIMock{},
		err:     nilPutItemInput,
	},
}

func TestAWSDocDBPartnerAdapter_SavePartner(t *testing.T) {
	for caseName, tCase := range saveTestCases {
		svc := NewAWSDocDBPartnerAdapter("test", tCase.conn)
		caseResult, err := svc.SavePartner(tCase.payload)
		if err != tCase.err {
			t.Errorf("case: %s\n expected: %+e\n got: %+e\n", caseName, tCase.err, err)
		}
		if err == nil && caseResult != nil && caseResult.ID != tCase.payload.ID {
			t.Errorf("case: %s\n expected: %+v\n got: %+v\n", caseName, *tCase.payload, *caseResult)
		}
	}
}

type getTestCase struct {
	payload  string
	expected *partners.Partner
	conn     *DynamoDBAPIMock
	err      error
}

var getTestCases = map[string]getTestCase{
	"Success": {
		payload:  "test",
		expected: &partners.Partner{ID: "test"},
		conn: &DynamoDBAPIMock{
			GetItemOutput: &dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					"id": {
						S: aws.String("test"),
					},
				},
			},
		},
		err: nil,
	},
	"Error": {
		payload:  "",
		conn:     &DynamoDBAPIMock{},
		expected: nil,
		err:      nilGetItemInput,
	},
}

func TestAWSDocDBPartnerAdapter_GetPartner(t *testing.T) {
	for caseName, tCase := range getTestCases {
		svc := NewAWSDocDBPartnerAdapter("test", tCase.conn)
		caseResult, err := svc.GetPartner(tCase.payload)
		if err != tCase.err {
			t.Errorf("case: %s\n expected: %+e\n got: %+e\n", caseName, tCase.err, err)
		}
		if err == nil && caseResult != nil && reflect.DeepEqual(*caseResult, *tCase.expected) {
			t.Errorf("case: %s\n expected: %+v\n got: %+v\n", caseName, tCase.payload, caseResult)
		}
	}
}

type searchTestCase struct {
	payload  *partners.Point
	expected []partners.Partner
	conn     *DynamoDBAPIMock
	err      error
}

var searchTestCases = map[string]searchTestCase{
	"Success": {
		payload:  &partners.Point{},
		expected: []partners.Partner{{ID: "test"}},
		conn: &DynamoDBAPIMock{
			ScanOutput: &dynamodb.ScanOutput{
				Items: []map[string]*dynamodb.AttributeValue{{
					"id": {
						S: aws.String("test"),
					},
				}},
			},
		},
		err: nil,
	},
	"Error": {
		payload:  nil,
		conn:     &DynamoDBAPIMock{},
		expected: nil,
		err:      nilScanInput,
	},
}

func TestAWSDocDBPartnerAdapter_SearchPartners(t *testing.T) {
	for caseName, tCase := range searchTestCases {
		svc := NewAWSDocDBPartnerAdapter("test", tCase.conn)
		caseResult, err := svc.SearchPartners(tCase.payload)
		if err != tCase.err {
			t.Errorf("case: %s\n expected: %+e\n got: %+e\n", caseName, tCase.err, err)
		}
		if caseResult == nil {
			t.Errorf("case: %s\n expected: %+v\n got: %+v\n", caseName, tCase.payload, caseResult)
		}
		if err == nil && caseResult != nil && reflect.DeepEqual(caseResult, tCase.expected) {
			t.Errorf("case: %s\n expected: %+v\n got: %+v\n", caseName, tCase.payload, caseResult)
		}
	}
}
