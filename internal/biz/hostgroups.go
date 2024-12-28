package biz

import (
	"appix/internal/data/repo"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type HostgroupsUsecase struct {
	txm       repo.TxManager
	hgrepo    repo.HostgroupsRepo
	hteamrepo repo.HostgroupTeamsRepo
	hprepo    repo.HostgroupProductsRepo
	htagrepo  repo.HostgroupTagsRepo
	hfrepo    repo.HostgroupFeaturesRepo

	clsrepo  repo.ClustersRepo
	dcrepo   repo.DatacentersRepo
	prdrepo  repo.ProductsRepo
	teamrepo repo.TeamsRepo
	ftrepo   repo.FeaturesRepo
	tagrepo  repo.TagsRepo
	envrepo  repo.EnvsRepo

	log *log.Helper
}

func NewHostgroupsUsecase(repo repo.HostgroupsRepo,
	hteamrepo repo.HostgroupTeamsRepo,
	hprepo repo.HostgroupProductsRepo,
	htagrepo repo.HostgroupTagsRepo,
	hfrepo repo.HostgroupFeaturesRepo,
	clsrepo repo.ClustersRepo,
	dcrepo repo.DatacentersRepo,
	envrepo repo.EnvsRepo,
	ftrepo repo.FeaturesRepo,
	tagrepo repo.TagsRepo,
	teamrepo repo.TeamsRepo,
	prdrepo repo.ProductsRepo,
	logger log.Logger,
	txm repo.TxManager) *HostgroupsUsecase {

	return &HostgroupsUsecase{
		hgrepo:    repo,
		hteamrepo: hteamrepo,
		hprepo:    hprepo,
		htagrepo:  htagrepo,
		hfrepo:    hfrepo,
		clsrepo:   clsrepo,
		dcrepo:    dcrepo,
		prdrepo:   prdrepo,
		teamrepo:  teamrepo,
		ftrepo:    ftrepo,
		tagrepo:   tagrepo,
		envrepo:   envrepo,
		log:       log.NewHelper(logger),
		txm:       txm,
	}
}

func (s *HostgroupsUsecase) validate(isNew bool, hgs []*Hostgroup) error {
	for _, hg := range hgs {
		if err := hg.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

const hgPropCluster = "cluster"
const hgPropDatacenter = "datacenter"
const hgPropEnv = "env"
const hgPropProduct = "product"
const hgPropTeam = "team"
const hgPropFeature = "feature"
const hgPropTag = "tag"
const hgPropShareProduct = "shareProduct"
const hgPropShareTeam = "shareTeam"

func hostgroupPropFilter(hgs []*Hostgroup, prop string) repo.CountFilter {
	var ids []uint32
	switch prop {
	case hgPropCluster:
		for _, a := range hgs {
			ids = append(ids, a.ClusterId)
		}
		ids = DedupSliceUint32(ids)
		return &repo.ClustersFilter{
			Ids: ids,
		}
	case hgPropDatacenter:
		for _, a := range hgs {
			ids = append(ids, a.DatacenterId)
		}
		ids = DedupSliceUint32(ids)
		return &repo.DatacentersFilter{
			Ids: ids,
		}
	case hgPropEnv:
		for _, a := range hgs {
			ids = append(ids, a.EnvId)
		}
		ids = DedupSliceUint32(ids)
		return &repo.EnvsFilter{
			Ids: ids,
		}
	case hgPropProduct:
		for _, a := range hgs {
			ids = append(ids, a.ProductId)
		}
		ids = DedupSliceUint32(ids)
		return &repo.ProductsFilter{
			Ids: ids,
		}
	case hgPropTeam:
		for _, a := range hgs {
			ids = append(ids, a.TeamId)
		}
		ids = DedupSliceUint32(ids)
		return &repo.TeamsFilter{
			Ids: ids,
		}
	case hgPropFeature:
		for _, a := range hgs {
			ids = append(ids, a.FeaturesId...)
		}
		ids = DedupSliceUint32(ids)
		return &repo.FeaturesFilter{
			Ids: ids,
		}
	case hgPropTag:
		for _, a := range hgs {
			ids = append(ids, a.TagsId...)
		}
		ids = DedupSliceUint32(ids)
		return &repo.TagsFilter{
			Ids: ids,
		}
	case hgPropShareProduct:
		for _, a := range hgs {
			ids = append(ids, a.ShareProductsId...)
		}
		ids = DedupSliceUint32(ids)
		return &repo.ProductsFilter{
			Ids: ids,
		}
	case hgPropShareTeam:
		for _, a := range hgs {
			ids = append(ids, a.ShareTeamsId...)
		}
		ids = DedupSliceUint32(ids)
		return &repo.TeamsFilter{
			Ids: ids,
		}
	}
	return nil
}
func (s *HostgroupsUsecase) validateProps(
	ctx context.Context,
	tx repo.TX,
	hgs []*Hostgroup) error {
	type propsCount struct {
		name    string
		ids     repo.CountFilter
		countFn func(context.Context, repo.TX, repo.CountFilter) (int64, error)
	}
	counters := []propsCount{
		{hgPropCluster, hostgroupPropFilter(hgs, hgPropCluster), s.clsrepo.CountClusters},
		{hgPropDatacenter, hostgroupPropFilter(hgs, hgPropDatacenter), s.dcrepo.CountDatacenters},
		{hgPropEnv, hostgroupPropFilter(hgs, hgPropEnv), s.envrepo.CountEnvs},
		{hgPropProduct, hostgroupPropFilter(hgs, hgPropProduct), s.prdrepo.CountProducts},
		{hgPropTeam, hostgroupPropFilter(hgs, hgPropTeam), s.teamrepo.CountTeams},
		{hgPropFeature, hostgroupPropFilter(hgs, hgPropFeature), s.ftrepo.CountFeatures},
		{hgPropTag, hostgroupPropFilter(hgs, hgPropTag), s.tagrepo.CountTags},
		{hgPropShareProduct, hostgroupPropFilter(hgs, hgPropShareProduct), s.prdrepo.CountProducts},
		{hgPropShareTeam, hostgroupPropFilter(hgs, hgPropShareTeam), s.teamrepo.CountTeams},
	}
	for _, counter := range counters {
		if counter.ids == nil {
			return fmt.Errorf("invalid %s", counter.name)
		}
		if count, err := counter.countFn(ctx, tx, counter.ids); err != nil {
			return err
		} else {
			if count == 0 {
				return fmt.Errorf("invalid %s", counter.name)
			}
		}
	}
	return nil
}

func (s *HostgroupsUsecase) createM2MProps(
	ctx context.Context, tx repo.TX, hgid uint32, ids []uint32, prop string) error {

	switch prop {
	case hgPropTag:
		_hg_tags := make([]*repo.HostgroupTag, 0, len(ids))
		for i, id := range ids {
			_hg_tags[i] = &repo.HostgroupTag{
				HostgroupID: hgid,
				TagID:       id,
			}
		}
		return s.htagrepo.CreateHostgroupTags(ctx, tx, _hg_tags)
	case hgPropFeature:
		_hg_features := make([]*repo.HostgroupFeature, 0, len(ids))
		for i, id := range ids {
			_hg_features[i] = &repo.HostgroupFeature{
				HostgroupID: hgid,
				FeatureID:   id,
			}
		}
		return s.hfrepo.CreateHostgroupFeatures(ctx, tx, _hg_features)
	case hgPropShareProduct:
		_hg_products := make([]*repo.HostgroupProduct, 0, len(ids))
		for i, id := range ids {
			_hg_products[i] = &repo.HostgroupProduct{
				HostgroupID: hgid,
				ProductID:   id,
			}
		}
		return s.hprepo.CreateHostgroupProducts(ctx, tx, _hg_products)
	case hgPropShareTeam:
		_hg_teams := make([]*repo.HostgroupTeam, 0, len(ids))
		for i, id := range ids {
			_hg_teams[i] = &repo.HostgroupTeam{
				HostgroupID: hgid,
				TeamID:      id,
			}
		}
		return s.hteamrepo.CreateHostgroupTeams(ctx, tx, _hg_teams)
	}
	return fmt.Errorf("createProps invalid prop %s", prop)
}

// CreateHostgroups is
func (s *HostgroupsUsecase) CreateHostgroups(ctx context.Context, hgs []*Hostgroup) error {
	if err := s.validate(true, hgs); err != nil {
		return err
	}
	dbhgs, err := ToDBHostgroups(hgs)
	if err != nil {
		return err
	}
	return s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.validateProps(ctx, tx, hgs); err != nil {
			return err
		}

		if err := s.hgrepo.CreateHostgroups(ctx, tx, dbhgs); err != nil {
			return err
		}

		for _, hg := range hgs {
			if err := s.createM2MProps(ctx, tx, hg.Id, hg.TagsId, hgPropTag); err != nil {
				return err
			}
			if err := s.createM2MProps(ctx, tx, hg.Id, hg.FeaturesId, hgPropFeature); err != nil {
				return err
			}
			if err := s.createM2MProps(ctx, tx, hg.Id, hg.ShareProductsId, hgPropShareProduct); err != nil {
				return err
			}
			if err := s.createM2MProps(ctx, tx, hg.Id, hg.ShareTeamsId, hgPropShareTeam); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *HostgroupsUsecase) listM2MProps(
	ctx context.Context, tx repo.TX, hgids []uint32, prop string) (interface{}, error) {

	switch prop {
	case hgPropTag:
		return s.htagrepo.ListHostgroupTags(ctx, tx, &repo.HostgroupTagsFilter{
			HostgroupIds: hgids,
		})
	case hgPropFeature:
		return s.hfrepo.ListHostgroupFeatures(ctx, tx, &repo.HostgroupFeaturesFilter{
			HostgroupIds: hgids,
		})
	case hgPropShareProduct:
		return s.hprepo.ListHostgroupProducts(ctx, tx, &repo.HostgroupProductsFilter{
			HostgroupIds: hgids,
		})
	case hgPropShareTeam:
		return s.hteamrepo.ListHostgroupTeams(ctx, tx, &repo.HostgroupTeamsFilter{
			HostgroupIds: hgids,
		})
	}
	return nil, fmt.Errorf("listProps invalid prop %s", prop)
}

func (s *HostgroupsUsecase) deleteM2MProps(
	ctx context.Context, tx repo.TX, ids []uint32, prop string) error {

	switch prop {
	case hgPropTag:
		return s.htagrepo.DeleteHostgroupTags(ctx, tx, ids)
	case hgPropFeature:
		return s.hfrepo.DeleteHostgroupFeatures(ctx, tx, ids)
	case hgPropShareProduct:
		return s.hprepo.DeleteHostgroupProducts(ctx, tx, ids)
	case hgPropShareTeam:
		return s.hteamrepo.DeleteHostgroupTeams(ctx, tx, ids)
	}
	return fmt.Errorf("deleteProps invalid prop %s", prop)
}

func (s *HostgroupsUsecase) deleteM2MPropsByHostgroup(
	ctx context.Context, tx repo.TX, hgid uint32, prop string) error {

	switch prop {
	case hgPropTag:
		return s.htagrepo.DeleteHostgroupTags(ctx, tx, []uint32{hgid})
	case hgPropFeature:
		return s.hfrepo.DeleteHostgroupFeatures(ctx, tx, []uint32{hgid})
	case hgPropShareProduct:
		return s.hprepo.DeleteHostgroupProducts(ctx, tx, []uint32{hgid})
	case hgPropShareTeam:
		return s.hteamrepo.DeleteHostgroupTeams(ctx, tx, []uint32{hgid})
	}
	return fmt.Errorf("deleteProps invalid prop %s", prop)
}

func (s *HostgroupsUsecase) handleM2MProps(
	ctx context.Context, tx repo.TX, hgid uint32, ids []uint32, prop string) error {

	oldItems, err := s.listM2MProps(ctx, tx, []uint32{hgid}, prop)
	if err != nil {
		return err
	}
	var oldIds []uint32

	switch items := oldItems.(type) {
	case []*repo.AppTag:
		for _, item := range items {
			oldIds = append(oldIds, item.Id)
		}
	case []*repo.AppFeature:
		for _, item := range items {
			oldIds = append(oldIds, item.Id)
		}
	case []*repo.AppHostgroup:
		for _, item := range items {
			oldIds = append(oldIds, item.Id)
		}
	}

	toDelIds := DiffSliceUint32(oldIds, ids)
	toNewids := DiffSliceUint32(ids, oldIds)

	if len(toDelIds) > 0 {
		if err := s.deleteM2MProps(ctx, tx, toDelIds, prop); err != nil {
			return err
		}
	}
	if len(toNewids) > 0 {
		if err := s.createM2MProps(ctx, tx, hgid, toNewids, prop); err != nil {
			return err
		}
	}
	return nil
}

// UpdateHostgroups is
func (s *HostgroupsUsecase) UpdateHostgroups(ctx context.Context, hgs []*Hostgroup) error {
	if err := s.validate(false, hgs); err != nil {
		return err
	}
	_hgs, err := ToDBHostgroups(hgs)
	if err != nil {
		return err
	}

	return s.txm.RunInTX(func(tx repo.TX) error {
		if s.validateProps(ctx, tx, hgs); err != nil {
			return err
		}
		if err := s.hgrepo.UpdateHostgroups(ctx, tx, _hgs); err != nil {
			return err
		}

		for _, hg := range hgs {
			if err := s.handleM2MProps(ctx, tx,
				hg.Id, hg.TagsId, hgPropTag); err != nil {
				return err
			}
			if err := s.handleM2MProps(ctx, tx,
				hg.Id, hg.FeaturesId, hgPropFeature); err != nil {
				return err
			}
			if err := s.handleM2MProps(ctx, tx,
				hg.Id, hg.ShareProductsId, hgPropShareProduct); err != nil {
				return err
			}
			if err := s.handleM2MProps(ctx, tx,
				hg.Id, hg.ShareTeamsId, hgPropShareTeam); err != nil {
				return err
			}
		}
		return nil
	})
}

// DeleteHostgroups is
func (s *HostgroupsUsecase) DeleteHostgroups(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}
	return s.txm.RunInTX(func(tx repo.TX) error {
		for _, id := range ids {
			if err := s.deleteM2MPropsByHostgroup(ctx, tx, id, hgPropTag); err != nil {
				return err
			}
			if err := s.deleteM2MPropsByHostgroup(ctx, tx, id, hgPropFeature); err != nil {
				return err
			}
			if err := s.deleteM2MPropsByHostgroup(ctx, tx, id, hgPropShareProduct); err != nil {
				return err
			}
			if err := s.deleteM2MPropsByHostgroup(ctx, tx, id, hgPropShareTeam); err != nil {
				return err
			}
		}
		return s.hgrepo.DeleteHostgroups(ctx, tx, ids)
	})
}

// GetHostgroups is
func (s *HostgroupsUsecase) GetHostgroups(ctx context.Context, id uint32) (*Hostgroup, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	hg, err := s.hgrepo.GetHostgroups(ctx, id)
	if err != nil {
		return nil, err
	}
	bizhg, e := ToBizHostgroup(hg)
	if e != nil {
		return nil, e
	}
	if err := s.attachM2MProps(ctx, bizhg); err != nil {
		return nil, err
	}
	return bizhg, nil
}

func (s *HostgroupsUsecase) attachM2MProps(ctx context.Context, hg *Hostgroup) error {
	// tags id
	_tags, err := s.listM2MProps(ctx, nil, []uint32{hg.Id}, hgPropTag)
	if err != nil {
		return err
	}
	for _, tag := range _tags.([]*repo.AppTag) {
		hg.TagsId = append(hg.TagsId, tag.Id)
	}
	// features id
	_fts, err := s.listM2MProps(ctx, nil, []uint32{hg.Id}, hgPropFeature)
	if err != nil {
		return err
	}
	for _, ft := range _fts.([]*repo.AppFeature) {
		hg.FeaturesId = append(hg.FeaturesId, ft.Id)
	}

	// share products id
	_prds, err := s.listM2MProps(ctx, nil, []uint32{hg.Id}, hgPropShareProduct)
	if err != nil {
		return err
	}
	for _, prd := range _prds.([]*repo.AppHostgroup) {
		hg.ShareProductsId = append(hg.ShareProductsId, prd.Id)
	}
	// share teams id
	_teams, err := s.listM2MProps(ctx, nil, []uint32{hg.Id}, hgPropShareTeam)
	if err != nil {
		return err
	}
	for _, team := range _teams.([]*repo.AppHostgroup) {
		hg.ShareTeamsId = append(hg.ShareTeamsId, team.Id)
	}

	return nil
}

// ListHostgroups is
func (s *HostgroupsUsecase) ListHostgroups(
	ctx context.Context, filter *ListHostgroupsFilter) ([]*Hostgroup, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}

	dbFilter := ToDBHostgroupsFilter(filter)

	processInitIds := func(filterIds []uint32, prop string) error {
		if len(filterIds) == 0 {
			return nil
		}
		var items interface{}
		var err error
		switch prop {
		case hgPropTag:
			items, err = s.htagrepo.ListHostgroupTags(ctx, nil, &repo.HostgroupTagsFilter{
				Ids: filter.TagsId})
		case hgPropFeature:
			items, err = s.hfrepo.ListHostgroupFeatures(ctx, nil, &repo.HostgroupFeaturesFilter{
				Ids: filter.FeaturesId})
		case hgPropShareProduct:
			items, err = s.hprepo.ListHostgroupProducts(ctx, nil, &repo.HostgroupProductsFilter{
				Ids: filter.ShareProductsId})
		case hgPropShareTeam:
			items, err = s.hteamrepo.ListHostgroupTeams(ctx, nil, &repo.HostgroupTeamsFilter{
				Ids: filter.ShareTeamsId})
		}
		if err != nil {
			return err
		}

		var hg_ids []uint32
		switch v := items.(type) {
		case []*repo.HostgroupTag:
			for _, item := range v {
				hg_ids = append(hg_ids, item.HostgroupID)
			}
		case []*repo.HostgroupFeature:
			for _, item := range v {
				hg_ids = append(hg_ids, item.HostgroupID)
			}
		case []*repo.HostgroupProduct:
			for _, item := range v {
				hg_ids = append(hg_ids, item.HostgroupID)
			}
		case []*repo.HostgroupTeam:
			for _, item := range v {
				hg_ids = append(hg_ids, item.HostgroupID)
			}
		}
		if len(hg_ids) == 0 {
			return fmt.Errorf("ListHostgroups no hostgroup with %s", prop)
		}
		if len(dbFilter.Ids) == 0 {
			dbFilter.Ids = hg_ids
		} else {
			dbFilter.Ids = IntersectSliceUint32(dbFilter.Ids, hg_ids)
			if len(dbFilter.Ids) == 0 {
				return fmt.Errorf("ListHostgroups no hostgroup intersect with %s", prop)
			}
		}
		return nil
	}

	if filter != nil {
		// tags id
		if err := processInitIds(filter.TagsId, hgPropTag); err != nil {
			return nil, err
		}
		// features id
		if err := processInitIds(filter.FeaturesId, hgPropFeature); err != nil {
			return nil, err
		}
		// share products id
		if err := processInitIds(filter.ShareProductsId, hgPropShareProduct); err != nil {
			return nil, err
		}
		// share teams id
		if err := processInitIds(filter.ShareTeamsId, hgPropShareTeam); err != nil {
			return nil, err
		}
	}
	hg, err := s.hgrepo.ListHostgroups(ctx, nil, dbFilter)
	if err != nil {
		return nil, err
	}
	bizhg, err := ToBizHostgroups(hg)
	if err != nil {
		return nil, err
	}
	for _, hg := range bizhg {
		if err := s.attachM2MProps(ctx, hg); err != nil {
			return nil, err
		}
	}

	return bizhg, nil
}
