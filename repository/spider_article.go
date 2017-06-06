package repository

import (
	"myweb/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"errors"
	"myweb/constant"
)

type SpiderArticleRepository struct {
	Id int
	Title string
	Author string
	Source_web string
	Source_url string
	Content string
	Status int
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
	sql := fmt.Sprintf("SELECT * FROM spider_article WHERE status in (0,10) ORDER BY create_time DESC,id DESC LIMIT %d,%d",currentPage * pageSize,pageSize)
	_ , err := model.Raw(sql).QueryRows(&list)
	if err != nil {
		return nil,models.EmptyData
	}
	return &list,nil
}

func (this *SpiderArticleRepository) Count() (int ,error)  {
	models := orm.NewOrm()
	var res  []orm.Params

	sql := fmt.Sprint("SELECT count(1) as ct FROM spider_article WHERE status in (0,10)")
	_,err := models.Raw(sql).Values(&res)
	if err != nil {
		return 0,err
	}
	ct := fmt.Sprint(res[0]["ct"])
	count ,_ := strconv.Atoi(ct)
	return count,nil
}

func (this *SpiderArticleRepository) Edit(id int) (bool,error) {
	model := orm.NewOrm()
	ar := models.SpiderArticle{Id:id}
	if model.Read(&ar) == nil {
		if this.Title != "" {
			ar.Title = this.Title
		}
		if this.Author != "" {
			ar.Author = this.Author
		}
		if this.Content != "" {
			ar.Content = this.Content
		}
		if this.Source_web != "" {
			ar.Source_web = this.Source_web
		}
		if this.Source_url != "" {
			ar.Source_url = this.Source_url
		}
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

func (this *SpiderArticleRepository) Delete(id int) (bool,error) {
	model := orm.NewOrm()
	ar := models.SpiderArticle{Id:id}
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

/**
   放入黑名单
 */
func (this *SpiderArticleRepository) Blacklist(id int) (bool,error) {
	model := orm.NewOrm()
	ar := models.SpiderArticle{Id:id}
	if model.Read(&ar) == nil {
		ar.Status = constant.SPIDER_ARTICLE_BLACKLIST
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

func (this *SpiderArticleRepository) SetState(id int,state int) (bool,error) {
	model := orm.NewOrm()
	ar := models.SpiderArticle{Id:id}
	if model.Read(&ar) == nil {
		//文章下线
		ar.Status = state
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
