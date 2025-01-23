package biz

import (
	"opspillar/internal/data"
	"opspillar/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type ClustersUsecase struct {
	csrepo    repo.ClustersRepo
	authzrepo repo.AuthzRepo
	log       *log.Helper
	txm       repo.TxManager
	required  []requiredBy
}

func NewClustersUsecase(
	repo repo.ClustersRepo,
	authzrepo repo.AuthzRepo,
	hgrepo repo.HostgroupsRepo,
	logger log.Logger, txm repo.TxManager) *ClustersUsecase {
	return &ClustersUsecase{
		csrepo:    repo,
		authzrepo: authzrepo,
		log:       log.NewHelper(logger),
		txm:       txm,
		required: []requiredBy{
			{inst: hgrepo, name: "hostgroup"},
		},
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

// enforce only Enforce `cluster` resource instead of `cluster instance`
func (s *ClustersUsecase) enforce(ctx context.Context, tx repo.TX) error {

	user := ctx.Value(data.CtxUserName).(string)
	ires := repo.NewResource4Sv1("clusters", "", "", "")
	can, err := s.authzrepo.Enforce(ctx, tx, &repo.AuthenRequest{
		Sub:      user,
		Resource: ires,
		Action:   repo.ActWrite,
	})
	if err != nil {
		return err
	}
	if !can {
		return fmt.Errorf("PermissionDenied")
	}
	return nil
}

// CreateClusters is
func (s *ClustersUsecase) CreateClusters(ctx context.Context, cs []*Cluster) error {
	if err := s.validate(true, cs); err != nil {
		return err
	}
	_cs, err := ToDBClusters(cs)
	if err != nil {
		return err
	}
	err = s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		return s.csrepo.CreateClusters(ctx, tx, _cs)
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateClusters is
func (s *ClustersUsecase) UpdateClusters(ctx context.Context, cs []*Cluster) error {
	if err := s.validate(false, cs); err != nil {
		return err
	}
	_cs, err := ToDBClusters(cs)
	if err != nil {
		return err
	}
	err = s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		return s.csrepo.UpdateClusters(ctx, tx, _cs)
	})
	if err != nil {
		return err
	}
	return nil
}

// TODO need authz
// DeleteClusters is
func (s *ClustersUsecase) DeleteClusters(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}

	return s.txm.RunInTX(func(tx repo.TX) error {

		if err := s.enforce(ctx, tx); err != nil {
			return err
		}

		for _, r := range s.required {
			c, err := r.inst.CountRequire(ctx, tx, repo.RequireCluster, ids)
			if err != nil {
				return err
			}
			if c > 0 {
				return fmt.Errorf("Cluster is required by %s", r.name)
			}
		}
		return s.csrepo.DeleteClusters(ctx, tx, ids)
	})
}

// GetClusters is
func (s *ClustersUsecase) GetClusters(ctx context.Context, id uint32) (*Cluster, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	_cs, err := s.csrepo.GetClusters(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToBizCluster(_cs)
}

// ListClusters is
func (s *ClustersUsecase) ListClusters(ctx context.Context,
	filter *ListClustersFilter) ([]*Cluster, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	_cs, err := s.csrepo.ListClusters(ctx, nil, ToDBClustersFilter(filter))
	if err != nil {
		return nil, err
	}
	return ToBizClusters(_cs)
}
