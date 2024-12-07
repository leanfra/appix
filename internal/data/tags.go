package data

import (
	"appix/internal/biz"
	"context"

	"gorm.io/gorm"
)

type TagsRepoImpl struct {
	data *Data
}

func NewTagsRepoImpl(data *Data) (biz.TagsRepo, error) {

	if data == nil {
		return nil, ErrEmptyDatabase
	}

	data.db.AutoMigrate(&Tag{})

	return &TagsRepoImpl{
		data: data,
	}, nil
}

// CreateTags is
func (d *TagsRepoImpl) CreateTags(ctx context.Context, tags []biz.Tag) error {

	db_tags, err := NewTags(tags)
	if err != nil {
		return err
	}
	r := d.data.db.Create(db_tags)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// UpdateTags is
func (d *TagsRepoImpl) UpdateTags(ctx context.Context, tags []biz.Tag) error {
	db_tags, err := NewTags(tags)
	if err != nil {
		return err
	}
	r := d.data.db.Save(db_tags)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// DeleteTags is
func (d *TagsRepoImpl) DeleteTags(ctx context.Context, ids []int64) error {

	r := d.data.db.Where("id in (?)", ids).Delete(&Tag{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// GetTags is
func (d *TagsRepoImpl) GetTags(ctx context.Context, id int64) (*biz.Tag, error) {

	tag := &Tag{}
	r := d.data.db.First(tag, id)
	if r.Error != nil {
		return nil, r.Error
	}
	return NewBizTag(tag)
}

// ListTags is
func (d *TagsRepoImpl) ListTags(ctx context.Context,
	filter *biz.ListTagsFilter) ([]biz.Tag, error) {

	tags := []Tag{}

	var r *gorm.DB
	query := d.data.db
	if filter != nil {
		var offset int
		if filter.Page > 0 && filter.PageSize > 0 {
			offset = int((filter.Page - 1) * filter.PageSize)
		}

		for _, pair := range filter.Filters {
			query = query.Where("key =? AND value =?", pair.Key, pair.Value)
		}
		r = query.Offset(offset).Limit(int(filter.PageSize)).Find(&tags)
	} else {
		r = query.Find(&tags)
	}

	if r.Error != nil {
		return nil, r.Error
	}
	return NewBizTags(tags)
}
