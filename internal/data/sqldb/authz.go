package sqldb

import (
	"appix/internal/conf"
	"appix/internal/data/repo"
	"context"
	"errors"
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthzRepoGorm struct {
	data *DataGorm
	log  *log.Helper
	conf *conf.Authz
}

func NewAuthzRepoGorm(conf *conf.Authz, data *DataGorm, logger log.Logger) (repo.AuthzRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}
	return &AuthzRepoGorm{
		data: data,
		conf: conf,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *AuthzRepoGorm) createEnforcer(ctx context.Context, tx repo.TX) (*casbin.Enforcer, error) {

	db := d.data.WithTX(tx).WithContext(ctx)
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	return casbin.NewEnforcer(d.conf.ModelFile, adapter)
}

func (d *AuthzRepoGorm) CreateRule(ctx context.Context, tx repo.TX, rule *repo.Rule) error {

	enforcer, err := d.createEnforcer(ctx, tx)
	if err != nil {
		return errors.Join(fmt.Errorf("CreatePolicy failed"), err)
	}
	res := rule.Resource.ResourceStr()
	s, err := enforcer.AddPolicy(rule.Sub, res, rule.Action)
	if err != nil {
		return errors.Join(fmt.Errorf("CreatePolicy failed"), err)
	}
	if !s {
		d.log.Warnf("CreatePolicy existed %v", rule)
	}

	return nil
}

func (d *AuthzRepoGorm) DeleteRule(ctx context.Context, tx repo.TX, rule *repo.Rule) error {
	enforcer, err := d.createEnforcer(ctx, tx)
	if err != nil {
		return errors.Join(fmt.Errorf("DeletePolicy failed"), err)
	}

	res := rule.Resource.ResourceStr()
	s, err := enforcer.RemovePolicy(rule.Sub, res, rule.Action)
	if err != nil {
		return errors.Join(fmt.Errorf("DeletePolicy failed with value %v", s), err)
	}

	return nil
}

func (d *AuthzRepoGorm) ListRule(ctx context.Context, tx repo.TX, filter *repo.RuleFilter) ([]*repo.Rule, error) {
	enforcer, err := d.createEnforcer(ctx, tx)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("ListPolicy failed"), err)
	}

	var _rules [][]string
	if filter == nil || filter.Sub == "" {
		_rules, err = enforcer.GetPolicy()
	} else {
		if filter.Sub != "" {
			_rules, err = enforcer.GetFilteredPolicy(0, filter.Sub)
		} else {
			_rules, err = enforcer.GetPolicy()
		}
	}
	if err != nil {
		return nil, errors.Join(fmt.Errorf("ListPolicy failed"), err)
	}

	rules := make([]*repo.Rule, len(_rules))
	for i, _rule := range _rules {
		ires := &repo.Resource4Sv1{}
		if err := ires.ParseStr(_rule[1]); err != nil {
			return nil, errors.Join(fmt.Errorf("ListRule failed"), err)
		}
		rules[i] = &repo.Rule{
			Sub:      _rule[0],
			Resource: ires,
			Action:   _rule[2],
		}
	}
	return rules, nil
}

func (d *AuthzRepoGorm) Enforce(ctx context.Context, tx repo.TX, request *repo.AuthenRequest) (bool, error) {
	enforcer, err := d.createEnforcer(ctx, tx)
	if err != nil {
		return false, errors.Join(fmt.Errorf("Enforce failed"), err)
	}
	return enforcer.Enforce(request.Sub, request.Resource.ResourceStr(), request.Action)
}

func (d *AuthzRepoGorm) CreateGroup(ctx context.Context, tx repo.TX, group *repo.Group) error {
	enforcer, err := d.createEnforcer(ctx, tx)
	if err != nil {
		return errors.Join(fmt.Errorf("CreateGroup failed"), err)
	}
	existed, err := enforcer.AddGroupingPolicy(group.User, group.Role)
	if err != nil {
		return errors.Join(fmt.Errorf("CreateGroup failed"), err)
	}
	if !existed {
		d.log.Warnf("CreateGroup existed %v", group)
	}
	return nil
}

func (d *AuthzRepoGorm) DeleteGroup(ctx context.Context, tx repo.TX, group *repo.Group) error {
	enforcer, err := d.createEnforcer(ctx, tx)
	if err != nil {
		return errors.Join(fmt.Errorf("DeleteGroup failed"), err)
	}
	s, err := enforcer.RemoveGroupingPolicy(group.User, group.Role)
	if err != nil || !s {
		return errors.Join(fmt.Errorf("DeleteGroup failed with value %v", s), err)
	}
	return nil
}

func (d *AuthzRepoGorm) ListGroup(ctx context.Context, tx repo.TX, filter *repo.GroupFilter) ([]*repo.Group, error) {
	enforcer, err := d.createEnforcer(ctx, tx)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("ListGroup failed"), err)
	}

	var _groups [][]string
	if filter == nil || filter.User == "" && filter.Role == "" {
		_groups, err = enforcer.GetGroupingPolicy()
	} else if filter.User != "" {
		_groups, err = enforcer.GetFilteredGroupingPolicy(0, filter.User)
	} else if filter.Role != "" {
		_groups, err = enforcer.GetFilteredGroupingPolicy(1, filter.Role)
	}

	if err != nil {
		return nil, errors.Join(fmt.Errorf("ListGroup failed"), err)
	}

	groups := make([]*repo.Group, len(_groups))
	for i, _group := range _groups {
		if len(_group) < 2 {
			d.log.Errorf("Invalid grouping policy: %v", _group)
			continue
		}
		group := &repo.Group{
			User: _group[0],
			Role: _group[1],
		}
		groups[i] = group
	}

	return groups, nil
}
