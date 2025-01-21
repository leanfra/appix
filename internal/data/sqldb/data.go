package sqldb

import (
	"appix/internal/conf"
	"appix/internal/data/repo"
	"errors"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DataGorm .
type DataGorm struct {
	DB     *gorm.DB
	driver string
}

// NewDataGorm .
func NewDataGorm(c *conf.Data, logger log.Logger) (*DataGorm, func(), error) {
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

	return &DataGorm{
		DB:     _db,
		driver: driver,
	}, cleanup, nil
}

func (d *DataGorm) WithTX(tx repo.TX) *gorm.DB {
	if tx == nil {
		return d.DB
	}
	return tx.GetDB().(*gorm.DB)
}

// ColumnExistsForSQLite 用于检测SQLite数据库中指定表是否存在某一列
func (d *DataGorm) ColumnExists(tableName string, columnName string) (bool, error) {
	var count int
	if d.driver == "sqlite" {
		// 在SQLite中，通过查询sqlite_master表获取表结构相关信息来判断列是否存在
		query := fmt.Sprintf("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='%s' AND sql LIKE '%%%s%%'", tableName, columnName)
		d.DB.Raw(query).Scan(&count)
	} else if d.driver == "mysql" {
		query := fmt.Sprintf("SELECT COUNT(*) FROM information_schema.COLUMNS WHERE table_name = '%s' AND column_name = '%s'", tableName, columnName)
		d.DB.Raw(query).Scan(&count)
	} else {
		return false, ErrUnsupportedDatabaseDriver
	}
	return count > 0, nil
}

const FilterKVSplit = ":"

var ErrUnsupportedDatabaseDriver = errors.New("unsupportedDatabaseDriver")
var ErrEmptyDatabase = errors.New("emptyDatabaseSource")
var ErrNoRowsAffected = errors.New("noRowsAffected")
var ErrMissingRecords = errors.New("missing records")
var ErrMissingTags = errors.New("missing Tag")
var ErrMissingTeams = errors.New("missing Team")
var ErrMissingProducts = errors.New("missing Product")
var ErrMissingFeatures = errors.New("missing Feature")
var ErrMissingEnvs = errors.New("missing Env")
var ErrMissingDatacenters = errors.New("missing Datacenter")
var ErrMissingClusters = errors.New("missing Cluster")
var ErrRequireFeatureIds = errors.New("missing FeatureIds")

func validateData(data *DataGorm) error {
	if data == nil || data.DB == nil {
		return ErrEmptyDatabase
	}
	return nil
}

func initTable(db *gorm.DB, model interface{}, table string) error {
	m := db.Migrator()

	if !m.HasTable(table) {
		log.Warnf("missing table %s", table)
		return db.AutoMigrate(model)
	}

	if !m.HasColumn(&repo.Application{}, "owner_id") {
		m.AddColumn(&repo.Application{}, "owner_id")
	}

	if !m.HasColumn(&repo.Team{}, "leader_id") {
		m.AddColumn(&repo.Team{}, "leader_id")
	}

	log.Infof("exists table %s", table)
	return nil
}

// buildOrLike build or conditions.
//
//	return: "key LIKE ? OR key LIKE ? ..."
func buildOrLike(key string, count int) string {
	var builder strings.Builder
	for i := 0; i < count; i++ {
		if i > 0 {
			builder.WriteString(" OR ")
		}
		builder.WriteString(fmt.Sprintf("%s LIKE ?", key))
	}
	return builder.String()
}

// buildOrKV build or conditions with key and value.
//
//	return: "(k=? AND v=?) OR (k=? AND v=?) ... "
//	return: [k1, v1, k2, v2, ...]
func buildOrKV(kname string, vname string, kvstr []string) (string, []interface{}) {
	var kv []interface{}
	var builder strings.Builder
	for i := 0; i < len(kvstr); i++ {
		if i > 0 {
			builder.WriteString(" OR ")
		}
		_kvs := strings.Split(kvstr[i], FilterKVSplit)
		if len(_kvs) == 2 {
			kv = append(kv, _kvs[0], _kvs[1])
			builder.WriteString(fmt.Sprintf("( %s = ? AND  %s = ? )", kname, vname))
		}
	}
	return builder.String(), kv
}
