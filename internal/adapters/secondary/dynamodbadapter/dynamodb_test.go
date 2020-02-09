package dynamodbadapter

import (
	"testing"
)

func TestCreateDynamoSess(t *testing.T) {
	sess := CreateDynamoSess()
	if sess == nil {
		t.Error("Session isn't a DynamoDBAPI")
	}
}
