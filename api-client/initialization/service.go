package initialization

import "github.com/offer10/byte-douyin/api-client/service"

func SetupService() {
	service.UserConn()
	service.FavoriteConn()
}
