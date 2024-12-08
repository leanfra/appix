package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type TagsRepo interface {
	CreateTags(ctx context.Context, tags []Tag) error
	UpdateTags(ctx context.Context, tags []Tag) error
	DeleteTags(ctx context.Context, ids []int64) error
	GetTags(ctx context.Context, id int64) (*Tag, error)
	ListTags(ctx context.Context, filter *ListTagsFilter) ([]Tag, error)
}

type TagsUsecase struct {
	repo TagsRepo
	log  *log.Helper
}

func NewTagsUsecase(repo TagsRepo, logger log.Logger) *TagsUsecase {
	return &TagsUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *TagsUsecase) validate(isNew bool, tags []Tag) error {
	for _, t := range tags {
		if err := t.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateTags is
func (s *TagsUsecase) CreateTags(ctx context.Context, tags []Tag) error {
	// validate tags
	if err := s.validate(true, tags); err != nil {
		return err
	}

	return s.repo.CreateTags(ctx, tags)
}

// UpdateTags is
func (s *TagsUsecase) UpdateTags(ctx context.Context, tags []Tag) error {

	if err := s.validate(false, tags); err != nil {
		return err
	}
	return s.repo.UpdateTags(ctx, tags)
}

// DeleteTags is
func (s *TagsUsecase) DeleteTags(ctx context.Context, ids []int64) error {

	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteTags(ctx, ids)
}

// GetTags is
func (s *TagsUsecase) GetTags(ctx context.Context, id int64) (*Tag, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	return s.repo.GetTags(ctx, id)
}

// ListTags is
func (s *TagsUsecase) ListTags(ctx context.Context, filter *ListTagsFilter) ([]Tag, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	return s.repo.ListTags(ctx, filter)
}
