package biz

import (
	"context"
	"fmt"

	"appix/internal/data"
	"appix/internal/data/repo"

	"github.com/go-kratos/kratos/v2/log"
)

type TeamsUsecase struct {
	teamRepo  repo.TeamsRepo
	authzrepo repo.AuthzRepo
	txm       repo.TxManager
	log       *log.Helper
	required  []requiredBy
}

func NewTeamsUsecase(
	teamrepo repo.TeamsRepo,
	authzrepo repo.AuthzRepo,
	hgrepo repo.HostgroupsRepo,
	htrepo repo.HostgroupTeamsRepo,
	apprepo repo.ApplicationsRepo,
	logger log.Logger,
	txm repo.TxManager) *TeamsUsecase {

	return &TeamsUsecase{
		teamRepo:  teamrepo,
		authzrepo: authzrepo,
		log:       log.NewHelper(logger),
		txm:       txm,
		required: []requiredBy{
			{inst: hgrepo, name: "hostgroup"},
			{inst: apprepo, name: "app"},
			{inst: htrepo, name: "hostgroup_team"},
		},
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

func (s *TeamsUsecase) enforce(ctx context.Context, tx repo.TX) error {
	curUser := ctx.Value(data.CtxUserName).(string)
	ires := repo.NewResource4Sv1("team", "", "", "")
	can, err := s.authzrepo.Enforce(ctx, tx, &repo.AuthenRequest{
		Sub:      curUser,
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

// CreateTeams is
func (s *TeamsUsecase) CreateTeams(ctx context.Context, teams []*Team) error {
	if err := s.validate(true, teams); err != nil {
		return err
	}

	_teams, e := ToDBTeams(teams)
	if e != nil {
		return e
	}

	err := s.txm.RunInTX(
		func(tx repo.TX) error {
			if err := s.enforce(ctx, tx); err != nil {
				return err
			}
			if e := s.teamRepo.CreateTeams(ctx, tx, _teams); e != nil {
				return e
			}
			return nil
		})
	return err
}

// UpdateTeams is
func (s *TeamsUsecase) UpdateTeams(ctx context.Context, teams []*Team) error {
	if err := s.validate(false, teams); err != nil {
		return err
	}
	_teams, e := ToDBTeams(teams)
	if e != nil {
		return e
	}
	//return s.teamRepo.UpdateTeams(ctx, _teams)
	err := s.txm.RunInTX(
		func(tx repo.TX) error {
			if err := s.enforce(ctx, tx); err != nil {
				return err
			}
			if e := s.teamRepo.UpdateTeams(ctx, tx, _teams); e != nil {
				return e
			}
			return nil
		})
	return err
}

// DeleteTeams is
func (s *TeamsUsecase) DeleteTeams(ctx context.Context,
	ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}

	return s.txm.RunInTX(
		func(tx repo.TX) error {

			if err := s.enforce(ctx, tx); err != nil {
				return err
			}

			for _, r := range s.required {
				c, err := r.inst.CountRequire(ctx, tx, repo.RequireTeam, ids)
				if err != nil {
					return err
				}
				if c > 0 {
					return fmt.Errorf("some %s requires", r.name)
				}
			}
			if e := s.teamRepo.DeleteTeams(ctx, tx, ids); e != nil {
				return e
			}
			return nil
		})
}

// GetTeams is
func (s *TeamsUsecase) GetTeams(ctx context.Context, id uint32) (*Team, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	dbt, e := s.teamRepo.GetTeams(ctx, id)
	if e != nil {
		return nil, e
	}
	return ToBizTeam(dbt)
}

// ListTeams is
func (s *TeamsUsecase) ListTeams(ctx context.Context,
	filter *ListTeamsFilter) ([]*Team, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	teams, e := s.teamRepo.ListTeams(ctx, nil, ToDBTeamsFilter(filter))

	if e != nil {
		return nil, e
	}
	return ToBizTeams(teams)
}
