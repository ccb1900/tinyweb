package upload

import (
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/upload/contract"
	"github.com/ccb1900/tinyweb/upload/driver/local"
	"mime/multipart"
)
var u contract.IUpload
func Init()  {
	switch config.Get("upload.driver") {
	case "local":
		u = local.New()
	}
}

func Upload(file *multipart.FileHeader, dst string) error  {
	return u.Upload(file,dst)
}
