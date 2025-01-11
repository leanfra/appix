package data

import (
	"appix/internal/data/sqldb"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	sqldb.NewDataGorm,
	sqldb.NewTxManagerGorm,
	sqldb.NewFeaturesRepoGorm,
	sqldb.NewTagsRepoGorm,
	sqldb.NewTeamsRepoGorm,
	sqldb.NewProductsRepoGorm,
	sqldb.NewEnvsRepoGorm,
	sqldb.NewClustersRepoGorm,
	sqldb.NewDatacentersRepoGorm,
	sqldb.NewHostgroupsRepoGorm,
	sqldb.NewApplicationsRepoGorm,
	sqldb.NewAppTagsRepoGorm,
	sqldb.NewAppFeaturesRepoGorm,
	sqldb.NewAppHostgroupsRepoGorm,
	sqldb.NewHostgroupTeamsRepoGorm,
	sqldb.NewHostgroupProductsRepoGorm,
	sqldb.NewHostgroupTagsRepoGorm,
	sqldb.NewHostgroupFeaturesRepoGorm,
	sqldb.NewAdminRepoGorm,
	sqldb.NewAuthzRepoGorm,
	NewJwtMemRepo,
)
