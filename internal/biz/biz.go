package biz

import (
	"errors"
	"strings"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewGreeterUsecase,
	NewTagsUsecase,
	NewFeaturesUsecase,
	NewTeamsUsecase,
	NewProductsUsecase,
	NewEnvsUsecase,
	NewClustersUsecase,
	NewDatacentersUsecase,
	NewHostgroupsUsecase,
	NewApplicationsUsecase,
)

const MaxFilterValues = 10
const DefaultPageSize = 50
const FilterKVSplit = ":"

var ErrFilterValuesExceedMax = errors.New("filter values exceeded max number")
var ErrFilterKVInvalid = errors.New("filter KV invalid format")

func filterKvValidate(kvstr string) error {
	kv := strings.Split(kvstr, FilterKVSplit)
	if len(kv) != 2 {
		return ErrFilterKVInvalid
	}
	return nil
}
