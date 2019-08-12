package models

import "github.com/astaxie/beego/orm"

type MatchInformation struct {
	Id          	int `orm:"column(id);auto"`
	UuId 			string
	Type            int
	LikeFields 		string
}

func (t *MatchInformation) TableName() string {
	return "match_information"
}

func init() {
	orm.RegisterModel(new(MatchInformation))
}

func GetMatchInformationById(coLikeUserList *[]MatchInformation, matchType int)error {
		o := orm.NewOrm()
		_, err := o.QueryTable(new(MusicInformation).TableName()).Filter("Type", matchType).OrderBy("-id").All(coLikeUserList)
		return err
}

func GetTargetUidById(TargetiIdList *[]LikeRelationInformation, userId string) error{
	o := orm.NewOrm()
	_, err := o.QueryTable(new(LikeRelationInformation).TableName()).Filter("UserId", userId).All(TargetiIdList)
	return err
}