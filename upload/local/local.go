package local

import (
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/helper"
	"github.com/ccb1900/tinyweb/log"
	"io"
	"mime/multipart"
	"os"
)

func upload(file *multipart.FileHeader, dst string) error  {
	rootPath := config.Get("upload.local.path")

	if !helper.IsFileExist(rootPath) {
		err := os.MkdirAll(rootPath, 0777)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err

}
