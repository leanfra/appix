package service

import (
	"context"
	"fmt"

	pb "appix/api/appix/v1"

	"github.com/go-kratos/kratos/v2/log"

	//  TODO: modify project name
	biz "appix/internal/biz"
)

type DatacentersService struct {
	pb.UnimplementedDatacentersServer
	usecase *biz.DatacentersUsecase
	log     *log.Helper
}

func NewDatacentersService(uc *biz.DatacentersUsecase, logger log.Logger) *DatacentersService {
	return &DatacentersService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

func toBizDatacenter(c *pb.Datacenter) (biz.Datacenter, error) {
	if c == nil {
		return biz.Datacenter{}, fmt.Errorf("invalidDatacenter")
	}
	return biz.Datacenter{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
	}, nil
}

func toBizDatacenters(cs []*pb.Datacenter) ([]biz.Datacenter, error) {
	bizClusters := make([]biz.Datacenter, len(cs))
	for i, c := range cs {
		bizCluster, err := toBizDatacenter(c)
		if err != nil {
			return nil, err
		}
		bizClusters[i] = bizCluster
	}
	return bizClusters, nil
}

func (s *DatacentersService) CreateDatacenters(ctx context.Context, req *pb.CreateDatacentersRequest) (*pb.CreateDatacentersReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	bizDatacenters, err := toBizDatacenters(req.Datacenters)
	if err == nil {
		err = s.usecase.CreateDatacenters(ctx, bizDatacenters)
	}
	reply := &pb.CreateDatacentersReply{
		Action:  "CreateDatacenters",
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

func (s *DatacentersService) UpdateDatacenters(ctx context.Context,
	req *pb.UpdateDatacentersRequest) (*pb.UpdateDatacentersReply, error) {

	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	bizDatacenters, err := toBizDatacenters(req.Datacenters)
	if err == nil {
		err = s.usecase.UpdateDatacenters(ctx, bizDatacenters)
	}
	reply := &pb.UpdateDatacentersReply{
		Action:  "UpdateDatacenters",
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

func (s *DatacentersService) DeleteDatacenters(ctx context.Context,
	req *pb.DeleteDatacentersRequest) (*pb.DeleteDatacentersReply, error) {

	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	err := s.usecase.DeleteDatacenters(ctx, req.Ids)
	reply := &pb.DeleteDatacentersReply{
		Action:  "DeleteDatacenters",
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

func (s *DatacentersService) GetDatacenters(ctx context.Context,
	req *pb.GetDatacentersRequest) (*pb.GetDatacentersReply, error) {

	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	datacenter, err := s.usecase.GetDatacenters(ctx, req.Id)
	reply := &pb.GetDatacentersReply{
		Action:  "GetDatacenters",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Datacenter = &pb.Datacenter{
			Id:          datacenter.Id,
			Name:        datacenter.Name,
			Description: datacenter.Description,
		}
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, err
}

func (s *DatacentersService) ListDatacenters(ctx context.Context,
	req *pb.ListDatacentersRequest) (*pb.ListDatacentersReply, error) {

	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	var filter *biz.ListDatacentersFilter
	if req.Filter != nil {
		filter = &biz.ListDatacentersFilter{
			Page:     req.Filter.Page,
			PageSize: req.Filter.PageSize,
		}
		filter.Filters = make([]biz.DatacenterFilter, len(req.Filter.Filters))
		for i, f := range req.Filter.Filters {
			filter.Filters[i].Name = f.Name
		}
	}

	datacenters, err := s.usecase.ListDatacenters(ctx, filter)
	reply := &pb.ListDatacentersReply{
		Action:  "listDatacenters",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Datacenters = make([]*pb.Datacenter, len(datacenters))
		for i, d := range datacenters {
			reply.Datacenters[i] = &pb.Datacenter{
				Id:          d.Id,
				Name:        d.Name,
				Description: d.Description,
			}
		}
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, err
}
