package repository

import (
	"myweb/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type SpiderArticleRepository struct {
	Id int
	Title string
	Author string
	Content string
	Create_time string
}

func (this *SpiderArticleRepository) GetInfoById() *SpiderArticleRepository {
	model := orm.NewOrm()
	var info SpiderArticleRepository
	sql := fmt.Sprintf("SELECT * FROM spider_article WHERE id = %d",this.Id)
	model.Raw(sql).QueryRow(&info)
	return &info
}


func (this *SpiderArticleRepository) List(currentPage int,pageSize int) (*[]SpiderArticleRepository,error) {
	model := orm.NewOrm()
	var list []SpiderArticleRepository

	//当前页从0 开始
	sql := fmt.Sprintf("SELECT * FROM spider_article ORDER BY create_time DESC LIMIT %d,%d",currentPage * pageSize,pageSize)
	_ , err := model.Raw(sql).QueryRows(&list)
	if err != nil {
		return nil,models.EmptyData
	}
	return &list,nil
}

func (this *SpiderArticleRepository) Count() (int ,error)  {
	models := orm.NewOrm()
	var res  []orm.Params

	sql := fmt.Sprint("SELECT count(1) as ct FROM spider_article")
	_,err := models.Raw(sql).Values(&res)
	if err != nil {
		return 0,err
	}
	ct := fmt.Sprint(res[0]["ct"])
	count ,_ := strconv.Atoi(ct)
	return count,nil
}

