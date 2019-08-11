package services

import (
	"yooyin/models"
)

type MusicTypeService struct {
	BaseService
}

func (this *MusicTypeService) GetMusicTypes() ([]*models.MusicType, error) {
	var musicTypes []*models.MusicType
	_, err := this.orm().QueryTable("music_types").All(&musicTypes)
	return musicTypes, err
}
