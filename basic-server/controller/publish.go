package controller

import (
	"context"
	"github.com/offer10/byte-douyin/basic-server/service"
	"github.com/offer10/byte-douyin/database/model"
	"github.com/offer10/byte-douyin/pb"
)

type PublishServerImpl struct {
	pb.UnimplementedPublishServiceServer
}

var PublishService = service.PublishService{}

func (p *PublishServerImpl) Action(ctx context.Context, req *pb.PublishActionRequest) (*pb.PublishActionResponse, error) {
	m := model.Video{
		AuthorID: req.AuthorID,
		Title:    req.Title,
		PlayURL:  req.PlayUrl,
		CoverURL: req.CoverUrl,
	}
	id, err := PublishService.Create(&m)
	if err != nil {
		return nil, err
	}

	reply := &pb.PublishActionResponse{
		VideoID: int64(id),
	}
	return reply, nil
}

func (p *PublishServerImpl) List(ctx context.Context, req *pb.PublishListRequest) (*pb.PublishListResponse, error) {
	resp, err := PublishService.List(req.UserID)
	if err != nil {
		return nil, err
	}

	list := make([]*pb.Video, 0)
	for _, video := range resp {
		list = append(list, &pb.Video{
			Id:            video.ID,
			AuthorId:      video.AuthorID,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
		})
	}

	reply := &pb.PublishListResponse{
		List: list,
	}
	return reply, nil
}

func (p *PublishServerImpl) BatchGet(ctx context.Context, req *pb.PublishBatchGetRequest) (*pb.PublishBatchGetResponse, error) {
	resp, err := PublishService.BatchGet(req.Ids)
	if err != nil {
		return nil, err
	}

	list := make([]*pb.Video, 0)
	for _, video := range resp {
		list = append(list, &pb.Video{
			Id:            video.ID,
			AuthorId:      video.AuthorID,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
		})
	}

	reply := &pb.PublishBatchGetResponse{
		List: list,
	}
	return reply, nil
}
