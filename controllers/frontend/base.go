package frontend

import (
	"github.com/astaxie/beego"
	"myweb/helper"
	"fmt"
)



type BaseController struct {
	beego.Controller
}

func (this *BaseController) init()  {
	this.Layout = "frontend/layout/layout.html"
}

func (this *BaseController) Prepare()  {

	//设置uuid cookie
	co,bl := this.GetSecureCookie(beego.AppConfig.String("cookie.secure"),"uuid")
	if bl == false {
		this.SetSecureCookie(beego.AppConfig.String("cookie.secure"),
			"uuid",
			helper.GetUUID(), 30*24*60*60, "/",beego.AppConfig.String("cookie.domain"), false, true)
	}
	fmt.Println(co)
}
