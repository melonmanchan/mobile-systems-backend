package aws

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"

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

func (u Uploader) uploadToAWS(file io.Reader, key string, contentType string) (putOutput *s3manager.UploadOutput, err error) {
	uploadResult, err := u.Upload(&s3manager.UploadInput{
		Bucket:      &u.cfg.Bucket,
		ContentType: &contentType,
		Key:         &key,
		Body:        file,
	})

	return uploadResult, err
}

func (u Uploader) UploadAvatar(file io.Reader) (url string, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contentType := http.DetectContentType(buf.Bytes())

	types, err := mime.ExtensionsByType(contentType)

	if err != nil {
		return "", err
	}

	name := "asdasdasd" + types[0]

	_, err = u.uploadToAWS(file, "avatars/"+name, contentType)

	if err != nil {
		return "", err
	}

	return u.getAvatarURL(name), err
}
