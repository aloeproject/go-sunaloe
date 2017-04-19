package admin

import (
	"myweb/repository"
	"fmt"
)



type AuthController struct {
	BaseController
}

func (this *AuthController) Login() {
	this.TplName = "backend/auth/login.html"

	this.Render()
}

func (this *AuthController) Login_post()  {
	password := this.GetString("password")
	email := this.GetString("email")
	rep := repository.UserRepository{}
	user,err := rep.Login(email,password)
	if err!=nil || user.Id == 0 {
		fmt.Println("登录失败，错误信息为:",err)
	} else {
		this.SetSession(repository.UserSessionKey,user)
		//登录成功跳转
		this.Redirect(LoginSuccessRedirect,302)
		return
	}
}

func (this *AuthController) Reg(){
	this.TplName = "backend/auth/reg.html"
	this.Render()
}

func (this *AuthController) Reg_post()  {
	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")

	rep := repository.UserRepository{}
	state,err := rep.Register(username,email,password)
	if state == true {
		fmt.Println("注册成功")
	} else {
		fmt.Println("失败:原因:",err)
	}
}

