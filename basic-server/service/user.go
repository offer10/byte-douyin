package service

import (
	"github.com/offer10/byte-douyin/basic-server/conf"
	"github.com/offer10/byte-douyin/database/model"
	"gorm.io/gorm"
)

type IUserService interface {
	FindByUsername(user *model.User, username string) error
	Create(user *model.User) error
	Get(userId int64) (user *model.User, err error)
	IsFollow(userId int64, seeId int64) (isFollow bool, err error)
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

func (u UserService) Get(userId int64) (user *model.User, err error) {
	if err := conf.MySQL.Model(&model.User{}).
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	user.FollowCount = u.GetFollowCountByID(userId)
	user.FollowerCount = u.GetFollowerCountByID(userId)

	return user, nil
}

func (u UserService) IsFollow(userId int64, seeId int64) (isFollow bool, err error) {
	var count int64
	err = conf.MySQL.Model(&model.Relation{}).
		Where("user_id = ? AND follow_id = ?", userId, seeId).
		Count(&count).Error

	if err == nil && count > 0 {
		return true, err
	}

	return false, err
}

//获取关注数量
func (f UserService) GetFollowCountByID(userID int64) (count int64) {
	conf.MySQL.Model(&model.Relation{}).
		Where("user_id = ?", userID).
		Count(&count)
	return count
}

//获取粉丝数量
func (f UserService) GetFollowerCountByID(userID int64) (count int64) {
	conf.MySQL.Model(&model.Relation{}).
		Where("follow_id = ?", userID).
		Count(&count)
	return count
}
