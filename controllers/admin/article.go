package admin

import (
	"myweb/repository"
	"fmt"
	"myweb/library"
	"myweb/models"
)

var PAGE_SIZE = 10

type ArticleController struct {
	BaseController
}

func (this *BaseController) Index()  {
	this.init()
	rep := repository.ArticleRepository{}
	//获得当前页码
	thisPage,_ := this.GetInt("p",1)
	articleList,_ := rep.List(thisPage -1,PAGE_SIZE)
	count,_ := rep.Count()
	page := library.NewPage(count,thisPage,PAGE_SIZE,articleList)
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "backend/article/script.html"
	this.Data["article_list"] = articleList
	this.Data["page"] = page
	this.TplName = "backend/article/list.html"
	this.Render()
}

func (this *BaseController) Add()  {
	this.init()
	this.TplName = "backend/article/add.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "backend/article/require-umeditor.html"
	this.LayoutSections["HeadHtml"] = "backend/article/headhtml.html"

	this.Data["operation_msg"] = ""
	//页面标识添加
	this.Data["is_add"] = true

	title := this.GetString("title")
	content := this.GetString("content")
	if title != "" && content != "" {
		re := repository.ArticleRepository{Category_id:0, Title:title, Content:content,Title_img:""}
		if ok,err:=re.Add();ok {
			this.Data["operation_msg"] = "添加成功"
		} else {
			this.Data["operation_msg"] = fmt.Sprintf("添加失败，错误原因:%s",err)
		}
	}
	this.Render()
}

func (this *BaseController) Edit(){
	this.init()
	this.TplName = "backend/article/add.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "backend/article/require-umeditor.html"
	this.LayoutSections["HeadHtml"] = "backend/article/headhtml.html"

	this.Data["operation_msg"] = ""
	//页面标识修改
	this.Data["is_add"] = false

	var articleInfo models.Article
	articleId,_ := this.GetInt("aid")
	rep := repository.ArticleRepository{Id:articleId}
	articleInfo = rep.GetInfoById()

	if this.IsPost() {
		title := this.GetString("title")
		content := this.GetString("content")
		re := repository.ArticleRepository{Category_id:0, Title:title, Content:content,Title_img:""}
		if ok,err:=re.Edit(articleId);ok {
			this.Data["operation_msg"] = "修改成功"
			rep := repository.ArticleRepository{Id:articleId}
			articleInfo = rep.GetInfoById()
		} else {
			this.Data["operation_msg"] = fmt.Sprintf("修改失败，错误原因:%s",err)
		}
	}
	this.Data["article_info"] = articleInfo
	this.Render()
}

func (this *BaseController) Del(){
	if this.IsPost() {
		articleId,_ := this.GetInt("aid")
		re := repository.ArticleRepository{}
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
