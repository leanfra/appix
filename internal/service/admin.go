package service

import (
	"context"

	pb "appix/api/appix/v1"

	"github.com/go-kratos/kratos/v2/log"

	//  TODO: modify project name
	biz "appix/internal/biz"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	usecase *biz.AdminUsecase
	log     *log.Helper
}

func NewAdminService(uc *biz.AdminUsecase, logger log.Logger) *AdminService {
	return &AdminService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

func toBizUsers(users []*pb.User) ([]*biz.User, error) {
	bizUsers := make([]*biz.User, 0, len(users))
	for _, user := range users {
		bizUser, err := toBizUser(user)
		if err != nil {
			return nil, err
		}
		bizUsers = append(bizUsers, bizUser)
	}
	return bizUsers, nil
}

func toBizUser(user *pb.User) (*biz.User, error) {
	return &biz.User{
		Id:       user.Id,
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Token:    user.Token,
	}, nil
}

func toPbUser(user *biz.User) *pb.User {
	return &pb.User{
		Id:       user.Id,
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Token:    user.Token,
	}
}

func toPbUsers(users []*biz.User) []*pb.User {
	pbUsers := make([]*pb.User, 0, len(users))
	for _, user := range users {
		pbUsers = append(pbUsers, toPbUser(user))
	}
	return pbUsers
}

func ToBizUsersFilter(filter *pb.ListUsersRequest) *biz.ListUsersFilter {
	return &biz.ListUsersFilter{
		Page:      filter.Page,
		PageSize:  filter.PageSize,
		UserNames: filter.UserNames,
		Emails:    filter.Emails,
		Phones:    filter.Phones,
		Ids:       filter.Ids,
	}
}

func (s *AdminService) CreateUsers(ctx context.Context, req *pb.CreateUsersRequest) (*pb.CreateUsersReply, error) {
	bizUsers, err := toBizUsers(req.Users)
	if err != nil {
		return nil, err
	}

	reply := &pb.CreateUsersReply{
		Action:  "CreateUsers",
		Code:    0,
		Message: "success",
	}
	err = s.usecase.CreateUsers(ctx, bizUsers)
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}

	return reply, nil
}
func (s *AdminService) UpdateUsers(ctx context.Context, req *pb.UpdateUsersRequest) (*pb.UpdateUsersReply, error) {
	bizUsers, err := toBizUsers(req.Users)
	if err != nil {
		return nil, err
	}

	reply := &pb.UpdateUsersReply{
		Action:  "UpdateUsers",
		Code:    0,
		Message: "success",
	}
	err = s.usecase.UpdateUsers(ctx, bizUsers)
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}

	return reply, nil
}
func (s *AdminService) DeleteUsers(ctx context.Context, req *pb.DeleteUsersRequest) (*pb.DeleteUsersReply, error) {
	reply := &pb.DeleteUsersReply{
		Action:  "DeleteUsers",
		Code:    0,
		Message: "success",
	}
	err := s.usecase.DeleteUsers(ctx, nil, req.Ids)
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}

	return reply, nil
}
func (s *AdminService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersReply, error) {

	reply := &pb.GetUsersReply{
		Action:  "GetUsers",
		Code:    0,
		Message: "success",
	}
	user, err := s.usecase.GetUsers(ctx, uint32(req.Id))
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	reply.Users = toPbUser(user)

	return reply, nil
}

func (s *AdminService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersReply, error) {
	reply := &pb.ListUsersReply{
		Action:  "ListUsers",
		Code:    0,
		Message: "success",
	}
	users, err := s.usecase.ListUsers(ctx, nil, ToBizUsersFilter(req))
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	reply.Users = toPbUsers(users)

	return reply, nil
}

func (s *AdminService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	reply := &pb.LoginReply{
		Action:  "Login",
		Code:    0,
		Message: "success",
	}
	user, err := s.usecase.Login(ctx, req.UserName, req.Password)
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	reply.User = toPbUser(user)

	return reply, nil
}
func (s *AdminService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	reply := &pb.LogoutReply{
		Action:  "Logout",
		Code:    0,
		Message: "success",
	}

	err := s.usecase.Logout(ctx, req.Id)
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}

	return reply, nil
}
