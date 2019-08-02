package controllers

import "github.com/astaxie/beego"

type JsonReturnDataMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonResult(c *beego.Controller, code int, message string, data interface{}) {
	jsonRes := JsonReturnDataMessage{Code: code, Message: message, Data: data}
	c.Data["json"] = &jsonRes
	c.ServeJSON()
	c.StopRun()
}
