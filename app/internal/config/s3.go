package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type S3 struct {
	config *Config
}

func NewS3(config *Config) *S3 {
	return &S3{config}
}

func (s *S3) getConfig() *aws.Config {
	return &aws.Config{
		Endpoint: &s.config.S3EndpointUrl,
		Credentials: credentials.NewStaticCredentials(
			s.config.S3AccessKey,
			s.config.S3SecretKey,
			s.config.S3AccessKey,
		),
		Region: aws.String("ru-1"),
	}
}

func (s *S3) getSession(c *aws.Config) *session.Session {
	return session.Must(session.NewSession(c))
}

func (s *S3) Upload(pathToS3 string, file *os.File) (*s3manager.UploadOutput, error) {
	c := s.getConfig()
	sess := s.getSession(c)
	uploader := s3manager.NewUploader(sess)
	return uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.config.S3BucketName),
		Key:    aws.String(pathToS3),
		Body:   file,
	})
}

func (s *S3) Delete(pathToS3 string) error {
	c := s.getConfig()
	sess := s.getSession(c)
	svc := s3.New(sess)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.config.S3BucketName),
		Key:    aws.String(pathToS3),
	}
	_, err := svc.DeleteObject(input)
	if err != nil {
		return err
	}
	return nil
}
