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
)

var ErrRequestNil = errors.New("requestIsNil")
