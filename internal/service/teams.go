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

func toBizTeam(t *pb.Team) (*biz.Team, error) {
	return &biz.Team{
		Id:          t.Id,
		Code:        t.Code,
		Description: t.Description,
		Leader:      t.Leader,
		Name:        t.Name,
	}, nil
}

func toBizTeams(ts []*pb.Team) ([]*biz.Team, error) {
	var biz_teams = make([]*biz.Team, len(ts))
	for i, t := range ts {
		biz_team, err := toBizTeam(t)
		if err != nil {
			return nil, err
		}
		biz_teams[i] = biz_team
	}
	return biz_teams, nil
}

func (s *TeamsService) CreateTeams(ctx context.Context, req *pb.CreateTeamsRequest) (*pb.CreateTeamsReply, error) {

	if req == nil {
		return nil, ErrRequestNil
	}

	ts, err := toBizTeams(req.Teams)
	if err == nil {
		err = s.usecase.CreateTeams(ctx, ts)
	}

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
	ts, err := toBizTeams(req.Teams)
	if err == nil {
		err = s.usecase.UpdateTeams(ctx, ts)
	}
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
		reply.Team = toPbTeam(t)
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
			Ids:      req.Filter.Ids,
			Codes:    req.Filter.Codes,
			Names:    req.Filter.Names,
			Leaders:  req.Filter.Leaders,
		}

	}

	ts, err := s.usecase.ListTeams(ctx, filter)
	reply := &pb.ListTeamsReply{
		Action:  "ListTeams",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Teams = toPbTeams(ts)
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, err
}

func toPbTeam(t *biz.Team) *pb.Team {
	if t == nil {
		return nil
	}
	return &pb.Team{
		Id:          t.Id,
		Code:        t.Code,
		Description: t.Description,
		Leader:      t.Leader,
		Name:        t.Name,
	}
}

func toPbTeams(ts []*biz.Team) []*pb.Team {
	var res []*pb.Team
	for _, t := range ts {
		if t != nil {
			res = append(res, &pb.Team{
				Id:          t.Id,
				Code:        t.Code,
				Description: t.Description,
				Leader:      t.Leader,
				Name:        t.Name,
			})
		}
	}
	return res
}
