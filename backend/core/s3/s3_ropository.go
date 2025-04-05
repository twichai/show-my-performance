package s3Core

import "mime/multipart"

type S3Repository interface {
	UploadFile(file []*multipart.FileHeader, fileName string) (string, error)
}
