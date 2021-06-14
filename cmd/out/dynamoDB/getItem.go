package out_dynamodb

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

type QueryData struct {
	TableName string
	KeyName   string
	Value     string
}

func UnmarshalDynamoAttributes(atributeMap map[string]*dynamodb.AttributeValue) map[string]interface{} {

	var f map[string]interface{}
	err := dynamodbattribute.UnmarshalMap(atributeMap, &f)
	if err != nil {
		panic(err)
	}
	return f
}

func GetItem(svc *dynamodb.DynamoDB, query QueryData) (map[string]*dynamodb.AttributeValue, error) {

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(query.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			query.KeyName: {
				S: aws.String(query.Value),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find '" + query.Value + "'"
		log.Println(msg)
		return nil, errors.New(msg)

	}

	return result.Item, nil
}
