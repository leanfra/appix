package biz

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type ClustersUsecase struct {
	repo repo.ClustersRepo
	log  *log.Helper
	txm  repo.TxManager
}

func NewClustersUsecase(repo repo.ClustersRepo, logger log.Logger, txm repo.TxManager) *ClustersUsecase {
	return &ClustersUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		txm:  txm,
	}
}

func (s *ClustersUsecase) validate(isNew bool, cs []*Cluster) error {
	for _, c := range cs {
		if err := c.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateClusters is
func (s *ClustersUsecase) CreateClusters(ctx context.Context, cs []*Cluster) error {
	if err := s.validate(true, cs); err != nil {
		return err
	}
	_cs, err := NewClusters(cs)
	if err != nil {
		return err
	}
	return s.repo.CreateClusters(ctx, _cs)
}

// UpdateClusters is
func (s *ClustersUsecase) UpdateClusters(ctx context.Context, cs []*Cluster) error {
	if err := s.validate(false, cs); err != nil {
		return err
	}
	_cs, err := NewClusters(cs)
	if err != nil {
		return err
	}
	return s.repo.UpdateClusters(ctx, _cs)
}

// DeleteClusters is
func (s *ClustersUsecase) DeleteClusters(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}

	return s.repo.DeleteClusters(ctx, ids)
}

// GetClusters is
func (s *ClustersUsecase) GetClusters(ctx context.Context, id uint32) (*Cluster, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	_cs, err := s.repo.GetClusters(ctx, id)
	if err != nil {
		return nil, err
	}
	return NewBizCluster(_cs)
}

// ListClusters is
func (s *ClustersUsecase) ListClusters(ctx context.Context,
	filter *ListClustersFilter) ([]*Cluster, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	_cs, err := s.repo.ListClusters(ctx, nil, NewClustersFilter(filter))
	if err != nil {
		return nil, err
	}
	return NewBizClusters(_cs)
}
