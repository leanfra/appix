package data

import (
	"appix/internal/biz"
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TagsRepoImpl struct {
	db *gorm.DB
}

func NewTagsRepoImpl(dsn string) (*TagsRepoImpl, error) {

	var _db *gorm.DB
	var err error
	if len(dsn) > 0 {
		if dsn[:5] == "mysql" {
			// 连接 MySQL 数据库
			_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		} else if dsn[:6] == "sqlite" {
			// 连接 SQLite 数据库
			_db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		} else {
			return nil, fmt.Errorf("UnsupportedDatabaseDSN")
		}
	} else {
		return nil, fmt.Errorf("EmptyDSN")
	}

	if err != nil {
		return nil, err
	}
	return &TagsRepoImpl{
		db: _db,
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
func (d *TagsRepoImpl) GetTags(ctx context.Context, id string) (biz.Tag, error) {
	// TODO database operations

	return biz.Tag{}, nil
}

// ListTags is
func (d *TagsRepoImpl) ListTags(ctx context.Context,
	filter *biz.ListTagsFilter) ([]biz.Tag, error) {
	// TODO database operations

	return nil, nil
}
