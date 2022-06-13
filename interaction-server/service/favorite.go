package service

import (
	"github.com/offer10/byte-douyin/basic-server/conf"
	"github.com/offer10/byte-douyin/database/model"
	"gorm.io/gorm"
)

type IFavoriteService interface {
	Get(fav *model.Favorite) error
	Delete(fav *model.Favorite) error
	Create(fav *model.Favorite) error
	UpdateLike(fav *model.Favorite) error
	UpdateDisLike(fav *model.Favorite) error
	List(userID int64) (videoIds []int64, err error)
	IsFav(userID int64, videoId int64) (isFav bool, err error)
}
type FavoriteService struct{}

func NewFavoriteService() IFavoriteService {
	return FavoriteService{}
}

func (u FavoriteService) Get(fav *model.Favorite) error {
	if err := conf.MySQL.Model(&model.Favorite{}).
		Where("user_id = ? AND video_id = ?", fav.UserId, fav.VideoId).
		First(&fav).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

func (u FavoriteService) Delete(fav *model.Favorite) error {
	if err := conf.MySQL.
		Where("user_id = ? AND video_id = ?", fav.UserId, fav.VideoId).
		Delete(&model.Favorite{}).Error; err != nil {
		return err
	}
	return nil
}

func (u FavoriteService) Create(fav *model.Favorite) error {
	if err := conf.MySQL.Create(&fav).Error; err != nil {
		return err
	}
	return nil
}

func (u FavoriteService) UpdateLike(fav *model.Favorite) error {
	if err := conf.MySQL.Model(&model.Video{}).
		Where("id = ?", fav.VideoId).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (u FavoriteService) UpdateDisLike(fav *model.Favorite) error {
	if err := conf.MySQL.Model(&model.Video{}).
		Where("id = ? AND favorite_count > ?", fav.VideoId, 0).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (u FavoriteService) List(userID int64) (videoIds []int64, err error) {
	if err := conf.MySQL.Model(&model.Favorite{}).
		Select("video_id").
		Where("user_id = ?", userID).
		Find(&videoIds).Error; err != nil {
		return nil, err
	}

	return videoIds, err
}

func (u FavoriteService) IsFav(UserID int64, videoId int64) (isFav bool, err error) {
	var count int64
	if err := conf.MySQL.Model(&model.Favorite{}).
		Where("user_id = ? AND video_id = ?", UserID, videoId).
		Count(&count).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, err
}
