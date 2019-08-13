package models

import "github.com/astaxie/beego/orm"


func GetTargetUidById(TargetiIdList *[]LikeRelationInformation, userId string) error{
	o := orm.NewOrm()
	_, err := o.QueryTable(new(LikeRelationInformation).TableName()).Filter("UserId", userId).All(TargetiIdList)
	return err
}