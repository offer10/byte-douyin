package initialization

import (
	"github.com/offer10/byte-douyin/api-client/service"
	"github.com/offer10/byte-douyin/basic-server/initialization"
)

func SetupService() {
	initialization.RegisterOSS()
	service.UserConn()
	service.FavoriteConn()
	service.CommentConn()
	service.RelationConn()
	service.PublishConn()
	service.FeedConn()
}
