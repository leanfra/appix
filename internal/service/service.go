package service

import (
	"errors"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewTagsService,
	NewFeaturesService,
	NewTeamsService,
	NewProductsService,
	NewEnvsService,
	NewClustersService,
	NewDatacentersService,
	NewHostgroupsService,
	NewApplicationsService,
	NewAdminService,
)

var ErrRequestNil = errors.New("requestIsNil")
