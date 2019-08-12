package controllers

import (
	"yooyin/services"
)

type MusicStyleController struct {
	BaseController
}

func (this *MusicStyleController) GetMusicStyles() {
	MusicStyleService := new(services.MusicStyleService)
	musicStyles, _ := MusicStyleService.GetMusicStyles()
	this.JsonResponse(0, "ok", musicStyles)
}
