package biz_test

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockTXManager struct {
	mock.Mock
}

func (m *MockTXManager) RunInTX(fn func(repo.TX) error) error {
	return fn(nil)
}

// Mock AdminRepo
type MockAdminRepo struct {
	mock.Mock
}

func (m *MockAdminRepo) CreateUsers(ctx context.Context, tx repo.TX, users []*repo.User) error {
	args := m.Called(ctx, tx, users)
	return args.Error(0)
}

func (m *MockAdminRepo) UpdateUsers(ctx context.Context, tx repo.TX, user []*repo.User) error {
	args := m.Called(ctx, tx, user)
	return args.Error(0)
}

func (m *MockAdminRepo) DeleteUsers(ctx context.Context, tx repo.TX, user []uint32) error {
	args := m.Called(ctx, tx, user)
	return args.Error(0)
}

func (m *MockAdminRepo) GetUsers(ctx context.Context, tx repo.TX, id uint32) (*repo.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.User), args.Error(1)
}

func (m *MockAdminRepo) ListUsers(ctx context.Context, tx repo.TX, filter *repo.UsersFilter) ([]*repo.User, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.User), args.Error(1)
}

func (m *MockAdminRepo) CountUsers(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAdminRepo) Logout(ctx context.Context, id uint32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Mock AuthzRepo
type MockAuthzRepo struct {
	mock.Mock
}

func (m *MockAuthzRepo) CreateRule(ctx context.Context, tx repo.TX, rule *repo.Rule) error {
	args := m.Called(ctx, tx, rule)
	return args.Error(0)
}

func (m *MockAuthzRepo) UpdateRule(ctx context.Context, tx repo.TX, rule *repo.Rule) error {
	args := m.Called(ctx, tx, rule)
	return args.Error(0)
}

func (m *MockAuthzRepo) DeleteRule(ctx context.Context, tx repo.TX, rule *repo.Rule) error {
	args := m.Called(ctx, tx, rule)
	return args.Error(0)
}

func (m *MockAuthzRepo) GetRule(ctx context.Context, id uint32) (*repo.Rule, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Rule), args.Error(1)
}

func (m *MockAuthzRepo) ListRule(ctx context.Context, tx repo.TX, filter *repo.RuleFilter) ([]*repo.Rule, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Rule), args.Error(1)
}

// CreateGroup
func (m *MockAuthzRepo) CreateGroup(ctx context.Context, tx repo.TX, group *repo.Group) error {
	args := m.Called(ctx, tx, group)
	return args.Error(0)
}

// DeleteGroup
func (m *MockAuthzRepo) DeleteGroup(ctx context.Context, tx repo.TX, group *repo.Group) error {
	args := m.Called(ctx, tx, group)
	return args.Error(0)
}

// Enforce
func (m *MockAuthzRepo) Enforce(ctx context.Context, tx repo.TX, request *repo.AuthenRequest) (bool, error) {
	args := m.Called(ctx, tx, request)
	return args.Get(0).(bool), args.Error(1)
}

// ListGroup
func (m *MockAuthzRepo) ListGroup(ctx context.Context, tx repo.TX, filter *repo.GroupFilter) ([]*repo.Group, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Group), args.Error(1)
}

func (m *MockAuthzRepo) ListRules(ctx context.Context, tx repo.TX, filter *repo.RuleFilter) ([]*repo.Rule, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Rule), args.Error(1)
}

type MockTeamsRepo struct {
	mock.Mock
}

func (m *MockTeamsRepo) CreateTeams(ctx context.Context, tx repo.TX, teams []*repo.Team) error {
	args := m.Called(ctx, tx, teams)
	for _, t := range teams {
		fmt.Printf("%v\n", t)
	}
	return args.Error(0)
}

func (m *MockTeamsRepo) UpdateTeams(ctx context.Context, tx repo.TX, teams []*repo.Team) error {
	args := m.Called(ctx, tx, teams)
	return args.Error(0)
}

func (m *MockTeamsRepo) DeleteTeams(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockTeamsRepo) GetTeams(ctx context.Context, id uint32) (*repo.Team, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Team), args.Error(1)
}

func (m *MockTeamsRepo) ListTeams(ctx context.Context, tx repo.TX, filter *repo.TeamsFilter) ([]*repo.Team, error) {
	args := m.Called(ctx, tx, filter)
	rt := args.Get(0)
	if rt == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Team), nil
}

func (m *MockTeamsRepo) CountTeams(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTeamsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockTagsRepo struct {
	mock.Mock
}

func (m *MockTagsRepo) CreateTags(ctx context.Context, tx repo.TX, tags []*repo.Tag) error {
	args := m.Called(ctx, tags)
	return args.Error(0)
}

func (m *MockTagsRepo) UpdateTags(ctx context.Context, tx repo.TX, tags []*repo.Tag) error {
	args := m.Called(ctx, tags)
	return args.Error(0)
}

func (m *MockTagsRepo) DeleteTags(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockTagsRepo) GetTags(ctx context.Context, id uint32) (*repo.Tag, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Tag), args.Error(1)
}

func (m *MockTagsRepo) ListTags(ctx context.Context, tx repo.TX, filter *repo.TagsFilter) ([]*repo.Tag, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Tag), args.Error(1)
}

func (m *MockTagsRepo) CountTags(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTagsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockProductsRepo struct {
	mock.Mock
}

func (m *MockProductsRepo) CreateProducts(ctx context.Context, tx repo.TX, ps []*repo.Product) error {
	args := m.Called(ctx, tx, ps)
	return args.Error(0)
}

func (m *MockProductsRepo) UpdateProducts(ctx context.Context, tx repo.TX, ps []*repo.Product) error {
	args := m.Called(ctx, tx, ps)
	return args.Error(0)
}

func (m *MockProductsRepo) DeleteProducts(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockProductsRepo) GetProducts(ctx context.Context, id uint32) (*repo.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Product), args.Error(1)
}

func (m *MockProductsRepo) ListProducts(ctx context.Context, tx repo.TX, filter *repo.ProductsFilter) ([]*repo.Product, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Product), args.Error(1)
}

func (m *MockProductsRepo) CountProducts(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockProductsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockHostgroupFeaturesRepo struct {
	mock.Mock
}

func (m *MockHostgroupFeaturesRepo) CreateHostgroupFeatures(ctx context.Context, tx repo.TX, hgfs []*repo.HostgroupFeature) error {
	args := m.Called(ctx, tx, hgfs)
	return args.Error(0)
}

func (m *MockHostgroupFeaturesRepo) UpdateHostgroupFeatures(ctx context.Context, tx repo.TX, hgfs []*repo.HostgroupFeature) error {
	args := m.Called(ctx, tx, hgfs)
	return args.Error(0)
}

func (m *MockHostgroupFeaturesRepo) DeleteHostgroupFeatures(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockHostgroupFeaturesRepo) GetHostgroupFeatures(ctx context.Context, id uint32) (*repo.HostgroupFeature, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.HostgroupFeature), args.Error(1)
}

func (m *MockHostgroupFeaturesRepo) ListHostgroupMatchFeatures(
	ctx context.Context, tx repo.TX, filter *repo.HostgroupMatchFeaturesFilter) ([]uint32, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]uint32), args.Error(1)
}

func (m *MockHostgroupFeaturesRepo) ListHostgroupFeatures(
	ctx context.Context, tx repo.TX, filter *repo.HostgroupFeaturesFilter) ([]*repo.HostgroupFeature, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.HostgroupFeature), args.Error(1)
}

func (m *MockHostgroupFeaturesRepo) CountHostgroupFeatures(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockHostgroupFeaturesRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockHostgroupTagsRepo struct {
	mock.Mock
}

func (m *MockHostgroupTagsRepo) CreateHostgroupTags(ctx context.Context, tx repo.TX, hgt []*repo.HostgroupTag) error {
	args := m.Called(ctx, tx, hgt)
	return args.Error(0)
}

func (m *MockHostgroupTagsRepo) UpdateHostgroupTags(ctx context.Context, tx repo.TX, hgt []*repo.HostgroupTag) error {
	args := m.Called(ctx, tx, hgt)
	return args.Error(0)
}

func (m *MockHostgroupTagsRepo) DeleteHostgroupTags(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockHostgroupTagsRepo) GetHostgroupTags(ctx context.Context, id uint32) (*repo.HostgroupTag, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.HostgroupTag), args.Error(1)
}

func (m *MockHostgroupTagsRepo) ListHostgroupTags(ctx context.Context, tx repo.TX, filter *repo.HostgroupTagsFilter) ([]*repo.HostgroupTag, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.HostgroupTag), args.Error(1)
}

func (m *MockHostgroupTagsRepo) CountHostgroupTags(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockHostgroupTagsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockHostgroupProductsRepo struct {
	mock.Mock
}

func (m *MockHostgroupProductsRepo) CreateHostgroupProducts(ctx context.Context, tx repo.TX, hgp []*repo.HostgroupProduct) error {
	args := m.Called(ctx, tx, hgp)
	return args.Error(0)
}

func (m *MockHostgroupProductsRepo) UpdateHostgroupProducts(ctx context.Context, tx repo.TX, hgp []*repo.HostgroupProduct) error {
	args := m.Called(ctx, hgp)
	return args.Error(0)
}

func (m *MockHostgroupProductsRepo) DeleteHostgroupProducts(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockHostgroupProductsRepo) GetHostgroupProducts(ctx context.Context, id uint32) (*repo.HostgroupProduct, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.HostgroupProduct), args.Error(1)
}

func (m *MockHostgroupProductsRepo) ListHostgroupProducts(ctx context.Context, tx repo.TX, filter *repo.HostgroupProductsFilter) ([]*repo.HostgroupProduct, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.HostgroupProduct), args.Error(1)
}

func (m *MockHostgroupProductsRepo) CountHostgroupProducts(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockHostgroupProductsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockHostgroupTeamsRepo struct {
	mock.Mock
}

func (m *MockHostgroupTeamsRepo) CreateHostgroupTeams(ctx context.Context,
	tx repo.TX,
	hgt []*repo.HostgroupTeam) error {
	args := m.Called(ctx, tx, hgt)
	return args.Error(0)
}

func (m *MockHostgroupTeamsRepo) UpdateHostgroupTeams(ctx context.Context,
	tx repo.TX,
	hgt []*repo.HostgroupTeam) error {
	args := m.Called(ctx, tx, hgt)
	return args.Error(0)
}

func (m *MockHostgroupTeamsRepo) DeleteHostgroupTeams(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockHostgroupTeamsRepo) GetHostgroupTeams(ctx context.Context, id uint32) (*repo.HostgroupTeam, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.HostgroupTeam), args.Error(1)
}

func (m *MockHostgroupTeamsRepo) ListHostgroupTeams(ctx context.Context, tx repo.TX, filter *repo.HostgroupTeamsFilter) ([]*repo.HostgroupTeam, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.HostgroupTeam), args.Error(1)
}

func (m *MockHostgroupTeamsRepo) CountHostgroupTeams(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockHostgroupTeamsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockHostgroupsRepo struct {
	mock.Mock
}

func (m *MockHostgroupsRepo) CreateHostgroups(ctx context.Context, tx repo.TX, hg []*repo.Hostgroup) error {
	args := m.Called(ctx, tx, hg)
	return args.Error(0)
}

func (m *MockHostgroupsRepo) UpdateHostgroups(ctx context.Context, tx repo.TX, hg []*repo.Hostgroup) error {
	args := m.Called(ctx, tx, hg)
	return args.Error(0)
}

func (m *MockHostgroupsRepo) DeleteHostgroups(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockHostgroupsRepo) GetHostgroups(ctx context.Context, id uint32) (*repo.Hostgroup, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Hostgroup), args.Error(1)
}

func (m *MockHostgroupsRepo) ListHostgroups(ctx context.Context, tx repo.TX, filter *repo.HostgroupsFilter) ([]*repo.Hostgroup, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Hostgroup), args.Error(1)
}

func (m *MockHostgroupsRepo) CountHostgroups(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockHostgroupsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockFeaturesRepo struct {
	mock.Mock
}

func (m *MockFeaturesRepo) CreateFeatures(ctx context.Context, tx repo.TX, f []*repo.Feature) error {
	args := m.Called(ctx, tx, f)
	return args.Error(0)
}

func (m *MockFeaturesRepo) UpdateFeatures(ctx context.Context, tx repo.TX, f []*repo.Feature) error {
	args := m.Called(ctx, tx, f)
	return args.Error(0)
}

func (m *MockFeaturesRepo) DeleteFeatures(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, ids)
	return args.Error(0)
}

func (m *MockFeaturesRepo) GetFeatures(ctx context.Context, id uint32) (*repo.Feature, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Feature), args.Error(1)
}

func (m *MockFeaturesRepo) ListFeatures(ctx context.Context, tx repo.TX, filter *repo.FeaturesFilter) ([]*repo.Feature, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Feature), args.Error(1)
}

func (m *MockFeaturesRepo) CountFeatures(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockFeaturesRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockEnvsRepo struct {
	mock.Mock
}

func (m *MockEnvsRepo) CreateEnvs(ctx context.Context, tx repo.TX, e []*repo.Env) error {
	args := m.Called(ctx, tx, e)
	return args.Error(0)
}

func (m *MockEnvsRepo) UpdateEnvs(ctx context.Context, tx repo.TX, e []*repo.Env) error {
	args := m.Called(ctx, tx, e)
	return args.Error(0)
}

func (m *MockEnvsRepo) DeleteEnvs(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockEnvsRepo) GetEnvs(ctx context.Context, id uint32) (*repo.Env, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Env), args.Error(1)
}

func (m *MockEnvsRepo) ListEnvs(ctx context.Context, tx repo.TX, filter *repo.EnvsFilter) ([]*repo.Env, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Env), args.Error(1)
}

func (m *MockEnvsRepo) CountEnvs(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockEnvsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockDatacentersRepo struct {
	mock.Mock
}

func (m *MockDatacentersRepo) CreateDatacenters(ctx context.Context, tx repo.TX, d []*repo.Datacenter) error {
	args := m.Called(ctx, d)
	return args.Error(0)
}

func (m *MockDatacentersRepo) UpdateDatacenters(ctx context.Context, tx repo.TX, d []*repo.Datacenter) error {
	args := m.Called(ctx, d)
	return args.Error(0)
}

func (m *MockDatacentersRepo) DeleteDatacenters(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockDatacentersRepo) GetDatacenters(ctx context.Context, id uint32) (*repo.Datacenter, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Datacenter), args.Error(1)
}

func (m *MockDatacentersRepo) ListDatacenters(ctx context.Context, tx repo.TX, filter *repo.DatacentersFilter) ([]*repo.Datacenter, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Datacenter), args.Error(1)
}

func (m *MockDatacentersRepo) CountDatacenters(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDatacentersRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockClustersRepo struct {
	mock.Mock
}

func (m *MockClustersRepo) CreateClusters(ctx context.Context, tx repo.TX, c []*repo.Cluster) error {
	args := m.Called(ctx, tx, c)
	return args.Error(0)
}

func (m *MockClustersRepo) UpdateClusters(ctx context.Context, tx repo.TX, c []*repo.Cluster) error {
	args := m.Called(ctx, tx, c)
	return args.Error(0)
}

func (m *MockClustersRepo) DeleteClusters(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockClustersRepo) GetClusters(ctx context.Context, id uint32) (*repo.Cluster, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Cluster), args.Error(1)
}

func (m *MockClustersRepo) ListClusters(ctx context.Context, tx repo.TX, filter *repo.ClustersFilter) ([]*repo.Cluster, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Cluster), args.Error(1)
}

func (m *MockClustersRepo) CountClusters(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockClustersRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

type MockAppHostgroupsRepo struct {
	mock.Mock
}

func (m *MockAppHostgroupsRepo) CreateAppHostgroups(ctx context.Context, tx repo.TX, a []*repo.AppHostgroup) error {
	args := m.Called(ctx, tx, a)
	return args.Error(0)
}

func (m *MockAppHostgroupsRepo) UpdateAppHostgroups(ctx context.Context, tx repo.TX, a []*repo.AppHostgroup) error {
	args := m.Called(ctx, a)
	return args.Error(0)
}

func (m *MockAppHostgroupsRepo) DeleteAppHostgroups(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockAppHostgroupsRepo) GetAppHostgroups(ctx context.Context, id uint32) (*repo.AppHostgroup, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.AppHostgroup), args.Error(1)
}

func (m *MockAppHostgroupsRepo) ListAppHostgroups(ctx context.Context, tx repo.TX, filter *repo.AppHostgroupsFilter) ([]*repo.AppHostgroup, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.AppHostgroup), args.Error(1)
}

func (m *MockAppHostgroupsRepo) CountAppHostgroups(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAppHostgroupsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAppHostgroupsRepo) DeleteAppHostgroupsByAppId(ctx context.Context, tx repo.TX, appID []uint32) error {
	args := m.Called(ctx, tx, appID)
	return args.Error(0)
}

type MockAppFeaturesRepo struct {
	mock.Mock
}

func (m *MockAppFeaturesRepo) CreateAppFeatures(ctx context.Context, tx repo.TX, a []*repo.AppFeature) error {
	args := m.Called(ctx, tx, a)
	return args.Error(0)
}

func (m *MockAppFeaturesRepo) UpdateAppFeatures(ctx context.Context, tx repo.TX, a []*repo.AppFeature) error {
	args := m.Called(ctx, tx, a)
	return args.Error(0)
}

func (m *MockAppFeaturesRepo) DeleteAppFeatures(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockAppFeaturesRepo) GetAppFeatures(ctx context.Context, id uint32) (*repo.AppFeature, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.AppFeature), args.Error(1)
}

func (m *MockAppFeaturesRepo) ListAppFeatures(ctx context.Context, tx repo.TX, filter *repo.AppFeaturesFilter) ([]*repo.AppFeature, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.AppFeature), args.Error(1)
}

func (m *MockAppFeaturesRepo) CountAppFeatures(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAppFeaturesRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAppFeaturesRepo) DeleteAppFeaturesByAppId(ctx context.Context, tx repo.TX, appID []uint32) error {
	args := m.Called(ctx, tx, appID)
	return args.Error(0)
}

type MockAppTagsRepo struct {
	mock.Mock
}

func (m *MockAppTagsRepo) CreateAppTags(ctx context.Context, tx repo.TX, a []*repo.AppTag) error {
	args := m.Called(ctx, tx, a)
	return args.Error(0)
}

func (m *MockAppTagsRepo) UpdateAppTags(ctx context.Context, tx repo.TX, a []*repo.AppTag) error {
	args := m.Called(ctx, tx, a)
	return args.Error(0)
}

func (m *MockAppTagsRepo) DeleteAppTags(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}

func (m *MockAppTagsRepo) GetAppTags(ctx context.Context, id uint32) (*repo.AppTag, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.AppTag), args.Error(1)
}

func (m *MockAppTagsRepo) ListAppTags(ctx context.Context, tx repo.TX, filter *repo.AppTagsFilter) ([]*repo.AppTag, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.AppTag), args.Error(1)
}

func (m *MockAppTagsRepo) CountAppTags(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAppTagsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAppTagsRepo) DeleteAppTagsByAppId(ctx context.Context, tx repo.TX, appID []uint32) error {
	args := m.Called(ctx, tx, appID)
	return args.Error(0)
}

type MockApplicationsRepo struct {
	mock.Mock
}

func (m *MockApplicationsRepo) CreateApplications(ctx context.Context,
	tx repo.TX, a []*repo.Application) error {
	args := m.Called(ctx, tx, a)
	return args.Error(0)
}
func (m *MockApplicationsRepo) UpdateApplications(ctx context.Context,
	tx repo.TX,
	a []*repo.Application) error {
	args := m.Called(ctx, tx, a)
	return args.Error(0)
}
func (m *MockApplicationsRepo) DeleteApplications(ctx context.Context, tx repo.TX, ids []uint32) error {
	args := m.Called(ctx, tx, ids)
	return args.Error(0)
}
func (m *MockApplicationsRepo) GetApplications(ctx context.Context, id uint32) (*repo.Application, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Application), args.Error(1)
}
func (m *MockApplicationsRepo) ListApplications(ctx context.Context, tx repo.TX, filter *repo.ApplicationsFilter) ([]*repo.Application, error) {
	args := m.Called(ctx, tx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*repo.Application), args.Error(1)
}
func (m *MockApplicationsRepo) CountApplications(ctx context.Context, tx repo.TX, filter repo.CountFilter) (int64, error) {
	args := m.Called(ctx, tx, filter)
	return args.Get(0).(int64), args.Error(1)
}
func (m *MockApplicationsRepo) CountRequire(ctx context.Context, tx repo.TX, need repo.RequireType, ids []uint32) (int64, error) {
	args := m.Called(ctx, tx, need, ids)
	return args.Get(0).(int64), args.Error(1)
}
