package service

import (
	"github.com/offer10/byte-douyin/database/model"
	"github.com/offer10/byte-douyin/relation-server/conf"
)

type IRelationService interface {
	AddFollow(rel *model.Relation) error
	UnFollow(rel *model.Relation) error
	GetFollowByID(userID int64) (followIds []int64, err error)
	GetFollowCountByID(userID int64) (count int64)
	GetFollowerByID(userID int64) (followerIds []int64, err error)
	GetFollowerCountByID(userID int64) (count int64)
}
type RelationService struct{}

func NewRelationService() IRelationService {
	return RelationService{}
}

//添加关注
func (f RelationService) AddFollow(rel *model.Relation) (err error) {
	if err = conf.MySQL.Create(&rel).Error; err != nil {
		return err
	}
	return nil
}

//取消关注
func (f RelationService) UnFollow(rel *model.Relation) (err error) {
	if err = conf.MySQL.Where("user_id = ? AND follow_id", rel.UserId, rel.FollowId).
		Delete(&model.Relation{}).Error; err != nil {
		return err
	}
	return nil
}

//获取关注列表
func (f RelationService) GetFollowByID(userID int64) (followIds []int64, err error) {
	if err := conf.MySQL.Model(&model.Relation{}).
		Select("follow_id").
		Where("user_id = ?", userID).
		Find(&followIds).Error; err != nil {
		return nil, err
	}
	return followIds, err

}

//获取关注数量
func (f RelationService) GetFollowCountByID(userID int64) (count int64) {
	conf.MySQL.Model(&model.Relation{}).
		Where("userid = ?", userID).
		Count(&count)
	return count
}

//获取粉丝列表
func (f RelationService) GetFollowerByID(userID int64) (followerIds []int64, err error) {
	if err := conf.MySQL.Model(&model.Relation{}).
		Select("user_id").
		Where("follow_id = ?", userID).
		Find(&followerIds).Error; err != nil {
		return nil, err
	}
	return followerIds, err
}

//获取粉丝数量
func (f RelationService) GetFollowerCountByID(userID int64) (count int64) {
	conf.MySQL.Model(&model.Relation{}).
		Where("follow_id = ?", userID).
		Count(&count)
	return count
}
