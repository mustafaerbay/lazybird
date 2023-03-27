package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaerbay/lazybird/initializers"
	"github.com/mustafaerbay/lazybird/models"
)



func RolesCreate(c *gin.Context) {
	// Get data
	var body struct {
		Name  string
		Permissions []models.Permission
	}
	c.Bind(&body)
	// create a post
	role := models.Role{
		Name: body.Name,
		Permissions: body.Permissions,
	}
	// pass pointer of data to Create

	result := initializers.DB.Create(&role)
	//
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"role": role,
	})
}

func RoleIndex(c *gin.Context) {
	var roles []models.Role
	// Get the posts

	initializers.DB.Find(&roles)
	for i := range roles {
		initializers.DB.Model(&roles[i]).Association("Permissions").Find(&roles[i].Permissions)
	}
	//respond
	c.JSON(200, gin.H{
		"roles": roles,
	})
}

func RolesShow(c *gin.Context) {
	// get id of url
	id := c.Param("id")
	var role models.Role
	// Get the posts

	initializers.DB.First(&role, id)

	//respond
	c.JSON(200, gin.H{
		"post": role,
	})
}

func RolessUpdate(c *gin.Context) {
	//get the id of url
	id := c.Param("id")

	// get the data off req body
	var body struct {
		Name  string
		Permissions []models.Permission
	}
	c.Bind(&body)

	var role models.Role
	initializers.DB.First(&role, id)

	// find the post were updating
	initializers.DB.Model(&role).Updates(models.Role{
		Name: body.Name,
		Permissions: body.Permissions,
	})

	// db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	//respond
	c.JSON(200, gin.H{
		"post": role,
	})
}

func RolesDelete(c *gin.Context) {
	id := c.Param("id")

	// find the post
	var role models.Role
	initializers.DB.First(&role, id)

	// delete the post
	initializers.DB.Delete(&models.Role{}, id)

	c.JSON(200, gin.H{
		"role": "deleted successfully",
	})
}


