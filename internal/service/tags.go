package service

import (
	"context"

	pb "appix/api/appix/v1"
	biz "appix/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type TagsService struct {
	pb.UnimplementedTagsServer
	usecase *biz.TagsUsecase
	log     *log.Helper
}

func NewTagsService(uc *biz.TagsUsecase, logger log.Logger) *TagsService {
	return &TagsService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

// TODO if we need process ctx timeout?

func toBizTag(t *pb.Tag) (*biz.Tag, error) {
	if t == nil {
		return nil, nil
	}
	return &biz.Tag{
		Id:    t.Id,
		Key:   t.Key,
		Value: t.Value,
	}, nil
}

func toBizTags(ts []*pb.Tag) ([]*biz.Tag, error) {
	tags := make([]*biz.Tag, len(ts))
	for i, t := range ts {
		biz_tag, err := toBizTag(t)
		if err != nil {
			return nil, err
		}
		tags[i] = biz_tag
	}
	return tags, nil
}

func (s *TagsService) CreateTags(ctx context.Context, req *pb.CreateTagsRequest) (*pb.CreateTagsReply, error) {

	if req == nil {
		return nil, ErrRequestNil
	}

	tags, err := toBizTags(req.Tags)
	if err == nil {
		err = s.usecase.CreateTags(ctx, tags)
	}

	reply := &pb.CreateTagsReply{
		Action:  "CreateTags",
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

func (s *TagsService) UpdateTags(ctx context.Context, req *pb.UpdateTagsRequest) (*pb.UpdateTagsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}

	tags, err := toBizTags(req.Tags)
	if err == nil {
		err = s.usecase.UpdateTags(ctx, tags)
	}

	reply := &pb.UpdateTagsReply{
		Action:  "UpdateTags",
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

func (s *TagsService) DeleteTags(ctx context.Context, req *pb.DeleteTagsRequest) (*pb.DeleteTagsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
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
		return reply, nil
	}

	return reply, nil
}

func (s *TagsService) GetTags(ctx context.Context, req *pb.GetTagsRequest) (*pb.GetTagsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
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
	return reply, nil
}

func (s *TagsService) ListTags(ctx context.Context, req *pb.ListTagsRequest) (*pb.ListTagsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	var filter = biz.DefaultTagsFilter()
	if req.Filter != nil {
		filter = &biz.ListTagsFilter{
			Ids:  req.Filter.Ids,
			Keys: req.Filter.Keys,
			Kvs:  req.Filter.Kvs,
		}
		if req.Filter.PageSize > 0 {
			filter.PageSize = req.Filter.PageSize
		}
		if req.Filter.Page > 0 {
			filter.Page = req.Filter.Page
		}
	}

	s.log.Infow("function", "listTags", "filter", filter)

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

	return reply, nil
}
