package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/offer10/byte-douyin/api-client/response"

	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/request"
	"github.com/offer10/byte-douyin/api-client/service"
	"github.com/offer10/byte-douyin/pb"
)

type ICommentController interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type CommentController struct{}

func NewCommentController() ICommentController {
	return CommentController{}
}

func (u CommentController) Action(ctx *gin.Context) {
	payload := request.CommentActionRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	payload.UserId = GetLoginUserId(ctx)
	if payload.ActionType == 1 {
		resp, err := service.CommentClient.Action(ctx, &pb.CommentActionRequest{
			VideoID:     payload.VideoId,
			UserID:      payload.UserId,
			ActionType:  payload.ActionType,
			CommentText: payload.CommentText,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":       err.Error(),
				"status_code": http.StatusBadRequest,
				"status_msg":  nil,
			})
			return
		}
		user, _ := GetUser(ctx, resp.Comment.UserId, GetLoginUserId(ctx))
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "",
			"comment": response.Comment{
				Id:         resp.Comment.Id,
				Content:    resp.Comment.Content,
				CreateDate: resp.Comment.CreateDate,
				User:       user,
			},
		})
	} else {
		_, err := service.CommentClient.Action(ctx, &pb.CommentActionRequest{
			VideoID:    payload.VideoId,
			UserID:     payload.UserId,
			ActionType: payload.ActionType,
			CommentID:  payload.CommentID,
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

}

func (u CommentController) List(ctx *gin.Context) {
	payload := request.CommentListRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.CommentClient.List(ctx, &pb.CommentListRequest{
		VideoID: payload.VideoId,
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
	commentList := response.CommentList{}
	for _, comment := range resp.List {
		user, _ := GetUser(ctx, comment.UserId, GetLoginUserId(ctx))
		commentList = append(commentList, response.Comment{
			Id:         comment.Id,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
			User:       user,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   nil,
		"comment_list": commentList,
	})
}

func GetLoginUserId(ctx *gin.Context) int64 {
	return ctx.GetInt64("user_id")
}

func GetUser(ctx context.Context, userId int64, seeId int64) (user response.User, err error) {
	resp, err := service.UserClient.Get(ctx, &pb.UserGetRequest{
		UserID: userId,
		SeeId:  seeId,
	})
	if err != nil {
		return response.User{}, err
	}
	user = response.User{
		Id:            resp.Id,
		Name:          resp.Name,
		FollowerCount: resp.FollowerCount,
		FollowCount:   resp.FollowCount,
		IsFollow:      resp.IsFollow,
	}

	return user, nil
}
