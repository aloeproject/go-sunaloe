package library

import (
	"github.com/astaxie/beego"
	"fmt"
)

func ShortArticleContent(content string) string{
	var ct string
	length := 200
	if len(content) >= length {
		ct = fmt.Sprintf("%s...",beego.Substr(content,0,length))
	} else {
		ct = content
	}
	return ct
}
