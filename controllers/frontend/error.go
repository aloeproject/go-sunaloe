package frontend

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.TplName = "frontend/layout/error-404.html"
	this.Render()
}

