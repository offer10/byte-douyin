package initialization

import (
	"github.com/gin-gonic/gin"
	"github.com/offer10/byte-douyin/api-client/router"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/douyin")
	router.RegisterUserRouter(apiGroup.Group("/user"))
	router.RegisterFavoriteRouter(apiGroup.Group("/favorite"))

	return r
}
