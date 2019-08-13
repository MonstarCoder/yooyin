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
	_, _ = o.Delete(m, "uuid", "type")
	id, err = o.Insert(m)
	return
}

func GetUserLikeMusicInfoByUuId(uuid string, infos *[]UserLikeMusicInfo) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(UserLikeMusicInfo)).Filter("uuid", uuid).OrderBy("-id").All(infos)
	return err
}

func GetAllUserLikeMusicInfo(info *[]UserLikeMusicInfo) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable(new(UserLikeMusicInfo)).All(info)
	return
}
