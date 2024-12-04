package service

import (
	"context"
	"errors"

	pb "appix/api/appix/v1"
	biz "appix/internal/biz"
)

type TagsService struct {
	pb.UnimplementedTagsServer
	usecase biz.TagsUsecase
}

func NewTagsService(uc biz.TagsUsecase) *TagsService {
	return &TagsService{
		usecase: uc,
	}
}

// TODO if we need process ctx timeout?

func (s *TagsService) CreateTags(ctx context.Context, req *pb.CreateTagsRequest) (*pb.CreateTagsReply, error) {

	if req == nil {
		return nil, errors.New("req is nil")
	}

	tags := make([]biz.Tag, len(req.Tags))
	for i, tag := range req.Tags {
		tags[i] = biz.Tag{
			Key:   tag.Key,
			Value: tag.Value,
		}
	}
	err := s.usecase.CreateTags(ctx, tags)

	reply := &pb.CreateTagsReply{
		Action:  "CreateTags",
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

func (s *TagsService) UpdateTags(ctx context.Context, req *pb.UpdateTagsRequest) (*pb.UpdateTagsReply, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}

	tags := make([]biz.Tag, len(req.Tags))
	for i, tag := range req.Tags {
		tags[i] = biz.Tag{
			Key:   tag.Key,
			Value: tag.Value,
		}
	}
	err := s.usecase.UpdateTags(ctx, tags)

	reply := &pb.UpdateTagsReply{
		Action:  "UpdateTags",
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

func (s *TagsService) DeleteTags(ctx context.Context, req *pb.DeleteTagsRequest) (*pb.DeleteTagsReply, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}

	err := s.usecase.DeleteTags(ctx, req.Ids)

	reply := &pb.DeleteTagsReply{
		Action:  "DeleteTags",
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

func (s *TagsService) GetTags(ctx context.Context, req *pb.GetTagsRequest) (*pb.GetTagsReply, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}

	tag, err := s.usecase.GetTags(ctx, req.Id)

	reply := &pb.GetTagsReply{
		Action:  "GetTags",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Tag = &pb.Tag{
			Id:    tag.Id,
			Key:   tag.Key,
			Value: tag.Value,
		}
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, err
}

func (s *TagsService) ListTags(ctx context.Context, req *pb.ListTagsRequest) (*pb.ListTagsReply, error) {
	if req == nil {
		return nil, errors.New("req is nil")
	}
	var filter *biz.ListTagsFilter
	if req.Filter != nil {
		filter = &biz.ListTagsFilter{
			PageSize: req.Filter.PageSize,
			Page:     req.Filter.Page,
		}

		filter.Filters = make([]biz.TagFilter, len(req.Filter.Filters))
		for i, f := range req.Filter.Filters {
			filter.Filters[i].Key = f.Key
			filter.Filters[i].Value = f.Value
		}
	}

	tags, err := s.usecase.ListTags(ctx, filter)

	reply := &pb.ListTagsReply{
		Action:  "ListTags",
		Code:    0,
		Message: "success",
	}

	if err == nil {
		reply.Tags = make([]*pb.Tag, len(tags))
		for i, tag := range tags {
			reply.Tags[i] = &pb.Tag{
				Id:    tag.Id,
				Key:   tag.Key,
				Value: tag.Value,
			}
		}
		return reply, nil
	}

	reply.Code = 1
	reply.Message = err.Error()

	return reply, err
}
