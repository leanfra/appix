package data

import (
	"appix/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type ApplicationsRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewApplicationsRepoImpl(data *Data, logger log.Logger) (biz.ApplicationsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}

	if err := initTable(data.db, &Application{}, applicationTable); err != nil {
		return nil, err
	}

	return &ApplicationsRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *ApplicationsRepoImpl) checkClassExists(tx *gorm.DB, apps []*biz.Application) error {

	var allTagsId []uint32
	var allFeaturesId []uint32
	var allProductsId []uint32
	var allTeamsId []uint32
	var allEnvsId []uint32
	var allClustersId []uint32
	var allDatacentersId []uint32

	for _, app := range apps {
		allTagsId = append(allTagsId, app.TagsId...)
		allFeaturesId = append(allFeaturesId, app.FeaturesId...)
		allProductsId = append(allProductsId, app.ProductId)
		allTeamsId = append(allTeamsId, app.TeamId)
		allEnvsId = append(allEnvsId, app.EnvId)
		allClustersId = append(allClustersId, app.ClusterId)
		allDatacentersId = append(allDatacentersId, app.DatacenterId)
	}

	allTagsId = DedupSliceUint32(allTagsId)
	allFeaturesId = DedupSliceUint32(allFeaturesId)
	allProductsId = DedupSliceUint32(allProductsId)
	allTeamsId = DedupSliceUint32(allTeamsId)
	allClustersId = DedupSliceUint32(allClustersId)
	allEnvsId = DedupSliceUint32(allEnvsId)
	allDatacentersId = DedupSliceUint32(allDatacentersId)

	if r := existsRecords(tx, &Tag{}, allTagsId); r != nil {
		tx.Rollback()
		return ErrMissingTags
	}
	if r := existsRecords(tx, &Feature{}, allFeaturesId); r != nil {
		tx.Rollback()
		return ErrMissingFeatures
	}
	if r := existsRecords(tx, &Product{}, allProductsId); r != nil {
		tx.Rollback()
		return ErrMissingProducts
	}
	if r := existsRecords(tx, &Team{}, allTeamsId); r != nil {
		tx.Rollback()
		return ErrMissingTeams
	}
	if r := existsRecords(tx, &Cluster{}, allClustersId); r != nil {
		tx.Rollback()
		return ErrMissingClusters
	}
	if r := existsRecords(tx, &Env{}, allEnvsId); r != nil {
		tx.Rollback()
		return ErrMissingEnvs
	}
	if r := existsRecords(tx, &Datacenter{}, allDatacentersId); r != nil {
		tx.Rollback()
		return ErrMissingDatacenters
	}
	return nil
}

// CreateApplications is
func (d *ApplicationsRepoImpl) CreateApplications(
	ctx context.Context, apps []*biz.Application) error {

	tx := d.data.db.WithContext(ctx).Begin()
	if e := d.checkClassExists(tx, apps); e != nil {
		tx.Rollback()
		return e
	}

	for _, app := range apps {
		db_app, _ := NewApplication(app)
		if e := tx.Create(db_app).Error; e != nil {
			tx.Rollback()
			d.log.Errorf("create application failed. Error: %v", e)
			return e
		}

		if e := createClass(tx, &ResTag{}, app.TagsId, applicationType, db_app.Id); e != nil {
			tx.Rollback()
			return e
		}

		if e := createClass(tx, &ResFeature{}, app.FeaturesId, applicationType, db_app.Id); e != nil {
			tx.Rollback()
			return e
		}
	}

	if r := tx.Commit(); r != nil {
		tx.Rollback()
		return r.Error
	}

	return nil
}

// UpdateApplications is
func (d *ApplicationsRepoImpl) UpdateApplications(
	ctx context.Context, apps []*biz.Application) error {

	tx := d.data.db.Begin()
	if e := d.checkClassExists(tx, apps); e != nil {
		tx.Rollback()
		return e
	}

	for _, app := range apps {
		db_app, _ := NewApplication(app)

		old_tag_ids, e := listClassIds(tx, &ResTag{}, "tag_id", applicationType, db_app.Id)
		if e != nil {
			tx.Rollback()
			d.log.Errorf("get old tag id failed. Error: %v", e)
			return e
		}

		old_feature_ids, e := listClassIds(tx, &ResFeature{}, "feature_id", applicationType, db_app.Id)
		if e != nil {
			tx.Rollback()
			d.log.Errorf("get old feature id failed. Error: %v", e)
			return e
		}

		new_tag_ids := DiffUint32(app.TagsId, old_tag_ids)
		del_tag_ids := DiffUint32(old_tag_ids, app.TagsId)

		new_feature_ids := DiffUint32(app.FeaturesId, old_feature_ids)
		del_feature_ids := DiffUint32(old_feature_ids, app.FeaturesId)

		if e := tx.Save(db_app).Error; e != nil {
			tx.Rollback()
			d.log.Errorf("update application failed. Error: %v", e)
			return e
		}

		if e := updateClass(tx, &ResTag{}, "tag_id",
			del_tag_ids, new_tag_ids, applicationType, db_app.Id); e != nil {
			tx.Rollback()
			return e
		}

		if e := updateClass(tx, &ResFeature{}, "feature_id",
			del_feature_ids, new_feature_ids, applicationType, db_app.Id); e != nil {
			tx.Rollback()
			return e
		}
	}
	if r := tx.Commit(); r != nil {
		tx.Rollback()
		return r.Error
	}

	return nil
}

// DeleteApplications is
func (d *ApplicationsRepoImpl) DeleteApplications(
	ctx context.Context, ids []uint32) error {

	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Application{})
	if r.Error != nil {
		return r.Error
	}

	return nil
}

// GetApplications is
func (d *ApplicationsRepoImpl) GetApplications(
	ctx context.Context, id uint32) (*biz.Application, error) {

	app := &Application{}
	r := d.data.db.WithContext(ctx).Where("id = ?", id).First(app)
	if r.Error != nil {
		return nil, r.Error
	}

	return NewBizApplication(app)
}

// ListApplications is
func (d *ApplicationsRepoImpl) ListApplications(
	ctx context.Context, filter *biz.ListApplicationsFilter) ([]*biz.Application, error) {

	var apps []*Application
	query := d.data.db.WithContext(ctx)

	if filter != nil {
		// TODO set default pagesize in biz
		if filter.Page > 0 && filter.PageSize > 0 {
			offset := int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}
		if len(filter.Ids) > 0 {
			query = query.Where("id in (?)", filter.Ids)
		}
		if len(filter.Names) > 0 {
			nameConditions := buildOrLike("name", len(filter.Names))
			query = query.Where(nameConditions, filter.Names)
		}
		if len(filter.Clusters) > 0 {
			query = query.
				Where(
					"cluster_id in (select id from cluster where name in (?))",
					filter.Clusters)
		}
		if len(filter.Datacenters) > 0 {
			query = query.
				Where(
					"datacenter_id in (select id from datacenter where name in (?))",
					filter.Datacenters)
		}
		if len(filter.Envs) > 0 {
			query = query.Where("env_id in (select id from env where name in (?))",
				filter.Envs)
		}
		if len(filter.Products) > 0 {
			query = query.Where(
				"product_id in (select id from product where name in (?))",
				filter.Products)
		}
		if len(filter.Teams) > 0 {
			query = query.Where(
				"team_id in (select id from team where name in (?))",
				filter.Teams)
		}
		if filter.IsStateful != biz.IsStatefulNone {
			query = query.Where("is_stateful = ?", filter.IsStateful)
		}

		if len(filter.Tags) > 0 {
			tagsOr, kvs := buildOrKV("key", "value", filter.Tags)

			subquery := d.data.db.WithContext(ctx).
				Model(&ResTag{}).
				Select("res_id").
				Where("res_type = ?", applicationType).
				Where(tagsOr, kvs)
			query = query.Where("id in (?)", subquery)
		}

		if len(filter.Features) > 0 {
			featuresOr, kvs := buildOrKV("name", "value", filter.Features)

			subquery := d.data.db.WithContext(ctx).
				Model(&ResFeature{}).
				Select("res_id").
				Where("res_type = ?", applicationType).
				Where(featuresOr, kvs)
			query = query.Where("id in (?)", subquery)
		}
	}

	r := query.Find(&apps)
	if r.Error != nil {
		return nil, r.Error
	}

	bizApps := make([]*biz.Application, len(apps))
	var e error
	for i, app := range apps {
		bizApps[i], e = NewBizApplication(app)
		if e != nil {
			return nil, e
		}
		_features_id, err := listClassIds(d.data.db, &ResFeature{}, "feature_id", applicationType, app.Id)
		if err != nil {
			return nil, err
		}
		_tags_id, err := listClassIds(d.data.db, &ResTag{}, "tag_id", applicationType, app.Id)
		if err != nil {
			return nil, err
		}
		bizApps[i].FeaturesId = _features_id
		bizApps[i].TagsId = _tags_id
	}

	return bizApps, nil
}
