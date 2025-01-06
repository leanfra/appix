package sqldb

import (
	"appix/internal/data/repo"
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	//  TODO: modify project name
	// biz "appix/internal/biz"
)

type AdminRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewAdminRepoGorm(data *DataGorm, logger log.Logger) (repo.AdminRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.DB, &repo.User{}, repo.UserTable); err != nil {
		return nil, err
	}
	return &AdminRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateUsers is
func (d *AdminRepoGorm) CreateUsers(ctx context.Context, users []*repo.User) error {
	if len(users) == 0 {
		return nil
	}
	return d.data.WithTX(nil).WithContext(ctx).Create(users).Error

}

// UpdateUsers is
func (d *AdminRepoGorm) UpdateUsers(ctx context.Context, users []*repo.User) error {
	if len(users) == 0 {
		return nil
	}
	return d.data.WithTX(nil).WithContext(ctx).Updates(users).Error
}

// DeleteUsers is
func (d *AdminRepoGorm) DeleteUsers(ctx context.Context, tx repo.TX, ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Delete(&repo.User{}, ids).Error
}

// GetUsers is
func (d *AdminRepoGorm) GetUsers(ctx context.Context, id uint32) (*repo.User, error) {
	var user repo.User
	if err := d.data.WithTX(nil).WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ListUsers is
func (d *AdminRepoGorm) ListUsers(ctx context.Context, tx repo.TX, filter *repo.UsersFilter) ([]*repo.User, error) {
	query := d.data.WithTX(tx).WithContext(ctx)
	if len(filter.UserName) > 0 {
		query = query.Where("user_name IN ?", filter.UserName)
	}
	if len(filter.Email) > 0 {
		query = query.Where("email IN ?", filter.Email)
	}
	if len(filter.Phone) > 0 {
		query = query.Where("phone IN ?", filter.Phone)
	}
	users := make([]*repo.User, 0)
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}

// Login is
func (d *AdminRepoGorm) Login(ctx context.Context, username string, password string) (*repo.User, error) {
	var user repo.User
	if err := d.data.WithTX(nil).WithContext(ctx).Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("password is incorrect")
	}

	return &user, nil
}

// Logout is
func (d *AdminRepoGorm) Logout(ctx context.Context, id uint32) error {
	return nil
}
