package service

import (
	"context"
	"fmt"

	pb "appix/api/appix/v1"

	"github.com/go-kratos/kratos/v2/log"

	biz "appix/internal/biz"
)

type EnvsService struct {
	pb.UnimplementedEnvsServer
	usecase *biz.EnvsUsecase
	log     *log.Helper
}

func NewEnvsService(uc *biz.EnvsUsecase, logger log.Logger) *EnvsService {
	return &EnvsService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

func toBizEnv(env *pb.Env) (*biz.Env, error) {
	if env == nil {
		return nil, nil
	}
	return &biz.Env{
		Description: env.Description,
		Id:          env.Id,
		Name:        env.Name,
	}, nil
}

func toBizEnvs(envs []*pb.Env) ([]*biz.Env, error) {
	_bizenvs := make([]*biz.Env, len(envs))
	var err error
	for i, e := range envs {
		if _bizenvs[i], err = toBizEnv(e); err != nil {
			return nil, err
		}
	}
	return _bizenvs, nil
}

func (s *EnvsService) CreateEnvs(ctx context.Context, req *pb.CreateEnvsRequest) (*pb.CreateEnvsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	_bizenvs, err := toBizEnvs(req.Envs)

	if err == nil {
		err = s.usecase.CreateEnvs(ctx, _bizenvs)
	}

	reply := &pb.CreateEnvsReply{
		Action:  "CreateEnvs",
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

func (s *EnvsService) UpdateEnvs(ctx context.Context, req *pb.UpdateEnvsRequest) (*pb.UpdateEnvsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	_bizenvs, err := toBizEnvs(req.Envs)
	if err == nil {
		err = s.usecase.UpdateEnvs(ctx, _bizenvs)
	}
	reply := &pb.UpdateEnvsReply{
		Action:  "UpdateEnvs",
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

func (s *EnvsService) DeleteEnvs(ctx context.Context, req *pb.DeleteEnvsRequest) (*pb.DeleteEnvsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	err := s.usecase.DeleteEnvs(ctx, req.Ids)
	reply := &pb.DeleteEnvsReply{
		Action:  "DeleteEnvs",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
	}
	return reply, nil
}

func (s *EnvsService) GetEnvs(ctx context.Context, req *pb.GetEnvsRequest) (*pb.GetEnvsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	env, err := s.usecase.GetEnvs(ctx, req.Id)
	reply := &pb.GetEnvsReply{
		Action:  "GetEnvs",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Env = toPbEnv(env)

		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, nil
}

func (s *EnvsService) ListEnvs(ctx context.Context, req *pb.ListEnvsRequest) (*pb.ListEnvsReply, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	var filter = biz.DefaultEnvFilter()
	if req.Filter != nil {
		filter = &biz.ListEnvsFilter{
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

	envs, err := s.usecase.ListEnvs(ctx, filter)
	reply := &pb.ListEnvsReply{
		Action:  "ListEnvs",
		Code:    0,
		Message: "success",
	}
	if err == nil {
		reply.Envs = toPbEnvs(envs)
		return reply, nil
	}
	reply.Code = 1
	reply.Message = err.Error()
	return reply, nil
}

func toPbEnv(env *biz.Env) *pb.Env {
	if env == nil {
		return nil
	}
	return &pb.Env{
		Id:          env.Id,
		Name:        env.Name,
		Description: env.Description,
	}
}

func toPbEnvs(envs []*biz.Env) []*pb.Env {
	if envs == nil {
		return nil
	}
	var res []*pb.Env
	for _, e := range envs {
		if e != nil {
			res = append(res, toPbEnv(e))
		}
	}
	return res
}
