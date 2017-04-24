package repository

import (
	"github.com/astaxie/beego/orm"
	"myweb/models"
	"myweb/vendor"
	"errors"
	"myweb/helper"
)

const UserSessionKey = "user"

type UserRepository struct {

}

func (this *UserRepository) GetInfoById(uid int) (models.UserAdmin,error)  {
	model := orm.NewOrm()
	user := models.UserAdmin{Id:uid}
	err := model.Read(&user)
	if err == nil {
		return user,nil
	}
	return user,err
}




func (this *UserRepository) Login(email,password string) (models.UserAdmin,error){
	var user models.UserAdmin
	o := orm.NewOrm()
	qs := o.QueryTable(user)
	err := qs.Filter("email",email).One(&user)
	if err == orm.ErrNoRows {
		return user,errors.New("不存在此用户")
	}
	state := vendor.CheckPassword(password,user.Password_hash)
	if state == false {
		err = errors.New("密码错误")
		user = models.UserAdmin{}
	}

	return user,err
}

func (this *UserRepository) Register(username,email,password string) (bool,error){
	var user models.UserAdmin
	o := orm.NewOrm()
	qs := o.QueryTable(user)
	err := qs.Filter("email",email).One(&user)
	if err != nil && err != orm.ErrNoRows {
		return false,err
	}
	if user.Id != 0 {
		return false,errors.New("已经存在用户")
	}

	newUser := new(models.UserAdmin)
	newUser.Username = username
	newUser.Email = email
	newUser.Password_hash = vendor.GetPasswordHash(password)
	newUser.Create_time = helper.GetNowDate()
	newUser.Update_time = helper.GetNowDate()
	row,err := o.Insert(newUser)
	if row != 0 {
		return true,nil
	} else {
		return false,err
	}
}
