package routers

import (
	ad "myweb/controllers/admin"
	fd "myweb/controllers/frontend"
	im "myweb/controllers/websocket"
	"github.com/astaxie/beego"
)

func init() {

	//前台
	beego.Router("/", &fd.IndexController{},"get:Index")
	beego.Router("/d/:id([0-9]+)", &fd.IndexController{},"get:Detail")
	beego.Router("/im",&im.IndexController{},"get:WSocket")

	//后台
	//登录和注册
	beego.Router("/admin", &ad.IndexController{},"get:Index")
	beego.Router("/auth/login", &ad.AuthController{},"get:Login")
	beego.Router("/auth/login_post", &ad.AuthController{},"post:Login_post")
	beego.Router("/auth/reg", &ad.AuthController{},"get:Reg")
	beego.Router("/auth/reg_post", &ad.AuthController{},"post:Reg_post")
	//文章
	beego.Router("/admin/article", &ad.ArticleController{},"get:Index")
	beego.Router("/admin/article/add", &ad.ArticleController{},"post,get:Add")
	beego.Router("/admin/article/edit", &ad.ArticleController{},"post,get:Edit")
	beego.Router("/admin/article/del", &ad.ArticleController{},"post:Del")
}
