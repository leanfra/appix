package biz

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type TagsUsecase struct {
	repo     repo.TagsRepo
	txm      repo.TxManager
	log      *log.Helper
	required []requiredBy
}

func NewTagsUsecase(repo repo.TagsRepo,
	logger log.Logger,
	apptagrepo repo.AppTagsRepo,
	hgtagrepo repo.HostgroupTagsRepo,
	txm repo.TxManager) *TagsUsecase {

	return &TagsUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		txm:  txm,
		required: []requiredBy{
			{inst: apptagrepo, name: "app_tag"},
			{inst: hgtagrepo, name: "hostgroup_tag"},
		},
	}
}

func (s *TagsUsecase) validate(isNew bool, tags []*Tag) error {
	for _, t := range tags {
		if err := t.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateTags is
func (s *TagsUsecase) CreateTags(ctx context.Context, tags []*Tag) error {
	// validate tags
	if err := s.validate(true, tags); err != nil {
		return err
	}

	_tags, e := ToDBTags(tags)
	if e != nil {
		return e
	}

	return s.repo.CreateTags(ctx, _tags)
}

// UpdateTags is
func (s *TagsUsecase) UpdateTags(ctx context.Context, tags []*Tag) error {

	if err := s.validate(false, tags); err != nil {
		return err
	}
	_tags, e := ToDBTags(tags)
	if e != nil {
		return e
	}
	return s.repo.UpdateTags(ctx, _tags)
}

// DeleteTags is
func (s *TagsUsecase) DeleteTags(ctx context.Context, ids []uint32) error {

	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.txm.RunInTX(func(tx repo.TX) error {
		for _, r := range s.required {
			c, err := r.inst.CountRequire(ctx, tx, repo.RequireTag, ids)
			if err != nil {
				return err
			}
			if c > 0 {
				return fmt.Errorf("some %s requires", r.name)
			}
		}
		if e := s.repo.DeleteTags(ctx, tx, ids); e != nil {
			return e
		}
		return nil
	})
}

// GetTags is
func (s *TagsUsecase) GetTags(ctx context.Context, id uint32) (*Tag, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	t, e := s.repo.GetTags(ctx, id)
	if e != nil {
		return nil, e
	}
	return ToBizTag(t)
}

// ListTags is
func (s *TagsUsecase) ListTags(ctx context.Context,
	filter *ListTagsFilter) ([]*Tag, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	_ts, e := s.repo.ListTags(ctx, nil, ToDBTagsFilter(filter))
	if e != nil {
		return nil, e
	}
	return ToBizTags(_ts)
}
