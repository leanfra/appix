package biz_test

import (
	"appix/internal/biz"
	"appix/internal/data/repo"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateApp(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	apprepo := new(MockApplicationsRepo)
	atagrepo := new(MockAppTagsRepo)
	afrepo := new(MockAppFeaturesRepo)
	ahgrepo := new(MockAppHostgroupsRepo)
	clsrepo := new(MockClustersRepo)
	dcrepo := new(MockDatacentersRepo)
	prdrepo := new(MockProductsRepo)
	teamrepo := new(MockTeamsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	hgrepo := new(MockHostgroupsRepo)

	usecase := biz.NewApplicationsUsecase(
		apprepo, atagrepo, afrepo, ahgrepo, clsrepo,
		dcrepo, prdrepo, teamrepo, ftrepo, tagrepo,
		hgrepo, nil, txm)

	// 测试字段验证
	bad_field := []*biz.Application{
		{0, "", "desc", "web", false, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
		{0, "name", "desc", "", false, 0, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
		{0, "name", "desc", "web", false, 1, 0, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
		{0, "name", "desc", "web", false, 1, 1, 0, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
		{0, "name", "desc", "web", false, 1, 1, 1, 0, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
	}

	for _, bc := range bad_field {
		err := usecase.CreateApplications(ctx, []*biz.Application{bc})
		t.Logf("bad field: %v", err)
		assert.Error(t, err)
	}

	app := []*biz.Application{
		{0, "test-app", "desc", "web", false, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
	}

	// 测试集群验证
	clscall := clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err := usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()

	// 测试数据中心验证
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall := dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()

	// 测试产品验证
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall := prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	prdcall.Unset()

	// 测试团队验证
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall := teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	prdcall.Unset()
	tcall.Unset()

	// 测试特性验证
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall := ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()

	// 测试标签验证
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall := tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()

	// 测试主机组验证
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall := hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()

	// 测试创建应用失败
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	appcall := apprepo.On("CreateApplications", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create application fail"))

	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	appcall.Unset()

	// 测试创建app-tag fail
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	appcall = apprepo.On("CreateApplications", ctx, mock.Anything, mock.Anything).
		Return(nil)
	atagcall := atagrepo.On("CreateAppTags", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create app-tag fail"))

	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	appcall.Unset()
	atagcall.Unset()

	// 测试创建app-feature fail
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	appcall = apprepo.On("CreateApplications", ctx, mock.Anything, mock.Anything).
		Return(nil)
	atagcall = atagrepo.On("CreateAppTags", ctx, mock.Anything, mock.Anything).
		Return(nil)
	afcall := afrepo.On("CreateAppFeatures", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create app-feature fail"))

	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	appcall.Unset()
	atagcall.Unset()
	afcall.Unset()

	// 测试创建app-hostgroup fail
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	appcall = apprepo.On("CreateApplications", ctx, mock.Anything, mock.Anything).
		Return(nil)
	atagcall = atagrepo.On("CreateAppTags", ctx, mock.Anything, mock.Anything).
		Return(nil)
	afcall = afrepo.On("CreateAppFeatures", ctx, mock.Anything, mock.Anything).
		Return(nil)
	ahgcall := ahgrepo.On("CreateAppHostgroups", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create app-hostgroup fail"))

	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	appcall.Unset()
	atagcall.Unset()
	afcall.Unset()
	ahgcall.Unset()
}
