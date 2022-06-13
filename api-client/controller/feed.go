package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/response"
	"github.com/offer10/byte-douyin/api-client/service"
	"github.com/offer10/byte-douyin/pb"
	"net/http"
	"strconv"
)

type IFeedController interface {
	Feed(ctx *gin.Context)
}
type FeedController struct {
}

func NewFeedController() IFeedController {
	return FeedController{}
}
func (u FeedController) Feed(ctx *gin.Context) {
	var latestTime, nextTime int64
	latestTime_ := ctx.Query("latest_time")
	if latestTime_ != "" {
		latestTime, _ = strconv.ParseInt(latestTime_, 10, 64)
	} else {
		latestTime = 0
	}
	resp, err := service.FeedClient.Feed(ctx, &pb.FeedRequest{
		LatestTime: latestTime,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  nil,
		})
	}
	videoList := response.VideoList{}
	for _, video := range resp.List {
		user, _ := GetUser(ctx, video.AuthorId, 0)
		videoList = append(videoList, response.Video{
			Author:        user,
			Id:            video.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			IsFavorite:    true,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
		"next_time":   nextTime,
		"video_list":  videoList,
	})
}
