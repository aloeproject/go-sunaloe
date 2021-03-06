package repository

import (
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"fmt"
	"strconv"
	"myweb/helper"
	"errors"
	"strings"
	"time"
	"myweb/constant"
)

type ArticleRepository struct {
	Id int
	Category_id int
	Category_name string
	Title string
	Content string
	Title_img string
	Status int
	Create_time string
	Update_time string
	Author string
	Article_source_url string
	Article_source_name string
	Article_source_type int
	St_Update_time time.Time
	Ed_Update_time time.Time
}

type ArticleList struct {
	Id int
	Category_id int
	Category_name string
	Title string
	Author string
	Content string
	Title_img string
	Article_source_url string
	Article_source_name string
	Article_source_type int
	Status int
	Create_time string
	Update_time string
}

func getTitleImg(str string) string {
	if strings.Index(str,".") == 0 {
		return str[1:]
	}
	return str
}

func (this *ArticleRepository) GetInfoById() *ArticleList {
	model := orm.NewOrm()
	var info ArticleList
	sql := fmt.Sprintf("SELECT a.Id as id,category_id,ifnull(c.name,'无') as category_name,title,author," +
		"content,article_source_type,article_source_url,article_source_name," +
		"title_img,status,a.create_time as create_time,a.update_time as update_time" +
		" FROM article a LEFT JOIN category c ON a.Category_id = c.id WHERE a.Id = %d",this.Id)
	model.Raw(sql).QueryRow(&info)
	return &info
}



func (this *ArticleRepository) List(currentPage int,pageSize int) (*[]ArticleList,error) {
	model := orm.NewOrm()
	var list []ArticleList
	var stWhere string
	var where string
	if this.Status == 0  {
		stWhere = fmt.Sprint("1,10")
	} else {
		stWhere = fmt.Sprint("10")
	}

	if this.St_Update_time.Unix() == -62135596800  {
		where = ""
	} else {
		where = fmt.Sprintf("AND update_time >= '%s' AND update_time <= '%s' ",fmt.Sprint(this.St_Update_time),fmt.Sprint(this.Ed_Update_time))
	}

	if this.Category_name != "" {
		where += fmt.Sprintf(" AND c.name = '%s' ",this.Category_name)
	}

	if this.Article_source_type != 0 {
		where += fmt.Sprintf(" AND a.article_source_type = '%d' ",this.Article_source_type)
	}

	//当前页从0 开始
	sql := fmt.Sprintf("SELECT a.Id as id,category_id,ifnull(c.name,'无') as category_name,title,author," +
		"content,title_img,status,a.create_time as create_time,a.update_time as update_time" +
		" FROM article a LEFT JOIN category c ON a.Category_id = c.id WHERE status IN (%s) %s ORDER BY create_time desc LIMIT %d,%d",stWhere,where,currentPage * pageSize,pageSize)
	_ , err := model.Raw(sql).QueryRows(&list)
	if err != nil {
		return nil,models.EmptyData
	}
	return &list,nil
}

func (this *ArticleRepository) Count() (int ,error)  {
	models := orm.NewOrm()
	var res  []orm.Params
	where := ""

	if this.St_Update_time.Unix() == -62135596800  {
		where = ""
	} else {
		where = fmt.Sprintf("AND update_time >= '%s' AND update_time <= '%s' ",fmt.Sprint(this.St_Update_time),fmt.Sprint(this.Ed_Update_time))
	}

	if this.Category_name != "" {
		where += fmt.Sprintf(" AND c.name = '%s' ",this.Category_name)
	}

	if this.Article_source_type != 0 {
		where += fmt.Sprintf(" AND a.article_source_type = '%d' ",this.Article_source_type)
	}

	sql := fmt.Sprintf("SELECT count(1) as ct FROM article a LEFT JOIN category c ON a.Category_id = c.id WHERE status IN (1,10) %s",where)
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
	ar.Article_source_url = this.Article_source_url
	if this.Article_source_type == 0 {
		ar.Article_source_type = constant.ARTICLE_SOURCE_NORMAL
	} else {
		ar.Article_source_type = this.Article_source_type
	}
	ar.Author = this.Author
	ar.Content = this.Content
	ar.Title_img = getTitleImg(this.Title_img)
	ar.Article_source_name = this.Article_source_name
	if this.Status == 0 {
		ar.Status = constant.ARTICLE_STATUS_NORMAL
	} else {
		ar.Status = this.Status
	}
	ar.Create_time = helper.GetNowDate()
	ar.Update_time = helper.GetNowDate()
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
			ar.Title_img = getTitleImg(this.Title_img)
		}
		ar.Update_time = helper.GetNowDate()
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
//最新文章
func (this *ArticleRepository) NewestArticle() (*[]models.Article,error)  {
	model := orm.NewOrm()

	var where string

	if this.Article_source_type != 0 {
		where += fmt.Sprintf(" AND article_source_type = '%d' ",this.Article_source_type)
	}

	//当前页从0 开始
	sql := fmt.Sprintf("SELECT * FROM article WHERE status = 10 %s ORDER BY update_time desc LIMIT 5",where)

	var list  []models.Article
	_ , err := model.Raw(sql).QueryRows(&list)
	if err != nil {
		return nil,models.EmptyData
	}
	return &list,nil
}

func (this *ArticleRepository) GetDateCategory() (*map[string]string,error)  {
	model := orm.NewOrm()
	//当前页从0 开始
	sql := fmt.Sprint("select left(update_time,7) as date_category from article group by left(update_time,7)")
	var res  []orm.Params
	_ , err := model.Raw(sql).Values(&res)
	if err != nil {
		return nil,models.EmptyData
	}
	ret := make(map[string]string)
	for _,v := range res {
		date := fmt.Sprint(v["date_category"])
		dataArr := strings.Split(date,"-")
		ret[date] = dataArr[0]+"年"+dataArr[1]+"月"
	}
	//ret := make(map[string]string)
	return &ret,nil
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

func (this *ArticleRepository) SetState(id int,state int) (bool,error) {
	model := orm.NewOrm()
	ar := models.Article{Id:id}
	if model.Read(&ar) == nil {
		//文章下线
		ar.Status = state
		ar.Update_time = helper.GetNowDate()
		num,err := model.Update(&ar)
		if err != nil {
			return false,err
		}
		if num != 0 {
			return true,nil
		}
	}
	return false,errors.New("不存在此文章")
}


func NextPrevArticle(aid int) (*[2]ArticleList) {
	model := orm.NewOrm()
	var ret [2]ArticleList
	//当前页从0 开始
	nextSql := fmt.Sprintf("SELECT id,title FROM article WHERE status = 10 AND id < %d ORDER BY id DESC LIMIT 1",aid)
	prevSql := fmt.Sprintf("SELECT id,title FROM article WHERE status = 10 AND id > %d ORDER BY id ASC LIMIT 1",aid)

	var nextList,prevList  ArticleList
	model.Raw(nextSql).QueryRow(&nextList)
	model.Raw(prevSql).QueryRow(&prevList)
	ret[0] = nextList
	ret[1] = prevList
	return &ret
}