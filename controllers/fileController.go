package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/mustafaerbay/lazybird/initializers"
)


func FileCreateS3(c *gin.Context) {
	BucketName := os.Getenv("AWS_S3_BUCKET_NAME")
	fmt.Println(BucketName)
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error retrieving file from form data: %s", err)
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Unable to process file upload."})
		return
	}
	f, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %s", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to process file upload."})
		return
	}
	defer f.Close()

	key := file.Filename
	uploader := manager.NewUploader(initializers.S3Client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Body:   f,
		Key:    &key,
		ACL:    "public-read",
	})
	if err != nil {
		log.Printf("Error uploading file to S3: %s", err)
		return
	}
	c.JSON(200, gin.H{
		"message": "File uploaded successfully!",
		"output":  result.Location,
	})
}

func FileDeleteS3(c *gin.Context) {
	var body struct {
		Key    string
		Bucket string
	}
	c.Bind(&body)
	result, err := initializers.S3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(body.Bucket),
		Key:    aws.String(body.Key),
	})
	if err != nil {
		log.Printf("Error deleting object from bucket %s : %s", body.Bucket, err)
	}

	c.JSON(200, gin.H{
		"post":   "deleted successfully",
		"result": result,
	})
}