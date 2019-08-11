package controllers

import (
	"yooyin/services"
)

type UserLikeTypeController struct {
	BaseController
}

func (this *UserLikeTypeController) GetUserLikeTypes() {
	UserLikeTypeService := new(services.UserLikeTypeService)
	userLikeTypes, _ := UserLikeTypeService.GetMusicTypes()
	this.JsonResponse(0, "ok", userLikeTypes)
}
