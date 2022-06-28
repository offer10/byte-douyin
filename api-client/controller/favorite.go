package controller

import (
	"github.com/offer10/byte-douyin/api-client/response"
	"net/http"

	"context"
	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/request"
	"github.com/offer10/byte-douyin/api-client/service"
	"github.com/offer10/byte-douyin/pb"
)

type IFavoriteController interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type FavoriteController struct{}

func NewFavoriteController() IFavoriteController {
	return FavoriteController{}
}

func (u FavoriteController) Action(ctx *gin.Context) {
	payload := request.FavoriteActionRequest{}
	payload.UserId = GetLoginUserId(ctx)
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	_, err := service.FavoriteClient.Action(ctx, &pb.FavoriteActionRequest{
		UserID:     payload.UserId,
		VideoID:    payload.VideoId,
		ActionType: payload.ActionType,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
	})
}

func (u FavoriteController) List(ctx *gin.Context) {
	payload := request.FavoriteListRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp1, err := service.FavoriteClient.List(ctx, &pb.FavoriteListRequest{
		UserID: payload.UserId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  nil,
			"error":       err.Error(),
		})
		return
	}

	resp2, _ := service.PublishClient.BatchGet(ctx, &pb.PublishBatchGetRequest{
		Ids: resp1.List,
	})

	if resp2 == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  nil,
			"video_list":  nil,
		})
		return
	}

	// 组装数据返回
	videoList := response.VideoList{}
	for _, video := range resp2.List {
		user, _ := GetUser(ctx, video.AuthorId, 0)
		isFav, _ := GetIsFav(ctx, GetLoginUserId(ctx), video.Id)
		videoList = append(videoList, response.Video{
			Author:        user,
			Id:            video.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			IsFavorite:    isFav,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"video_list":  videoList,
	})
}

func GetIsFav(ctx context.Context, userId int64, videoId int64) (isFav bool, err error) {
	resp, err := service.FavoriteClient.IsFav(ctx, &pb.FavoriteIsFavRequest{
		UserID:  userId,
		VideoID: videoId,
	})
	if err != nil {
		return false, err
	}

	return resp.IsFav, nil
}
