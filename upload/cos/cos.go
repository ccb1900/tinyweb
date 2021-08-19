package cos

import (
	"context"
	"fmt"
	"github.com/ccb1900/tinyweb/config"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func start()  {
	u, _ := url.Parse("https://"+config.Get("upload.cos.bucket")+".cos."+config.Get("upload.cos.region")+".myqcloud.com")
	su, _ := url.Parse("https://cos."+config.Get("region")+".myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Get("upload.cos.id"),
			SecretKey: config.Get("upload.cos.key"),
		},
	})

	name := "hello/world"
	f := strings.NewReader("test hello world")
	r, err := client.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	bb,e := ioutil.ReadAll(r.Body)
	log.Println("sss",string(bb),e)
	opt := &cos.BucketGetOptions{
		Prefix:  "hello",
		//MaxKeys: 3,
	}
	v, _, err := client.Bucket.Get(context.Background(), opt)
	if err != nil {
		log.Println("eee")
		panic(err)
	}

	for _, c := range v.Contents {
		fmt.Printf("%s, %d\n", c.Key, c.Size)
	}

	fmt.Println(client.Object.GetObjectURL("hello/world"))
}
