package services

import (
	"yooyin/models"
)

type MusicStyleService struct {
	BaseService
}

func (this *MusicStyleService) GetMusicStyles() ([]*models.MusicStyle, error) {
	var musicTypes []*models.MusicStyle
	_, err := this.orm().QueryTable(new(models.MusicStyle)).All(&musicTypes)
	return musicTypes, err
}
