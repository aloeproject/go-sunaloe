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
	beego.Router("/about", &fd.AboutController{},"get:Index")
	beego.Router("/d/:id([0-9]+)", &fd.IndexController{},"get:Detail")
	beego.Router("/d/:id([0-9]+).html", &fd.IndexController{},"get:Detail")
	beego.Router("/date/:date([0-9]{4}-[0-9]{2})", &fd.IndexController{},"get:Index")
	beego.Router("/c/:category_name([\u4e00-\u9fa5_a-zA-Z0-9]+)", &fd.IndexController{},"get:Index")
	beego.Router("/im",&im.IndexController{},"get:WSocket")

	//后台
	//登录和注册
	beego.Router("/admin", &ad.IndexController{},"get:Index")
	beego.Router("/auth/login", &ad.AuthController{},"get,post:Login")
	//beego.Router("/auth/reg", &ad.AuthController{},"get:Reg")
	//beego.Router("/auth/reg_post", &ad.AuthController{},"post:Reg_post")
	//文章
	beego.Router("/admin/article", &ad.ArticleController{},"get:Index")
	beego.Router("/admin/article/add", &ad.ArticleController{},"post,get:Add")
	beego.Router("/admin/article/edit", &ad.ArticleController{},"post,get:Edit")
	beego.Router("/admin/article/del", &ad.ArticleController{},"post:Del")
	//分类
	beego.Router("/admin/category", &ad.CategoryController{},"get:Index")
	beego.Router("/admin/category/add", &ad.CategoryController{},"post:Add")
}
