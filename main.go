package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/mustafaerbay/lazybird/controllers"
	"github.com/mustafaerbay/lazybird/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectToS3()
}
func main() {

	r := gin.Default()
	r.POST("/permissions", controllers.PermissionsCreate)
	r.GET("/permissions", controllers.PermissionIndex)

	r.POST("/roles", controllers.RolesCreate)
	r.GET("/roles", controllers.RoleIndex)

	r.POST("/posts", controllers.PostsCreate)
	
	r.PUT("/posts/:id", controllers.PostsUpdate)
	
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	
	r.POST("/s3/upload", controllers.FileCreateS3)
	r.DELETE("/s3/upload", controllers.FileDeleteS3)
	fmt.Println(os.Getenv("AWS_S3_BUCKET_NAME"))
	r.Run() // listen and serve on 0.0.0.0:8080
}