package admin

import (
	"myweb/repository"
	"fmt"
	"myweb/models"
	"myweb/helper"
//	"io/ioutil"
	"io"
	"os"
	"time"
	"log"
	"net/http"
)

var PAGE_SIZE = 10

func check(err error) {
	//检查并定义一个panic异常
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

type ArticleController struct {
	BaseController
}

func uploadImg(req *http.Request) (string) {
	f,h,err := req.FormFile("first_image")
	check(err)
	filename := h.Filename
	//时间戳做随机
	unix := fmt.Sprint(time.Now().Unix())
	filename = unix+"_"+filename
	dir,err := helper.GetUploadImageDir()
	file := dir+"/"+filename
	t,err :=os.Create(file)
	check(err)
	_,err = io.Copy(t,f)
	check(err)

	defer f.Close()
	defer t.Close()
	return file
}


func (this *ArticleController) Index()  {
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

func (this *ArticleController) Add()  {
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
	if this.Ctx.Input.IsPost() == true {
		//图片上传
		f,_,_ := this.Ctx.Request.FormFile("first_image")
		file := ""
		if f != nil {
			file = uploadImg(this.Ctx.Request)
		}

		if title != "" && content != "" {
			re := repository.ArticleRepository{Category_id:category, Title:title, Content:content,Title_img:file}
			if ok,err:=re.Add();ok {
				this.Data["operation_msg"] = "添加成功"
			} else {
				this.Data["operation_msg"] = fmt.Sprintf("添加失败，错误原因:%s",err)
			}
		} else {
			this.Data["operation_msg"] = "标题或者内容不能为空"
		}

	}
	this.Render()
}

func (this *ArticleController) Edit(){
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
	if this.Ctx.Input.IsPost() {
		title := this.GetString("title")
		content := this.GetString("content")
		categoryId,_ := this.GetInt("category")
		//图片上传
		f,_,_ := this.Ctx.Request.FormFile("first_image")
		file := ""
		if f != nil {
			file = uploadImg(this.Ctx.Request)
		}

		re := repository.ArticleRepository{Category_id:categoryId, Title:title, Content:content,Title_img:file}
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
	if this.Ctx.Input.IsPost() {
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
