package service

import (
	"errors"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewGreeterService,
	NewTagsService,
	NewFeaturesService,
	NewTeamsService,
	NewProductsService,
	NewEnvsService,
	NewClustersService,
	NewDatacentersService,
	NewHostgroupsService,
)

var ErrRequestNil = errors.New("requestIsNil")
