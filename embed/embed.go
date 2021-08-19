package embed

// 嵌入静态页面

import (
	"io/fs"
	"log"
"net/http"
)

type WebFs struct {
	fs http.FileSystem
}

func (w *WebFs) Open(name string) (http.File,error)  {
	log.Println(name)
	o, err := w.fs.Open(name)
	// 用于支持历史模式
	if err != nil {
		return w.fs.Open("index.html")
	}
	return o,nil
}

func NewWebFs(fsys fs.FS) *WebFs {
	return &WebFs{fs: http.FS(fsys)}
}

