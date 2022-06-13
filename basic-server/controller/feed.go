package controller

import (
	"context"
	"github.com/offer10/byte-douyin/basic-server/service"
	"github.com/offer10/byte-douyin/pb"
)

type FeedServiceImpl struct {
	pb.UnimplementedFeedServiceServer
}

var FeedService = service.FeedService{}

//实现对应接口
func (p *FeedServiceImpl) Feed(ctx context.Context, req *pb.FeedRequest) (*pb.FeedResponse, error) {
	resp, err := FeedService.GetVideos(req.LatestTime)
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

	reply := &pb.FeedResponse{
		List: list,
	}
	return reply, nil
}
