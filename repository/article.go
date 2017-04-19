package repository

import (
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"fmt"
	"strconv"
	"myweb/library"
	"errors"
)

type ArticleRepository struct {
	Id int
	Category_id int
	Title string
	Content string
	Title_img string
	Status int
}

func (this *ArticleRepository) GetInfoById() models.Article {
	model := orm.NewOrm()
	ar := models.Article{Id:this.Id}
	model.Read(&ar)
	return ar
}

func (this *ArticleRepository) List(currentPage int,pageSize int) ([]models.Article,error) {
	model := orm.NewOrm()
	var list []models.Article
	var stWhere string
	if this.Status == 0  {
		stWhere = fmt.Sprint("1,10")
	} else {
		stWhere = fmt.Sprint("10")
	}
	//当前页从0 开始
	sql := fmt.Sprintf("SELECT * FROM article WHERE status IN (%s) LIMIT %d,%d",stWhere,currentPage * pageSize,pageSize)
	_ , err := model.Raw(sql).QueryRows(&list)
	if err != nil {
		return nil,models.EmptyData
	}
	return list,nil
}

func (this *ArticleRepository) Count() (int ,error)  {
	models := orm.NewOrm()
	var res  []orm.Params
	sql := "SELECT count(1) as ct FROM article WHERE status IN (1,10)"
	_,err := models.Raw(sql).Values(&res)
	if err != nil {
		return 0,err
	}
	ct := fmt.Sprint(res[0]["ct"])
	count ,_ := strconv.Atoi(ct)
	return count,nil
}

func (this *ArticleRepository) Add() (bool,error) {
	model := orm.NewOrm()
	ar := new(models.Article)
	ar.User_id = 1
	ar.Category_id = this.Category_id
	ar.Title = this.Title
	ar.Content = this.Content
	ar.Title_img = this.Title_img
	ar.Status = 10
	ar.Create_time = library.GetNowDate()
	ar.Update_time = library.GetNowDate()
	_,err := model.Insert(ar)
	if err == nil {
		return true,nil
	} else {
		return false,err
	}
}

func (this *ArticleRepository) Edit(id int) (bool,error) {
	model := orm.NewOrm()
	ar := models.Article{Id:id}
	if model.Read(&ar) == nil {
		if this.Category_id != 0 {
			ar.Category_id = this.Category_id
		}
		if this.Title != "" {
			ar.Title = this.Title
		}
		if this.Content != "" {
			ar.Content = this.Content
		}
		if this.Title_img != ""{
			ar.Title_img = this.Title_img
		}
		ar.Update_time = library.GetNowDate()
		num,err := model.Update(&ar)
		if err != nil {
			return false,err
		}
		if num != 0 {
			return true,nil
		}
	}
	return false,errors.New("没有做修改")
}

func (this *ArticleRepository) Delete(id int) (bool,error) {
	model := orm.NewOrm()
	ar := models.Article{Id:id}
	if model.Read(&ar) == nil {
		num,err := model.Delete(&ar)
		if err != nil {
			return false,err
		}
		if num != 0 {
			return true,nil
		}
	}
	return false,errors.New("不存在此文章")
}
