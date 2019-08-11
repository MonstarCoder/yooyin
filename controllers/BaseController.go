package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) JsonResponse(code int, message string, data interface{}) {
	jsonRes := JsonReturnDataMessage{
		Code: code,
		Message: message,
		Data: data,
	}
	this.Data["json"] = &jsonRes
	this.ServeJSON()
	this.StopRun()
}
