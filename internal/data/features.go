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

type FeaturesRepoImpl struct {
	db *gorm.DB
}

func NewFeaturesRepoImpl(dsn string) (*FeaturesRepoImpl, error) {

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
	return &FeaturesRepoImpl{
		db: _db,
	}, nil
}


// CreateFeatures is
func (d *FeaturesRepoImpl) CreateFeatures(ctx context.Context) error {
	// TODO database operations

	return nil
}
// UpdateFeatures is
func (d *FeaturesRepoImpl) UpdateFeatures(ctx context.Context) error {
	// TODO database operations

	return nil
}
// DeleteFeatures is
func (d *FeaturesRepoImpl) DeleteFeatures(ctx context.Context) error {
	// TODO database operations

	return nil
}
// GetFeatures is
func (d *FeaturesRepoImpl) GetFeatures(ctx context.Context) error {
	// TODO database operations

	return nil
}
// ListFeatures is
func (d *FeaturesRepoImpl) ListFeatures(ctx context.Context) error {
	// TODO database operations

	return nil
}
