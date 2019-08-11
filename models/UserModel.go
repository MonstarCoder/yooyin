package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id 						int       `json:"id"`
	OpenId 				string    `json:"openId"`
	UnionId 		  string    `json:"unionId"`
	NickName      string    `json:"nickName"`
	AvatarUrl     string    `json:"avatarUrl"`
	Gender        int       `json:"gender"`
	Country       string    `json:"country"`
	Province      string    `json:"province"`
	City          string    `json:"city"`
	Language      string    `json:"language"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(User))
}
