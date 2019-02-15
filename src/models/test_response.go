package models

import (
	"github.com/jinzhu/gorm"
)

//TestResponse struct
type TestResponse struct {
	UserID         int `json:"user_id"`
	TestDetailID   uint
	TestQuestionID uint
	Response       string       `json:"response"`
	PassOrFail     string       `json:"pass_or_fail"`
	TestDetail     TestDetail   `json:"test_detail" gorm:"foreignkey:TestDetailID"`
	TestQuestion   TestQuestion `json:"test_question" gorm:"foreignkey:TestQuestionID"`
	gorm.Model
}
