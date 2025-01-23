package biz_test

import (
	"opspillar/internal/biz"
	"opspillar/internal/data"
	"opspillar/internal/data/repo"
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
	prdrepo := new(MockProductsRepo)
	teamrepo := new(MockTeamsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	hgrepo := new(MockHostgroupsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	adminrepo := new(MockAdminRepo)
	usecase := biz.NewApplicationsUsecase(
		apprepo, atagrepo, afrepo, ahgrepo,
		prdrepo, teamrepo, ftrepo, tagrepo,
		hgrepo, hfrepo, authzrepo, adminrepo, nil, txm)

	// 测试字段验证
	bad_field := []*biz.Application{
		{0, "", "desc", 10, false, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
		{0, "name", "desc", 0, false, 0, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
		{0, "name", "desc", 10, false, 1, 0, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
	}

	for _, bc := range bad_field {
		err := usecase.CreateApplications(ctx, []*biz.Application{bc})
		t.Logf("bad field: %v", err)
		assert.Error(t, err)
	}

	app := []*biz.Application{
		{0, "test-app", "desc", 10, false, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
	}

	ctx = context.WithValue(ctx, data.CtxUserName, "demouser")

	var err error
	// 测试产品验证
	prdcall := prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("prod: %v", err)
	prdcall.Unset()

	// 测试团队验证
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall := teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("team: %v", err)
	prdcall.Unset()
	tcall.Unset()

	// 测试特性验证
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall := ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("feature: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()

	// 测试标签验证
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall := tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("tag: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()

	// 测试主机组验证
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall := hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("hg: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()

	// test ownerId fail
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	admcall := adminrepo.On("CountUsers", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("ownerId: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	admcall.Unset()

	// 测试主机组匹配失败
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	admcall = adminrepo.On("CountUsers", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	hfcall := hfrepo.On("ListHostgroupMatchFeatures", ctx, mock.Anything, mock.Anything).Return([]uint32{}, nil)
	hgcall2 := hgrepo.On("ListHostgroups", ctx, mock.Anything, mock.Anything).Return([]*repo.Hostgroup{}, nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("hg match: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	admcall.Unset()
	hfcall.Unset()
	hgcall2.Unset()

	// 测试主机组匹配部分失败
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hfcall = hfrepo.On("ListHostgroupMatchFeatures", ctx, mock.Anything, mock.Anything).Return([]uint32{2, 3}, nil)
	admcall = adminrepo.On("CountUsers", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	hgcall2 = hgrepo.On("ListHostgroups", ctx, mock.Anything, mock.Anything).Return([]*repo.Hostgroup{
		{Id: 1},
		{Id: 2},
	}, nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("hg match partial: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	admcall.Unset()
	hfcall.Unset()
	hgcall2.Unset()

	// 测试创建应用失败
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	admcall = adminrepo.On("CountUsers", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	hfcall = hfrepo.On("ListHostgroupMatchFeatures", ctx, mock.Anything, mock.Anything).Return([]uint32{2, 3}, nil)
	hgcall2 = hgrepo.On("ListHostgroups", ctx, mock.Anything, mock.Anything).Return([]*repo.Hostgroup{
		{Id: 2},
		{Id: 3},
	}, nil)
	appcall := apprepo.On("CreateApplications", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create application fail"))

	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("app create: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	admcall.Unset()
	hfcall.Unset()
	hgcall2.Unset()
	appcall.Unset()

	// 测试创建app-tag fail
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	admcall = adminrepo.On("CountUsers", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	hfcall = hfrepo.On("ListHostgroupMatchFeatures", ctx, mock.Anything, mock.Anything).Return([]uint32{2, 3}, nil)
	hgcall2 = hgrepo.On("ListHostgroups", ctx, mock.Anything, mock.Anything).Return([]*repo.Hostgroup{
		{Id: 2},
		{Id: 3},
	}, nil)
	appcall = apprepo.On("CreateApplications", ctx, mock.Anything, mock.Anything).
		Return(nil)
	atagcall := atagrepo.On("CreateAppTags", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create app-tag fail"))

	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("app-tag create: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	admcall.Unset()
	hfcall.Unset()
	hgcall2.Unset()
	appcall.Unset()
	atagcall.Unset()

	// 测试创建app-feature fail
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hfcall = hfrepo.On("ListHostgroupMatchFeatures", ctx, mock.Anything, mock.Anything).Return([]uint32{2, 3}, nil)
	admcall = adminrepo.On("CountUsers", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	hgcall2 = hgrepo.On("ListHostgroups", ctx, mock.Anything, mock.Anything).Return([]*repo.Hostgroup{
		{Id: 2},
		{Id: 3},
	}, nil)
	appcall = apprepo.On("CreateApplications", ctx, mock.Anything, mock.Anything).
		Return(nil)
	atagcall = atagrepo.On("CreateAppTags", ctx, mock.Anything, mock.Anything).
		Return(nil)
	afcall := afrepo.On("CreateAppFeatures", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create app-feature fail"))

	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("app-feature create: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	admcall.Unset()
	hfcall.Unset()
	hgcall2.Unset()
	appcall.Unset()
	atagcall.Unset()
	afcall.Unset()

	// 测试创建app-hostgroup fail
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	admcall = adminrepo.On("CountUsers", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	hfcall = hfrepo.On("ListHostgroupMatchFeatures", ctx, mock.Anything, mock.Anything).Return([]uint32{2, 3}, nil)
	hgcall2 = hgrepo.On("ListHostgroups", ctx, mock.Anything, mock.Anything).Return([]*repo.Hostgroup{
		{Id: 2},
		{Id: 3},
	}, nil)
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
	t.Logf("app-hostgroup create: %v", err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	admcall.Unset()
	hgcall2.Unset()
	hfcall.Unset()
	appcall.Unset()
	atagcall.Unset()
	afcall.Unset()
	ahgcall.Unset()
}

func TestUpdateApp(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	apprepo := new(MockApplicationsRepo)
	atagrepo := new(MockAppTagsRepo)
	afrepo := new(MockAppFeaturesRepo)
	ahgrepo := new(MockAppHostgroupsRepo)
	prdrepo := new(MockProductsRepo)
	teamrepo := new(MockTeamsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	hgrepo := new(MockHostgroupsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	adminrepo := new(MockAdminRepo)

	usecase := biz.NewApplicationsUsecase(
		apprepo, atagrepo, afrepo, ahgrepo, prdrepo, teamrepo, ftrepo, tagrepo,
		hgrepo, hfrepo, authzrepo, adminrepo, nil, txm)

	// bad field
	bad_field := []*biz.Application{
		{0, "name", "desc", 10, false, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
		{10, "", "desc", 10, false, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
		{10, "name", "desc", 0, false, 0, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
		{10, "name", "desc", 10, false, 1, 0, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
	}

	for _, bc := range bad_field {
		err := usecase.UpdateApplications(ctx, []*biz.Application{bc})
		t.Logf("bad field: %v", err)
		assert.Error(t, err)
	}

	app := []*biz.Application{
		{10, "test-app", "desc", 10, false, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}},
	}

	ctx = context.WithValue(ctx, data.CtxUserName, "demouser")
	// validate enforce
	teamrepo.On("GetTeams", ctx, mock.Anything).Return(&repo.Team{ID: 1, Name: "test-team"}, nil)
	adminrepo.On("GetUsers", ctx, mock.Anything, mock.Anything).Return(&repo.User{Id: 10, UserName: "demouser"}, nil)

	efcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, nil)
	err := usecase.UpdateApplications(ctx, app)
	assert.Error(t, err)
	t.Logf("enforce: %v", err)
	efcall.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// 测试产品验证
	prdcall := prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateApplications(ctx, app)
	assert.Error(t, err)
	prdcall.Unset()

	// 测试团队验证
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall := teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateApplications(ctx, app)
	assert.Error(t, err)
	prdcall.Unset()
	tcall.Unset()

	// 测试特性验证
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall := ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateApplications(ctx, app)
	assert.Error(t, err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()

	// 测试标签验证
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall := tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateApplications(ctx, app)
	assert.Error(t, err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()

	// 测试主机组验证
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall := hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateApplications(ctx, app)
	assert.Error(t, err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()

	// test countUser
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	usercall := adminrepo.On("CountUsers", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateApplications(ctx, app)
	assert.Error(t, err)
	t.Log(err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	usercall.Unset()

	// 测试创建应用失败
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	hgcall = hgrepo.On("CountHostgroups", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	usercall = adminrepo.On("CountUsers", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	appcall := apprepo.On("UpdateApplications", ctx, mock.Anything, mock.Anything).
		Return(errors.New("update application fail"))

	err = usecase.UpdateApplications(ctx, app)
	assert.Error(t, err)
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	hgcall.Unset()
	appcall.Unset()
	usercall.Unset()

}

func TestAppHandleM2MProps(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	apprepo := new(MockApplicationsRepo)
	atagrepo := new(MockAppTagsRepo)
	afrepo := new(MockAppFeaturesRepo)
	ahgrepo := new(MockAppHostgroupsRepo)
	prdrepo := new(MockProductsRepo)
	teamrepo := new(MockTeamsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	hgrepo := new(MockHostgroupsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	adminrepo := new(MockAdminRepo)

	usecase := biz.NewApplicationsUsecase(
		apprepo, atagrepo, afrepo, ahgrepo,
		prdrepo, teamrepo, ftrepo, tagrepo,
		hgrepo, hfrepo, authzrepo, adminrepo, nil, txm)

	// app-tag
	atagFilter := &repo.AppTagsFilter{
		AppIds: []uint32{1},
	}
	newAppTagIds := []uint32{2, 3}
	oldAppTag := []*repo.AppTag{
		{Id: 1, AppID: 1, TagID: 1},
		{Id: 2, AppID: 1, TagID: 2},
	}
	toCreateAppTag := []*repo.AppTag{
		{AppID: 1, TagID: 3},
	}
	atagrepo.On("ListAppTags", ctx, mock.Anything, atagFilter).Return(oldAppTag, nil)
	atagrepo.On("DeleteAppTags", ctx, mock.Anything, []uint32{1}).Return(nil)
	atagrepo.On("CreateAppTags", ctx, mock.Anything, toCreateAppTag).Return(nil)

	err := usecase.HandleM2MProps(
		ctx, nil, uint32(1), newAppTagIds, "tag")
	assert.NoError(t, err)

	// app-feature
	afFilter := &repo.AppFeaturesFilter{
		AppIds: []uint32{1},
	}
	newAppFeatureIds := []uint32{2, 3}
	oldAppFeatures := []*repo.AppFeature{
		{Id: 1, AppID: 1, FeatureID: 1},
		{Id: 2, AppID: 1, FeatureID: 2},
	}
	toCreateAppFeature := []*repo.AppFeature{
		{AppID: 1, FeatureID: 3},
	}
	afrepo.On("ListAppFeatures", ctx, mock.Anything, afFilter).Return(oldAppFeatures, nil)
	afrepo.On("DeleteAppFeatures", ctx, mock.Anything, []uint32{1}).Return(nil)
	afrepo.On("CreateAppFeatures", ctx, mock.Anything, toCreateAppFeature).Return(nil)

	err = usecase.HandleM2MProps(
		ctx, nil, uint32(1), newAppFeatureIds, "feature")
	assert.NoError(t, err)

	// app-hostgroup
	ahFilter := &repo.AppHostgroupsFilter{
		AppIds: []uint32{1},
	}
	newAppHostgroupIds := []uint32{2, 3}
	oldAppHostgroups := []*repo.AppHostgroup{
		{Id: 1, AppID: 1, HostgroupID: 1},
		{Id: 2, AppID: 1, HostgroupID: 2},
	}
	toCreateAppHostgroup := []*repo.AppHostgroup{
		{AppID: 1, HostgroupID: 3},
	}
	ahgrepo.On("ListAppHostgroups", ctx, mock.Anything, ahFilter).Return(oldAppHostgroups, nil)
	ahgrepo.On("DeleteAppHostgroups", ctx, mock.Anything, []uint32{1}).Return(nil)
	ahgrepo.On("CreateAppHostgroups", ctx, mock.Anything, toCreateAppHostgroup).Return(nil)

	err = usecase.HandleM2MProps(
		ctx, nil, uint32(1), newAppHostgroupIds, "hostgroup")
	assert.NoError(t, err)

}

func TestDeleteApplications(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	apprepo := new(MockApplicationsRepo)
	atagrepo := new(MockAppTagsRepo)
	afrepo := new(MockAppFeaturesRepo)
	ahgrepo := new(MockAppHostgroupsRepo)
	prdrepo := new(MockProductsRepo)
	teamrepo := new(MockTeamsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	hgrepo := new(MockHostgroupsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	adminrepo := new(MockAdminRepo)

	usecase := biz.NewApplicationsUsecase(
		apprepo, atagrepo, afrepo, ahgrepo, prdrepo, teamrepo, ftrepo, tagrepo,
		hgrepo, hfrepo, authzrepo, adminrepo, nil, txm)

	ids := []uint32{1, 2}

	ctx = context.WithValue(ctx, data.CtxUserName, "demouser")
	apprepo.On("ListApplications", ctx, mock.Anything, mock.Anything).Return([]*repo.Application{
		{Id: 1, OwnerId: 10, TeamId: 11},
		{Id: 2, OwnerId: 10, TeamId: 12},
	}, nil)
	teamrepo.On("GetTeams", ctx, mock.Anything).Return(&repo.Team{ID: 1, Name: "test-team"}, nil)
	adminrepo.On("GetUsers", ctx, mock.Anything, mock.Anything).Return(&repo.User{Id: 10, UserName: "demouser"}, nil)

	// validate enforce
	efcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, nil)
	err := usecase.DeleteApplications(ctx, ids)
	assert.Error(t, err)
	t.Log(err)
	efcall.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// test delete app-tags
	rerr := errors.New("delete relations error")
	atagcall := atagrepo.On("DeleteAppTagsByAppId", ctx, mock.Anything, ids).Return(rerr)
	err = usecase.DeleteApplications(ctx, ids)
	assert.Error(t, err)
	atagcall.Unset()

	atagrepo.On("DeleteAppTagsByAppId", ctx, mock.Anything, ids).Return(nil)
	afcall := afrepo.On("DeleteAppFeaturesByAppId", ctx, mock.Anything, ids).Return(rerr)
	err = usecase.DeleteApplications(ctx, ids)
	assert.Error(t, err)
	afcall.Unset()

	atagrepo.On("DeleteAppTagsByAppId", ctx, mock.Anything, ids).Return(nil)
	afrepo.On("DeleteAppFeaturesByAppId", ctx, mock.Anything, ids).Return(nil)
	ahcall := ahgrepo.On("DeleteAppHostgroupsByAppId", ctx, mock.Anything, ids).Return(rerr)
	err = usecase.DeleteApplications(ctx, ids)
	assert.Error(t, err)
	ahcall.Unset()

	// Delete related records
	atagrepo.On("DeleteAppTagsByAppId", ctx, mock.Anything, ids).Return(nil)
	afrepo.On("DeleteAppFeaturesByAppId", ctx, mock.Anything, ids).Return(nil)
	ahgrepo.On("DeleteAppHostgroupsByAppId", ctx, mock.Anything, ids).Return(nil)

	// Delete apps
	apprepo.On("DeleteApplications", ctx, mock.Anything, ids).Return(nil)

	err = usecase.DeleteApplications(ctx, ids)
	assert.NoError(t, err)

}

func TestListApplications(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	txm.On("RunInTX", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		fn := args.Get(0).(func(repo.TX) error)
		fn(nil)
	})

	apprepo := new(MockApplicationsRepo)
	atagrepo := new(MockAppTagsRepo)
	afrepo := new(MockAppFeaturesRepo)
	ahgrepo := new(MockAppHostgroupsRepo)
	prdrepo := new(MockProductsRepo)
	teamrepo := new(MockTeamsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	hgrepo := new(MockHostgroupsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	adminrepo := new(MockAdminRepo)

	usecase := biz.NewApplicationsUsecase(
		apprepo, atagrepo, afrepo, ahgrepo,
		prdrepo, teamrepo, ftrepo, tagrepo,
		hgrepo, hfrepo, authzrepo, adminrepo, nil, txm)

	// Empty filter
	//filter := &biz.ListApplicationsFilter{}
	_apps := []*repo.Application{
		{
			Id:          1,
			Name:        "app1",
			Description: "desc1",
			TeamId:      1,
			ProductId:   1,
			IsStateful:  true,
			OwnerId:     10,
		},
		{
			Id:          2,
			Name:        "app2",
			Description: "desc2",
			TeamId:      2,
			ProductId:   2,
			IsStateful:  false,
			OwnerId:     10,
		},
	}
	biz_apps := []*biz.Application{
		{
			Id:           1,
			Name:         "app1",
			Description:  "desc1",
			OwnerId:      10,
			IsStateful:   true,
			TeamId:       1,
			ProductId:    1,
			TagsId:       []uint32{1, 2},
			FeaturesId:   []uint32{1, 2},
			HostgroupsId: []uint32{1, 2},
		},
		{
			Id:           2,
			Name:         "app2",
			Description:  "desc2",
			OwnerId:      10,
			IsStateful:   false,
			TeamId:       2,
			ProductId:    2,
			TagsId:       []uint32{1, 2},
			FeaturesId:   []uint32{1, 2},
			HostgroupsId: []uint32{1, 2},
		},
	}

	apprepo.On("ListApplications", ctx, mock.Anything, mock.Anything).Return(_apps, nil)
	atagcall := atagrepo.On("ListAppTags", ctx, mock.Anything, mock.Anything).Return([]*repo.AppTag{
		{Id: 1, AppID: 1, TagID: 1},
		{Id: 2, AppID: 1, TagID: 2},
	}, nil)
	afcall := afrepo.On("ListAppFeatures", ctx, mock.Anything, mock.Anything).Return([]*repo.AppFeature{
		{Id: 1, AppID: 1, FeatureID: 1},
		{Id: 2, AppID: 1, FeatureID: 2},
	}, nil)
	ahcall := ahgrepo.On("ListAppHostgroups", ctx, mock.Anything, mock.Anything).Return([]*repo.AppHostgroup{
		{Id: 1, AppID: 1, HostgroupID: 1},
		{Id: 2, AppID: 1, HostgroupID: 2},
	}, nil)
	apps, err := usecase.ListApplications(ctx, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(apps))
	assert.Equal(t, biz_apps[0], apps[0])
	assert.Equal(t, biz_apps[1], apps[1])

	filter := &biz.ListApplicationsFilter{
		Page:         1,
		PageSize:     10,
		TagsId:       []uint32{1, 2},
		FeaturesId:   []uint32{1, 2},
		HostgroupsId: []uint32{1, 2},
	}

	// app-tag filter nil
	atagcall.Unset()
	atagrepo.On("ListAppTags", ctx, mock.Anything, mock.Anything).Return([]*repo.AppTag{}, nil)
	_, err = usecase.ListApplications(ctx, filter)
	t.Logf("err: %v", err)
	assert.Error(t, err)

	// app-feature filter nil
	atagcall.Unset()
	atagrepo.On("ListAppTags", ctx, mock.Anything, mock.Anything).Return([]*repo.AppTag{
		{Id: 1, AppID: 1, TagID: 1},
		{Id: 2, AppID: 1, TagID: 2},
	}, nil)
	afcall.Unset()
	afcall = afrepo.On("ListAppFeatures", ctx, mock.Anything, mock.Anything).Return([]*repo.AppFeature{}, nil)
	_, err = usecase.ListApplications(ctx, filter)
	t.Logf("err: %v", err)
	assert.Error(t, err)

	// app-hostgroup filter nil
	afcall.Unset()
	afrepo.On("ListAppFeatures", ctx, mock.Anything, mock.Anything).Return([]*repo.AppFeature{
		{Id: 1, AppID: 1, FeatureID: 1},
		{Id: 2, AppID: 1, FeatureID: 2},
	}, nil)
	ahcall.Unset()
	ahgrepo.On("ListAppHostgroups", ctx, mock.Anything, mock.Anything).Return([]*repo.AppHostgroup{}, nil)
	_, err = usecase.ListApplications(ctx, filter)
	t.Logf("err: %v", err)
	assert.Error(t, err)

}
