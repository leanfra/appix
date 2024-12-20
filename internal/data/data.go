package data

import (
	"appix/internal/biz"
	"appix/internal/conf"
	"errors"
	"fmt"
	"strings"

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
	NewClustersRepoImpl,
	NewDatacentersRepoImpl,
	NewHostgroupsRepoImpl,
	NewApplicationsRepoImpl,
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
var ErrMissingRecords = errors.New("missing records")
var ErrMissingTags = errors.New("missing Tag")
var ErrMissingTeams = errors.New("missing Team")
var ErrMissingProducts = errors.New("missing Product")
var ErrMissingFeatures = errors.New("missing Feature")
var ErrMissingEnvs = errors.New("missing Env")
var ErrMissingDatacenters = errors.New("missing Datacenter")
var ErrMissingClusters = errors.New("missing Cluster")

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
func buildOrKV(kname string, vname string, kvstr []string) (string, []string) {
	var kv []string
	var builder strings.Builder
	for i := 0; i < len(kvstr); i++ {
		if i > 0 {
			builder.WriteString(" OR ")
		}
		_kvs := strings.Split(kvstr[i], biz.FilterKVSplit)
		if len(_kvs) == 2 {
			kv = append(kv, _kvs...)
			builder.WriteString(fmt.Sprintf("( %s = ? AND  %s = ? )", kname, vname))
		}
	}
	return builder.String(), kv
}

func DedupSliceUint32(s []uint32) []uint32 {
	if s == nil {
		return nil
	}
	var result []uint32
	m := make(map[uint32]struct{})
	for i := 0; i < len(s); i++ {
		if _, exists := m[s[i]]; !exists {
			m[s[i]] = struct{}{}
			result = append(result, s[i])

		}
	}
	return result
}

// DiffUint32 return (s1 - s2)
func DiffUint32(s1 []uint32, s2 []uint32) []uint32 {
	result := []uint32{}
	set2Map := make(map[uint32]bool)
	for _, v := range s2 {
		set2Map[v] = true
	}
	for _, v := range s1 {
		if _, ok := set2Map[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

func existsRecords(db *gorm.DB, mod interface{}, ids []uint32) error {
	var count int64
	if r := db.Model(mod).Where("id in (?)", ids).Count(&count); r.Error != nil {
		return r.Error
	}
	if count != int64(len(ids)) {
		return ErrNoRowsAffected
	}
	return nil
}
