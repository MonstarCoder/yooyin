package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserLikeType struct {
	Id 						int       `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (u *UserLikeType) TableName() string {
	return "user_like_types"
}

func init() {
	orm.RegisterModel(new(UserLikeType))
}
