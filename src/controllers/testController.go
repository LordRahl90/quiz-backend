package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/LordRahl90/little_quiz_backend/src/dataobjects"
	"github.com/LordRahl90/little_quiz_backend/src/models"
	"github.com/LordRahl90/little_quiz_backend/src/services"
	"github.com/LordRahl90/little_quiz_backend/src/utils"
)

//InitiateTest Initiates the test
func InitiateTest(w http.ResponseWriter, r *http.Request) {
	initiate := &dataobjects.InitiateRequest{}
	initiate.UserID = r.Context().Value("user").(uint)
	err := json.NewDecoder(r.Body).Decode(initiate)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Request"))
	}
	go func() {
		newTest := &models.TestDetail{
			UserID:      int(initiate.UserID),
			Subject:     initiate.Subject,
			Examination: "",
			Year:        "",
			Duration:    int(initiate.Duration),
		}

		models.GetDB().Create(newTest)
		initiate.TestDetailID = newTest.ID
	}()

	initiate.Subject = strings.ToLower(initiate.Subject)
	response := utils.GetQuestions(initiate)
	go services.SaveUserQuestion(initiate, response)

	utils.Respond(w, response)
}

//MarkTest funcion to mark the required test
func MarkTest(w http.ResponseWriter, r *http.Request) {
	myAnswers := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&myAnswers)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Response Format"))
		return
	}

	testDetail := myAnswers["test_detail_id"].(float64)
	userID := int(r.Context().Value("user").(uint))
	questionResponses := myAnswers["questions"].([]interface{})

	response := services.MarkExam(userID, int(testDetail), questionResponses)

	utils.Respond(w, response)
}
