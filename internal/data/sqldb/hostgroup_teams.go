package sqldb

import (
	"appix/internal/data/repo"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type HostgroupTeamsRepoGorm struct {
	data *DataGorm
	log  *log.Helper
}

func NewHostgroupTeamsRepoGorm(data *DataGorm, logger log.Logger) (repo.HostgroupTeamsRepo, error) {
	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.DB, &repo.HostgroupTeam{}, repo.HostgroupTeamTable); err != nil {
		return nil, err
	}
	return &HostgroupTeamsRepoGorm{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *HostgroupTeamsRepoGorm) CreateHostgroupTeams(ctx context.Context,
	tx repo.TX,
	hostgroups []*repo.HostgroupTeam) error {
	if len(hostgroups) == 0 {
		return nil
	}
	return d.data.DB.WithContext(ctx).Create(hostgroups).Error
}

func (d *HostgroupTeamsRepoGorm) UpdateHostgroupTeams(ctx context.Context,
	tx repo.TX,
	hostgroups []*repo.HostgroupTeam) error {
	if len(hostgroups) == 0 {
		return nil
	}
	return d.data.DB.WithContext(ctx).Updates(hostgroups).Error
}

func (d *HostgroupTeamsRepoGorm) DeleteHostgroupTeams(ctx context.Context,
	tx repo.TX,
	ids []uint32) error {
	if len(ids) == 0 {
		return nil
	}
	return d.data.DB.WithContext(ctx).Delete(&repo.HostgroupTeam{}, ids).Error
}
func (d *HostgroupTeamsRepoGorm) ListHostgroupTeams(ctx context.Context,
	tx repo.TX,
	filter *repo.HostgroupTeamsFilter) ([]*repo.HostgroupTeam, error) {

	query := d.data.WithTX(tx).WithContext(ctx).Model(&repo.HostgroupTeam{})
	if len(filter.Ids) > 0 {
		query = query.Where("id in (?)", filter.Ids)
	}
	if len(filter.HostgroupIds) > 0 {
		query = query.Where("hostgroup_id in (?)", filter.HostgroupIds)
	}
	if len(filter.TeamIds) > 0 {
		query = query.Where("feature_id in (?)", filter.TeamIds)
	}
	if len(filter.KVs) > 0 {
		s_q, kvs := buildOrKV("hostgroup_id", "feature_id", filter.KVs)
		query = query.Where(s_q, kvs)
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := int(filter.PageSize * (filter.Page - 1))
		query = query.Offset(offset).Limit(int(filter.PageSize))
	}

	var hostgroups []*repo.HostgroupTeam
	if err := query.Find(&hostgroups).Error; err != nil {
		return nil, err
	}

	return hostgroups, nil
}
