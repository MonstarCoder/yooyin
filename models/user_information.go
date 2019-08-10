package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type YooyinUser struct {
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Birthday      int    `json:"birthday"`
	Gender        int    `json:"gender"`
	Id            int    `json:"id"`
	//LastLoginIp   string `json:"last_login_ip"`
	//LastLoginTime int64  `json:"last_login_time"`
	//Mobile        string `json:"mobile"`
	//Password      string `json:"password"`
	//RegisterIp    string `json:"register_ip"`
	//RegisterTime  int64  `json:"register_time"`
	//UserLevelId   int    `json:"user_level_id"`
	//Username      string `json:"username"`
	WeixinOpenid  string `json:"weixin_openid"`
}

func init() {
	// set default database
	//orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/nideshop?charset=utf8mb4", 30)

	// register model
	orm.RegisterModel(new(YooyinUser))
}