package admin

import (
	"myweb/repository"
	"myweb/helper"
	"strconv"
)

type SogouSpiderController struct {
	BaseController
}

func (this SogouSpiderController)  Index() {
	this.init()

	rep := new(repository.SpiderArticleRepository)
	thisPage,_ := this.GetInt("p",1)
	articleList,_ := rep.List(thisPage - 1,PAGE_SIZE)
	count,_ := rep.Count()
	page := helper.NewPage(count,thisPage,PAGE_SIZE,articleList)
	this.Data["page"] = page
	this.Data["article_list"] = articleList
	this.TplName = "backend/spider/index.html"
	this.Render()
}


func (this SogouSpiderController)  Detail() {
	this.init()
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "backend/article/require-umeditor.html"
	this.LayoutSections["HeadHtml"] = "backend/article/headhtml.html"

	rep := new(repository.SpiderArticleRepository)
	aid := this.Ctx.Input.Param(":id")
	rep.Id,_ = strconv.Atoi(aid)
	articleInfo := rep.GetInfoById()
	this.Data["article_info"] = articleInfo
	this.TplName = "backend/spider/detail.html"
	this.Render()
}