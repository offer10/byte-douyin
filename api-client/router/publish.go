package router

import (
	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/controller"
	"github.com/offer10/byte-douyin/api-client/middleware"
)

func RegisterPublishRouter(r *gin.RouterGroup) {
	publishController := controller.NewPublishController()
	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("action/", publishController.Action)
		group.GET("list/", publishController.List)
	}
}
