package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type MusicStyle struct {
	Id 						int       `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (u *MusicStyle) TableName() string {
	return "music_styles"
}

func init() {
	orm.RegisterModel(new(MusicStyle))
}
