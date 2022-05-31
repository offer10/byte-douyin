package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pandalzy/byte-douyin/api-client/controller"
)

func RegisterUserRouter(r *gin.RouterGroup) {
	userController := controller.NewUserController()
	r.POST("login/", userController.Login)
}
