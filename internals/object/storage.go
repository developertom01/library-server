package object

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/library-server/config"
)

type ObjectStorage struct {
	Client *s3.S3
}

func NewObjectStorage() (*ObjectStorage, error) {
	config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(config.SPACES_KEY, config.SPACES_SECRET, ""),
		Endpoint:         aws.String(config.SPACES_BUCKET_ENDPOINT),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(false),
	}
	newSession, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(newSession)
	return &ObjectStorage{
		Client: s3Client,
	}, nil
}

func (os ObjectStorage) PutFile(file io.ReadSeeker, uploadType string, path string) (string, error) {
	input := &s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(config.SPACES_BUCKET_NAME),
		ACL:    aws.String("public-read"),
		Key:    aws.String(path),
		Metadata: map[string]*string{
			"upload-type": aws.String(uploadType),
		},
	}
	_, err := os.Client.PutObject(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v/%v", config.SPACES_BUCKET_ENDPOINT, path), nil
}
