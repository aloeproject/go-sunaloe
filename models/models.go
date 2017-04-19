package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"errors"
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

var EmptyData = errors.New("数据为空")

func init(){
	orm.RegisterModel(new(UserAdmin),new(Article))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:123456@/sunaloe_blog?charset=utf8",30)
}