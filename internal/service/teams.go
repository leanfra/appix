package service

import (
	"context"

	pb "appix/api/appix/v1"

	"github.com/go-kratos/kratos/v2/log"

	biz "appix/internal/biz"
)

type TeamsService struct {
	pb.UnimplementedTeamsServer
	usecase *biz.TeamsUsecase
	log     *log.Helper
}

func NewTeamsService(uc *biz.TeamsUsecase, logger log.Logger) *TeamsService {
	return &TeamsService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

func (s *TeamsService) CreateTeams(ctx context.Context, req *pb.CreateTeamsRequest) (*pb.CreateTeamsReply, error) {

	if req == nil {
		return nil, ErrRequestNil
	}

	ts := make([]biz.Team, len(req.Teams))
	for i, t := range req.Teams {
		ts[i] = biz.Team{
			Code:        t.Code,
			Description: t.Description,
			Leader:      t.Leader,
			Name:        t.Name,
		}
	}
	err := s.usecase.CreateTeams(ctx, ts)

	reply := &pb.CreateTeamsReply{
		Action:  "CreateTeams",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, err
	}

	return reply, err
}

func (s *TeamsService) UpdateTeams(ctx context.Context, req *pb.UpdateTeamsRequest) (*pb.UpdateTeamsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	ts := make([]biz.Team, len(req.Teams))
	for i, t := range req.Teams {
		ts[i] = biz.Team{
			Id:          t.Id,
			Code:        t.Code,
			Description: t.Description,
			Leader:      t.Leader,
			Name:        t.Name,
		}
	}
	err := s.usecase.UpdateTeams(ctx, ts)
	reply := &pb.UpdateTeamsReply{
		Action:  "UpdateTeams",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, err
	}
	return reply, err
}

func (s *TeamsService) DeleteTeams(ctx context.Context, req *pb.DeleteTeamsRequest) (*pb.DeleteTeamsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	err := s.usecase.DeleteTeams(ctx, req.Ids)
	reply := &pb.DeleteTeamsReply{
		Action:  "DeleteTeams",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, err
	}
	return reply, err
}

func (s *TeamsService) GetTeams(ctx context.Context, req *pb.GetTeamsRequest) (*pb.GetTeamsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	t, err := s.usecase.GetTeams(ctx, req.Id)
	reply := &pb.GetTeamsReply{
		Action:  "GetTeams",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Team = &pb.Team{
			Id:          t.Id,
			Code:        t.Code,
			Description: t.Description,
			Leader:      t.Leader,
			Name:        t.Name,
		}
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, err
}

func (s *TeamsService) ListTeams(ctx context.Context, req *pb.ListTeamsRequest) (*pb.ListTeamsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	var filter *biz.ListTeamsFilter
	if req.Filter != nil {
		filter = &biz.ListTeamsFilter{
			Page:     req.Filter.Page,
			PageSize: req.Filter.PageSize,
		}

		filter.Filters = make([]biz.TeamFilter, len(req.Filter.Filters))
		for i, f := range req.Filter.Filters {
			filter.Filters[i] = biz.TeamFilter{
				Code:   f.Code,
				Name:   f.Name,
				Leader: f.Leader,
			}
		}
	}

	ts, err := s.usecase.ListTeams(ctx, filter)
	reply := &pb.ListTeamsReply{
		Action:  "ListTeams",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Teams = make([]*pb.Team, len(ts))
		for i, t := range ts {
			reply.Teams[i] = &pb.Team{
				Id:          t.Id,
				Code:        t.Code,
				Description: t.Description,
				Leader:      t.Leader,
				Name:        t.Name,
			}
		}
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, err
}
