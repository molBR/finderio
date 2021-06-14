package util

import (
	"log"
	"strings"
)

func getItem(jsonString string) string {

	i := strings.Index(jsonString, string(rune(34)))
	log.Println(jsonString[i:])
	return jsonString
}

func DynamoUnmarshal(jsonString string) string {

	return getItem(jsonString)

}
