package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var s3Session *session.Session
var bucketName string

func Init() error {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("https://cloudfront.net/"),
	})

	if err != nil {
		return err
	}

	s3Session = sess

	bucketName = "d3i847k00a14br"

	return nil
}

func newUploader() *s3manager.Uploader {
	return s3manager.NewUploader(s3Session)
}

func newDownloader() *s3manager.Downloader {
	return s3manager.NewDownloader(s3Session)
}
