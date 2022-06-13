package router

import (
	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/controller"
	"github.com/offer10/byte-douyin/api-client/middleware"
)

func RegisterFeedRouter(r *gin.RouterGroup) {
	feedController := controller.NewFeedController()
	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.GET("", feedController.Feed)
	}
}
