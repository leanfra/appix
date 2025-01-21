package service

import (
	"context"
	"fmt"

	pb "appix/api/appix/v1"

	"github.com/go-kratos/kratos/v2/log"

	biz "appix/internal/biz"
)

type ApplicationsService struct {
	pb.UnimplementedApplicationsServer
	usecase *biz.ApplicationsUsecase
	log     *log.Helper
}

func NewApplicationsService(uc *biz.ApplicationsUsecase, logger log.Logger) *ApplicationsService {
	return &ApplicationsService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

func toBizApp(a *pb.Application) (*biz.Application, error) {
	if a == nil {
		return nil, nil
	}
	return &biz.Application{
		Id:           a.Id,
		Name:         a.Name,
		OwnerId:      a.OwnerId,
		Description:  a.Description,
		IsStateful:   a.IsStateful,
		ProductId:    a.ProductId,
		TeamId:       a.TeamId,
		FeaturesId:   a.FeaturesId,
		TagsId:       a.TagsId,
		HostgroupsId: a.HostgroupsId,
	}, nil
}

func toBizApps(apps []*pb.Application) ([]*biz.Application, error) {
	bizApps := make([]*biz.Application, len(apps))
	for i, a := range apps {
		bizApp, e := toBizApp(a)
		if e != nil {
			return nil, e
		}
		bizApps[i] = bizApp
	}
	return bizApps, nil
}

func toPbApp(a *biz.Application) (*pb.Application, error) {
	if a == nil {
		return nil, nil
	}
	return &pb.Application{
		Id:           a.Id,
		Name:         a.Name,
		Description:  a.Description,
		OwnerId:      a.OwnerId,
		IsStateful:   a.IsStateful,
		ProductId:    a.ProductId,
		TeamId:       a.TeamId,
		FeaturesId:   a.FeaturesId,
		TagsId:       a.TagsId,
		HostgroupsId: a.HostgroupsId,
	}, nil
}

func toPbApps(apps []*biz.Application) ([]*pb.Application, error) {
	pbApps := make([]*pb.Application, len(apps))
	for i, a := range apps {
		pbApp, e := toPbApp(a)
		if e != nil {
			return nil, e
		}
		pbApps[i] = pbApp
	}
	return pbApps, nil
}

func (s *ApplicationsService) CreateApplications(ctx context.Context, req *pb.CreateApplicationsRequest) (*pb.CreateApplicationsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	_bizApps, err := toBizApps(req.Apps)
	if err == nil {
		err = s.usecase.CreateApplications(ctx, _bizApps)
	}
	reply := &pb.CreateApplicationsReply{
		Action:  "createApplications",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	return reply, nil
}

func (s *ApplicationsService) UpdateApplications(ctx context.Context, req *pb.UpdateApplicationsRequest) (*pb.UpdateApplicationsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	_bizApps, err := toBizApps(req.Apps)
	if err == nil {
		err = s.usecase.UpdateApplications(ctx, _bizApps)
	}
	reply := &pb.UpdateApplicationsReply{
		Action:  "updateApplications",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	return reply, nil
}

func (s *ApplicationsService) DeleteApplications(ctx context.Context, req *pb.DeleteApplicationsRequest) (*pb.DeleteApplicationsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	err := s.usecase.DeleteApplications(ctx, req.Ids)
	reply := &pb.DeleteApplicationsReply{
		Action:  "DeleteApplications",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	return reply, nil
}
func (s *ApplicationsService) GetApplications(ctx context.Context, req *pb.GetApplicationsRequest) (*pb.GetApplicationsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	bizApp, err := s.usecase.GetApplications(ctx, req.Id)
	reply := &pb.GetApplicationsReply{
		Action:  "GetApplications",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		papp, _ := toPbApp(bizApp)
		reply.App = papp
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, nil
}

func (s *ApplicationsService) ListApplications(ctx context.Context, req *pb.ListApplicationsRequest) (*pb.ListApplicationsReply, error) {
	filter := biz.DefaultApplicationFilter()
	if req != nil {
		if len(req.Ids) > 0 {
			filter.Ids = req.Ids
		}
		if len(req.Names) > 0 {
			filter.Names = req.Names
		}
		if req.PageSize > 0 {
			filter.PageSize = req.PageSize
		}
		if req.Page > 0 {
			filter.Page = req.Page
		}
		if req.IsStateful != biz.IsStatefulNone {
			filter.IsStateful = req.IsStateful
		}
		if len(req.ProductsId) > 0 {
			filter.ProductsId = req.ProductsId
		}
		if len(req.TeamsId) > 0 {
			filter.TeamsId = req.TeamsId
		}
		if len(req.FeaturesId) > 0 {
			filter.FeaturesId = req.FeaturesId
		}
		if len(req.TagsId) > 0 {
			filter.TagsId = req.TagsId
		}
		if len(req.HostgroupsId) > 0 {
			filter.HostgroupsId = req.HostgroupsId
		}
	}

	apps, err := s.usecase.ListApplications(ctx, filter)
	reply := &pb.ListApplicationsReply{
		Action:  "listApplications",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Apps, _ = toPbApps(apps)
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, nil
}

func (s *ApplicationsService) MatchAppHostgroups(ctx context.Context, req *pb.MatchAppHostgroupsRequest) (*pb.MatchAppHostgroupsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	ids, err := s.usecase.MatchHostgroups(ctx, nil, &biz.MatchAppHostgroupsFilter{
		FeaturesId: req.FeaturesId,
		ProductId:  req.ProductId,
		TeamId:     req.TeamId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.MatchAppHostgroupsReply{
		Action:       "MatchAppHostgroups",
		Code:         0,
		Message:      "success",
		HostgroupsId: ids,
	}, nil
}
