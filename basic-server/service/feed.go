package service

import (
	"github.com/offer10/byte-douyin/basic-server/conf"
	"github.com/offer10/byte-douyin/database/model"
	"gorm.io/gorm"
	"time"
)

type IFeedService interface {
	GetVideos(latestTime int64) (list []model.Video, err error)
}
type FeedService struct {
}

func NewFeedService() IFeedService {
	return FeedService{}
}
func (u FeedService) GetVideos(latestTime int64) (list []model.Video, err error) {
	if latestTime != 0 {
		if err = conf.MySQL.Model(&model.Video{}).Limit(10).
			Select("id, author_id, cover_url, play_url, favorite_count, title,( SELECT count(*) FROM comments WHERE video_id = videos.id ) AS comment_count ").
			Where("created_at < ?",
				time.Unix(latestTime/1000+43200, 0).Format("2006-01-02 15:04:05")).
			Order("created_at desc").Find(&list).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return list, nil
			}
			return nil, err
		}
	} else {
		if err := conf.MySQL.Model(&model.Video{}).Limit(10).Order("created_at desc").
			Select("id, author_id, cover_url, play_url, favorite_count, title,( SELECT count(*) FROM comments WHERE video_id = videos.id ) AS comment_count ").
			Find(&list).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return list, nil
			}
			return nil, err
		}
	}
	return list, err
}
