package data

import (
	"appix/internal/conf"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewFeaturesRepoImpl,
	NewTagsRepoImpl,
	NewTeamsRepoImpl,
	NewProductsRepoImpl,
	NewEnvsRepoImpl,
)

// Data .
type Data struct {
	db *gorm.DB
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	dsn := c.GetDatabase().GetSource()

	driver := c.GetDatabase().GetDriver()
	var _db *gorm.DB
	var err error
	if len(dsn) > 0 {
		if driver == "mysql" {
			// 连接 MySQL 数据库
			_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		} else if driver == "sqlite" {
			// 连接 SQLite 数据库
			_db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		} else {
			return nil, cleanup, ErrUnsupportedDatabaseDriver
		}
	} else {
		return nil, cleanup, ErrEmptyDatabase
	}

	if err != nil {
		return nil, cleanup, err
	}

	return &Data{
		db: _db,
	}, cleanup, nil
}

var ErrUnsupportedDatabaseDriver = errors.New("unsupportedDatabaseDriver")
var ErrEmptyDatabase = errors.New("emptyDatabaseSource")
var ErrNoRowsAffected = errors.New("noRowsAffected")

func validateData(data *Data) error {
	if data == nil || data.db == nil {
		return ErrEmptyDatabase
	}
	return nil
}

func initTable(db *gorm.DB, model interface{}, table string) error {
	m := db.Migrator()

	if !m.HasTable(table) {
		return db.AutoMigrate(model)
	}
	return nil
}
