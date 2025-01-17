package biz_test

import (
	"appix/internal/biz"
	"appix/internal/data"
	"appix/internal/data/repo"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateHostgroup(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	hgrepo := new(MockHostgroupsRepo)
	htrepo := new(MockHostgroupTeamsRepo)
	htagrepo := new(MockHostgroupTagsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	clsrepo := new(MockClustersRepo)
	dcrepo := new(MockDatacentersRepo)
	envrepo := new(MockEnvsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	teamrepo := new(MockTeamsRepo)
	prdrepo := new(MockProductsRepo)
	ahrepo := new(MockAppHostgroupsRepo)

	usecase := biz.NewHostgroupsUsecase(
		hgrepo, htrepo, hprepo, htagrepo, hfrepo, clsrepo,
		dcrepo, envrepo, ftrepo, tagrepo, teamrepo,
		prdrepo, ahrepo, authzrepo, nil, txm)

	// bad field
	bad_field := []*biz.Hostgroup{
		{0, "", "desc", 1, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{0, "name", "desc", 0, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{0, "name", "desc", 1, 0, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{0, "name", "desc", 1, 1, 0, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{0, "name", "desc", 1, 1, 1, 0, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{0, "name", "desc", 1, 1, 1, 1, 0, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
	}

	for _, bc := range bad_field {
		err := usecase.CreateHostgroups(ctx, []*biz.Hostgroup{bc})
		t.Logf("bad field: %v", err)
		assert.Error(t, err)
	}
	// enforce fail
	teamrepo.On("GetTeams", ctx, mock.Anything, mock.Anything).Return(&repo.Team{
		2, "team2", "team2code", "team2l", "desc"}, nil)
	hg := []*biz.Hostgroup{
		{0, "name", "desc", 1, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
	}
	authcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything, mock.Anything).Return(false, errors.New("enforce fail"))
	err := usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	authcall.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// props count 0
	//// cluster
	clscall := clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()

	//// datacenter
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall := dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()

	//// env
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall := envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()

	//// product
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall := prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()

	//// team
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall := teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()

	//// feature
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall := ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("no feature: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()

	//// lack feature
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("lack feature: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()

	//// tag
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall := tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("no tag: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()

	//// lack tag
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("lack tag: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()

	//// share product
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall := prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(1), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("share product: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()

	//// share team
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall = prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)

	steamcall := teamrepo.On("CountTeams", ctx, mock.Anything,
		&repo.TeamsFilter{Ids: []uint32{2, 3}}).Return(int64(1), nil)
	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("share product: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()
	steamcall.Unset()

	// create hostgroup fail
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall = prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	steamcall = teamrepo.On("CountTeams", ctx, mock.Anything,
		&repo.TeamsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	hgcall := hgrepo.On("CreateHostgroups", ctx, mock.Anything, mock.Anything).Return(errors.New("create hostgroup fail"))

	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("create hostgroup: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()
	steamcall.Unset()
	hgcall.Unset()

	// create hostgroup-tag fail
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall = prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	steamcall = teamrepo.On("CountTeams", ctx, mock.Anything,
		&repo.TeamsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	hgcall = hgrepo.On("CreateHostgroups", ctx, mock.Anything, mock.Anything).
		Return(nil)
	htcall := htagrepo.On("CreateHostgroupTags", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create hostgroup-tag fail"))

	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("create hostgroup-tag: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()
	steamcall.Unset()
	hgcall.Unset()
	htcall.Unset()

	// create hostgroup-feature fail
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall = prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	steamcall = teamrepo.On("CountTeams", ctx, mock.Anything,
		&repo.TeamsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	hgcall = hgrepo.On("CreateHostgroups", ctx, mock.Anything, mock.Anything).
		Return(nil)
	htcall = htagrepo.On("CreateHostgroupTags", ctx, mock.Anything, mock.Anything).
		Return(nil)
	hfcall := hfrepo.On("CreateHostgroupFeatures", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create hostgroup-feature fail"))

	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("create hostgroup-feature: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()
	steamcall.Unset()
	hgcall.Unset()
	htcall.Unset()
	hfcall.Unset()

	// create hostgroup-share-product fail
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall = prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	steamcall = teamrepo.On("CountTeams", ctx, mock.Anything,
		&repo.TeamsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	hgcall = hgrepo.On("CreateHostgroups", ctx, mock.Anything, mock.Anything).
		Return(nil)
	htcall = htagrepo.On("CreateHostgroupTags", ctx, mock.Anything, mock.Anything).
		Return(nil)
	hfcall = hfrepo.On("CreateHostgroupFeatures", ctx, mock.Anything, mock.Anything).
		Return(nil)
	hpcall := hprepo.On("CreateHostgroupProducts", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create hostgroup-share-product fail"))

	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("create hostgroup-product: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()
	steamcall.Unset()
	hgcall.Unset()
	htcall.Unset()
	hfcall.Unset()
	hpcall.Unset()

	// create hostgroup-share-team fail
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall = prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	steamcall = teamrepo.On("CountTeams", ctx, mock.Anything,
		&repo.TeamsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	hgcall = hgrepo.On("CreateHostgroups", ctx, mock.Anything, mock.Anything).
		Return(nil)
	htcall = htagrepo.On("CreateHostgroupTags", ctx, mock.Anything, mock.Anything).
		Return(nil)
	hfcall = hfrepo.On("CreateHostgroupFeatures", ctx, mock.Anything, mock.Anything).
		Return(nil)
	hpcall = hprepo.On("CreateHostgroupProducts", ctx, mock.Anything, mock.Anything).
		Return(nil)
	htmcall := htrepo.On("CreateHostgroupTeams", ctx, mock.Anything, mock.Anything).
		Return(errors.New("create hostgroup-share-team fail"))

	err = usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("create hostgroup-team: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()
	steamcall.Unset()
	hgcall.Unset()
	htcall.Unset()
	hfcall.Unset()
	hpcall.Unset()
	htmcall.Unset()
}

func TestUpdateHostgroup(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	hgrepo := new(MockHostgroupsRepo)
	htrepo := new(MockHostgroupTeamsRepo)
	htagrepo := new(MockHostgroupTagsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	clsrepo := new(MockClustersRepo)
	dcrepo := new(MockDatacentersRepo)
	envrepo := new(MockEnvsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	teamrepo := new(MockTeamsRepo)
	prdrepo := new(MockProductsRepo)
	ahrepo := new(MockAppHostgroupsRepo)

	usecase := biz.NewHostgroupsUsecase(
		hgrepo, htrepo, hprepo, htagrepo, hfrepo, clsrepo,
		dcrepo, envrepo, ftrepo, tagrepo, teamrepo,
		prdrepo, ahrepo, authzrepo, nil, txm)

	// bad field
	bad_field := []*biz.Hostgroup{
		{1, "", "desc", 1, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{0, "name", "desc", 1, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{10, "name", "desc", 0, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{10, "name", "desc", 1, 0, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{10, "name", "desc", 1, 1, 0, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{10, "name", "desc", 1, 1, 1, 0, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
		{10, "name", "desc", 1, 1, 1, 1, 0, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
	}

	for _, bc := range bad_field {
		err := usecase.UpdateHostgroups(ctx, []*biz.Hostgroup{bc})
		t.Logf("bad field: %v", err)
		assert.Error(t, err)
	}

	hg := []*biz.Hostgroup{
		{11, "name", "desc", 1, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
	}

	teamrepo.On("GetTeams", ctx, mock.Anything, mock.Anything).Return(&repo.Team{
		2, "team2", "team2code", "team2l", "desc"}, nil)
	// enforce fail
	authcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, errors.New("enforce fail"))
	err := usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	authcall.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// props count 0
	//// cluster
	clscall := clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()

	//// datacenter
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall := dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()

	//// env
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall := envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()

	//// product
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall := prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()

	//// team
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall := teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()

	//// feature
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall := ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("no feature: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()

	//// lack feature
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("lack feature: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()

	//// tag
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall := tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("no tag: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()

	//// lack tag
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("lack tag: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()

	//// share product
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall := prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(1), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("share product: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()

	//// share team
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall = prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)

	steamcall := teamrepo.On("CountTeams", ctx, mock.Anything,
		&repo.TeamsFilter{Ids: []uint32{2, 3}}).Return(int64(1), nil)
	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("share product: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()
	steamcall.Unset()

	// repo hostgroup fail
	clscall = clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	dccall = dcrepo.On("CountDatacenters", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	envcall = envrepo.On("CountEnvs", ctx, mock.Anything, mock.Anything).Return(int64(1), nil)
	prdcall = prdrepo.On("CountProducts", ctx, mock.Anything, &repo.ProductsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	tcall = teamrepo.On("CountTeams", ctx, mock.Anything, &repo.TeamsFilter{Ids: []uint32{1}}).Return(int64(1), nil)
	fcall = ftrepo.On("CountFeatures", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	tgcall = tagrepo.On("CountTags", ctx, mock.Anything, mock.Anything).Return(int64(2), nil)
	sprdcall = prdrepo.On("CountProducts", ctx, mock.Anything,
		&repo.ProductsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	steamcall = teamrepo.On("CountTeams", ctx, mock.Anything,
		&repo.TeamsFilter{Ids: []uint32{2, 3}}).Return(int64(2), nil)
	hgcall := hgrepo.On("UpdateHostgroups", ctx, mock.Anything, mock.Anything).
		Return(errors.New("update hostgroup fail"))

	err = usecase.UpdateHostgroups(ctx, hg)
	assert.Error(t, err)
	t.Logf("update hostgroup: %v", err)
	clscall.Unset()
	dccall.Unset()
	envcall.Unset()
	prdcall.Unset()
	tcall.Unset()
	fcall.Unset()
	tgcall.Unset()
	sprdcall.Unset()
	steamcall.Unset()
	hgcall.Unset()

	// handle tag in TestHandleM2MProps
	// handle feature in TestHandleM2MProps
	// handle product in TestHandleM2MProps
	// handle team in TestHandleM2MProps
}

func TestHostgroupsHandleM2MProps(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	hgrepo := new(MockHostgroupsRepo)
	htrepo := new(MockHostgroupTeamsRepo)
	htagrepo := new(MockHostgroupTagsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	clsrepo := new(MockClustersRepo)
	dcrepo := new(MockDatacentersRepo)
	envrepo := new(MockEnvsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	teamrepo := new(MockTeamsRepo)
	prdrepo := new(MockProductsRepo)
	ahrepo := new(MockAppHostgroupsRepo)

	usecase := biz.NewHostgroupsUsecase(
		hgrepo, htrepo, hprepo, htagrepo, hfrepo, clsrepo,
		dcrepo, envrepo, ftrepo, tagrepo, teamrepo,
		prdrepo, ahrepo, authzrepo, nil, txm)

	// houstgroup-tag
	htagFilter := &repo.HostgroupTagsFilter{
		HostgroupIds: []uint32{1},
	}
	newHgTagIds := []uint32{2, 3}
	oldHgTag := []*repo.HostgroupTag{
		{Id: 1, HostgroupID: 1, TagID: 1},
		{Id: 2, HostgroupID: 1, TagID: 2},
	}
	toCreateHgTag := []*repo.HostgroupTag{
		{HostgroupID: 1, TagID: 3},
	}
	htagrepo.On("ListHostgroupTags", ctx, mock.Anything, htagFilter).Return(oldHgTag, nil)
	htagrepo.On("DeleteHostgroupTags", ctx, mock.Anything, []uint32{1}).Return(nil, nil).Once()
	htagrepo.On("CreateHostgroupTags", ctx, mock.Anything, toCreateHgTag).Return(nil, nil).Once()

	err := usecase.HandleM2MProps(
		ctx, nil, uint32(1), newHgTagIds, "tag")
	assert.NoError(t, err)

	// hostgroup-feature
	hfFilter := &repo.HostgroupFeaturesFilter{
		HostgroupIds: []uint32{1},
	}
	newHgFeatureIds := []uint32{2, 3}
	oldHgFeatures := []*repo.HostgroupFeature{
		{Id: 1, HostgroupID: 1, FeatureID: 1},
		{Id: 2, HostgroupID: 1, FeatureID: 2},
	}
	toCreateHgFeature := []*repo.HostgroupFeature{
		{HostgroupID: 1, FeatureID: 3},
	}

	hfrepo.On("ListHostgroupFeatures", ctx, mock.Anything, hfFilter).Return(oldHgFeatures, nil)
	hfrepo.On("DeleteHostgroupFeatures", ctx, mock.Anything, []uint32{1}).Return(nil, nil).Once()
	hfrepo.On("CreateHostgroupFeatures", ctx, mock.Anything, toCreateHgFeature).Return(nil, nil).Once()

	err = usecase.HandleM2MProps(
		ctx, nil, uint32(1), newHgFeatureIds, "feature")
	assert.NoError(t, err)

	// hostgroup-product
	hpFilter := &repo.HostgroupProductsFilter{
		HostgroupIds: []uint32{1},
	}
	newHgPrdIds := []uint32{2, 3}
	oldHgPrd := []*repo.HostgroupProduct{
		{Id: 1, HostgroupID: 1, ProductID: 1},
		{Id: 2, HostgroupID: 1, ProductID: 2},
	}
	toCreateHgPrd := []*repo.HostgroupProduct{
		{HostgroupID: 1, ProductID: 3},
	}
	hprepo.On("ListHostgroupProducts", ctx, mock.Anything, hpFilter).Return(oldHgPrd, nil)
	hprepo.On("DeleteHostgroupProducts", ctx, mock.Anything, []uint32{1}).Return(nil, nil).Once()
	hprepo.On("CreateHostgroupProducts", ctx, mock.Anything, toCreateHgPrd).Return(nil, nil).Once()

	err = usecase.HandleM2MProps(
		ctx, nil, uint32(1), newHgPrdIds, "shareProduct")
	assert.NoError(t, err)

	// hostgroup-team
	htFilter := &repo.HostgroupTeamsFilter{
		HostgroupIds: []uint32{1},
	}
	newHgTeamIds := []uint32{2, 3}
	oldHgTeams := []*repo.HostgroupTeam{
		{Id: 1, HostgroupID: 1, TeamID: 1},
		{Id: 2, HostgroupID: 1, TeamID: 2},
	}
	toCreateHgTeam := []*repo.HostgroupTeam{
		{HostgroupID: 1, TeamID: 3},
	}

	htrepo.On("ListHostgroupTeams", ctx, mock.Anything, htFilter).Return(oldHgTeams, nil)
	htrepo.On("DeleteHostgroupTeams", ctx, mock.Anything, []uint32{1}).Return(nil, nil).Once()
	htrepo.On("CreateHostgroupTeams", ctx, mock.Anything, toCreateHgTeam).Return(nil, nil).Once()

	err = usecase.HandleM2MProps(
		ctx, nil, uint32(1), newHgTeamIds, "shareTeam")
	assert.NoError(t, err)
}

func TestDeleteHostgroups(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	hgrepo := new(MockHostgroupsRepo)
	htrepo := new(MockHostgroupTeamsRepo)
	htagrepo := new(MockHostgroupTagsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	clsrepo := new(MockClustersRepo)
	dcrepo := new(MockDatacentersRepo)
	envrepo := new(MockEnvsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	teamrepo := new(MockTeamsRepo)
	prdrepo := new(MockProductsRepo)
	ahrepo := new(MockAppHostgroupsRepo)

	usecase := biz.NewHostgroupsUsecase(
		hgrepo, htrepo, hprepo, htagrepo, hfrepo, clsrepo,
		dcrepo, envrepo, ftrepo, tagrepo, teamrepo,
		prdrepo, ahrepo, authzrepo, nil, txm)

	teamrepo.On("GetTeams", ctx, mock.Anything, mock.Anything).Return(&repo.Team{
		2, "team2", "team2code", "team2l", "desc"}, nil)
	hgrepo.On("ListHostgroups", ctx, mock.Anything, mock.Anything).Return([]*repo.Hostgroup{
		{1, "hg1", "desc", 1, 1, 1, 1, 2}}, nil)
	// enforce fail
	authzcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, errors.New("enforce fail"))
	err := usecase.DeleteHostgroups(ctx, []uint32{1})
	assert.Error(t, err)
	t.Logf("delete hostgroup: %v", err)
	authzcall.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)
	// has required
	ahrepo.On("CountRequire", ctx, mock.Anything, repo.RequireHostgroup, []uint32{1}).Return(int64(1), nil)
	err = usecase.DeleteHostgroups(ctx, []uint32{1})
	assert.Error(t, err)
	t.Logf("delete hostgroup: %v", err)

}

func TestListHostgroup(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	hgrepo := new(MockHostgroupsRepo)
	htrepo := new(MockHostgroupTeamsRepo)
	htagrepo := new(MockHostgroupTagsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	clsrepo := new(MockClustersRepo)
	dcrepo := new(MockDatacentersRepo)
	envrepo := new(MockEnvsRepo)
	ftrepo := new(MockFeaturesRepo)
	tagrepo := new(MockTagsRepo)
	teamrepo := new(MockTeamsRepo)
	prdrepo := new(MockProductsRepo)
	ahrepo := new(MockAppHostgroupsRepo)

	usecase := biz.NewHostgroupsUsecase(
		hgrepo, htrepo, hprepo, htagrepo, hfrepo, clsrepo,
		dcrepo, envrepo, ftrepo, tagrepo, teamrepo,
		prdrepo, ahrepo, authzrepo, nil, txm)

	// fail
	query := &biz.ListHostgroupsFilter{
		Page:            1,
		PageSize:        100,
		TagsId:          []uint32{1},
		FeaturesId:      []uint32{1},
		ShareProductsId: []uint32{1},
		ShareTeamsId:    []uint32{1},
	}

	// hostgroup-tag filter empty
	htagcall := htagrepo.On("ListHostgroupTags", ctx, mock.Anything, &repo.HostgroupTagsFilter{
		Ids: []uint32{1},
	}).Return([]*repo.HostgroupTag{}, nil)
	_, err := usecase.ListHostgroups(ctx, query)
	htagcall.Unset()
	t.Logf("tag filter empty: %v", err)
	assert.Error(t, err)

	// hostgroup-feature filter empty
	htagcall = htagrepo.On("ListHostgroupTags", ctx, mock.Anything, &repo.HostgroupTagsFilter{
		Ids: []uint32{1},
	}).Return([]*repo.HostgroupTag{
		{HostgroupID: 1, TagID: 1, Id: 1},
	}, nil)
	hfcall := hfrepo.On("ListHostgroupFeatures", ctx, mock.Anything, &repo.HostgroupFeaturesFilter{
		Ids: []uint32{1},
	}).Return([]*repo.HostgroupFeature{}, nil)
	_, err = usecase.ListHostgroups(ctx, query)
	htagcall.Unset()
	hfcall.Unset()
	t.Logf("feature filter empty: %v", err)
	assert.Error(t, err)

	// hostgroup-share-product filter empty
	htagcall = htagrepo.On("ListHostgroupTags", ctx, mock.Anything, &repo.HostgroupTagsFilter{
		Ids: []uint32{1},
	}).Return([]*repo.HostgroupTag{
		{HostgroupID: 1, TagID: 1, Id: 1},
	}, nil)
	hfcall = hfrepo.On("ListHostgroupFeatures", ctx, mock.Anything,
		&repo.HostgroupFeaturesFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupFeature{
		{HostgroupID: 1, FeatureID: 1, Id: 1},
	}, nil)
	hpcall := hprepo.On("ListHostgroupProducts", ctx, mock.Anything,
		&repo.HostgroupProductsFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupProduct{}, nil)
	_, err = usecase.ListHostgroups(ctx, query)
	htagcall.Unset()
	hfcall.Unset()
	hpcall.Unset()
	t.Logf("share-product filter empty: %v", err)
	assert.Error(t, err)

	// hostgroup-share-team filter empty
	htagcall = htagrepo.On("ListHostgroupTags", ctx, mock.Anything, &repo.HostgroupTagsFilter{
		Ids: []uint32{1},
	}).Return([]*repo.HostgroupTag{
		{HostgroupID: 1, TagID: 1, Id: 1},
	}, nil)
	hfcall = hfrepo.On("ListHostgroupFeatures", ctx, mock.Anything,
		&repo.HostgroupFeaturesFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupFeature{
		{HostgroupID: 1, FeatureID: 1, Id: 1},
	}, nil)
	hpcall = hprepo.On("ListHostgroupProducts", ctx, mock.Anything,
		&repo.HostgroupProductsFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupProduct{
		{HostgroupID: 1, ProductID: 1, Id: 1},
	}, nil)
	htcall := htrepo.On("ListHostgroupTeams", ctx, mock.Anything,
		&repo.HostgroupTeamsFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupTeam{}, nil)
	_, err = usecase.ListHostgroups(ctx, query)
	htagcall.Unset()
	hfcall.Unset()
	hpcall.Unset()
	htcall.Unset()
	t.Logf("share-team filter empty: %v", err)
	assert.Error(t, err)

	// hostgroup repo fail
	htagcall = htagrepo.On("ListHostgroupTags", ctx, mock.Anything, &repo.HostgroupTagsFilter{
		Ids: []uint32{1},
	}).Return([]*repo.HostgroupTag{
		{HostgroupID: 1, TagID: 1, Id: 1},
	}, nil)
	hfcall = hfrepo.On("ListHostgroupFeatures", ctx, mock.Anything,
		&repo.HostgroupFeaturesFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupFeature{
		{HostgroupID: 1, FeatureID: 1, Id: 1},
	}, nil)
	hpcall = hprepo.On("ListHostgroupProducts", ctx, mock.Anything,
		&repo.HostgroupProductsFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupProduct{
		{HostgroupID: 1, ProductID: 1, Id: 1},
	}, nil)
	htcall = htrepo.On("ListHostgroupTeams", ctx, mock.Anything,
		&repo.HostgroupTeamsFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupTeam{
		{HostgroupID: 1, TeamID: 1, Id: 1},
	}, nil)
	hgcall := hgrepo.On("ListHostgroups", ctx, mock.Anything, &repo.HostgroupsFilter{
		Ids: []uint32{1},
	}).Return([]*repo.Hostgroup{}, errors.New("mock repo fail"))
	_, err = usecase.ListHostgroups(ctx, query)
	htagcall.Unset()
	hfcall.Unset()
	hpcall.Unset()
	htcall.Unset()
	hgcall.Unset()
	t.Logf("hostgroup repo fail: %v", err)
	assert.Error(t, err)

	// repo find zero
	htagcall = htagrepo.On("ListHostgroupTags", ctx, mock.Anything,
		&repo.HostgroupTagsFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupTag{
		{HostgroupID: 1, TagID: 1, Id: 1},
	}, nil)
	hfcall = hfrepo.On("ListHostgroupFeatures", ctx, mock.Anything,
		&repo.HostgroupFeaturesFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupFeature{
		{HostgroupID: 1, FeatureID: 1, Id: 1},
	}, nil)
	hpcall = hprepo.On("ListHostgroupProducts", ctx, mock.Anything,
		&repo.HostgroupProductsFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupProduct{
		{HostgroupID: 1, ProductID: 1, Id: 1},
	}, nil)
	htcall = htrepo.On("ListHostgroupTeams", ctx, mock.Anything,
		&repo.HostgroupTeamsFilter{
			Ids: []uint32{1},
		}).Return([]*repo.HostgroupTeam{
		{HostgroupID: 1, TeamID: 1, Id: 1},
	}, nil)
	hgcall = hgrepo.On("ListHostgroups", ctx, mock.Anything, &repo.HostgroupsFilter{
		Ids: []uint32{1},
	}).Return([]*repo.Hostgroup{}, nil)
	r, err := usecase.ListHostgroups(ctx, query)
	htagcall.Unset()
	hfcall.Unset()
	hpcall.Unset()
	htcall.Unset()
	hgcall.Unset()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(r))

	// empty return all
	want_hg_biz := []*biz.Hostgroup{
		{
			Id:              1,
			Name:            "test",
			Description:     "test",
			ClusterId:       uint32(1),
			DatacenterId:    uint32(1),
			EnvId:           uint32(1),
			ProductId:       uint32(1),
			TeamId:          uint32(1),
			FeaturesId:      []uint32{1, 2},
			TagsId:          []uint32{1, 2},
			ShareProductsId: []uint32{1, 2},
			ShareTeamsId:    []uint32{1, 2},
		},
	}

	want_hg_repo := []*repo.Hostgroup{
		{
			Id:           1,
			Name:         "test",
			Description:  "test",
			ClusterId:    uint32(1),
			DatacenterId: uint32(1),
			EnvId:        uint32(1),
			ProductId:    uint32(1),
			TeamId:       uint32(1),
		},
	}

	htagcall = htagrepo.On("ListHostgroupTags", ctx, mock.Anything, mock.Anything).
		Return(
			[]*repo.HostgroupTag{
				{HostgroupID: 1, TagID: 1, Id: 1},
				{HostgroupID: 1, TagID: 2, Id: 2},
			}, nil)
	hfcall = hfrepo.On("ListHostgroupFeatures", ctx, mock.Anything, mock.Anything).
		Return([]*repo.HostgroupFeature{
			{HostgroupID: 1, FeatureID: 1, Id: 1},
			{HostgroupID: 1, FeatureID: 2, Id: 2},
		}, nil)
	hpcall = hprepo.On("ListHostgroupProducts", ctx, mock.Anything, mock.Anything).
		Return([]*repo.HostgroupProduct{
			{HostgroupID: 1, ProductID: 1, Id: 1},
			{HostgroupID: 1, ProductID: 2, Id: 2},
		}, nil)
	htcall = htrepo.On("ListHostgroupTeams", ctx, mock.Anything, mock.Anything).
		Return([]*repo.HostgroupTeam{
			{HostgroupID: 1, TeamID: 1, Id: 1},
			{HostgroupID: 1, TeamID: 2, Id: 2},
		}, nil)
	hgcall = hgrepo.On("ListHostgroups", ctx, mock.Anything, &repo.HostgroupsFilter{
		Ids: []uint32{1},
	}).Return(want_hg_repo, nil)

	r, err = usecase.ListHostgroups(ctx, query)
	htagcall.Unset()
	hfcall.Unset()
	hpcall.Unset()
	htcall.Unset()
	hgcall.Unset()
	assert.NoError(t, err)
	assert.Equal(t, want_hg_biz, r)

}
