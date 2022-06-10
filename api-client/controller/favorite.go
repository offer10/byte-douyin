package controller

import (
	"github.com/offer10/byte-douyin/api-client/response"
	"net/http"

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
	if err := ctx.ShouldBindJSON(&payload); err != nil {
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
	resp, err := service.FavoriteClient.List(ctx, &pb.FavoriteListRequest{
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

	// 组装数据返回
	videoList := response.VideoList{}
	for _, id := range resp.List {
		videoList = append(videoList, response.Video{
			Id:         id,
			IsFavorite: true,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"video_list":  videoList,
	})
}
