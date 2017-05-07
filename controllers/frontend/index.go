package frontend

import (
	"myweb/repository"
	"myweb/helper"
	"strings"
	"time"
	"fmt"
	"strconv"
)

var PAGE_SIZE = 10

type IndexController struct {
	BaseController
}


func (this *IndexController) Index()  {
	this.init()
	this.PageKeyword(new(repository.ArticleList))
	rep := repository.ArticleRepository{Status:10}
	date := this.Ctx.Input.Param(":date")
	if date != "" {
		dateArr := strings.Split(date,"-")
		month,_ := strconv.Atoi(dateArr[1])
		year,_ := strconv.Atoi(dateArr[0])
		st := fmt.Sprint(dateArr[0],"-",fmt.Sprintf("%02d",month),"-01")
		var ed string
		if month == 12 {
			ed = fmt.Sprint(fmt.Sprintf("%d",year + 1),"-","01","-01")
		} else {
			ed = fmt.Sprint(fmt.Sprintf("%d",year),"-",fmt.Sprintf("%02d",month + 1),"-01")
		}
		rep.St_Update_time ,_ = time.Parse("2006-01-02",st)
		rep.Ed_Update_time ,_ = time.Parse("2006-01-02",ed)
	}
	category_name := this.Ctx.Input.Param(":category_name")
	if category_name != "" {
		rep.Category_name = category_name
	}
	thisPage,_ := this.GetInt("p",1)
	if thisPage <= 1 {
		thisPage = 1
	}
	//首页文章
	articleList,_ := rep.List(thisPage -1,PAGE_SIZE)
	count,_ := rep.Count()
	page := helper.NewPage(count,thisPage,PAGE_SIZE,articleList)
	//最新文章
	newArticle,_ := rep.NewestArticle()

	dateCategory,_ := rep.GetDateCategory()
	cateRep := repository.CategoryRepository{}
	articleCate,_ := cateRep.List(0,0)

	this.Data["page"] = page
	this.Data["article_list"] = articleList
	this.Data["newest_list"] = newArticle
	this.Data["date_category"] = dateCategory
	this.Data["category_list"] = articleCate
	this.TplName = "frontend/index/index.html"
	this.Render()
}

func (this *IndexController) Detail()  {
	this.init()
	aid := this.Ctx.Input.Param(":id")
	id,_ := helper.String2int(aid)
	rep := repository.ArticleRepository{Id:id}
	articleInfo := rep.GetInfoById()
	this.PageKeyword(articleInfo)
	if articleInfo.Id == 0 {
		this.Abort("404")
	}
	this.Data["article_info"] = articleInfo
	//最新文章
	newArticle,_ := rep.NewestArticle()

	dateCategory,_ := rep.GetDateCategory()
	cateRep := repository.CategoryRepository{}
	articleCate,_ := cateRep.List(0,0)

	this.Data["newest_list"] = newArticle
	this.Data["date_category"] = dateCategory
	this.Data["category_list"] = articleCate

	this.TplName = "frontend/index/detail.html"
	this.Render()
}