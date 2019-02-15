package models

import (
	"github.com/jinzhu/gorm"
)

//TestScore struct
type TestScore struct {
	Score        uint
	Percentage   float64
	TestDetailID uint
	TestDetail   TestDetail `gorm:"foreignkey:TestDetailID" json:"test_detail"`
	UserID       uint
	User         Account `gorm:"foreignkey:UserID" json:"user"`
	gorm.Model
}
