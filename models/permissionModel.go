package models

// import "gorm.io/gorm"

type Permission struct {
	// gorm.Model
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `gorm:"unique;not null"`
	Definition string
	Roles      []Role `gorm:"many2many:role_permissions;"`
}

// Role []*Role `gorm:"many2many:role_permissions;"`
