package service

import (
	"github.com/offer10/byte-douyin/basic-server/conf"
	"github.com/offer10/byte-douyin/database/model"
	"gorm.io/gorm"
)

type IPublishService interface {
	Create(m *model.Video) (id int64, err error)
	Get(id int64) (m *model.Video, err error)
	BatchGet(ids []int64) (list []model.Video, err error)
	List(userID int64) (list []model.Video, err error)
}
type PublishService struct{}

func NewPublishService() IPublishService {
	return PublishService{}
}

func (u PublishService) Create(m *model.Video) (id int64, err error) {
	if err := conf.MySQL.Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (u PublishService) Get(id int64) (m *model.Video, err error) {
	if err := conf.MySQL.Model(&model.Video{}).
		Where("id = ?", id).
		First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return m, nil
		}
		return nil, err
	}
	return m, nil
}

func (u PublishService) BatchGet(ids []int64) (list []model.Video, err error) {
	if err := conf.MySQL.Model(&model.Video{}).
		Where("id IN (?)", ids).
		Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return list, nil
		}
		return list, err
	}
	return list, nil
}

func (u PublishService) List(userID int64) (list []model.Video, err error) {
	if err := conf.MySQL.Model(&model.Video{}).
		Select("id, author_id, cover_url, play_url, favorite_count, title,( SELECT count(*) FROM comments WHERE video_id = videos.id ) AS comment_count ").
		Where("author_id = ?", userID).
		Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return list, nil
		}
		return nil, err
	}

	return list, err
}
