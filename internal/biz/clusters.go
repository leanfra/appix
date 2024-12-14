package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type ClustersRepo interface {
	CreateClusters(ctx context.Context, cs []Cluster) error
	UpdateClusters(ctx context.Context, cs []Cluster) error
	DeleteClusters(ctx context.Context, ids []int64) error
	GetClusters(ctx context.Context, id int64) (*Cluster, error)
	ListClusters(ctx context.Context, filter *ListClustersFilter) ([]Cluster, error)
}

type ClustersUsecase struct {
	repo ClustersRepo
	log  *log.Helper
}

func NewClustersUsecase(repo ClustersRepo, logger log.Logger) *ClustersUsecase {
	return &ClustersUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *ClustersUsecase) validate(isNew bool, cs []Cluster) error {
	for _, c := range cs {
		if err := c.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateClusters is
func (s *ClustersUsecase) CreateClusters(ctx context.Context, cs []Cluster) error {
	if err := s.validate(true, cs); err != nil {
		return err
	}
	return s.repo.CreateClusters(ctx, cs)
}

// UpdateClusters is
func (s *ClustersUsecase) UpdateClusters(ctx context.Context, cs []Cluster) error {
	if err := s.validate(false, cs); err != nil {
		return err
	}
	return s.repo.UpdateClusters(ctx, cs)
}

// DeleteClusters is
func (s *ClustersUsecase) DeleteClusters(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}

	return s.repo.DeleteClusters(ctx, ids)
}

// GetClusters is
func (s *ClustersUsecase) GetClusters(ctx context.Context, id int64) (*Cluster, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	return s.repo.GetClusters(ctx, id)
}

// ListClusters is
func (s *ClustersUsecase) ListClusters(ctx context.Context, filter *ListClustersFilter) ([]Cluster, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	return s.repo.ListClusters(ctx, filter)
}
