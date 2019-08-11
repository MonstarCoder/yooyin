package services

import (
	"yooyin/models"
)

type UserLikeTypeService struct {
	BaseService
}

func (this *UserLikeTypeService) GetMusicTypes() ([]*models.UserLikeType, error) {
	var userLikeTypes []*models.UserLikeType
	_, err := this.orm().QueryTable(new(models.UserLikeType)).All(&userLikeTypes)
	return userLikeTypes, err
}
