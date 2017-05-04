package main

import (
	_ "myweb/routers"
	"github.com/astaxie/beego"
	//"html/template"
	//"os"
	"myweb/helper"
	"github.com/astaxie/beego/orm"
	"myweb/controllers/frontend"
)

func main() {
	orm.Debug = true
	//默认模板后缀
	beego.AddTemplateExt("html")
	//静态目录设置
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")
	beego.SetStaticPath("/static","static")
	beego.SetStaticPath("/img","static/img")
	//定义错误页面
	beego.ErrorController(&frontend.ErrorController{})

	beego.AddFuncMap("ShortArticleContent", helper.ShortArticleContent)
	beego.Run()
}

