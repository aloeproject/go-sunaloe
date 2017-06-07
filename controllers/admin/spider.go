package admin

import (
	"myweb/repository"
	"myweb/helper"
	"strconv"
	"myweb/constant"
	"fmt"
)

var SPIDER_PAGE_SIZE = 20

type SpiderController struct {
	BaseController
}

func (this *SpiderController)  Index() {
	this.init()
	rep := new(repository.SpiderArticleRepository)
	keyword := this.Ctx.Input.Param(":keyword")
	selectWeb := this.Ctx.Input.Param(":web")
	rep.Keyword = keyword
	rep.Source_web = selectWeb
	thisPage,_ := this.GetInt("p",1)
	articleList,_ := rep.List(thisPage - 1,SPIDER_PAGE_SIZE)
	count,_ := rep.Count()
	keywordGroup := rep.GetKeywordGroup()
	sourceGroup := rep.GetSpiderWebGroup()
	page := helper.NewPage(count,thisPage,SPIDER_PAGE_SIZE,articleList)
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "backend/spider/script.html"
	this.Data["page"] = page
	fmt.Println(count)
	this.Data["article_list"] = articleList
	this.Data["keyword_group"] = keywordGroup
	this.Data["source_group"] = sourceGroup
	this.Data["select_keyword"] = keyword
	this.Data["select_web"] = selectWeb
	this.TplName = "backend/spider/index.html"
	this.Render()
}


func (this *SpiderController)  Detail() {
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

func (this *SpiderController) Edit()  {
	this.init()
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "backend/article/require-umeditor.html"
	this.LayoutSections["HeadHtml"] = "backend/article/headhtml.html"

	this.Data["operation_msg"] = ""

	//获得文章详情
	spiderRep := new(repository.SpiderArticleRepository)
	articleId := this.GetString("aid")
	spiderRep.Id,_ = strconv.Atoi(articleId)
	articleInfo := spiderRep.GetInfoById()

	if this.Ctx.Input.IsPost() {
		title := this.GetString("title")
		author := this.GetString("author")
		content := this.GetString("content")
		sourceName := this.GetString("source_name")
		sourceUrl := this.GetString("source_url")

		re := repository.SpiderArticleRepository{Title:title,Content:content,Author:author,Source_url:sourceUrl,Source_web:sourceName}
		if ok,err:=re.Edit(articleInfo.Id);ok {
			this.Data["operation_msg"] = "修改成功"
			rep := repository.SpiderArticleRepository{Id:articleInfo.Id}
			articleInfo = rep.GetInfoById()
		} else {
			this.Data["operation_msg"] = fmt.Sprintf("修改失败，错误原因:%s",err)
		}
	}
	this.Data["article_info"] = articleInfo
	this.TplName = "backend/spider/detail.html"
	this.Render()
}

func (this *SpiderController) MoveArticle(){
	//获得文章详情
	spiderRep := new(repository.SpiderArticleRepository)
	aid := this.GetString("aid")
	spiderRep.Id,_ = strconv.Atoi(aid)
	articleInfo := spiderRep.GetInfoById()

	//存入正常文章
	re := repository.ArticleRepository{
		Title:articleInfo.Title,
		Author:articleInfo.Author,
		Content:articleInfo.Content,
		Article_source_name:articleInfo.Source_web,
		Article_source_type:constant.ATICLE_SOURCE_SPIDER,
		Article_source_url:articleInfo.Source_url,
		Status:constant.ARTICLE_STATUS_DOWN,
	}

	st,err := re.Add()

	var ret JsonResponse
	if st == true {
		spiderRep.SetState(spiderRep.Id,constant.SPIDER_ARCLIE_MOVED)
		ret.Code = 200
	} else {
		ret.Code = 400
		ret.Msg = fmt.Sprint(err)
	}
	this.Data["json"] = &ret
	this.ServeJSON()
}

/*
  删除爬取文章
 */
func (this *SpiderController) Del(){
	if this.Ctx.Input.IsPost() {
		articleId,_ := this.GetInt("aid")
		re := repository.SpiderArticleRepository{}
		st,err := re.Delete(articleId)
		var ret JsonResponse
		if st == true {
			ret.Code = 200
		} else {
			ret.Code = 400
			ret.Msg = fmt.Sprint(err)
		}
		this.Data["json"] = &ret
		this.ServeJSON()
	}
	this.Abort("404")
}

/*
   爬取黑名单
 */

func (this *SpiderController) Blacklist(){
	if this.Ctx.Input.IsPost() {
		articleId,_ := this.GetInt("aid")
		re := repository.SpiderArticleRepository{}
		st,err := re.Blacklist(articleId)
		var ret JsonResponse
		if st == true {
			ret.Code = 200
		} else {
			ret.Code = 400
			ret.Msg = fmt.Sprint(err)
		}
		this.Data["json"] = &ret
		this.ServeJSON()
	}
	this.Abort("404")
}