package services

import (
	"fmt"

	"github.com/LordRahl90/little_quiz_backend/src/dataobjects"
	"github.com/LordRahl90/little_quiz_backend/src/models"
	"github.com/LordRahl90/little_quiz_backend/src/utils"
	"github.com/labstack/gommon/log"
)

//SaveUserQuestion - save user question and information
func SaveUserQuestion(initiateObject *dataobjects.InitiateRequest, data map[string]interface{}) {
	questions := data["questions"].([]utils.M)
	for _, v := range questions {
		go func(v interface{}) {
			question := v.(utils.M)

			// lets save question first.
			questionPayload := &models.TestQuestion{
				TestDetailID:   initiateObject.TestDetailID,
				Question:       question["question"].(string),
				Section:        question["section"].(string),
				Answer:         question["answer"].(string),
				Solution:       question["solution"].(string),
				TestQuestionID: uint(question["id"].(float64)),
			}

			models.GetDB().Create(questionPayload)

			options := question["option"].(map[string]interface{})
			optionPayload := &models.Option{
				QuestionID: questionPayload.ID,
				A:          options["a"].(string),
				B:          options["b"].(string),
				C:          options["c"].(string),
				D:          options["d"].(string),
			}
			models.GetDB().Create(optionPayload)
		}(v)
	}
}

//MarkExam - function to mark exam and return data
func MarkExam(userID int, testDetailID int, responses []interface{}) map[string]interface{} {
	passed := 0
	failed := 0

	testDetail := &models.TestDetail{}

	err := models.GetDB().Table("test_details").Where("id=?", testDetailID).First(testDetail).Error
	if err != nil {
		return utils.Message(false, "Test Information not found")
	}

	fmt.Println(testDetail)

	for _, v := range responses {
		question := v.(map[string]interface{})
		questionID := uint(question["id"].(float64))
		passOrFail := "fail"

		tempQuestion := &models.TestQuestion{}
		err = models.GetDB().Table("test_questions").Where("test_detail_id=? and test_question_id=?", testDetailID, questionID).First(&tempQuestion).Error
		if err != nil {
			log.Info("Question Missed, doesnt exist or deleted")
			continue
		}
		// fmt.Println(tempQuestion.Answer)
		answerResponse := question["answer"].(string)
		if answerResponse == "" || answerResponse != tempQuestion.Answer {
			failed++
		} else {
			passed++
			passOrFail = "pass"
		}

		//lets keep this response for future/audit use.
		testResponse := &models.TestResponse{
			UserID:         userID,
			TestDetailID:   uint(testDetailID),
			TestQuestionID: questionID,
			Response:       answerResponse,
			PassOrFail:     passOrFail,
		}

		go models.GetDB().Create(testResponse)
	}

	percentage := (float64(passed) / float64(len(responses))) * 100

	testScore := &models.TestScore{
		Score:        uint(passed),
		UserID:       uint(userID),
		TestDetailID: uint(testDetailID),
		Percentage:   percentage,
	}

	go models.GetDB().Create(testScore)

	resp := utils.Message(true, "Completed")
	resp["data"] = map[string]interface{}{
		"details": testDetail,
		"total":   len(responses),
		"passed":  passed,
		"failed":  failed,
	}

	return resp
}
