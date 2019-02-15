package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/LordRahl90/little_quiz_backend/src/dataobjects"
)

//Message - This function composes the response format.
func Message(status bool, message string) map[string]interface{} {
	response := make(map[string]interface{})
	response["status"] = status
	response["message"] = message
	return response
}

//Respond - This function returns the actual response as retrieved from the server
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getClient() *http.Client {
	client := &http.Client{}
	return client
}

//M - Just to make an array of useless interface.
type M map[string]interface{}

//GetQuestions - Retrieve Questions from the API.
func GetQuestions(request *dataobjects.InitiateRequest) map[string]interface{} {
	url := fmt.Sprintf("http://questions.aloc.ng/api/q/%d/?subject=%s", request.Questions, request.Subject)
	client := getClient()

	externalResponse, err := client.Get(url)
	if err != nil {
		return Message(false, "Connection Error while retrieving questions")
	}

	responseData, err := ioutil.ReadAll(externalResponse.Body)
	if err != nil {
		return Message(false, "Bad Response from the question server")
	}

	data := make(map[string]interface{})
	outputData := []M{}
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return Message(false, "Cannont Convert the response accordingly")
	}

	questionsArray := data["data"].([]interface{})
	for _, v := range questionsArray {
		newValue := v.(map[string]interface{})
		fmt.Println(newValue) //TODO: Initiate the saving at this point....
		newValue["answer"] = ""
		outputData = append(outputData, newValue)
	}

	response := Message(true, "Questions Loaded Successfully.")
	response["questions"] = outputData

	response["duration"] = request.Duration
	return response
}
