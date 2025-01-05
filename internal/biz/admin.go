package biz

import (
	"appix/internal/conf"
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminUsecase struct {
	repo repo.AdminRepo
	log  *log.Helper
	conf *conf.Admin
}

func NewAdminUsecase(repo repo.AdminRepo, logger log.Logger) *AdminUsecase {
	return &AdminUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *AdminUsecase) validate(isNew bool, users []*User) error {
	for _, user := range users {
		if err := user.Validate(&UserOptions{
			IsNew:                isNew,
			StrictPasswordPolicy: s.conf.StrictPasswordPolicy,
		}); err != nil {
			return err
		}
	}
	return nil
}

// CreateUsers is
func (s *AdminUsecase) CreateUsers(ctx context.Context, users []*User) error {
	if err := s.validate(true, users); err != nil {
		return err
	}

	repoUsers := ToRepoUsers(users)
	return s.repo.CreateUsers(ctx, repoUsers)
}

// UpdateUsers is
func (s *AdminUsecase) UpdateUsers(ctx context.Context, users []*User) error {
	if err := s.validate(false, users); err != nil {
		return err
	}
	repoUsers := ToRepoUsers(users)
	return s.repo.UpdateUsers(ctx, repoUsers)
}

// DeleteUsers is
func (s *AdminUsecase) DeleteUsers(ctx context.Context, tx repo.TX, ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return s.repo.DeleteUsers(ctx, tx, ids)
}

// GetUsers is
func (s *AdminUsecase) GetUsers(ctx context.Context, id uint32) (*User, error) {
	user, err := s.repo.GetUsers(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToBizUser(user), nil
}

// ListUsers is
func (s *AdminUsecase) ListUsers(ctx context.Context, tx repo.TX, filter *ListUsersFilter) ([]*User, error) {
	repoUsers, err := s.repo.ListUsers(ctx, tx, ToDBUsersFilter(filter))
	if err != nil {
		return nil, err
	}
	return ToBizUsers(repoUsers), nil
}

// Login is
func (s *AdminUsecase) Login(ctx context.Context, username, password string) (*User, error) {
	user, err := s.repo.Login(ctx, username, password)
	if err != nil {
		return nil, err
	}
	return ToBizUser(user), nil
}

// Logout is
func (s *AdminUsecase) Logout(ctx context.Context, id uint32) error {
	return s.repo.Logout(ctx, id)
}
