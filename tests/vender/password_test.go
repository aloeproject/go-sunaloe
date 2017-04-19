package vender

import (
	"testing"
	"myweb/vendor"
	"fmt"
)

func TestGetPasswordHash(t *testing.T) {
	p := vendor.Password{}
	str := p.GetPasswordHash("123123")
	fmt.Println(str)
}
