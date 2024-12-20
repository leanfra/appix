package data

import (
	"appix/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	//  TODO: modify project name
	// biz "appix/internal/biz"
)

type HostgroupsRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewHostgroupsRepoImpl(data *Data, logger log.Logger) (biz.HostgroupsRepo, error) {

	if err := validateData(data); err != nil {
		return nil, err
	}
	if err := initTable(data.db, &Hostgroup{}, hostgroupTable); err != nil {
		return nil, err
	}
	if err := initTable(data.db, &ResFeature{}, resFeatureTable); err != nil {
		return nil, err
	}
	if err := initTable(data.db, &ResTag{}, resTagTable); err != nil {
		return nil, err
	}
	if err := initTable(data.db, &ResTeam{}, resTeamTable); err != nil {
		return nil, err
	}
	if err := initTable(data.db, &ResProduct{}, resProductTable); err != nil {
		return nil, err
	}

	return &HostgroupsRepoImpl{
		data: data,
		log:  log.NewHelper(logger),
	}, nil
}

func (d *HostgroupsRepoImpl) checkClassExists(tx *gorm.DB, hgs []*biz.Hostgroup) error {

	var allTagsId []uint32
	var allFeaturesId []uint32
	var allProductsId []uint32
	var allTeamsId []uint32
	var allClustersId []uint32
	var allEnvsId []uint32
	var allDatacentersId []uint32
	for _, hg := range hgs {
		allTagsId = append(allTagsId, hg.TagsId...)
		allFeaturesId = append(allFeaturesId, hg.FeaturesId...)
		allProductsId = append(allProductsId, hg.ShareProductsId...)
		allProductsId = append(allProductsId, hg.ProductId)
		allTeamsId = append(allTeamsId, hg.ShareTeamsId...)
		allTeamsId = append(allTeamsId, hg.TeamId)
		allClustersId = append(allClustersId, hg.ClusterId)
		allEnvsId = append(allEnvsId, hg.EnvId)
		allDatacentersId = append(allDatacentersId, hg.DatacenterId)
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

// CreateHostgroups is
func (d *HostgroupsRepoImpl) CreateHostgroups(ctx context.Context, hgs []*biz.Hostgroup) error {

	tx := d.data.db.WithContext(ctx).Begin()
	if e := d.checkClassExists(tx, hgs); e != nil {
		tx.Rollback()
		return e
	}

	for _, hg := range hgs {

		// never return error
		db_hg, _ := NewHostgroup(hg)

		// insert hostgroup
		r := tx.Create(db_hg)
		if r.Error != nil {
			tx.Rollback()
			return r.Error
		}

		// insert res_teams
		if e := createClass(tx, &ResTeam{}, hg.ShareTeamsId, hostgroupType, db_hg.Id); e != nil {
			tx.Rollback()
			return e
		}

		// insert res_tags
		if e := createClass(tx, &ResTag{}, hg.TagsId, hostgroupType, db_hg.Id); e != nil {
			tx.Rollback()
			return e
		}

		// insert res_products
		if e := createClass(tx, &ResProduct{}, hg.ShareProductsId, hostgroupType, db_hg.Id); e != nil {
			tx.Rollback()
			return e
		}

		// insert res_features
		if e := createClass(tx, &ResFeature{}, hg.FeaturesId, hostgroupType, db_hg.Id); e != nil {
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

// UpdateHostgroups is
func (d *HostgroupsRepoImpl) UpdateHostgroups(ctx context.Context, hgs []*biz.Hostgroup) error {

	tx := d.data.db.WithContext(ctx).Begin()
	if e := d.checkClassExists(tx, hgs); e != nil {
		tx.Rollback()
		return e
	}

	for _, hg := range hgs {

		// never return error
		db_hg, _ := NewHostgroup(hg)

		// get old share team
		old_team_ids, e := listClassIds(tx, &ResTeam{}, "team_id", hostgroupType, hg.Id)
		if e != nil {
			tx.Rollback()
			d.log.Errorf("get old share team id failed. Error: %v", e)
			return e
		}
		// get old share product
		old_product_ids, e := listClassIds(tx, &ResProduct{}, "product_id", hostgroupType, hg.Id)
		if e != nil {
			tx.Rollback()
			d.log.Errorf("get old share product id failed. Error: %v", e)
			return e
		}

		// get old tag
		old_tag_ids, e := listClassIds(tx, &ResTag{}, "tag_id", hostgroupType, hg.Id)
		if e != nil {
			tx.Rollback()
			d.log.Errorf("get old share tag id failed. Error: %v", e)
			return e
		}

		// get old feature
		old_feature_ids, e := listClassIds(tx, &ResFeature{}, "feature_id", hostgroupType, hg.Id)
		if e != nil {
			tx.Rollback()
			d.log.Errorf("get old feature id failed. Error: %v", e)
			return e
		}

		// get new share team
		new_team_ids := DiffUint32(hg.ShareTeamsId, old_team_ids)
		// get deleted share team
		del_team_ids := DiffUint32(old_team_ids, hg.ShareTeamsId)

		// get new share product
		new_product_ids := DiffUint32(hg.ShareProductsId, old_product_ids)
		// get deleted share product
		del_product_ids := DiffUint32(old_product_ids, hg.ShareProductsId)

		// get new tag
		new_tag_ids := DiffUint32(hg.TagsId, old_tag_ids)
		// get deleted tag
		del_tag_ids := DiffUint32(old_tag_ids, hg.TagsId)

		// get new feature
		new_feature_ids := DiffUint32(hg.FeaturesId, old_feature_ids)
		// get delete feature
		del_feature_ids := DiffUint32(old_feature_ids, hg.FeaturesId)

		// update hostgroup
		r := tx.Save(db_hg)
		if r.Error != nil {
			tx.Rollback()
			return r.Error
		}

		// update res_team
		if e := updateClass(tx, &ResTeam{}, "team_id",
			del_team_ids, new_team_ids, hostgroupType, hg.Id); e != nil {

			tx.Rollback()
			return e
		}

		// update res_product
		if e := updateClass(tx, &ResProduct{}, "product_id",
			del_product_ids, new_product_ids, hostgroupType, hg.Id); e != nil {

			tx.Rollback()
			return e
		}

		// update res_tag
		if e := updateClass(tx, &ResTag{}, "tag_id",
			del_tag_ids, new_tag_ids, hostgroupType, hg.Id); e != nil {

			tx.Rollback()
			return e
		}

		// update res_feature
		if e := updateClass(tx, &ResFeature{}, "feature_id",
			del_feature_ids, new_feature_ids, hostgroupType, hg.Id); e != nil {

			tx.Rollback()
			return e
		}
	}
	if r := tx.Commit(); r != nil {
		return r.Error
	}
	return nil
}

// DeleteHostgroups is
func (d *HostgroupsRepoImpl) DeleteHostgroups(ctx context.Context, ids []uint32) error {

	r := d.data.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Hostgroup{})
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// GetHostgroups is
func (d *HostgroupsRepoImpl) GetHostgroups(ctx context.Context, id uint32) (*biz.Hostgroup, error) {

	hg := &Hostgroup{}
	r := d.data.db.WithContext(ctx).Where("id = ?", id).First(hg)
	if r.Error != nil {
		return nil, r.Error
	}
	return NewBizHostgroup(hg)
}

// ListHostgroups is
func (d *HostgroupsRepoImpl) ListHostgroups(ctx context.Context,
	filter *biz.ListHostgroupsFilter) ([]*biz.Hostgroup, error) {

	var hostgroups []*Hostgroup
	query := d.data.db.WithContext(ctx)

	if filter != nil {
		if filter.Page > 0 && filter.PageSize > 0 {
			offset := int((filter.Page - 1) * filter.PageSize)
			query = query.Offset(offset).Limit(int(filter.PageSize))
		}
		if len(filter.Names) > 0 {
			nameConditions := buildOrLike("name", len(filter.Names))
			query = query.Where(nameConditions, filter.Names)
		}
		if len(filter.Ids) > 0 {
			query = query.Where("id in (?)", filter.Ids)
		}
		if len(filter.Clusters) > 0 {
			query = query.Where(
				"cluster_id in (select id from cluster where name in (?))",
				filter.Clusters)
		}
		if len(filter.Datacenters) > 0 {
			query = query.Where(
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

		if len(filter.ShareProducts) > 0 {
			q_product_ids := d.data.db.WithContext(ctx).
				Model(&Product{}).
				Select("id").
				Where("name in (?)", filter.ShareProducts)
			q_product_res_hg_ids := d.data.db.WithContext(ctx).
				Model(&ResProduct{}).
				Select("res_id").
				Where("res_type = ? AND product_id in (?)", hostgroupType, q_product_ids)
			query = query.Where("id in (?)", q_product_res_hg_ids)
		}

		if len(filter.ShareTeams) > 0 {
			q_team_ids := d.data.db.WithContext(ctx).
				Model(&Team{}).Select("id").
				Where("name in (?)", filter.ShareTeams)
			q_team_res_hg_ids := d.data.db.WithContext(ctx).
				Model(&ResTeam{}).Select("res_id").
				Where("res_type = ? AND team_id in (?)", hostgroupType, q_team_ids)
			query = query.Where("id in (?)", q_team_res_hg_ids)
		}

		if len(filter.Tags) > 0 {
			kvOr, kvs := buildOrKV("key", "value", filter.Tags)
			subquery := d.data.db.WithContext(ctx).
				Model(&ResTag{}).
				Select("res_id").
				Where("res_type = ?", hostgroupType).
				Where(kvOr, kvs)
			query = query.Where("id in (?)", subquery)
		}

		if len(filter.Features) > 0 {
			featureOr, kvs := buildOrKV("name", "value", filter.Features)
			subquery := d.data.db.WithContext(ctx).
				Model(&ResFeature{}).
				Select("res_id").
				Where("res_type = ?", hostgroupType).
				Where(featureOr, kvs)
			query = query.Where("id in (?)", subquery)
		}
	}

	r := query.Find(&hostgroups)
	if r.Error != nil {
		return nil, r.Error
	}

	bizHgs := make([]*biz.Hostgroup, len(hostgroups))
	var e error
	for i, hg := range hostgroups {
		bizHgs[i], e = NewBizHostgroup(hg)
		if e != nil {
			return nil, e
		}
		_features_id, err := listClassIds(
			d.data.db, &ResFeature{}, "feature_id", hostgroupType, hg.Id)
		if err != nil {
			return nil, err
		}
		_tags_id, err := listClassIds(
			d.data.db, &ResTag{}, "tag_id", hostgroupType, hg.Id)
		if err != nil {
			return nil, err
		}
		bizHgs[i].FeaturesId = _features_id
		bizHgs[i].TagsId = _tags_id
	}

	return bizHgs, nil
}
