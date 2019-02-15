package models

import "github.com/jinzhu/gorm"

//Permission struct. This is to validate the permission information
type Permission struct {
	gorm.Model
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	RoleID      int
	Role        Role `json:"role" gorm:"foreignkey:RoleID"`
}
