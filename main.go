package main

import (
	_ "myweb/routers"
	"github.com/astaxie/beego"
	//"html/template"
	//"os"
	"myweb/helper"
	"github.com/astaxie/beego/orm"
)

func main() {
	//模板嵌套
	//s1,_ := template.ParseFiles("/usr/local/go/src/myweb/views/layout/header.tpl", "myweb/views/layout/content.tpl", "myweb/views/layout/footer.tpl")
	/*
	s1,_ := template.ParseFiles("views/layout/header.tpl","views/layout/content.tpl","views/layout/footer.tpl")
	s1.ExecuteTemplate(os.Stdout,"header",nil)
	s1.ExecuteTemplate(os.Stdout,"content",nil)
	s1.ExecuteTemplate(os.Stdout,"footer",nil)
	s1.Execute(os.Stdout,nil)*/
	orm.Debug = true
	//默认模板后缀
	beego.AddTemplateExt("html")
	//静态目录设置
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")
	beego.SetStaticPath("/static","static")
	beego.SetStaticPath("/img","static/img")

	beego.AddFuncMap("ShortArticleContent", helper.ShortArticleContent)
	beego.Run()
}

