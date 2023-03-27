package main

import (
	"github.com/mustafaerbay/lazybird/initializers"
	"github.com/mustafaerbay/lazybird/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// To create database tables
	initializers.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Post{},
		&models.Permission{},
	)
}
