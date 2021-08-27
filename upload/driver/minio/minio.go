package minio

import (
	"context"
	"github.com/ccb1900/tinyweb/log"
	"github.com/ccb1900/tinyweb/upload/contract"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
)

type Minio struct {
	M *minio.Client
}

func (m Minio) Upload(file *multipart.FileHeader, dst string) error {
	ctx := context.Background()
	bucketName:="lianyou"
	if err := m.M.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "cn-wuhan-1"});err != nil {
		exists, errBucketExists := m.M.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Debug("We already own %s\n", bucketName)
		} else {
			return err
		}
	}
	objName:="minio.go"
	objPath := "./"+objName
	info, err := m.M.FPutObject(ctx,bucketName,objName,objPath,minio.PutObjectOptions{
		ContentType: "application/zip",
	})
	if err != nil {
		return err
	}
	log.Debug("info",info)
	return nil
}

func New() contract.IUpload  {
	ak := "dev001001"
	sk :="dev001001"
	endpoint :="127.0.0.1:9012"
	c,err:=minio.New(endpoint,&minio.Options{
		Creds:        credentials.NewStaticV4(ak,sk,""),
		Secure:       false,
	})
	if err !=nil {
		log.Debug(err)
		return nil
	}
	log.Debug("connected...")
	return &Minio{
		M:c,
	}
}
