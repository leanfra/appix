package data

import (
	"appix/internal/biz"
	"context"
)

type TagsRepoImpl struct {
	data *Data
}

func NewTagsRepoImpl(data *Data) (biz.TagsRepo, error) {

	return &TagsRepoImpl{
		data: data,
	}, nil
}

// CreateTags is
func (d *TagsRepoImpl) CreateTags(ctx context.Context, tags []biz.Tag) error {
	// TODO database operations

	return nil
}

// UpdateTags is
func (d *TagsRepoImpl) UpdateTags(ctx context.Context, tags []biz.Tag) error {
	// TODO database operations

	return nil
}

// DeleteTags is
func (d *TagsRepoImpl) DeleteTags(ctx context.Context, ids []string) error {
	// TODO database operations

	return nil
}

// GetTags is
func (d *TagsRepoImpl) GetTags(ctx context.Context, id string) (*biz.Tag, error) {
	// TODO database operations

	return nil, nil
}

// ListTags is
func (d *TagsRepoImpl) ListTags(ctx context.Context,
	filter *biz.ListTagsFilter) ([]biz.Tag, error) {
	// TODO database operations

	return nil, nil
}
