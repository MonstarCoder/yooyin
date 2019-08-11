package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type MusicType struct {
	Id 						int       `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (u *MusicType) TableName() string {
	return "music_types"
}

func init() {
	orm.RegisterModel(new(MusicType))
}