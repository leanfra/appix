package service

import (
	"context"
	"fmt"

	pb "appix/api/appix/v1"
	biz "appix/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type FeaturesService struct {
	pb.UnimplementedFeaturesServer
	usecase *biz.FeaturesUsecase
	log     *log.Helper
}

func NewFeaturesService(uc *biz.FeaturesUsecase, logger log.Logger) *FeaturesService {
	return &FeaturesService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

func toBizFeature(feature *pb.Feature) (*biz.Feature, error) {

	if feature == nil {
		return nil, nil
	}

	return &biz.Feature{
		Id:    feature.Id,
		Name:  feature.Name,
		Value: feature.Value,
	}, nil

}

func toBizFeatures(features []*pb.Feature) ([]*biz.Feature, error) {
	_bizfeatures := make([]*biz.Feature, len(features))
	var err error
	for i, f := range features {
		if _bizfeatures[i], err = toBizFeature(f); err != nil {
			return nil, err
		}
	}
	return _bizfeatures, nil
}

func (s *FeaturesService) CreateFeatures(ctx context.Context, req *pb.CreateFeaturesRequest) (*pb.CreateFeaturesReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	features, err := toBizFeatures(req.Features)
	if err == nil {
		err = s.usecase.CreateFeatures(ctx, features)
	}

	reply := &pb.CreateFeaturesReply{
		Action:  "CreateFeatures",
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

func (s *FeaturesService) UpdateFeatures(ctx context.Context, req *pb.UpdateFeaturesRequest) (*pb.UpdateFeaturesReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	features, err := toBizFeatures(req.Features)
	if err == nil {
		err = s.usecase.UpdateFeatures(ctx, features)
	}

	reply := &pb.UpdateFeaturesReply{
		Action:  "UpdateFeatures",
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

func (s *FeaturesService) DeleteFeatures(ctx context.Context, req *pb.DeleteFeaturesRequest) (*pb.DeleteFeaturesReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	err := s.usecase.DeleteFeatures(ctx, req.Ids)
	reply := &pb.DeleteFeaturesReply{
		Action:  "DeleteFeatures",
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

func (s *FeaturesService) GetFeatures(ctx context.Context, req *pb.GetFeaturesRequest) (*pb.GetFeaturesReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	feature, err := s.usecase.GetFeatures(ctx, req.Id)

	reply := &pb.GetFeaturesReply{
		Action:  "GetFeatures",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Feature = &pb.Feature{
			Id:    feature.Id,
			Name:  feature.Name,
			Value: feature.Value,
		}
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, err
}

func (s *FeaturesService) ListFeatures(ctx context.Context,
	req *pb.ListFeaturesRequest) (*pb.ListFeaturesReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	var filter *biz.ListFeaturesFilter
	if req.Filter != nil {
		filter = &biz.ListFeaturesFilter{
			PageSize: req.Filter.PageSize,
			Page:     req.Filter.Page,
			Ids:      req.Filter.Ids,
			Names:    req.Filter.Names,
			Kvs:      req.Filter.Kvs,
		}

	}

	features, err := s.usecase.ListFeatures(ctx, filter)

	reply := &pb.ListFeaturesReply{
		Action:  "ListFeatures",
		Code:    0,
		Message: "success",
	}

	if err == nil {
		reply.Features = make([]*pb.Feature, len(features))
		for i, tag := range features {
			reply.Features[i] = &pb.Feature{
				Id:    tag.Id,
				Name:  tag.Name,
				Value: tag.Value,
			}
		}
		return reply, nil
	}

	reply.Code = 1
	reply.Message = err.Error()

	return reply, err

}
