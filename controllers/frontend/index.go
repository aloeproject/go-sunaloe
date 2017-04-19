package frontend

import (
	"myweb/repository"
	"myweb/library"
	"fmt"
)

var PAGE_SIZE = 10

type IndexController struct {
	BaseController
}

func (this *IndexController) Index()  {
	this.init()
	rep := repository.ArticleRepository{Status:10}
	thisPage,_ := this.GetInt("p",1)
	if thisPage <= 1 {
		thisPage = 1
	}
	articleList,_ := rep.List(thisPage -1,PAGE_SIZE)
	count,_ := rep.Count()
	page := library.NewPage(count,thisPage,PAGE_SIZE,articleList)
	this.Data["page"] = page
	this.Data["article_list"] = articleList
	fmt.Println(articleList)
	this.TplName = "frontend/index/index.html"
	this.Render()
}

func (this *IndexController) Detail()  {
	this.init()
	aid := this.Ctx.Input.Param(":id")
	id,_ := library.String2int(aid)
	rep := repository.ArticleRepository{Id:id}
	articleInfo := rep.GetInfoById()
	if articleInfo.Id == 0 {
		this.Abort("404")
	}
	this.Data["article_info"] = articleInfo

	this.TplName = "frontend/index/detail.html"
	this.Render()
}