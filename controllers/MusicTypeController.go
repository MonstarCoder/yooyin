package controllers

import (
	"yooyin/services"
)

type MusicTypeController struct {
	BaseController
}

func (this *MusicTypeController) GetMusicTypes() {
	MusicTypeService := new(services.MusicTypeService)
	musicTypes, _ := MusicTypeService.GetMusicTypes()
	this.JsonResponse(0, "ok", musicTypes)
}
