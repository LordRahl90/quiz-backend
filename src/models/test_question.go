package models

import (
	"github.com/jinzhu/gorm"
)

//Option struct
type Option struct {
	A          string       `json:"a"`
	B          string       `json:"b"`
	C          string       `json:"c"`
	D          string       `json:"d"`
	QuestionID uint         `json:"question_id"`
	Question   TestQuestion `gorm:"foreignkey:QuestionID" json:"question"`
	gorm.Model
}

//TestQuestion struct
type TestQuestion struct {
	Question       string `json:"question"`
	Option         string `json:"option"`
	Section        string `json:"section"`
	Answer         string `json:"answer"`
	Solution       string `json:"solution"`
	TestQuestionID uint   `json:"test_question_id"`
	TestDetailID   uint
	TestDetail     TestDetail `gorm:"foreignkey:TestDetailID" json:"test_detail"`
	gorm.Model
}
