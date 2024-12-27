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
		Owner:        a.Owner,
		Description:  a.Description,
		IsStateful:   a.IsStateful,
		ClusterId:    a.ClusterId,
		DatacenterId: a.DatacenterId,
		ProductId:    a.ProductId,
		TeamId:       a.TeamId,
		FeaturesId:   a.FeaturesId,
		TagsId:       a.TagsId,
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
		return reply, err
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
		return reply, err
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
		return reply, err
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
		reply.App = &pb.Application{
			Id:           bizApp.Id,
			Name:         bizApp.Name,
			Owner:        bizApp.Owner,
			Description:  bizApp.Description,
			IsStateful:   bizApp.IsStateful,
			ClusterId:    bizApp.ClusterId,
			DatacenterId: bizApp.DatacenterId,
			ProductId:    bizApp.ProductId,
			TeamId:       bizApp.TeamId,
			FeaturesId:   bizApp.FeaturesId,
			TagsId:       bizApp.TagsId,
		}
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, err
}

func (s *ApplicationsService) ListApplications(ctx context.Context, req *pb.ListApplicationsRequest) (*pb.ListApplicationsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	var filter = biz.DefaultApplicationFilter()
	if req.Filter != nil {
		filter = &biz.ListApplicationsFilter{
			Ids:           req.Filter.Ids,
			Names:         req.Filter.Names,
			IsStateful:    req.Filter.IsStateful,
			ClustersId:    req.Filter.ClustersId,
			DatacentersId: req.Filter.DatacentersId,
			ProductsId:    req.Filter.ProductsId,
			TeamsId:       req.Filter.TeamsId,
			FeaturesId:    req.Filter.FeaturesId,
			TagsId:        req.Filter.TagsId,
			HostgroupsId:  req.Filter.HostgroupsId,
			Page:          req.Filter.Page,
			PageSize:      req.Filter.PageSize,
		}
		if req.Filter.PageSize > 0 {
			filter.PageSize = req.Filter.PageSize
		}
		if req.Filter.Page > 0 {
			filter.Page = req.Filter.Page
		}
	}
	apps, err := s.usecase.ListApplications(ctx, filter)
	reply := &pb.ListApplicationsReply{
		Action:  "ListApplications",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Apps = make([]*pb.Application, len(apps))
		for i, a := range apps {
			reply.Apps[i] = &pb.Application{
				Id:           a.Id,
				Name:         a.Name,
				Owner:        a.Owner,
				Description:  a.Description,
				IsStateful:   a.IsStateful,
				ClusterId:    a.ClusterId,
				DatacenterId: a.DatacenterId,
				ProductId:    a.ProductId,
				TeamId:       a.TeamId,
				FeaturesId:   a.FeaturesId,
				TagsId:       a.TagsId,
				HostgroupsId: a.HostgroupsId,
			}
		}
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, err
}
