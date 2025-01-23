package biz

import (
	"opspillar/internal/data/repo"
	"errors"
	"regexp"
	"strings"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewTagsUsecase,
	NewFeaturesUsecase,
	NewTeamsUsecase,
	NewProductsUsecase,
	NewEnvsUsecase,
	NewClustersUsecase,
	NewDatacentersUsecase,
	NewHostgroupsUsecase,
	NewApplicationsUsecase,
	NewAdminUsecase,
)

const MaxFilterValues = 10
const DefaultPageSize = 50
const MaxPageSize = 200
const FilterKVSplit = ":"

var ErrFilterValuesExceedMax = errors.New("filter values exceeded max number")
var ErrFilterKVInvalid = errors.New("filter KV invalid format")
var ErrFilterInvalidPagesize = errors.New("filter invalid page size")
var ErrFilterInvalidPage = errors.New("filter invalid page")

func filterKvValidate(kvstr string) error {
	kv := strings.Split(kvstr, FilterKVSplit)
	if len(kv) != 2 {
		return ErrFilterKVInvalid
	}
	return nil
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
func DiffSliceUint32(s1 []uint32, s2 []uint32) []uint32 {
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

// IntersectSliceUint32 返回两个 []uint32 切片的交集
func IntersectSliceUint32(slice1, slice2 []uint32) []uint32 {
	if len(slice1) == 0 || len(slice2) == 0 {
		return []uint32{}
	}

	elemMap := make(map[uint32]struct{})
	for _, num := range slice1 {
		elemMap[num] = struct{}{}
	}

	intersection := []uint32{}
	for _, num := range slice2 {
		if _, exists := elemMap[num]; exists {
			intersection = append(intersection, num)
			delete(elemMap, num) // 删除已找到的元素以避免重复
		}
	}

	return intersection
}

type requiredBy struct {
	name string
	inst repo.RequireCounter
}

var NamePattern = regexp.MustCompile(`^[a-z][a-z0-9-]*[a-z0-9]$`)
var CodePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*[a-z0-9]$`)

const MaxNameLength = 255
const MaxCodeLength = 255

func ValidateName(name string) error {
	if len(name) > MaxNameLength {
		return errors.New("name too long")
	}
	if !NamePattern.MatchString(name) {
		return errors.New("name invalid")
	}
	return nil
}

func ValidateCode(code string) error {
	if len(code) > MaxCodeLength {
		return errors.New("code too long")
	}
	if !CodePattern.MatchString(code) {
		return errors.New("code invalid")
	}
	return nil
}
