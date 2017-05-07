package frontend

import (
	"github.com/astaxie/beego"
	"myweb/helper"
	"myweb/repository"
)


type BaseController struct {
	beego.Controller
}

func (this *BaseController) init()  {
	this.Layout = "frontend/layout/layout.html"
}

func (this *BaseController) PageKeyword(list *repository.ArticleList){
	keywords := ""
	controllerName,actionName := this.GetControllerAndAction()
	if controllerName == "IndexController" {
		keywords = "php,golang,go语言,mysql,mongodb,redis,爬虫,beego"
		if actionName == "Detail" {
			keywords = list.Category_name+","+list.Title
		}
	}
	this.Data["keywords"] = keywords
}

func (this *BaseController) Prepare()  {
	//设置uuid cookie
	_,bl := this.GetSecureCookie(beego.AppConfig.String("cookie.secure"),"uuid")
	if bl == false {
		this.SetSecureCookie(beego.AppConfig.String("cookie.secure"),
			"uuid",
			helper.GetUUID(), 30*24*60*60, "/",beego.AppConfig.String("cookie.domain"), false, true)
	}
}
