package sqldb

import (
	"opspillar/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type ApplicationsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewApplicationsRepoGorm(data *DataGorm, logger log.Logger) (repo.ApplicationsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.DB, &repo.Application{}, repo.ApplicationTable); err != nil {
		return nil, err
	}

	return &ApplicationsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateApplications is
func (d *ApplicationsRepoGorm) CreateApplications(
	ctx context.Context,
	tx repo.TX,
	apps []*repo.Application) error {

	r := d.data.WithTX(tx).WithContext(ctx).
		Model(&repo.Application{}).
		Create(apps)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// UpdateApplications is
func (d *ApplicationsRepoGorm) UpdateApplications(
	ctx context.Context,
	tx repo.TX,
	apps []*repo.Application) error {

	r := d.data.WithTX(tx).WithContext(ctx).
		Model(&repo.Application{}).
		Save(apps)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteApplications is
func (d *ApplicationsRepoGorm) DeleteApplications(
	ctx context.Context,
	tx repo.TX,
	ids []uint32) error {

	r := d.data.WithTX(tx).WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Application{})
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != int64(len(ids)) {
		return fmt.Errorf("delete failed. rows affected not equal wanted. affected %d. want %d", r.RowsAffected, len(ids))
	}

	return nil
}

// GetApplications is
func (d *ApplicationsRepoGorm) GetApplications(
	ctx context.Context, id uint32) (*repo.Application, error) {

	app := &repo.Application{}
	r := d.data.DB.WithContext(ctx).Where("id = ?", id).First(app)
	if r.Error != nil {
		return nil, r.Error
	}

	return app, nil
}

// ListApplications is
func (d *ApplicationsRepoGorm) ListApplications(ctx context.Context,
	tx repo.TX,
	filter *repo.ApplicationsFilter) ([]*repo.Application, error) {

	query := d.data.WithTX(tx).WithContext(ctx).Model(&repo.Application{})
	if filter != nil {
		if len(filter.Ids) > 0 {
			query = query.Where("id in (?)", filter.Ids)
		}
		if len(filter.Names) > 0 {
			s_q := buildOrLike("name", len(filter.Names))
			params := make([]interface{}, len(filter.Names))
			for i, v := range filter.Names {
				params[i] = "%" + v + "%"
			}
			query = query.Where(s_q, params...)
		}
		if len(filter.ProductsId) > 0 {
			query = query.Where("product_id in (?)", filter.ProductsId)
		}
		if len(filter.TeamsId) > 0 {
			query = query.Where("team_id in (?)", filter.TeamsId)
		}
		if filter.IsStateful != "" {
			if filter.IsStateful == repo.IsStatefulTrue {
				query = query.Where("is_stateful = ?", true)
			} else if filter.IsStateful == repo.IsStatefulFalse {
				query = query.Where("is_stateful = ?", false)
			}
		}
		if filter.Page > 0 && filter.PageSize > 0 {
			offset := int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}
	}
	var apps []*repo.Application
	if err := query.Find(&apps).Error; err != nil {
		return nil, err
	}

	return apps, nil
}

func (d *ApplicationsRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	var condition string
	switch need {
	case repo.RequireCluster:
		condition = "cluster_id in (?)"
	case repo.RequireDatacenter:
		condition = "datacenter_id in (?)"
	case repo.RequireProduct:
		condition = "product_id in (?)"
	case repo.RequireTeam:
		condition = "team_id in (?)"
	case repo.RequireUser:
		condition = "owner_id in (?)"
	default:
		return 0, repo.ErrorRequireIds
	}

	var count int64
	r := d.data.WithTX(tx).WithContext(ctx).Model(&repo.Application{}).
		Where(condition, ids).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}

	return count, nil
}
