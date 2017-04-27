package admin

import (
	"fmt"
	"myweb/repository"
	"myweb/helper"
)

type CategoryController struct {
	BaseController
}
//分类首页
func (this *CategoryController) Index() {
	this.init()
	this.LayoutSections = make(map[string]string)
	rep := repository.CategoryRepository{}
	//获得当前页码
	thisPage,_ := this.GetInt("p",1)
	categoryList,_ := rep.List(thisPage -1,PAGE_SIZE)
	count,_ := rep.Count()
	page := helper.NewPage(count,thisPage,PAGE_SIZE,categoryList)
	this.Data["page"] = page
	this.Data["category_list"] = categoryList
	this.LayoutSections["Scripts"] = "backend/category/script.html"
	this.TplName = "backend/category/index.html"
	this.Render()
}

func (this *CategoryController) Add()  {
	name := this.GetString("name")
	var ret JsonResponse
	if name == "" {
		ret.Code = 400
		ret.Msg = fmt.Sprint("类别名称不能为空")
		this.Data["json"] = &ret
		this.ServeJSON()
		return
	}

	rep := repository.CategoryRepository{Name:name}
	_,err := rep.Add()
	if err == nil {
		ret.Code = 200
		ret.Msg = ""
	} else {
		ret.Code = 400
		ret.Msg = fmt.Sprint(err)
	}
	this.Data["json"] = &ret
	this.ServeJSON()
}
