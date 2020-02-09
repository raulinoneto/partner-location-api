package dynamodbadapter

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	tableName = os.Getenv("PARTNERS_TABLE_NAME")
	dbUrl     = os.Getenv("DB_URL")
)

func GetLocalSess() *dynamodb.DynamoDB {
	db := getSession()
	if hasTable(db) != nil {
		if err := migrateLocal(db); err != nil {
			panic(fmt.Errorf("dynamodb.createDynamoSess: %e", err))
		}
	}
	return db
}

func getSession() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(dbUrl),
	})
	if err != nil {
		panic(fmt.Errorf("dynamodb.createDynamoSess: %e", err))
	}
	return dynamodb.New(sess)
}

func hasTable(db *dynamodb.DynamoDB) error {
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}
	_, err := db.DescribeTable(input)
	if err != nil {
		return fmt.Errorf("dynamodb.createDynamoSess: %s", err.Error())
	}
	return nil
}

func migrateLocal(db *dynamodb.DynamoDB) error {
	definitions := []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("id"),
			AttributeType: aws.String("S"),
		},
	}

	keySchema := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("id"),
			KeyType:       aws.String("HASH"),
		},
	}
	pThroughtput := &dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(10),
		WriteCapacityUnits: aws.Int64(10),
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions:  definitions,
		KeySchema:             keySchema,
		ProvisionedThroughput: pThroughtput,
		TableName:             aws.String(tableName),
	}
	_, err := db.CreateTable(input)
	return err
}
