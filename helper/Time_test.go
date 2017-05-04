package helper

import (
	"testing"
	"fmt"
)

func TestGetUploadImageDir(t *testing.T) {
	st := GetUploadImageDir()
	fmt.Println(st)
}
