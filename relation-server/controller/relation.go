package controller

import (
	"context"
	"github.com/offer10/byte-douyin/database/model"
	"github.com/offer10/byte-douyin/pb"
	"github.com/offer10/byte-douyin/relation-server/service"
)

type RelationServiceImpl struct {
	pb.UnimplementedRelationServiceServer
}

var RelationService = service.RelationService{}

//实现对应接口
func (p *RelationServiceImpl) Action(ctx context.Context, req *pb.RelationActionRequest) (*pb.RelationActionResponse, error) {
	rel := model.Relation{
		UserId:   req.UserID,
		FollowId: req.FollowID,
	}
	if req.ActionType == 1 {
		if err := RelationService.AddFollow(&rel); err != nil {
			return nil, err
		}
	} else if req.ActionType == 2 {
		if err := RelationService.UnFollow(&rel); err != nil {
			return nil, err
		}
		RelationService.UnFollow(&rel)
	}
	return &pb.RelationActionResponse{}, nil
}

func (p *RelationServiceImpl) FollowList(ctx context.Context, req *pb.RelationFollowListRequest) (*pb.RelationFollowListResponse, error) {
	followIds, err := RelationService.GetFollowByID(req.UserID)
	if err != nil {
		return nil, err
	}
	reply := &pb.RelationFollowListResponse{
		FollowIDList: followIds,
	}
	return reply, nil
}
func (p *RelationServiceImpl) FollowerList(ctx context.Context, req *pb.RelationFollowerListRequest) (*pb.RelationFollowerListResponse, error) {
	followerIds, err := RelationService.GetFollowerByID(req.UserID)
	if err != nil {
		return nil, err
	}
	reply := &pb.RelationFollowerListResponse{
		FollowerIDList: followerIds,
	}
	return reply, nil
}
