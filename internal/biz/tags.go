package biz

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type TagsUsecase struct {
	repo repo.TagsRepo
	txm  repo.TxManager
	log  *log.Helper
}

func NewTagsUsecase(repo repo.TagsRepo, logger log.Logger, txm repo.TxManager) *TagsUsecase {
	return &TagsUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		txm:  txm,
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
	return s.repo.DeleteTags(ctx, ids)
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
