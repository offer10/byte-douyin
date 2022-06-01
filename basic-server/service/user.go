package service

import (
	"github.com/offer10/byte-douyin/basic-server/conf"
	"github.com/offer10/byte-douyin/database/model"
)

type IUserService interface {
	FindByUsername(user *model.User, username string) error
	Create(user *model.User) error
}
type UserService struct{}

func NewUserService() IUserService {
	return UserService{}
}

func (u UserService) FindByUsername(user *model.User, username string) error {
	if err := conf.MySQL.Model(&model.User{}).
		Where("username = ?", username).
		First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Create(user *model.User) error {
	if err := conf.MySQL.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
