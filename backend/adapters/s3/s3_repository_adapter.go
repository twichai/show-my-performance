package s3Adapter

import (
	"mime/multipart"
)

type s3Repository struct {
	bucketName string
	region     string
	accessKey  string
}

// UploadFile implements core.S3Repository.
func (s *s3Repository) UploadFile(file []*multipart.FileHeader, fileName string) (string, error) {
	mockURL := "https://images.unsplash.com/photo-1742407795182-144225af8ebe?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
	return mockURL, nil
}

func NewS3Repository(bucketName string, region string) *s3Repository {
	return &s3Repository{bucketName: bucketName, region: region, accessKey: "your-access-key"}
}
