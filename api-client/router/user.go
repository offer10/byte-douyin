package router

import (
	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/controller"
	"github.com/offer10/byte-douyin/api-client/middleware"
)

func RegisterUserRouter(r *gin.RouterGroup) {
	userController := controller.NewUserController()
	r.POST("login/", userController.Login)
	r.POST("register/", userController.Register)

	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.GET("", userController.Info)
	}
}
