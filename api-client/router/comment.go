package router

import (
	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/controller"
	"github.com/offer10/byte-douyin/api-client/middleware"
)

func RegisterCommentRouter(r *gin.RouterGroup) {
	commentController := controller.NewCommentController()
	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("action/", commentController.Action)
		group.GET("list/", commentController.List)
	}
}
