package models

// import "gorm.io/gorm"

type User struct {
	// gorm.Model
	ID       uint `json:"id" gorm:"primary_key"`
	Name     string
	Lastname string
	Email    string `gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	RoleID	 uint  `gorm:"foreign_key"`
}
