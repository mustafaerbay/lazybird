package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaerbay/lazybird/initializers"
	"github.com/mustafaerbay/lazybird/models"
)



func PermissionsCreate(c *gin.Context) {
	// Get data
	var body struct {
		Name  string
		Definition string
	}
	c.Bind(&body)
	// create a post
	permission := models.Permission{
		Name: body.Name,
		Definition: body.Definition,
	}
	// pass pointer of data to Create

	result := initializers.DB.Create(&permission)
	//
	if result.Error != nil {
		// c.Status(400)
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"permission": permission,
	})
}

func PermissionIndex(c *gin.Context) {
	var permissions []models.Permission
	// Get the posts

	initializers.DB.Find(&permissions)

	//respond
	c.JSON(200, gin.H{
		"posts": permissions,
	})
}

func PermissionsShow(c *gin.Context) {
	// get id of url
	id := c.Param("id")
	var permission models.Permission
	// Get the posts

	initializers.DB.First(&permission, id)

	//respond
	c.JSON(200, gin.H{
		"post": permission,
	})
}

func PermissionUpdate(c *gin.Context) {
	//get the id of url
	id := c.Param("id")

	// get the data off req body
	var body struct {
		Name  string
		Definition string
	}
	c.Bind(&body)

	var permission models.Permission
	initializers.DB.First(&permission, id)

	// find the post were updating
	initializers.DB.Model(&permission).Updates(models.Permission{
		Name: body.Name,
		Definition: body.Definition,
	})

	// db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	//respond
	c.JSON(200, gin.H{
		"post": permission,
	})
}

func PermissionsDelete(c *gin.Context) {
	id := c.Param("id")

	// find the post
	var permission models.Permission
	initializers.DB.First(&permission, id)

	// delete the post
	initializers.DB.Delete(&models.Permission{}, id)

	c.JSON(200, gin.H{
		"permission": "deleted successfully",
	})
}


