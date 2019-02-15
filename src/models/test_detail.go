package models

import (
	"github.com/jinzhu/gorm"
)

//TestDetail struct
type TestDetail struct {
	UserID      int    `json:"user_id"`
	Subject     string `json:"subject"`
	Examination string `json:"examination"`
	Year        string `json:"year"`
	Duration    int    `json:"duration"`
	gorm.Model
	TestQuestions []TestQuestion `json:"test_questions" gorm:"foreignkey:TestDetailID"`
}
