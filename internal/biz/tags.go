package biz

import (
	"context"
)

type TagsRepo interface {
	CreateTags(ctx context.Context, tags []Tag) error
	UpdateTags(ctx context.Context, tags []Tag) error
	DeleteTags(ctx context.Context, ids []string) error
	GetTags(ctx context.Context, id string) (Tag, error)
	ListTags(ctx context.Context, filter *ListTagsFilter) ([]Tag, error)
}

type TagsUsecase struct {
	repo TagsRepo
}

func NewTagsUsecase(repo TagsRepo) *TagsUsecase {
	return &TagsUsecase{
		repo: repo,
	}
}

// CreateTags is
func (s *TagsUsecase) CreateTags(ctx context.Context, tags []Tag) error {
	return s.repo.CreateTags(ctx, tags)
}

// UpdateTags is
func (s *TagsUsecase) UpdateTags(ctx context.Context, tags []Tag) error {
	return s.repo.UpdateTags(ctx, tags)
}

// DeleteTags is
func (s *TagsUsecase) DeleteTags(ctx context.Context, ids []string) error {
	return s.repo.DeleteTags(ctx, ids)
}

// GetTags is
func (s *TagsUsecase) GetTags(ctx context.Context, id string) (Tag, error) {
	return s.repo.GetTags(ctx, id)
}

// ListTags is
func (s *TagsUsecase) ListTags(ctx context.Context, filter *ListTagsFilter) ([]Tag, error) {
	return s.repo.ListTags(ctx, filter)
}
