package main

import (
	_ "./routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	_ = orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			beego.AppConfig.String("databaseUserName"),
			beego.AppConfig.String("databasePasswd"),
			beego.AppConfig.String("databaseHost"),
			beego.AppConfig.String("databasePort"),
			beego.AppConfig.String("databaseName"),
		))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
