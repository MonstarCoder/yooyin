package services

import (
	"github.com/astaxie/beego/orm"
	"yooyin/models"
)

type userService struct {}

var UserService userService

func (this *userService) orm() orm.Ormer {
	var o orm.Ormer
	if o == nil {
		o = orm.NewOrm()
	}
	return o
}

func (this *userService) LoginUser(user *models.User) error {
	_, _, err := this.orm().ReadOrCreate(user, "OpenId")
	return err
}
