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
	Keyword string
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

	where := ""

	if this.Keyword != "" {
		where += fmt.Sprintf(" AND keyword = '%s' ",this.Keyword)
	}

	if this.Source_web != "" {
		where += fmt.Sprintf(" AND source_web = '%s' ",this.Source_web)
	}

	//当前页从0 开始
	sql := fmt.Sprintf("SELECT * FROM spider_article WHERE status in (%d,%d) %s ORDER BY create_time DESC,id DESC LIMIT %d,%d",
		constant.SPIDER_ARTCLIE_NORMAL,constant.SPIDER_ARCLIE_MOVED,
		where,
		currentPage * pageSize,pageSize)
	_ , err := model.Raw(sql).QueryRows(&list)
	if err != nil {
		return nil,models.EmptyData
	}
	return &list,nil
}

func (this *SpiderArticleRepository) Count() (int ,error)  {
	models := orm.NewOrm()
	var res  []orm.Params

	where := ""

	if this.Keyword != "" {
		where += fmt.Sprintf(" AND keyword = '%s' ",this.Keyword)
	}

	if this.Source_web != "" {
		where += fmt.Sprintf(" AND source_web = '%s' ",this.Source_web)
	}

	sql := fmt.Sprintf("SELECT count(1) as ct,keyword FROM spider_article WHERE status in (%d,%d) %s",
		constant.SPIDER_ARTCLIE_NORMAL,constant.SPIDER_ARCLIE_MOVED,where)
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

/*
    keyword = 数量
 */
func (this *SpiderArticleRepository) GetKeywordGroup() map[string]int {
	models := orm.NewOrm()
	var res  []orm.Params
	ret := make(map[string]int)

	where := ""


	if this.Source_web != "" {
		where += fmt.Sprintf(" AND source_web = '%s' ",this.Source_web)
	}

	sql := fmt.Sprintf("SELECT count(1) as ct,keyword FROM spider_article WHERE status in (%d,%d) %s GROUP BY keyword",
		constant.SPIDER_ARTCLIE_NORMAL,constant.SPIDER_ARCLIE_MOVED,where)
	models.Raw(sql).Values(&res)
	for _,item := range res {
		ct,_ := strconv.Atoi(item["ct"].(string))
		kword := item["keyword"].(string)
		ret[kword] = ct
	}
	return ret
}

func (this *SpiderArticleRepository) GetSpiderWebGroup() map[string]int {
	models := orm.NewOrm()
	var res  []orm.Params
	ret := make(map[string]int)
	sql := fmt.Sprintf("SELECT count(1) as ct,source_web FROM spider_article WHERE status in (%d,%d) GROUP BY source_web",
		constant.SPIDER_ARTCLIE_NORMAL,constant.SPIDER_ARCLIE_MOVED)
	models.Raw(sql).Values(&res)
	for _,item := range res {
		ct,_ := strconv.Atoi(item["ct"].(string))
		source_web := item["source_web"].(string)
		ret[source_web] = ct
	}
	return ret
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
