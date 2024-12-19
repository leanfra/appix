package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type TeamsRepo interface {
	CreateTeams(ctx context.Context, teams []*Team) error
	UpdateTeams(ctx context.Context, teams []*Team) error
	DeleteTeams(ctx context.Context, ids []uint32) error
	GetTeams(ctx context.Context, id uint32) (*Team, error)
	ListTeams(ctx context.Context, filter *ListTeamsFilter) ([]*Team, error)
}

type TeamsUsecase struct {
	repo TeamsRepo
	log  *log.Helper
}

func NewTeamsUsecase(repo TeamsRepo, logger log.Logger) *TeamsUsecase {
	return &TeamsUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *TeamsUsecase) validate(isNew bool, teams []*Team) error {
	for _, t := range teams {
		if err := t.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateTeams is
func (s *TeamsUsecase) CreateTeams(ctx context.Context, teams []*Team) error {
	if err := s.validate(true, teams); err != nil {
		return err
	}

	return s.repo.CreateTeams(ctx, teams)
}

// UpdateTeams is
func (s *TeamsUsecase) UpdateTeams(ctx context.Context, teams []*Team) error {
	if err := s.validate(false, teams); err != nil {
		return err
	}
	return s.repo.UpdateTeams(ctx, teams)
}

// DeleteTeams is
func (s *TeamsUsecase) DeleteTeams(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteTeams(ctx, ids)
}

// GetTeams is
func (s *TeamsUsecase) GetTeams(ctx context.Context, id uint32) (*Team, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	return s.repo.GetTeams(ctx, id)
}

// ListTeams is
func (s *TeamsUsecase) ListTeams(ctx context.Context, filter *ListTeamsFilter) ([]*Team, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	return s.repo.ListTeams(ctx, filter)
}
