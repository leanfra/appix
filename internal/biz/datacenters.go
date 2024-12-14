package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type DatacentersRepo interface {
	CreateDatacenters(ctx context.Context, dcs []Datacenter) error
	UpdateDatacenters(ctx context.Context, dcs []Datacenter) error
	DeleteDatacenters(ctx context.Context, ids []int64) error
	GetDatacenters(ctx context.Context, id int64) (*Datacenter, error)
	ListDatacenters(ctx context.Context, filter *ListDatacentersFilter) ([]Datacenter, error)
}

type DatacentersUsecase struct {
	repo DatacentersRepo
	log  *log.Helper
}

func NewDatacentersUsecase(repo DatacentersRepo, logger log.Logger) *DatacentersUsecase {
	return &DatacentersUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *DatacentersUsecase) validate(isNew bool, dcs []Datacenter) error {
	for _, d := range dcs {
		if err := d.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

// CreateDatacenters is
func (s *DatacentersUsecase) CreateDatacenters(ctx context.Context, dcs []Datacenter) error {

	if err := s.validate(true, dcs); err != nil {
		return err
	}
	return s.repo.CreateDatacenters(ctx, dcs)
}

// UpdateDatacenters is
func (s *DatacentersUsecase) UpdateDatacenters(ctx context.Context, dcs []Datacenter) error {
	if err := s.validate(false, dcs); err != nil {
		return err
	}
	return s.repo.UpdateDatacenters(ctx, dcs)
}

// DeleteDatacenters is
func (s *DatacentersUsecase) DeleteDatacenters(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.repo.DeleteDatacenters(ctx, ids)
}

// GetDatacenters is
func (s *DatacentersUsecase) GetDatacenters(ctx context.Context, id int64) (*Datacenter, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	return s.repo.GetDatacenters(ctx, id)
}

// ListDatacenters is
func (s *DatacentersUsecase) ListDatacenters(ctx context.Context,
	filter *ListDatacentersFilter) ([]Datacenter, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	return s.repo.ListDatacenters(ctx, filter)
}
