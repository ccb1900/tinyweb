package helper

import "testing"

func TestIsFileExist(t *testing.T) {
	t.Log(IsFileExist("helper.go"))
	t.Log(Md5("hello"))
}
