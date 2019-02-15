package models

import (
	"github.com/jinzhu/gorm"
)

//UserRole - Struct to manage user roles.
type UserRole struct {
	UserID int  `json:"user_id"`
	RoleID int  `json:"role_id"`
	Role   Role `gorm:"foreignkey:RoleID" json:"role"`
	gorm.Model
}
