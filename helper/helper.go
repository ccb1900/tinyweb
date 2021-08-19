package helper

import (
	"crypto/md5"
	"fmt"
	"os"
)

// IsFileExist 文件是否存在
func IsFileExist(name string) bool  {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// Md5 计算MD5
func Md5(s string) string  {
	return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}
