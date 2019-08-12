package controllers

import (
	"github.com/astaxie/beego"
	"yooyin/models"
	"yooyin/services"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) GetUser() *models.User {
	openId := this.GetSession("openId")

	if openId == nil {
		this.Ctx.Abort(403, "Forbidden")
	}

	user, err := new(services.UserService).GetUserByOpenId(openId.(string))

	if err != nil {
		return nil
	}

	return user
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
