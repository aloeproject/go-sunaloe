package frontend

import "github.com/astaxie/beego"



type BaseController struct {
	beego.Controller
}

func (this *BaseController) init()  {
	this.Layout = "frontend/layout/layout.html"
}

