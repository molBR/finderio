package service

import (
	"finderio/cmd/out/dynamoDB"
	"finderio/cmd/setup"
	"finderio/cmd/util"
	"log"
)

func Service(jsonString string, confSetup *setup.ConfSetup) (map[string]interface{}, error) {

	objectQuery := util.GraphqlUnmarshal(jsonString)
	log.Println(objectQuery)
	queryData := out_dynamodb.QueryData{objectQuery.TableName, objectQuery.HashKey, objectQuery.HashValue}
	queryResult, err := out_dynamodb.GetItem(confSetup.DynamoClient, queryData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	objectResult := out_dynamodb.UnmarshalDynamoAttributes(queryResult)
	log.Println(queryResult)
	returnableObject := make(map[string]interface{})
	for _, value := range objectQuery.Items {
		returnableObject[value] = objectResult[value]
	}
	log.Println(returnableObject)
	return returnableObject, nil
}
