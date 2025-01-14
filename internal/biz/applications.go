package biz

import (
	"appix/internal/data"
	"appix/internal/data/repo"
	"context"
	"errors"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

const IsStatefulTrue = "true"
const IsStatefulFalse = "false"
const IsStatefulNone = ""

type ApplicationsUsecase struct {
	apprepo   repo.ApplicationsRepo
	atagrepo  repo.AppTagsRepo
	afrepo    repo.AppFeaturesRepo
	ahgrepo   repo.AppHostgroupsRepo
	prdrepo   repo.ProductsRepo
	teamrepo  repo.TeamsRepo
	ftrepo    repo.FeaturesRepo
	tagrepo   repo.TagsRepo
	hgrepo    repo.HostgroupsRepo
	hfrepo    repo.HostgroupFeaturesRepo
	authzrepo repo.AuthzRepo
	adminrepo repo.AdminRepo
	log       *log.Helper
	txm       repo.TxManager
}

func NewApplicationsUsecase(
	apprepo repo.ApplicationsRepo,
	atagrepo repo.AppTagsRepo,
	afrepo repo.AppFeaturesRepo,
	ahgrepo repo.AppHostgroupsRepo,
	prdrepo repo.ProductsRepo,
	teamrepo repo.TeamsRepo,
	ftrepo repo.FeaturesRepo,
	tagrepo repo.TagsRepo,
	hgrepo repo.HostgroupsRepo,
	hfrepo repo.HostgroupFeaturesRepo,
	authzrepo repo.AuthzRepo,
	adminrepo repo.AdminRepo,
	logger log.Logger,
	txm repo.TxManager) *ApplicationsUsecase {

	return &ApplicationsUsecase{
		apprepo:   apprepo,
		atagrepo:  atagrepo,
		afrepo:    afrepo,
		ahgrepo:   ahgrepo,
		prdrepo:   prdrepo,
		teamrepo:  teamrepo,
		ftrepo:    ftrepo,
		tagrepo:   tagrepo,
		hgrepo:    hgrepo,
		hfrepo:    hfrepo,
		authzrepo: authzrepo,
		adminrepo: adminrepo,
		log:       log.NewHelper(logger),
		txm:       txm,
	}
}

func (s *ApplicationsUsecase) validate(isNew bool, apps []*Application) error {
	for _, a := range apps {
		if err := a.Validate(isNew); err != nil {
			return err
		}
	}
	return nil
}

const appPropProduct = "product"
const appPropTeam = "team"
const appPropFeature = "feature"
const appPropTag = "tag"
const appPropHostgroup = "hostgroup"
const appPropOwner = "user"

func appPropFilter(apps []*Application, prop string) repo.CountFilter {
	var ids []uint32
	switch prop {
	case appPropProduct:
		for _, a := range apps {
			ids = append(ids, a.ProductId)
		}
		ids = DedupSliceUint32(ids)
		return &repo.ProductsFilter{
			Ids: ids,
		}
	case appPropTeam:
		for _, a := range apps {
			ids = append(ids, a.TeamId)
		}
		ids = DedupSliceUint32(ids)
		return &repo.TeamsFilter{
			Ids: ids,
		}
	case appPropFeature:
		for _, a := range apps {
			ids = append(ids, a.FeaturesId...)
		}
		ids = DedupSliceUint32(ids)
		return &repo.FeaturesFilter{
			Ids: ids,
		}
	case appPropTag:
		for _, a := range apps {
			ids = append(ids, a.TagsId...)
		}
		ids = DedupSliceUint32(ids)
		return &repo.TagsFilter{
			Ids: ids,
		}
	case appPropHostgroup:
		for _, a := range apps {
			ids = append(ids, a.HostgroupsId...)
		}
		ids = DedupSliceUint32(ids)
		return &repo.HostgroupsFilter{
			Ids: ids,
		}
	case appPropOwner:
		for _, a := range apps {
			ids = append(ids, a.OwnerId)
		}
		ids = DedupSliceUint32(ids)
		return &repo.UsersFilter{
			Ids: ids,
		}
	}
	return nil
}

func (s *ApplicationsUsecase) validateProps(
	ctx context.Context,
	tx repo.TX,
	apps []*Application) error {
	type propsCount struct {
		name    string
		ids     repo.CountFilter
		countFn func(context.Context, repo.TX, repo.CountFilter) (int64, error)
	}
	counters := []propsCount{
		{appPropProduct, appPropFilter(apps, appPropProduct), s.prdrepo.CountProducts},
		{appPropTeam, appPropFilter(apps, appPropTeam), s.teamrepo.CountTeams},
		{appPropFeature, appPropFilter(apps, appPropFeature), s.ftrepo.CountFeatures},
		{appPropTag, appPropFilter(apps, appPropTag), s.tagrepo.CountTags},
		{appPropHostgroup, appPropFilter(apps, appPropHostgroup), s.hgrepo.CountHostgroups},
		{appPropOwner, appPropFilter(apps, appPropOwner), s.adminrepo.CountUsers},
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

// MatchHostgroups match hostgroups with application's features.
// hostgroups's features must be a superset of application's features.
// return hostgroup ids
func (s *ApplicationsUsecase) MatchHostgroups(
	ctx context.Context,
	tx repo.TX,
	filter *MatchAppHostgroupsFilter) (ids []uint32, err error) {

	hfs, err := s.hfrepo.ListHostgroupMatchFeatures(ctx, tx, &repo.HostgroupMatchFeaturesFilter{
		FeatureIds: filter.FeaturesId,
	})
	if err != nil {
		return nil, err
	}
	hostgroupfilter := &repo.HostgroupsFilter{
		Ids:        hfs,
		ProductsId: []uint32{filter.ProductId},
		TeamsId:    []uint32{filter.TeamId},
	}
	hgs, err := s.hgrepo.ListHostgroups(ctx, tx, hostgroupfilter)
	if err != nil {
		return nil, err
	}
	for _, hg := range hgs {
		ids = append(ids, hg.Id)
	}
	return ids, nil
}

func (s *ApplicationsUsecase) validateHostgroupMatch(
	ctx context.Context,
	tx repo.TX,
	app *Application) error {

	matched_hgs, err := s.MatchHostgroups(ctx, tx, &MatchAppHostgroupsFilter{
		FeaturesId: app.FeaturesId,
		ProductId:  app.ProductId,
		TeamId:     app.TeamId,
	})
	if err != nil {
		return fmt.Errorf("MatchHostgroups error. %w", err)
	}
	if len(matched_hgs) == 0 {
		return fmt.Errorf("no hostgroup matched application %s", app.Name)
	}
	// Convert matched_hgs to map for faster lookup
	matchedMap := make(map[uint32]bool)
	for _, id := range matched_hgs {
		matchedMap[id] = true
	}

	// Check if all app hostgroups are in matched hostgroups
	for _, hgid := range app.HostgroupsId {
		if !matchedMap[hgid] {
			return fmt.Errorf(
				"hostgroup %d not match application %s", hgid, app.Name)
		}
	}
	return nil
}

func (s *ApplicationsUsecase) enforce(ctx context.Context, tx repo.TX, apps []*Application) error {
	curUserName := ctx.Value(data.UserName).(string)
	for _, app := range apps {
		team, err := s.teamrepo.GetTeams(ctx, app.TeamId)
		if err != nil {
			return err
		}
		owner, err := s.adminrepo.GetUsers(ctx, tx, app.OwnerId)
		if err != nil {
			return err
		}
		ires := repo.NewResource4Sv1("app", team.Name, app.Name, owner.UserName)
		can, err := s.authzrepo.Enforce(ctx, tx, &repo.AuthenRequest{
			Sub:      curUserName,
			Resource: ires,
			Action:   repo.ActWrite,
		})
		if err != nil {
			return err
		}
		if !can {
			return errors.New("no permission")
		}
	}
	return nil
}

// CreateApplications is
func (s *ApplicationsUsecase) CreateApplications(ctx context.Context, apps []*Application) error {
	if err := s.validate(true, apps); err != nil {
		return err
	}

	return s.txm.RunInTX(func(tx repo.TX) error {

		if err := s.validateProps(ctx, tx, apps); err != nil {
			return err
		}

		// insert
		for _, app := range apps {
			if err := s.validateHostgroupMatch(ctx, tx, app); err != nil {
				return err
			}

			dbapp, err := ToDBApplication(app)
			if err != nil {
				return err
			}
			// create app and get id
			if err := s.apprepo.CreateApplications(ctx, tx, []*repo.Application{dbapp}); err != nil {
				return err
			}

			// create app-tags
			if err := s.createProps(ctx, tx, dbapp.Id, app.TagsId, appPropTag); err != nil {
				return err
			}

			// create app-features
			if err := s.createProps(ctx, tx, dbapp.Id, app.FeaturesId, appPropFeature); err != nil {
				return err
			}

			// create app-hostgroups
			if err := s.createProps(ctx, tx, dbapp.Id, app.HostgroupsId, appPropHostgroup); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *ApplicationsUsecase) listProps(
	ctx context.Context, tx repo.TX, appids []uint32, prop string) (interface{}, error) {

	switch prop {
	case appPropTag:
		return s.atagrepo.ListAppTags(ctx, tx, &repo.AppTagsFilter{
			AppIds: appids,
		})
	case appPropFeature:
		return s.afrepo.ListAppFeatures(ctx, tx, &repo.AppFeaturesFilter{
			AppIds: appids,
		})
	case appPropHostgroup:
		return s.ahgrepo.ListAppHostgroups(ctx, tx, &repo.AppHostgroupsFilter{
			AppIds: appids,
		})
	}
	return nil, fmt.Errorf("listProps invalid prop %s", prop)
}

func (s *ApplicationsUsecase) deleteProps(
	ctx context.Context, tx repo.TX, ids []uint32, prop string) error {

	switch prop {
	case appPropTag:
		return s.atagrepo.DeleteAppTags(ctx, tx, ids)
	case appPropFeature:
		return s.afrepo.DeleteAppFeatures(ctx, tx, ids)
	case appPropHostgroup:
		return s.ahgrepo.DeleteAppHostgroups(ctx, tx, ids)
	}
	return fmt.Errorf("deleteProps invalid prop %s", prop)
}

func (s *ApplicationsUsecase) createProps(
	ctx context.Context, tx repo.TX, appid uint32, ids []uint32, prop string) error {

	switch prop {
	case appPropTag:
		_app_tags := make([]*repo.AppTag, len(ids))
		for i, id := range ids {
			_app_tags[i] = &repo.AppTag{
				AppID: appid,
				TagID: id,
			}
		}
		return s.atagrepo.CreateAppTags(ctx, tx, _app_tags)
	case appPropFeature:
		_app_features := make([]*repo.AppFeature, len(ids))
		for i, id := range ids {
			_app_features[i] = &repo.AppFeature{
				AppID:     appid,
				FeatureID: id,
			}
		}
		return s.afrepo.CreateAppFeatures(ctx, tx, _app_features)
	case appPropHostgroup:
		_app_hostgroups := make([]*repo.AppHostgroup, len(ids))
		for i, id := range ids {
			_app_hostgroups[i] = &repo.AppHostgroup{
				AppID:       appid,
				HostgroupID: id,
			}
		}
		return s.ahgrepo.CreateAppHostgroups(ctx, tx, _app_hostgroups)
	}
	return fmt.Errorf("createProps invalid prop %s", prop)
}

func (s *ApplicationsUsecase) HandleM2MProps(
	ctx context.Context, tx repo.TX, appid uint32, ids []uint32, prop string) error {

	oldItems, err := s.listProps(ctx, tx, []uint32{appid}, prop)
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
		if err := s.deleteProps(ctx, tx, toDelIds, prop); err != nil {
			return err
		}
	}
	if len(toNewids) > 0 {
		if err := s.createProps(ctx, tx, appid, toNewids, prop); err != nil {
			return err
		}
	}
	return nil
}

// UpdateApplications is
func (s *ApplicationsUsecase) UpdateApplications(ctx context.Context, apps []*Application) error {
	if err := s.validate(false, apps); err != nil {
		return err
	}
	_apps, err := ToDBApplications(apps)
	if err != nil {
		return err
	}

	return s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforce(ctx, tx, apps); err != nil {
			return err
		}
		if err := s.validateProps(ctx, tx, apps); err != nil {
			return err
		}

		// update
		if err := s.apprepo.UpdateApplications(ctx, tx, _apps); err != nil {
			return err
		}

		for _, a := range apps {
			// tags
			if err := s.HandleM2MProps(ctx, tx, a.Id, a.TagsId, appPropTag); err != nil {
				return err
			}

			// features
			if err := s.HandleM2MProps(ctx, tx, a.Id, a.FeaturesId, appPropFeature); err != nil {
				return err
			}

			// hostgroups
			if err := s.HandleM2MProps(ctx, tx, a.Id, a.HostgroupsId, appPropHostgroup); err != nil {
				return err
			}
		}

		return nil
	})

}

// DeleteApplications is
func (s *ApplicationsUsecase) DeleteApplications(ctx context.Context, ids []uint32) error {
	if len(ids) == 0 {
		return fmt.Errorf("EmptyIds")
	}

	return s.txm.RunInTX(func(tx repo.TX) error {
		apps, err := s.apprepo.ListApplications(ctx, tx, &repo.ApplicationsFilter{
			Ids: ids,
		})
		if err != nil {
			return err
		}
		bizapps, err := ToBizApplications(apps)
		if err != nil {
			return err
		}
		if err := s.enforce(ctx, tx, bizapps); err != nil {
			return err
		}
		// delete props

		if err := s.atagrepo.DeleteAppTagsByAppId(ctx, tx, ids); err != nil {
			return err
		}
		if err := s.afrepo.DeleteAppFeaturesByAppId(ctx, tx, ids); err != nil {
			return err
		}
		if err := s.ahgrepo.DeleteAppHostgroupsByAppId(ctx, tx, ids); err != nil {
			return err
		}
		// delete app
		return s.apprepo.DeleteApplications(ctx, tx, ids)
	})
}

// GetApplications is
func (s *ApplicationsUsecase) GetApplications(ctx context.Context, id uint32) (*Application, error) {
	if id <= 0 {
		return nil, fmt.Errorf("InvalidId")
	}
	_app, err := s.apprepo.GetApplications(ctx, id)
	if err != nil {
		return nil, err
	}
	bapp, e := ToBizApplication(_app)
	if e != nil {
		return nil, e
	}

	if err := s.attachM2MProps(ctx, bapp); err != nil {
		return nil, err
	}

	return bapp, nil
}

func (s *ApplicationsUsecase) attachM2MProps(ctx context.Context, app *Application) error {
	// tags id
	_tags, err := s.listProps(ctx, nil, []uint32{app.Id}, appPropTag)
	if err != nil {
		return err
	}
	for _, tag := range _tags.([]*repo.AppTag) {
		app.TagsId = append(app.TagsId, tag.TagID)
	}
	// features id
	_fts, err := s.listProps(ctx, nil, []uint32{app.Id}, appPropFeature)
	if err != nil {
		return err
	}
	for _, ft := range _fts.([]*repo.AppFeature) {
		app.FeaturesId = append(app.FeaturesId, ft.FeatureID)
	}

	// hostgroups id
	_hgs, err := s.listProps(ctx, nil, []uint32{app.Id}, appPropHostgroup)
	if err != nil {
		return err
	}
	for _, hg := range _hgs.([]*repo.AppHostgroup) {
		app.HostgroupsId = append(app.HostgroupsId, hg.HostgroupID)
	}
	return nil
}

// ListApplications is
func (s *ApplicationsUsecase) ListApplications(
	ctx context.Context,
	filter *ListApplicationsFilter) ([]*Application, error) {

	if filter != nil {
		if err := filter.Validate(); err != nil {
			return nil, err
		}
	}

	dbFilter := ToDBApplicationsFilter(filter)

	processInitIds := func(filterIds []uint32, prop string) error {
		if len(filterIds) == 0 {
			return nil
		}
		var items interface{}
		var err error
		var app_ids []uint32
		switch prop {
		case appPropTag:
			items, err = s.atagrepo.ListAppTags(ctx, nil, &repo.AppTagsFilter{
				Ids: filterIds})
			if err != nil {
				return fmt.Errorf("ListApplications listAppTags error. %w", err)
			}
			for _, item := range items.([]*repo.AppTag) {
				app_ids = append(app_ids, item.AppID)
			}
		case appPropFeature:
			items, err = s.afrepo.ListAppFeatures(ctx, nil, &repo.AppFeaturesFilter{
				Ids: filterIds})
			if err != nil {
				return fmt.Errorf("ListApplications listAppFeatures error. %w", err)
			}
			for _, item := range items.([]*repo.AppFeature) {
				app_ids = append(app_ids, item.AppID)
			}
		case appPropHostgroup:
			items, err = s.ahgrepo.ListAppHostgroups(ctx, nil, &repo.AppHostgroupsFilter{
				Ids: filterIds})
			if err != nil {
				return fmt.Errorf("ListApplications listAppHostgroups error. %w", err)
			}
			for _, item := range items.([]*repo.AppHostgroup) {
				app_ids = append(app_ids, item.AppID)
			}
		default:
			return fmt.Errorf("ListApplications invalid prop %s", prop)
		}

		if len(app_ids) == 0 {
			return fmt.Errorf("ListApplications no app with %s", prop)
		}

		if len(dbFilter.Ids) == 0 {
			dbFilter.Ids = app_ids
		} else {
			dbFilter.Ids = IntersectSliceUint32(dbFilter.Ids, app_ids)
			if len(dbFilter.Ids) == 0 {
				return fmt.Errorf("ListApplications no app intersect with %s", prop)
			}
		}
		return nil
	}

	if filter != nil {

		// tags id
		if err := processInitIds(filter.TagsId, appPropTag); err != nil {
			return nil, err
		}

		// features id
		if err := processInitIds(filter.FeaturesId, appPropFeature); err != nil {
			return nil, err
		}

		// hostgroups id
		if err := processInitIds(filter.HostgroupsId, appPropHostgroup); err != nil {
			return nil, err
		}
	}

	_apps, err := s.apprepo.ListApplications(ctx, nil, dbFilter)
	if err != nil {
		return nil, err
	}
	bapps, err := ToBizApplications(_apps)
	if err != nil {
		return nil, err
	}
	for _, app := range bapps {
		if err := s.attachM2MProps(ctx, app); err != nil {
			return nil, err
		}
	}
	return bapps, nil
}
