package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type AppTagsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewAppTagsRepoGorm(data *DataGorm, logger log.Logger) (repo.AppTagsRepo, error) {
	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.DB, &repo.AppTag{}, repo.AppTagTable); err != nil {
		return nil, err
	}
	return &AppTagsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *AppTagsRepoGorm) CreateAppTags(ctx context.Context,
	tx repo.TX,
	ats []*repo.AppTag) error {

	r := d.data.WithTX(tx).WithContext(ctx).Create(ats)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func (d *AppTagsRepoGorm) UpdateAppTags(ctx context.Context,
	tx repo.TX,
	apps []*repo.AppTag) error {

	r := d.data.WithTX(tx).WithContext(ctx).Save(apps)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func (d *AppTagsRepoGorm) DeleteAppTags(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {

	r := d.data.WithTX(tx).WithContext(ctx).Delete(&repo.AppTag{}, ids)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func (d *AppTagsRepoGorm) ListAppTags(ctx context.Context,
	tx repo.TX,
	filter *repo.AppTagsFilter) ([]*repo.AppTag, error) {

	query := d.data.WithTX(tx).WithContext(ctx)
	if len(filter.Ids) > 0 {
		query = query.Where("id in (?)", filter.Ids)
	}
	if len(filter.AppIds) > 0 {
		query = query.Where("app_id in (?)", filter.AppIds)
	}
	if len(filter.TagIds) > 0 {
		query = query.Where("tag_id in (?)", filter.TagIds)
	}
	if len(filter.KVs) > 0 {
		s_q, kvs := buildOrKV("app_id", "tag_id", filter.KVs)
		query = query.Where(s_q, kvs...)
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := int(filter.PageSize * (filter.Page - 1))
		query = query.Offset(offset).Limit(int(filter.PageSize))
	}

	var db_ats []*repo.AppTag

	r := query.Find(&db_ats)
	if r.Error != nil {
		return nil, r.Error
	}
	return db_ats, nil
}

func (d *AppTagsRepoGorm) DeleteAppTagsByAppId(ctx context.Context,
	tx repo.TX,
	appids []uint32) error {

	r := d.data.WithTX(tx).WithContext(ctx).Where("app_id in (?)", appids).Delete(&repo.AppTag{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func (d *AppTagsRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	var condition string
	switch need {
	case repo.RequireApp:
		condition = "app_id in (?)"
	case repo.RequireTag:
		condition = "tag_id in (?)"
	default:
		return 0, nil
	}

	var count int64
	r := d.data.WithTX(tx).WithContext(ctx).Model(&repo.AppTag{}).
		Where(condition, ids).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}

	return count, nil
}
