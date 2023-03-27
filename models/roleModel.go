package models

// import "gorm.io/gorm"

type Role struct {
	// gorm.Model
	ID          uint         `json:"id" gorm:"primary_key"`
	Name        string       `gorm:"unique;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
