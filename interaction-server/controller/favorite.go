package controller

import (
	"context"

	"github.com/offer10/byte-douyin/database/model"
	"github.com/offer10/byte-douyin/interaction-server/service"
	"github.com/offer10/byte-douyin/pb"
)

type FavoriteServerImpl struct {
	pb.UnimplementedFavoriteServiceServer
}

var FavoriteService = service.FavoriteService{}

func (p *FavoriteServerImpl) Action(ctx context.Context, req *pb.FavoriteActionRequest) (*pb.FavoriteActionResponse, error) {
	fav := model.Favorite{
		UserId:  req.UserID,
		VideoId: req.VideoID,
	}
	// 点赞
	if req.ActionType == 1 {
		if err := FavoriteService.Get(&fav); err != nil {
			return nil, err
		}
		// 无点赞记录，插入一条
		if fav.ID == 0 {
			if err := FavoriteService.Create(&fav); err != nil {
				return nil, err
			}
			if err := FavoriteService.UpdateLike(&fav); err != nil {
				return nil, err
			}
		}
		// 取消点赞
	} else if req.ActionType == 2 {
		if err := FavoriteService.Delete(&fav); err != nil {
			return nil, err
		}
		if err := FavoriteService.UpdateDisLike(&fav); err != nil {
			return nil, err
		}
	}

	return &pb.FavoriteActionResponse{}, nil
}

func (p *FavoriteServerImpl) List(ctx context.Context, req *pb.FavoriteListRequest) (*pb.FavoriteListResponse, error) {
	videoIds, err := FavoriteService.List(req.UserID)
	if err != nil {
		return nil, err
	}

	reply := &pb.FavoriteListResponse{
		List: videoIds,
	}
	return reply, nil
}
