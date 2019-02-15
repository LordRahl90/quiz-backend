package models

import (
	"github.com/jinzhu/gorm"
)

//Role Struct. To Manage each user role management
type Role struct {
	gorm.Model
	Name        string       `json:"name"`
	DisplayName string       `json:"display_name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions"`
	UserRole    []UserRole   `gorm:"foreignkey:RoleID"`
}
