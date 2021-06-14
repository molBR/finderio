package util

import (
	"log"
	"strings"
)

type QueryObject struct {
	TableName string
	HashKey   string
	HashValue string
	Items     []string
}

func getTableName(jsonString string) (string, string) {
	i := strings.Index(jsonString, "query")
	jsonString = jsonString[i:]
	i = strings.Index(jsonString, "{")
	jsonString = jsonString[i+3:]
	i = 0
	for jsonString[i] == ' ' {
		i++
	}
	auxString := jsonString[i:]
	i = 0
	for auxString[i] != ' ' {
		i++
	}
	tableName := auxString[:i]
	log.Println("TABLE NAME", tableName)
	jsonString = auxString[i:]
	return jsonString, tableName
}

func getHashKey(jsonString string) (string, string) {
	i := strings.Index(jsonString, "(")
	jsonString = jsonString[i+1:]
	i = strings.Index(jsonString, ":")
	hashKey := jsonString[:i]
	log.Println("HASH KEY", hashKey)
	jsonString = jsonString[i+3:]
	return jsonString, hashKey
}

func getHashValue(jsonString string) (string, string) {
	i := strings.Index(jsonString, ")")
	hashValue := jsonString[:i-2]
	log.Println("HASH VALUE", hashValue)
	jsonString = jsonString[i:]
	return jsonString, hashValue
}

func getItems(jsonString string, valueArray []string) (string, []string) {
	i := strings.Index(jsonString, "n")
	jsonString = jsonString[i+1:]
	i = 0
	for jsonString[i] == ' ' {
		i++
	}
	if jsonString[i] == '}' {
		return jsonString, valueArray
	}
	jsonString = jsonString[i:]
	i = strings.Index(jsonString, "n")
	value := jsonString[:i-1]
	log.Println("VALOR:", value)
	valueArray = append(valueArray, value)
	return getItems(jsonString, valueArray)

}

func GraphqlUnmarshal(jsonString string) *QueryObject {

	jsonString, tableName := getTableName(jsonString)
	jsonString, hashKey := getHashKey(jsonString)
	jsonString, hashValue := getHashValue(jsonString)
	var valueArray []string
	jsonString, valueArray = getItems(jsonString, valueArray)
	return &QueryObject{TableName: tableName,
		HashKey: hashKey, HashValue: hashValue, Items: valueArray}

}
