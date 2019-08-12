package models

import (
	"github.com/astaxie/beego/orm"
)

type UserLikeMusicInfo struct {
	Id         int    `orm:"column(id);auto"`
	Uuid       string `orm:"column(uuid)"`
	Type       int    `orm:"column(type)"`
	LikeFields string `orm:"column(like_fields)" description:"用户喜欢的歌曲信息，json存储，由前端定义"`
}

func (t *UserLikeMusicInfo) TableName() string {
	return "user_like_music"
}

func init() {
	orm.RegisterModel(new(UserLikeMusicInfo))
}

func AddUserLikeMusicInfo(m *UserLikeMusicInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetUserLikeMusicInfoByUuId(uuid string) (v *UserLikeMusicInfo, err error) {
	o := orm.NewOrm()
	v = &UserLikeMusicInfo{Uuid: uuid}
	err = o.QueryTable(v).Filter("uuid", uuid).OrderBy("-id").One(v)
	return v, err
}

func GetAllUserLikeMusicInfo(info *[]UserLikeMusicInfo) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable(new(UserLikeMusicInfo)).All(info)
	return
}
