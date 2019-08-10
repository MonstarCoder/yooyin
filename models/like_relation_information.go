package models

import "github.com/astaxie/beego/orm"

type LikeRelationInformation struct {
	Id           int
	UserId       string
	TargetUserId string
	IsLiked      bool
}

func (t *LikeRelationInformation) TableName() string {
 	return "like_relation_information"
}

func init() {
 	orm.RegisterModel(new(LikeRelationInformation))
}

func AddLikeRelationInformation(m *LikeRelationInformation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetLikeRelastionInformationByUserId(likeRelationList *[]LikeRelationInformation, userId string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(MusicInformation).TableName()).Filter("user_id", userId).OrderBy("-id").All(likeRelationList)
	return err
}
