package models

import "github.com/astaxie/beego/orm"

type ContactInformation struct {
	Id            int
	UserId        string
	ContactUserId string
}

func (t *ContactInformation) TableName() string {
	return "contact_information"
}

func init() {
	orm.RegisterModel(new(ContactInformation))
}

func AddContactInformation(m *ContactInformation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
