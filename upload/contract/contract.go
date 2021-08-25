package contract

import "mime/multipart"

type IUpload interface {
	Upload(file *multipart.FileHeader, dst string) error
}
