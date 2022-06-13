package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/request"
	"github.com/offer10/byte-douyin/api-client/response"
	"github.com/offer10/byte-douyin/api-client/service"
	"github.com/offer10/byte-douyin/database/model"
	"github.com/offer10/byte-douyin/pb"
	"github.com/offer10/byte-douyin/relation-server/conf"
	"net/http"
)

type IRelationController interface {
	Action(ctx *gin.Context)
	FollowList(ctx *gin.Context)
	FollowerList(ctx *gin.Context)
}
type RelationController struct{}

func NewRelationController() IRelationController {
	return RelationController{}
}

func (u RelationController) Action(ctx *gin.Context) {
	payload := request.RelationActionRequest{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	_, err := service.RelationClient.Action(ctx, &pb.RelationActionRequest{
		UserID:     payload.UserId,
		FollowID:   payload.FollowId,
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

func (u RelationController) FollowList(ctx *gin.Context) {
	payload := request.RelationFollowListRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.RelationClient.FollowList(ctx, &pb.RelationFollowListRequest{
		UserID: payload.UserId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  nil,
		})
	}

	followList := response.UserList{}
	for _, id := range resp.GetFollowIDList() {
		user, _ := GetUser(ctx, id, GetLoginUserId(ctx))
		followList = append(followList, user)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
		"user_list":   followList,
	})
}

func (u RelationController) FollowerList(ctx *gin.Context) {
	payload := request.RelationFollowerListRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.RelationClient.FollowerList(ctx, &pb.RelationFollowerListRequest{
		UserID: payload.UserId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  nil,
		})
	}

	followerList := response.UserList{}
	for _, id := range resp.GetFollowerIDList() {
		user, _ := GetUser(ctx, id, GetLoginUserId(ctx))
		followerList = append(followerList, user)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
		"user_list":   followerList,
	})
}

//
func GetUserByID(req *request.RelationFollowListRequest) (*response.User, error) {
	var user model.User
	res := &response.User{}
	conf.MySQL.First(&user, req.UserId)
	res.Name = user.Username
	res.Id = int64(user.ID)
	res.FollowCount = user.FollowCount
	res.FollowerCount = user.FollowerCount
	res.IsFollow = true
	return res, nil
}
