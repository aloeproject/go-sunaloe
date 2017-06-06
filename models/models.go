package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
)

type UserAdmin struct {
	Id int
	Username string
	Email string
	Password_hash string
	Description string
	Create_time string
	Update_time string
}

func (a *UserAdmin) TableName() string {
	return "user_admin"
}

type Article struct {
	Id int
	User_id int
	Category_id int
	Title string
	Content string
	Title_img string
	Author string
	Article_source_url string
	Article_source_name string
	Article_source_type int
	Status int
	Create_time string
	Update_time string
}

type Category struct {
	Id int
	Name string
	Create_time string
}

type ArticleClick struct {
	Id int
	Aid int
	Ip string
	Gid string
	Create_time string
}

type SpiderArticle struct {
	Id int
	Article_hash string
	Title string
	Keyword string
	Author string
	Content string
	Source_web string
	Source_url string
	Status int
	Create_time string
}

func (a *ArticleClick) TableName() string {
	return "article_click"
}

func (a *Category) TableName() string {
	return "category"
}

func (a *SpiderArticle) TableName() string {
	return "spider_article"
}

var EmptyData = errors.New("数据为空")

func init(){
	orm.RegisterModel(new(UserAdmin),new(Article),new(Category),new(ArticleClick),new(SpiderArticle))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	dbuser := beego.AppConfig.String("dbuser")
	dbpasswd := beego.AppConfig.String("dbpasswd")
	dbname := beego.AppConfig.String("dbname")
	dataSource := fmt.Sprintf("%s:%s@/%s?charset=utf8",dbuser,dbpasswd,dbname)
	orm.RegisterDataBase("default","mysql",dataSource,30)
}