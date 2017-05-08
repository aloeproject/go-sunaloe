package repository

import (
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"fmt"
	"strconv"
	"myweb/helper"
)

type ArticleClickRepository struct {
	Aid int
}

func (this *ArticleClickRepository) Save(aid int,ip string,gid string) (bool,error) {
	model := orm.NewOrm()
	obj := new(models.ArticleClick)
	obj.Aid = aid
	obj.Ip = ip
	obj.Gid = gid
	obj.Create_time = helper.GetNowDate()
	_,err := model.Insert(obj)
	if err == nil {
		return true,nil
	} else {
		return false,err
	}
}

func (this *ArticleClickRepository) GetClickCount() (int,error) {
	models := orm.NewOrm()
	var res  []orm.Params
	sql := fmt.Sprintf("SELECT count(1) as ct FROM article_click WHERE aid = %d",this.Aid)
	_,err := models.Raw(sql).Values(&res)
	if err != nil {
		return 0,err
	}
	ct := fmt.Sprint(res[0]["ct"])
	count ,_ := strconv.Atoi(ct)
	return count,nil
}
