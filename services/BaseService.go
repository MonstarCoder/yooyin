package services

import (
	"github.com/astaxie/beego/orm"
)

type BaseService struct {}

func (this *BaseService) orm() orm.Ormer {
	var o orm.Ormer
	if o == nil {
		o = orm.NewOrm()
	}
	return o
}
