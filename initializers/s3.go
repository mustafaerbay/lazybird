package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var S3Client *s3.Client

func ConnectToS3() {
	var err error
	bucketName := os.Getenv("AWS_S3_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")
	fmt.Println("#@############################")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	// create s3 client
	S3Client = s3.NewFromConfig(cfg)

	_, err = S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
        CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})
	if err != nil {
		// log.Default().Output(1,"BUCKET CREATED SUCCESSFULLYYYYYY")
		log.Printf("Error creating bucket %s , %s ",bucketName, err)
		// fmt.Println("Error creating bucket:", err)
		// os.Exit(1)
	}

	
	response, err := S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})
	if err != nil {
		fmt.Println("Error listing objects in bucket:", err)
		os.Exit(1)
	}

	// Print object names
	for _, obj := range response.Contents {
		fmt.Println(*obj.Key)
	}

	// bucketname := "test321123ascdf12edfsacf"

}
