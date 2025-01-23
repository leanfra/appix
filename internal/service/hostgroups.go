package service

import (
	"context"

	pb "opspillar/api/opspillar/v1"

	"github.com/go-kratos/kratos/v2/log"

	biz "opspillar/internal/biz"
)

type HostgroupsService struct {
	pb.UnimplementedHostgroupsServer
	usecase *biz.HostgroupsUsecase
	log     *log.Helper
}

func NewHostgroupsService(uc *biz.HostgroupsUsecase, logger log.Logger) *HostgroupsService {
	return &HostgroupsService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

func toBizHostgroup(p *pb.Hostgroup) (*biz.Hostgroup, error) {
	if p == nil {
		return nil, nil
	}
	return &biz.Hostgroup{
		Id:              p.Id,
		Name:            p.Name,
		Description:     p.Description,
		ClusterId:       p.ClusterId,
		DatacenterId:    p.DatacenterId,
		EnvId:           p.EnvId,
		ProductId:       p.ProductId,
		TeamId:          p.TeamId,
		FeaturesId:      p.FeaturesId,
		TagsId:          p.TagsId,
		ShareProductsId: p.ShareProductsId,
		ShareTeamsId:    p.ShareTeamsId,
	}, nil
}

func toBizHostgroups(ps []*pb.Hostgroup) ([]*biz.Hostgroup, error) {
	if ps == nil {
		return nil, nil
	}
	bizPs := make([]*biz.Hostgroup, len(ps))
	for i, p := range ps {
		bizP, err := toBizHostgroup(p)
		if err != nil {
			return nil, err
		}
		bizPs[i] = bizP
	}
	return bizPs, nil
}

func (s *HostgroupsService) CreateHostgroups(ctx context.Context, req *pb.CreateHostgroupsRequest) (*pb.CreateHostgroupsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	bizHostgroups, err := toBizHostgroups(req.Hostgroups)
	if err == nil {
		err = s.usecase.CreateHostgroups(ctx, bizHostgroups)
	}
	reply := &pb.CreateHostgroupsReply{
		Action:  "CreateHostgroups",
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

func (s *HostgroupsService) UpdateHostgroups(ctx context.Context, req *pb.UpdateHostgroupsRequest) (*pb.UpdateHostgroupsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	bizHostgroups, err := toBizHostgroups(req.Hostgroups)
	if err == nil {
		err = s.usecase.UpdateHostgroups(ctx, bizHostgroups)
	}
	reply := &pb.UpdateHostgroupsReply{
		Action:  "UpdateHostgroups",
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

func (s *HostgroupsService) DeleteHostgroups(ctx context.Context, req *pb.DeleteHostgroupsRequest) (*pb.DeleteHostgroupsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	err := s.usecase.DeleteHostgroups(ctx, req.Ids)
	reply := &pb.DeleteHostgroupsReply{
		Action:  "DeleteHostgroups",
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

func (s *HostgroupsService) GetHostgroups(ctx context.Context, req *pb.GetHostgroupsRequest) (*pb.GetHostgroupsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	bizhostgroups, err := s.usecase.GetHostgroups(ctx, req.Id)
	reply := &pb.GetHostgroupsReply{
		Action:  "GetHostgroups",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	reply.Hostgroup = toPbHostgroup(bizhostgroups)
	return reply, nil
}

func (s *HostgroupsService) ListHostgroups(ctx context.Context, req *pb.ListHostgroupsRequest) (*pb.ListHostgroupsReply, error) {
	filter := biz.DefaultHostgroupFilter()
	if req != nil {
		if len(req.Names) > 0 {
			filter.Names = req.Names
		}
		if len(req.Ids) > 0 {
			filter.Ids = req.Ids
		}
		if len(req.ClustersId) > 0 {
			filter.ClustersId = req.ClustersId
		}
		if len(req.DatacentersId) > 0 {
			filter.DatacentersId = req.DatacentersId
		}
		if len(req.EnvsId) > 0 {
			filter.EnvsId = req.EnvsId
		}
		if len(req.ProductsId) > 0 {
			filter.ProductsId = req.ProductsId
		}
		if len(req.ShareProductsId) > 0 {
			filter.ShareProductsId = req.ShareProductsId
		}
		if len(req.ShareTeamsId) > 0 {
			filter.ShareTeamsId = req.ShareTeamsId
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
		if req.PageSize > 0 {
			filter.PageSize = req.PageSize
		}
		if req.Page > 0 {
			filter.Page = req.Page
		}
	}
	hgs, err := s.usecase.ListHostgroups(ctx, filter)
	reply := &pb.ListHostgroupsReply{
		Action:  "ListHostgroups",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	reply.Hostgroups = toPbHostgroups(hgs)
	return reply, nil
}

func toPbHostgroup(bizHostgroup *biz.Hostgroup) *pb.Hostgroup {
	if bizHostgroup == nil {
		return nil
	}
	return &pb.Hostgroup{
		Id:              bizHostgroup.Id,
		Name:            bizHostgroup.Name,
		Description:     bizHostgroup.Description,
		ClusterId:       bizHostgroup.ClusterId,
		DatacenterId:    bizHostgroup.DatacenterId,
		EnvId:           bizHostgroup.EnvId,
		ProductId:       bizHostgroup.ProductId,
		TeamId:          bizHostgroup.TeamId,
		FeaturesId:      bizHostgroup.FeaturesId,
		TagsId:          bizHostgroup.TagsId,
		ShareProductsId: bizHostgroup.ShareProductsId,
		ShareTeamsId:    bizHostgroup.ShareTeamsId,
	}
}

func toPbHostgroups(bizHostgroups []*biz.Hostgroup) []*pb.Hostgroup {
	if bizHostgroups == nil {
		return nil
	}
	pbHostgroups := make([]*pb.Hostgroup, len(bizHostgroups))
	for i, bizHostgroup := range bizHostgroups {
		pbHostgroups[i] = toPbHostgroup(bizHostgroup)
	}
	return pbHostgroups
}
