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
	//原创页面
	beego.Router("/o", &fd.OriginalController{},"get:Index")
	//精彩网文
	beego.Router("/w", &fd.SpiderController{},"get:Index")

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
	beego.Router("/admin/article/down", &ad.ArticleController{},"post:Down")
	beego.Router("/admin/article/up", &ad.ArticleController{},"post:Up")
	beego.Router("/admin/article/del", &ad.ArticleController{},"post:Del")
	beego.Router("/admin/spider", &ad.SpiderController{},"get:Index")
	beego.Router("/admin/spider/w/:web([0-9a-zA-Z_]+|^[\u4E00-\u9FFF]+)", &ad.SpiderController{},"get:Index")
	beego.Router("/admin/spider/w/:web([0-9a-zA-Z_]+|^[\u4E00-\u9FFF]+)/s/:keyword([0-9a-zA-Z_]+|^[\u4E00-\u9FFF]+)", &ad.SpiderController{},"get:Index")
	beego.Router("/admin/spider/s/:keyword([0-9a-zA-Z_]+)", &ad.SpiderController{},"get:Index")
	beego.Router("/admin/spider/:id([0-9]+)", &ad.SpiderController{},"get:Detail")
	beego.Router("/admin/spider/move", &ad.SpiderController{},"post:MoveArticle")
	beego.Router("/admin/spider/del", &ad.SpiderController{},"post:Del")
	beego.Router("/admin/spider/edit", &ad.SpiderController{},"post:Edit")
	beego.Router("/admin/spider/black_list", &ad.SpiderController{},"post:Blacklist")

	//文章上传
	beego.Router("/admin/article/upload_image", &ad.ArticleController{},"get,post:ImageUpload")
	//分类
	beego.Router("/admin/category", &ad.CategoryController{},"get:Index")
	beego.Router("/admin/category/add", &ad.CategoryController{},"post:Add")
}
