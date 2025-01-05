package service

import (
	"context"
	"fmt"

	pb "appix/api/appix/v1"

	"github.com/go-kratos/kratos/v2/log"

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

func toBizDatacenter(c *pb.Datacenter) (*biz.Datacenter, error) {
	if c == nil {
		return nil, nil
	}
	return &biz.Datacenter{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
	}, nil
}

func toBizDatacenters(cs []*pb.Datacenter) ([]*biz.Datacenter, error) {
	bizClusters := make([]*biz.Datacenter, len(cs))
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
		return reply, nil
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
		return reply, nil
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
		return reply, nil
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
		reply.Datacenter = toPbDatacenter(datacenter)
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, nil
}

func (s *DatacentersService) ListDatacenters(ctx context.Context,
	req *pb.ListDatacentersRequest) (*pb.ListDatacentersReply, error) {

	filter := biz.DefaultDatacentersFilter()
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
	}

	datacenters, err := s.usecase.ListDatacenters(ctx, filter)
	reply := &pb.ListDatacentersReply{
		Action:  "listDatacenters",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Datacenters = toPbDatacenters(datacenters)
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, nil
}

func toPbDatacenter(d *biz.Datacenter) *pb.Datacenter {
	if d == nil {
		return nil
	}
	return &pb.Datacenter{
		Id:          d.Id,
		Name:        d.Name,
		Description: d.Description,
	}
}

func toPbDatacenters(ds []*biz.Datacenter) []*pb.Datacenter {
	if ds == nil {
		return nil
	}
	var reply []*pb.Datacenter
	for _, d := range ds {
		if d != nil {
			reply = append(reply, toPbDatacenter(d))
		}
	}
	return reply
}
