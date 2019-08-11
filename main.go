package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "yooyin/routers"
)

func init() {
	_ = orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			beego.AppConfig.String("database::user"),
			beego.AppConfig.String("database::password"),
			beego.AppConfig.String("database::host"),
			beego.AppConfig.String("database::port"),
			beego.AppConfig.String("database::database"),
		))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
