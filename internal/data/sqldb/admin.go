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
func (d *AdminRepoGorm) CreateUsers(ctx context.Context, tx repo.TX, users []*repo.User) error {
	if len(users) == 0 {
		return nil
	}
	return d.data.WithTX(tx).WithContext(ctx).Create(users).Error

}

// UpdateUsers is
func (d *AdminRepoGorm) UpdateUsers(ctx context.Context, tx repo.TX, users []*repo.User) error {
	if len(users) == 0 {
		return nil
	}

	for _, user := range users {
		r := d.data.WithTX(tx).WithContext(ctx).Where("id=?", user.Id).Select("Email", "Phone", "Password").Updates(user)
		if r.Error != nil {
			return r.Error
		}
		if r.RowsAffected != int64(len(users)) {
			return errors.New("update failed")
		}
	}
	return nil
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

// Logout is
func (d *AdminRepoGorm) Logout(ctx context.Context, id uint32) error {
	return nil
}
