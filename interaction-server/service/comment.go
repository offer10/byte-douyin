package service

import (
	"github.com/offer10/byte-douyin/basic-server/conf"
	"github.com/offer10/byte-douyin/database/model"
)

type ICommentService interface {
	Delete(com *model.Comment) error
	Create(com *model.Comment) error
	List(videoID int64) (list []model.Comment, err error)
}
type CommentService struct{}

func NewCommentService() ICommentService {
	return CommentService{}
}

func (u CommentService) Delete(com *model.Comment) error {
	if err := conf.MySQL.
		Where("id = ?", com.ID).
		Delete(&model.Comment{}).Error; err != nil {
		return err
	}
	return nil
}

func (u CommentService) Create(com *model.Comment) error {
	if err := conf.MySQL.Create(&com).Error; err != nil {
		return err
	}
	return nil
}

func (u CommentService) List(videoID int64) (list []model.Comment, err error) {
	if err := conf.MySQL.Model(&model.Comment{}).
		Where("video_id = ?", videoID).
		Find(&list).Error; err != nil {
		return nil, err
	}

	return list, err
}
