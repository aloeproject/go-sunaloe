package admin

import (
	"myweb/repository"
	"fmt"
	"myweb/models"
	"myweb/helper"
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
	page := helper.NewPage(count,thisPage,PAGE_SIZE,articleList)
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
	categRep := repository.CategoryRepository{}
	cateList,_ := categRep.List(0,0)
	this.Data["category_list"] = cateList

	title := this.GetString("title")
	content := this.GetString("content")
	category,_ := this.GetInt("category")
	if title != "" && content != "" {
		re := repository.ArticleRepository{Category_id:category, Title:title, Content:content,Title_img:""}
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
		categoryId,_ := this.GetInt("category")
		re := repository.ArticleRepository{Category_id:categoryId, Title:title, Content:content,Title_img:""}
		if ok,err:=re.Edit(articleId);ok {
			this.Data["operation_msg"] = "修改成功"
			rep := repository.ArticleRepository{Id:articleId}
			articleInfo = rep.GetInfoById()
		} else {
			this.Data["operation_msg"] = fmt.Sprintf("修改失败，错误原因:%s",err)
		}
	}
	cate := new(repository.CategoryRepository)
	this.Data["article_info"] = articleInfo
	categoryList,_ := cate.List(0,0)
	this.Data["category_list"] = categoryList
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
