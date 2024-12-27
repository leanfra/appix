package biz

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type DatacentersUsecase struct {
	repo repo.DatacentersRepo
	log  *log.Helper
	txm  repo.TxManager
}

func NewDatacentersUsecase(repo repo.DatacentersRepo, logger log.Logger, txm repo.TxManager) *DatacentersUsecase {
	return &DatacentersUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		txm:  txm,
	}
}

func (s *DatacentersUsecase) validate(isNew bool, dcs []*Datacenter) error {
	for _, d := range dcs {
		if err := d.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateDatacenters is
func (s *DatacentersUsecase) CreateDatacenters(ctx context.Context, dcs []*Datacenter) error {

	if err := s.validate(true, dcs); err != nil {
		return err
	}

	_dcs, err := NewDatacenters(dcs)
	if err != nil {
		return err
	}
	return s.repo.CreateDatacenters(ctx, _dcs)
}

// UpdateDatacenters is
func (s *DatacentersUsecase) UpdateDatacenters(ctx context.Context, dcs []*Datacenter) error {
	if err := s.validate(false, dcs); err != nil {
		return err
	}
	_dcs, err := NewDatacenters(dcs)
	if err != nil {
		return err
	}
	return s.repo.UpdateDatacenters(ctx, _dcs)
}

// DeleteDatacenters is
func (s *DatacentersUsecase) DeleteDatacenters(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteDatacenters(ctx, ids)
}

// GetDatacenters is
func (s *DatacentersUsecase) GetDatacenters(ctx context.Context, id uint32) (*Datacenter, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	_dsc, err := s.repo.GetDatacenters(ctx, id)
	if err != nil {
		return nil, err
	}
	return NewBizDatacenter(_dsc)
}

// ListDatacenters is
func (s *DatacentersUsecase) ListDatacenters(ctx context.Context,
	filter *ListDatacentersFilter) ([]*Datacenter, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	_dcs, err := s.repo.ListDatacenters(ctx, nil, NewDatacentersFilter(filter))
	if err != nil {
		return nil, err
	}
	return NewBizDatacenters(_dcs)
}
