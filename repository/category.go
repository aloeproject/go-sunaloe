package repository

import (
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"myweb/helper"
	"fmt"
	"strconv"
)

type CategoryRepository struct {
	Name string
}

func (this *CategoryRepository) Add() (bool,error) {
	model := orm.NewOrm()
	ar := new(models.Category)
	ar.Name = this.Name
	ar.Create_time = helper.GetNowDate()
	_,err := model.Insert(ar)

	if err == nil {
		return true,nil
	} else {
		return false,err
	}
}

func (this *CategoryRepository) Count() (int ,error)  {
	models := orm.NewOrm()
	var res  []orm.Params
	sql := "SELECT count(1) as ct FROM category"
	_,err := models.Raw(sql).Values(&res)
	if err != nil {
		return 0,err
	}
	ct := fmt.Sprint(res[0]["ct"])
	count ,_ := strconv.Atoi(ct)
	return count,nil
}

func (this *CategoryRepository) List(currentPage,pageSize int) (*[]models.Category,error)  {
	model := orm.NewOrm()
	var list []models.Category
	//当前页从0 开始
	sql := fmt.Sprintf("SELECT * FROM category LIMIT %d,%d",currentPage * pageSize,pageSize)
	_ , err := model.Raw(sql).QueryRows(&list)
	if nil != err {
		return nil,models.EmptyData
	} else {
		return &list,nil
	}
}