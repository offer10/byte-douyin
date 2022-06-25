package controller

import (
	"context"

	"github.com/offer10/byte-douyin/database/model"
	"github.com/offer10/byte-douyin/interaction-server/service"
	"github.com/offer10/byte-douyin/pb"
)

type CommentServerImpl struct {
	pb.UnimplementedCommentServiceServer
}

var CommentService = service.CommentService{}

func (p *CommentServerImpl) Action(ctx context.Context, req *pb.CommentActionRequest) (*pb.CommentActionResponse, error) {
	if req.ActionType == 1 {
		com := model.Comment{
			UserId:  req.UserID,
			VideoId: req.VideoID,
			Content: req.CommentText,
		}
		if err := CommentService.Create(&com); err != nil {
			return nil, err
		}
		return &pb.CommentActionResponse{
			Comment: &pb.Comment{
				Id:         com.ID,
				UserId:     com.UserId,
				VideoId:    com.VideoId,
				Content:    com.Content,
				CreateDate: com.CreatedAt.Format("2006-01-02"),
			},
		}, nil
	} else {

		com := model.Comment{
			UserId:  req.UserID,
			VideoId: req.VideoID,
			BaseModel: model.BaseModel{
				ID: req.CommentID},
		}
		if err := CommentService.Delete(&com); err != nil {
			return nil, err
		}
		return &pb.CommentActionResponse{}, nil
	}

}

func (p *CommentServerImpl) List(ctx context.Context, req *pb.CommentListRequest) (*pb.CommentListResponse, error) {
	list, err := CommentService.List(req.VideoID)
	if err != nil {
		return nil, err
	}
	resList := make([]*pb.Comment, 0)

	for _, comment := range list {
		resList = append(resList, &pb.Comment{
			Id:         comment.ID,
			UserId:     comment.UserId,
			VideoId:    comment.VideoId,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("2006-01-02"),
		})
	}

	reply := &pb.CommentListResponse{
		List: resList,
	}
	return reply, nil
}
