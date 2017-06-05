package frontend

import (
	"github.com/astaxie/beego"
	"myweb/helper"
	"myweb/repository"
)


type BaseController struct {
	beego.Controller
	UUID string
}

func (this *BaseController) init()  {
	controllerName,actionName := this.GetControllerAndAction()
	this.Data["controllerName"] = controllerName
	this.Data["actionName"] = actionName
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
	uuid,bl := this.GetSecureCookie(beego.AppConfig.String("cookie.secure"),"uuid")
	if bl == false {
		uuid = helper.GetUUID()
		this.SetSecureCookie(beego.AppConfig.String("cookie.secure"),
			"uuid",
			uuid, 30*24*60*60, "/",beego.AppConfig.String("cookie.domain"), false, true)
	}
	this.UUID = uuid

}
