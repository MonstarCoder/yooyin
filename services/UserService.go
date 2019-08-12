package services

import (
	"yooyin/models"
)

type UserService struct {
	BaseService
}

func (this *UserService) LoginUser(user *models.User) error {
	_, _, err := this.orm().ReadOrCreate(user, "OpenId")
	return err
}

func (this *UserService) GetUserByOpenId(openId string) (*models.User, error) {
	user := &models.User{ OpenId: openId }
	err := this.orm().Read(user, "OpenId")
	return user, err
}
