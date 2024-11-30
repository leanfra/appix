package data

import (
	"context"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
    "gorm.io/gorm"

	//  TODO: modify project name
	// biz "appix/internal/biz"
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
            return nil, errors.New("Unsupported database DSN format")
        }
    } else {
        return nil, errors.New("DSN is not provided")
    }

	if err != nil {
		return nil, err
	}
	return &TagsRepoImpl{
		db: _db,
	}, nil
}


// CreateTags is
func (d *TagsRepoImpl) CreateTags(ctx context.Context) error {
	// TODO database operations

	return nil
}
// UpdateTags is
func (d *TagsRepoImpl) UpdateTags(ctx context.Context) error {
	// TODO database operations

	return nil
}
// DeleteTags is
func (d *TagsRepoImpl) DeleteTags(ctx context.Context) error {
	// TODO database operations

	return nil
}
// GetTags is
func (d *TagsRepoImpl) GetTags(ctx context.Context) error {
	// TODO database operations

	return nil
}
// ListTags is
func (d *TagsRepoImpl) ListTags(ctx context.Context) error {
	// TODO database operations

	return nil
}
