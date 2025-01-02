package service

import (
	"context"
	"fmt"

	pb "appix/api/appix/v1"

	"github.com/go-kratos/kratos/v2/log"

	biz "appix/internal/biz"
)

type ClustersService struct {
	pb.UnimplementedClustersServer
	usecase *biz.ClustersUsecase
	log     *log.Helper
}

func NewClustersService(uc *biz.ClustersUsecase, logger log.Logger) *ClustersService {
	return &ClustersService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

func toBizCluster(c *pb.Cluster) (*biz.Cluster, error) {
	if c == nil {
		return nil, nil
	}
	return &biz.Cluster{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
	}, nil
}

func toBizClusters(cs []*pb.Cluster) ([]*biz.Cluster, error) {
	bizClusters := make([]*biz.Cluster, len(cs))
	for i, c := range cs {
		bizCluster, err := toBizCluster(c)
		if err != nil {
			return nil, err
		}
		bizClusters[i] = bizCluster
	}
	return bizClusters, nil
}

func (s *ClustersService) CreateClusters(ctx context.Context, req *pb.CreateClustersRequest) (*pb.CreateClustersReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	_bizclusters, err := toBizClusters(req.Clusters)
	if err == nil {
		err = s.usecase.CreateClusters(ctx, _bizclusters)
	}

	reply := &pb.CreateClustersReply{
		Action:  "CreateClusters",
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

func (s *ClustersService) UpdateClusters(ctx context.Context, req *pb.UpdateClustersRequest) (*pb.UpdateClustersReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	_bizclusters, err := toBizClusters(req.Clusters)
	if err == nil {
		err = s.usecase.UpdateClusters(ctx, _bizclusters)
	}
	reply := &pb.UpdateClustersReply{
		Action:  "UpdateClusters",
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

func (s *ClustersService) DeleteClusters(ctx context.Context, req *pb.DeleteClustersRequest) (*pb.DeleteClustersReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	err := s.usecase.DeleteClusters(ctx, req.Ids)
	reply := &pb.DeleteClustersReply{
		Action:  "DeleteClusters",
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

func (s *ClustersService) GetClusters(ctx context.Context, req *pb.GetClustersRequest) (*pb.GetClustersReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	bizCluster, err := s.usecase.GetClusters(ctx, req.Id)
	reply := &pb.GetClustersReply{
		Action:  "GetClusters",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Cluster = &pb.Cluster{
			Id:          bizCluster.Id,
			Name:        bizCluster.Name,
			Description: bizCluster.Description,
		}
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, nil
}

func (s *ClustersService) ListClusters(ctx context.Context, req *pb.ListClustersRequest) (*pb.ListClustersReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	filter := biz.DefaultClusterFilter()
	if req.Filter != nil {
		filter = &biz.ListClustersFilter{
			Ids:   req.Filter.Ids,
			Names: req.Filter.Names,
		}
		if req.Filter.PageSize > 0 {
			filter.PageSize = req.Filter.PageSize
		}
		if req.Filter.Page > 0 {
			filter.Page = req.Filter.Page
		}
	}

	cs, err := s.usecase.ListClusters(ctx, filter)
	reply := &pb.ListClustersReply{
		Action:  "listClusters",
		Code:    0,
		Message: "success",
	}

	if err == nil {
		reply.Clusters = make([]*pb.Cluster, len(cs))
		for i, c := range cs {
			reply.Clusters[i] = &pb.Cluster{
				Id:          c.Id,
				Name:        c.Name,
				Description: c.Description,
			}
		}
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, nil
}
