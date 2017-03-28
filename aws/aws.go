package aws

import (
	"fmt"
	"os"

	"../config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Uploader ...
type Uploader struct {
	*s3manager.Uploader
	cfg config.S3Config
}

// BuildAWSUploader ...
func BuildAWSUploader(cfg config.S3Config) Uploader {
	s3Uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String(cfg.Region)}))
	return Uploader{s3Uploader, cfg}
}

func (u Uploader) getS3URL(fileName string) string {
	return fmt.Sprintf("https://s3.%s.amazonaws.com/%s/%s", u.cfg.Region, u.cfg.Bucket, fileName)
}

func (u Uploader) getAvatarURL(fileName string) string {
	return u.getS3URL("avatars/" + fileName)
}

func (u Uploader) uploadToAWS(fileName string, key string, contentType string) (putOutput *s3manager.UploadOutput, err error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	uploadResult, err := u.Upload(&s3manager.UploadInput{
		Bucket:      &u.cfg.Bucket,
		ContentType: &contentType,
		Key:         &key,
		Body:        file,
	})

	return uploadResult, err
}
