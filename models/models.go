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
	Status int
	Create_time string
	Update_time string
}

type Category struct {
	Id int
	Name string
	Create_time string
}

func (a *Category) TableName() string {
	return "category"
}

var EmptyData = errors.New("数据为空")

func init(){
	orm.RegisterModel(new(UserAdmin),new(Article),new(Category))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	dbuser := beego.AppConfig.String("dbuser")
	dbpasswd := beego.AppConfig.String("dbpasswd")
	dbname := beego.AppConfig.String("dbname")
	dataSource := fmt.Sprintf("%s:%s@/%s?charset=utf8",dbuser,dbpasswd,dbname)
	orm.RegisterDataBase("default","mysql",dataSource,30)
}