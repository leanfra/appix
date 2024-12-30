package biz_test

import (
	"appix/internal/biz"
	"appix/internal/data/repo"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateApp(t *testing.T) {
	ctx := context.Background()
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
		prdrepo, ahrepo, nil, txm)

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

	// props count 0
	hg := []*biz.Hostgroup{
		{0, "name", "desc", 1, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
	}
	//// cluster
	clscall := clsrepo.On("CountClusters", ctx, mock.Anything, mock.Anything).Return(int64(0), nil)
	err := usecase.CreateHostgroups(ctx, hg)
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

	/* // create hostgroup fail
	hg := []*biz.Hostgroup{
		{0, "name", "desc", 1, 1, 1, 1, 1, []uint32{1, 2}, []uint32{1, 2}, []uint32{2, 3}, []uint32{2, 3}},
	}
	call := hgrepo.On("CreateHostgroups", ctx, mock.Anything, mock.Anything).Return(errors.New("create hostgroup fail"))
	err := usecase.CreateHostgroups(ctx, hg)
	assert.Error(t, err)
	call.Unset() */
}
