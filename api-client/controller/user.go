package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pandalzy/byte-douyin/api-client/request"
	"github.com/pandalzy/byte-douyin/api-client/service"
	"github.com/pandalzy/byte-douyin/pb"
)

type IUserController interface {
	Login(ctx *gin.Context)
}

type UserController struct{}

func NewUserController() IUserController {
	return UserController{}
}

func (c UserController) Login(ctx *gin.Context) {
	payload := request.UserLoginRequest{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := service.UserClient.Login(ctx, &pb.UserLoginRequest{
		Username: payload.Username,
		Password: payload.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user_id":     resp.UserID,
		"token":       resp.Token,
	})
}
