package in_graphql

import (
	"encoding/json"
	"finderio/cmd/service"
	"finderio/cmd/setup"
	"io/ioutil"
	"log"
	"net/http"
)

type errorResponse struct {
	message string
}

func CreateServer(confSetup *setup.ConfSetup) {

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		body, _ := ioutil.ReadAll(r.Body)
		jsonString := string(body)
		result, err := service.Service(jsonString, confSetup)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			eJson := make(map[string]interface{})
			eJson["ERROR_MESSAGE"] = err.Error()
			encoder.Encode(eJson)
		} else {
			encoder.Encode(result)
		}
		return

	})
	log.Println("Server running on PORT: 8080")
	http.ListenAndServe(":8080", nil)

}
