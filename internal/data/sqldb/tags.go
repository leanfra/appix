package sqldb

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type TagsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewTagsRepoGorm(data *DataGorm, logger log.Logger) (repo.TagsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.DB, &repo.Tag{}, repo.TagTable); err != nil {
		return nil, err
	}

	// add description column if not exists
	exists, err := data.ColumnExists(repo.TagTable, "description")
	if err != nil {
		return nil, err
	}
	if !exists {
		data.DB.AutoMigrate(&repo.Tag{})
	}

	return &TagsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

// CreateTags is
func (d *TagsRepoGorm) CreateTags(ctx context.Context, tags []*repo.Tag) error {

	r := d.data.DB.WithContext(ctx).Create(tags)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// UpdateTags is
func (d *TagsRepoGorm) UpdateTags(ctx context.Context, tags []*repo.Tag) error {
	r := d.data.DB.WithContext(ctx).Save(tags)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteTags is
func (d *TagsRepoGorm) DeleteTags(ctx context.Context,
	tx repo.TX, ids []uint32) error {

	r := d.data.DB.WithContext(ctx).Where("id in (?)", ids).Delete(&repo.Tag{})
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != int64(len(ids)) {
		return fmt.Errorf("delete not equal expected. want %d. affected %d", len(ids), r.RowsAffected)
	}
	return nil
}

// GetTags is
func (d *TagsRepoGorm) GetTags(ctx context.Context, id uint32) (*repo.Tag, error) {

	tag := &repo.Tag{}
	r := d.data.DB.WithContext(ctx).First(tag, id)
	if r.Error != nil {
		return nil, r.Error
	}
	return tag, nil
}

// ListTags is
func (d *TagsRepoGorm) ListTags(ctx context.Context,
	tx repo.TX,
	filter *repo.TagsFilter) ([]*repo.Tag, error) {

	tags := []*repo.Tag{}

	var r *gorm.DB
	query := d.data.WithTX(tx).WithContext(ctx)
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}
		if len(filter.Ids) > 0 {
			query = query.Where("id in (?)", filter.Ids)
		}
		if len(filter.Keys) > 0 {
			keyConditions := buildOrLike("key", len(filter.Keys))
			params := make([]interface{}, len(filter.Keys))
			for i, v := range filter.Keys {
				params[i] = "%" + v + "%"
			}
			query = query.Where(keyConditions, params...)
		}
		if len(filter.Kvs) > 0 {
			kvConditions, kvs := buildOrKV("key", "value", filter.Kvs)
			query = query.Where(kvConditions, kvs...)
		}
	}
	r = query.Find(&tags)

	if r.Error != nil {
		return nil, r.Error
	}
	return tags, nil
}

func (d *TagsRepoGorm) CountTags(ctx context.Context,
	tx repo.TX,
	filter repo.CountFilter) (int64, error) {

	var count int64
	query := d.data.WithTX(tx).WithContext(ctx)
	if filter != nil {
		if len(filter.GetIds()) > 0 {
			query = query.Where("id in (?)", filter.GetIds())
		}
	}
	r := query.Model(&repo.Tag{}).Count(&count)
	if r.Error != nil {
		return 0, r.Error
	}
	return count, nil
}

func (d *TagsRepoGorm) CountRequire(ctx context.Context,
	tx repo.TX,
	need repo.RequireType,
	ids []uint32) (int64, error) {

	if len(ids) == 0 {
		return 0, repo.ErrorRequireIds
	}

	// require nothing
	return 0, nil

}
