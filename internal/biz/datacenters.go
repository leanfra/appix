package biz

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type DatacentersUsecase struct {
	repo     repo.DatacentersRepo
	log      *log.Helper
	txm      repo.TxManager
	required []requiredBy
}

func NewDatacentersUsecase(
	repo repo.DatacentersRepo,
	hgrepo repo.HostgroupsRepo,
	logger log.Logger, txm repo.TxManager) *DatacentersUsecase {
	return &DatacentersUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		txm:  txm,
		required: []requiredBy{
			{inst: hgrepo, name: "hostgroup"},
		},
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

	_dcs, err := ToDBDatacenters(dcs)
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
	_dcs, err := ToDBDatacenters(dcs)
	if err != nil {
		return err
	}
	return s.repo.UpdateDatacenters(ctx, _dcs)
}

// DeleteDatacenters is
func (s *DatacentersUsecase) DeleteDatacenters(ctx context.Context, tx repo.TX, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.txm.RunInTX(func(tx repo.TX) error {
		for _, r := range s.required {
			c, err := r.inst.CountRequire(ctx, tx, repo.RequireDatacenter, ids)
			if err != nil {
				return err
			}

			if c > 0 {
				return fmt.Errorf("Datacenter is required by %s", r.name)
			}
		}
		return s.repo.DeleteDatacenters(ctx, tx, ids)
	})
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
	return ToBizDatacenter(_dsc)
}

// ListDatacenters is
func (s *DatacentersUsecase) ListDatacenters(ctx context.Context,
	filter *ListDatacentersFilter) ([]*Datacenter, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	_dcs, err := s.repo.ListDatacenters(ctx, nil, ToDBDatacentersFilter(filter))
	if err != nil {
		return nil, err
	}
	return ToBizDatacenters(_dcs)
}
