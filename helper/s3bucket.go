package helper

import (
	"log"
	"mime/multipart"
	"time"

	"github.com/Zenk41/sipencari-rest-api/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

func UploadToS3(c echo.Context, folder string, filename string, src multipart.File) (string, error) {
	SECRET_KEY := util.GetEnv("AWS_S3_BUCKET_SECRET_KEY")
	KEY_ID := util.GetEnv("AWS_S3_BUCKET_KEY_ID")
	REGION := util.GetEnv("AWS_S3_REGION")
	BUCKET_NAME := util.GetEnv("AWS_S3_BUCKET_NAME")

	configS3 := &aws.Config{
		Region:      aws.String(REGION),
		Credentials: credentials.NewStaticCredentials(KEY_ID, SECRET_KEY, ""),
	}
	s3Session := session.New(configS3)
	uploader := s3manager.NewUploader(s3Session)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(folder + filename),
		Body:   src,
	})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return result.Location, nil
}

func MultipleUploadS3(c echo.Context, files []*multipart.FileHeader, path string) ([]string, error) {

	SECRET_KEY := util.GetEnv("AWS_S3_BUCKET_SECRET_KEY")
	KEY_ID := util.GetEnv("AWS_S3_BUCKET_KEY_ID")
	REGION := util.GetEnv("AWS_S3_REGION")
	BUCKET_NAME := util.GetEnv("AWS_S3_BUCKET_NAME")

	configS3 := &aws.Config{
		Region:      aws.String(REGION),
		Credentials: credentials.NewStaticCredentials(KEY_ID, SECRET_KEY, ""),
	}

	var urlImages []string
	var err error

	for i := len(files) - 1; i >= 0; i-- {
		src, err := files[i].Open()

		if err != nil {
			log.Println(err)
			return urlImages, err
		}

		fileName := time.Now().String() + ".png"

		s3Session := session.New(configS3)
		uploader := s3manager.NewUploader(s3Session)

		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(BUCKET_NAME),
			Key:    aws.String(path + fileName),
			Body:   src,
		})

		if err != nil {
			log.Fatal(err)
			return urlImages, err
		}

		url := result.Location

		urlImages = append(urlImages, url)
		defer src.Close()
	}
	return urlImages, err
}
