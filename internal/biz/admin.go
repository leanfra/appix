package biz

import (
	"appix/internal/conf"
	"appix/internal/data/repo"
	"appix/internal/middleware"
	"context"
	"errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminUsecase struct {
	adminRepo repo.AdminRepo
	tokenRepo repo.TokenRepo
	authzRepo repo.AuthzRepo
	log       *log.Helper
	conf      *conf.Admin
}

func NewAdminUsecase(
	conf *conf.Admin,
	adminRepo repo.AdminRepo,
	tokenRepo repo.TokenRepo,
	authzRepo repo.AuthzRepo,
	logger log.Logger,
) *AdminUsecase {

	authzRepo.CreateGroup(context.Background(), nil, &repo.Group{
		User: conf.AdminUser,
		Role: "admin",
	})

	return &AdminUsecase{
		adminRepo: adminRepo,
		tokenRepo: tokenRepo,
		authzRepo: authzRepo,
		log:       log.NewHelper(logger),
		conf:      conf,
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
	return s.adminRepo.CreateUsers(ctx, repoUsers)
}

// UpdateUsers is
func (s *AdminUsecase) UpdateUsers(ctx context.Context, users []*User) error {
	if err := s.validate(false, users); err != nil {
		return err
	}
	repoUsers := ToRepoUsers(users)
	return s.adminRepo.UpdateUsers(ctx, repoUsers)
}

// DeleteUsers is
func (s *AdminUsecase) DeleteUsers(ctx context.Context, tx repo.TX, ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return s.adminRepo.DeleteUsers(ctx, tx, ids)
}

// GetUsers is
func (s *AdminUsecase) GetUsers(ctx context.Context, id uint32) (*User, error) {
	user, err := s.adminRepo.GetUsers(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToBizUser(user), nil
}

// ListUsers is
func (s *AdminUsecase) ListUsers(ctx context.Context, tx repo.TX, filter *ListUsersFilter) ([]*User, error) {
	repoUsers, err := s.adminRepo.ListUsers(ctx, tx, ToDBUsersFilter(filter))
	if err != nil {
		return nil, err
	}
	return ToBizUsers(repoUsers), nil
}

// Login is
func (s *AdminUsecase) Login(ctx context.Context, username, password string) (*User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password is empty")
	}
	var user *repo.User
	var err error
	if username == s.conf.AdminUser && password == s.conf.AdminPassword {
		user = &repo.User{
			Id:       1,
			UserName: s.conf.AdminUser,
		}
	} else {
		user, err = s.adminRepo.Login(ctx, username, password)
		if err != nil {
			return nil, err
		}
	}

	_idstr := strconv.Itoa(int(user.Id))

	token, err := s.tokenRepo.CreateToken(ctx, repo.TokenClaims{
		"id":   _idstr,
		"name": user.UserName,
	})
	if err != nil {
		return nil, err
	}

	bizUser := ToBizUser(user)
	bizUser.Token = token

	return bizUser, nil
}

// Logout is
func (s *AdminUsecase) Logout(ctx context.Context, id uint32) error {
	tokenStr := ctx.Value(middleware.UserTokenKey).(string)
	claims, err := s.tokenRepo.ValidateToken(ctx, tokenStr)
	if err != nil {
		return errors.Join(errors.New("logout failed"), err)
	}
	// Convert to uint32
	_idClaim, ok := claims["id"].(string)
	if !ok {
		return errors.Join(errors.New("invalid token claims 1"), err)
	}
	_id := strconv.Itoa(int(id))
	if _id != _idClaim {
		return errors.Join(errors.New("invalid token claims 3"), err)
	}
	err = s.tokenRepo.DeleteToken(ctx, tokenStr)
	if err != nil {
		return errors.Join(errors.New("logout failed"), err)
	}

	return s.adminRepo.Logout(ctx, id)
}
