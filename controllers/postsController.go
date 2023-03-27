package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaerbay/lazybird/initializers"
	"github.com/mustafaerbay/lazybird/models"
)



func PostsCreate(c *gin.Context) {
	// Get data
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Title,
	}
	// pass pointer of data to Create

	result := initializers.DB.Create(&post)
	//
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	// Get the posts

	initializers.DB.Find(&posts)

	//respond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id of url
	id := c.Param("id")
	var post models.Post
	// Get the posts

	initializers.DB.First(&post, id)

	//respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//get the id of url
	id := c.Param("id")

	// get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	// find the post were updating
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: post.Body,
	})

	// db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	//respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	// find the post
	var post models.Post
	initializers.DB.First(&post, id)

	// delete the post
	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"post": "deleted successfully",
	})
}


