package repo

import (
	"context"
	"errors"
)

type TX interface {
	// Create(model interface{}) error
	GetDB() interface{}
	Error(err error) bool
}

type TxManager interface {
	RunInTX(fn func(tx TX) error) error
}

type CountFilter interface {
	GetIds() []uint32
}

type RequireType string

const (
	RequireTeam       RequireType = "team"
	RequireTag        RequireType = "tag"
	RequireProduct    RequireType = "product"
	RequireHostgroup  RequireType = "hostgroup"
	RequireFeature    RequireType = "feature"
	RequireEnv        RequireType = "env"
	RequireDatacenter RequireType = "datacenter"
	RequireCluster    RequireType = "cluster"
	RequireApp        RequireType = "app"
)

var ErrorRequireIds = errors.New("invalid require ids")

type CountRequire interface {
	CountRequire(ctx context.Context, tx TX, need RequireType, ids []uint32) (int64, error)
}

type ApplicationsRepo interface {
	CountRequire
	CreateApplications(ctx context.Context, tx TX, apps []*Application) error
	UpdateApplications(ctx context.Context, tx TX, apps []*Application) error
	DeleteApplications(ctx context.Context, tx TX, ids []uint32) error
	GetApplications(ctx context.Context, id uint32) (*Application, error)
	ListApplications(ctx context.Context, tx TX, filter *ApplicationsFilter) ([]*Application, error)
	//CountApplications(ctx context.Context, tx TX, filter *ApplicationsFilter) (int64, error)
}

type AppTagsRepo interface {
	CountRequire
	CreateAppTags(ctx context.Context, tx TX, apps []*AppTag) error
	UpdateAppTags(ctx context.Context, tx TX, apps []*AppTag) error
	DeleteAppTags(ctx context.Context, tx TX, ids []uint32) error
	DeleteAppTagsByAppId(ctx context.Context, tx TX, appids []uint32) error
	ListAppTags(ctx context.Context, tx TX, filter *AppTagsFilter) ([]*AppTag, error)
	//CountAppTags(ctx context.Context, tx TX, filter *AppTagsFilter) (int64, error)
}

type AppFeaturesRepo interface {
	CountRequire
	CreateAppFeatures(ctx context.Context, tx TX, apps []*AppFeature) error
	UpdateAppFeatures(ctx context.Context, tx TX, apps []*AppFeature) error
	DeleteAppFeatures(ctx context.Context, tx TX, ids []uint32) error
	DeleteAppFeaturesByAppId(ctx context.Context, tx TX, appids []uint32) error
	ListAppFeatures(ctx context.Context, tx TX, filter *AppFeaturesFilter) ([]*AppFeature, error)
}

type AppHostgroupsRepo interface {
	CountRequire
	CreateAppHostgroups(ctx context.Context, tx TX, apps []*AppHostgroup) error
	UpdateAppHostgroups(ctx context.Context, tx TX, apps []*AppHostgroup) error
	DeleteAppHostgroups(ctx context.Context, tx TX, ids []uint32) error
	DeleteAppHostgroupsByAppId(ctx context.Context, tx TX, appids []uint32) error
	ListAppHostgroups(ctx context.Context, tx TX, filter *AppHostgroupsFilter) ([]*AppHostgroup, error)
}

type ClustersRepo interface {
	CountRequire
	CreateClusters(ctx context.Context, cs []*Cluster) error
	UpdateClusters(ctx context.Context, cs []*Cluster) error
	DeleteClusters(ctx context.Context, ids []uint32) error
	GetClusters(ctx context.Context, id uint32) (*Cluster, error)
	ListClusters(ctx context.Context, tx TX, filter *ClustersFilter) ([]*Cluster, error)
	CountClusters(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}

type DatacentersRepo interface {
	CountRequire
	CreateDatacenters(ctx context.Context, dcs []*Datacenter) error
	UpdateDatacenters(ctx context.Context, dcs []*Datacenter) error
	DeleteDatacenters(ctx context.Context, ids []uint32) error
	GetDatacenters(ctx context.Context, id uint32) (*Datacenter, error)
	ListDatacenters(ctx context.Context, tx TX, filter *DatacentersFilter) ([]*Datacenter, error)
	CountDatacenters(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}

type EnvsRepo interface {
	CountRequire
	CreateEnvs(ctx context.Context, envs []*Env) error
	UpdateEnvs(ctx context.Context, envs []*Env) error
	DeleteEnvs(ctx context.Context, ids []uint32) error
	GetEnvs(ctx context.Context, id uint32) (*Env, error)
	ListEnvs(ctx context.Context, tx TX, filter *EnvsFilter) ([]*Env, error)
	CountEnvs(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}

type FeaturesRepo interface {
	CountRequire
	CreateFeatures(ctx context.Context, features []*Feature) error
	UpdateFeatures(ctx context.Context, features []*Feature) error
	DeleteFeatures(ctx context.Context, ids []uint32) error
	GetFeatures(ctx context.Context, id uint32) (*Feature, error)
	ListFeatures(ctx context.Context, tx TX, filter *FeaturesFilter) ([]*Feature, error)
	CountFeatures(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}

type HostgroupsRepo interface {
	CountRequire
	CreateHostgroups(ctx context.Context, tx TX, hgs []*Hostgroup) error
	UpdateHostgroups(ctx context.Context, tx TX, hgs []*Hostgroup) error
	DeleteHostgroups(ctx context.Context, tx TX, ids []uint32) error
	GetHostgroups(ctx context.Context, id uint32) (*Hostgroup, error)
	ListHostgroups(ctx context.Context, tx TX, filter *HostgroupsFilter) ([]*Hostgroup, error)
	CountHostgroups(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
type HostgroupTeamsRepo interface {
	CountRequire
	CreateHostgroupTeams(ctx context.Context, tx TX, hfs []*HostgroupTeam) error
	UpdateHostgroupTeams(ctx context.Context, tx TX, hfs []*HostgroupTeam) error
	DeleteHostgroupTeams(ctx context.Context, tx TX, ids []uint32) error
	ListHostgroupTeams(ctx context.Context, tx TX,
		filter *HostgroupTeamsFilter) ([]*HostgroupTeam, error)
}
type HostgroupProductsRepo interface {
	CountRequire
	CreateHostgroupProducts(ctx context.Context, tx TX, hfs []*HostgroupProduct) error
	UpdateHostgroupProducts(ctx context.Context, tx TX, hfs []*HostgroupProduct) error
	DeleteHostgroupProducts(ctx context.Context, tx TX, ids []uint32) error
	ListHostgroupProducts(ctx context.Context, tx TX,
		filter *HostgroupProductsFilter) ([]*HostgroupProduct, error)
}

type HostgroupTagsRepo interface {
	CountRequire
	CreateHostgroupTags(ctx context.Context, tx TX, hfs []*HostgroupTag) error
	UpdateHostgroupTags(ctx context.Context, tx TX, hfs []*HostgroupTag) error
	DeleteHostgroupTags(ctx context.Context, tx TX, ids []uint32) error
	ListHostgroupTags(ctx context.Context, tx TX,
		filter *HostgroupTagsFilter) ([]*HostgroupTag, error)
}

type HostgroupFeaturesRepo interface {
	CountRequire
	CreateHostgroupFeatures(ctx context.Context, tx TX, hfs []*HostgroupFeature) error
	UpdateHostgroupFeatures(ctx context.Context, tx TX, hfs []*HostgroupFeature) error
	DeleteHostgroupFeatures(ctx context.Context, tx TX, ids []uint32) error
	ListHostgroupFeatures(ctx context.Context, tx TX,
		filter *HostgroupFeaturesFilter) ([]*HostgroupFeature, error)
}

type ProductsRepo interface {
	CountRequire
	CreateProducts(ctx context.Context, ps []*Product) error
	UpdateProducts(ctx context.Context, ps []*Product) error
	DeleteProducts(ctx context.Context, ids []uint32) error
	GetProducts(ctx context.Context, id uint32) (*Product, error)
	ListProducts(ctx context.Context, tx TX, filter *ProductsFilter) ([]*Product, error)
	CountProducts(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}

type TagsRepo interface {
	CountRequire
	CreateTags(ctx context.Context, tags []*Tag) error
	UpdateTags(ctx context.Context, tags []*Tag) error
	DeleteTags(ctx context.Context, ids []uint32) error
	GetTags(ctx context.Context, id uint32) (*Tag, error)
	ListTags(ctx context.Context, tx TX, filter *TagsFilter) ([]*Tag, error)
	CountTags(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}

type TeamsRepo interface {
	CountRequire
	CreateTeams(ctx context.Context, teams []*Team) error
	UpdateTeams(ctx context.Context, teams []*Team) error
	DeleteTeams(ctx context.Context, tx TX, ids []uint32) error
	GetTeams(ctx context.Context, id uint32) (*Team, error)
	ListTeams(ctx context.Context, tx TX, filter *TeamsFilter) ([]*Team, error)
	CountTeams(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
