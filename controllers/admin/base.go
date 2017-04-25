package admin

import (
	"github.com/astaxie/beego"
	//"fmt"
	"fmt"
	"myweb/repository"
)

var  NoLoginRedirect = "/auth/login"
var  LoginSuccessRedirect = "/admin"

type JsonResponse struct{
	Code int `200`
	Msg string ``
	Body map[string]interface{} ``
}



type BaseController struct {
	beego.Controller
	RouterPattern string
}

func (this *BaseController) init()  {

	this.Layout = "backend/layout/layout.html"
}

func (this *BaseController) IsPost() bool {
	return this.Ctx.Input.IsPost()
}
//检查页面
func (this *BaseController) checkPage(url string){
	state := true
	switch url {
	case "/auth/login","/auth/reg","/auth/login_post","/auth/reg_post":
		state = true
	default:
		if nil != this.GetSession(repository.UserSessionKey) {
			state = true
		}
	}
	if state == false {
		fmt.Println(state)
		this.Redirect(NoLoginRedirect,302)
		return
	}
}



func (this *BaseController) Prepare()  {
	url := this.Data["RouterPattern"]
	this.RouterPattern = fmt.Sprint(url)
	this.checkPage(this.RouterPattern)
}
