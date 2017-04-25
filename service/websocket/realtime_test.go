package service

import (
	"testing"
	"time"
	"fmt"
	"github.com/gorilla/websocket"
)

func Test_getRealTotalCount(t *testing.T)  {

	st := int64(time.Now().Unix())
	ed := int64(time.Now().Unix() + 1800)
	re := getRealTotalCount(st,ed)
	fmt.Println(re)
}
