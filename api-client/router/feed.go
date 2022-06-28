package router

import (
	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/controller"
)

func RegisterFeedRouter(r *gin.RouterGroup) {
	feedController := controller.NewFeedController()
	group := r.Group("")
	group.GET("", feedController.Feed)
}
