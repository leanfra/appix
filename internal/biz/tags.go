package biz

import (
	"appix/internal/data"
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type TagsUsecase struct {
	tagsrepo  repo.TagsRepo
	authzrepo repo.AuthzRepo
	txm       repo.TxManager
	log       *log.Helper
	required  []requiredBy
}

func NewTagsUsecase(repo repo.TagsRepo,
	authzrepo repo.AuthzRepo,
	logger log.Logger,
	apptagrepo repo.AppTagsRepo,
	hgtagrepo repo.HostgroupTagsRepo,
	txm repo.TxManager) *TagsUsecase {

	return &TagsUsecase{
		tagsrepo:  repo,
		authzrepo: authzrepo,
		log:       log.NewHelper(logger),
		txm:       txm,
		required: []requiredBy{
			{inst: apptagrepo, name: "app_tag"},
			{inst: hgtagrepo, name: "hostgroup_tag"},
		},
	}
}

func (s *TagsUsecase) validate(isNew bool, tags []*Tag) error {
	for _, t := range tags {
		if err := t.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

func (s *TagsUsecase) enforce(ctx context.Context, tx repo.TX) error {
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

// CreateTags is
func (s *TagsUsecase) CreateTags(ctx context.Context, tags []*Tag) error {
	// validate tags
	if err := s.validate(true, tags); err != nil {
		return err
	}

	_tags, e := ToDBTags(tags)
	if e != nil {
		return e
	}

	err := s.txm.RunInTX(
		func(tx repo.TX) error {
			if err := s.enforce(ctx, tx); err != nil {
				return err
			}
			if e := s.tagsrepo.CreateTags(ctx, tx, _tags); e != nil {
				return e
			}
			return nil
		})
	return err
}

// UpdateTags is
func (s *TagsUsecase) UpdateTags(ctx context.Context, tags []*Tag) error {

	if err := s.validate(false, tags); err != nil {
		return err
	}
	_tags, e := ToDBTags(tags)
	if e != nil {
		return e
	}
	//return s.tagsrepo.UpdateTags(ctx, _tags)
	err := s.txm.RunInTX(
		func(tx repo.TX) error {
			if err := s.enforce(ctx, tx); err != nil {
				return err
			}
			if e := s.tagsrepo.UpdateTags(ctx, tx, _tags); e != nil {
				return e
			}
			return nil
		})
	return err
}

// DeleteTags is
func (s *TagsUsecase) DeleteTags(ctx context.Context, ids []uint32) error {

	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx); err != nil {
			return err
		}
		for _, r := range s.required {
			c, err := r.inst.CountRequire(ctx, tx, repo.RequireTag, ids)
			if err != nil {
				return err
			}
			if c > 0 {
				return fmt.Errorf("some %s requires", r.name)
			}
		}
		if e := s.tagsrepo.DeleteTags(ctx, tx, ids); e != nil {
			return e
		}
		return nil
	})
}

// GetTags is
func (s *TagsUsecase) GetTags(ctx context.Context, id uint32) (*Tag, error) {
	if id <= 0 {
		return nil, fmt.Errorf("EmptyId")
	}
	t, e := s.tagsrepo.GetTags(ctx, id)
	if e != nil {
		return nil, e
	}
	return ToBizTag(t)
}

// ListTags is
func (s *TagsUsecase) ListTags(ctx context.Context,
	filter *ListTagsFilter) ([]*Tag, error) {
	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}
	_ts, e := s.tagsrepo.ListTags(ctx, nil, ToDBTagsFilter(filter))
	if e != nil {
		return nil, e
	}
	return ToBizTags(_ts)
}
