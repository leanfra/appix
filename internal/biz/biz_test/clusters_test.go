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

func TestCreateClusters(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	clsrepo := new(MockClustersRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewClustersUsecase(
		clsrepo,
		hgrepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Cluster{
		{Name: "name space"},
		{Name: "nameUpper"},
		{Name: "0-name"},
		{Name: "-name"},
		{Name: "name-"},
		{Name: "name_upperline"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.CreateClusters(ctx, []*biz.Cluster{bc})
		assert.Error(t, err)
	}

	// repo error
	prd := []*biz.Cluster{
		{Name: "name"}}
	call := clsrepo.On("CreateClusters", ctx, mock.Anything).Return(errors.New("repo error"))
	err := usecase.CreateClusters(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Cluster{
		{Name: "name"},
		{Name: "name-1"},
		{Name: "name"},
		{Name: "name"},
	}
	clsrepo.On("CreateClusters", ctx, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.CreateClusters(ctx, []*biz.Cluster{gc})
		assert.NoError(t, err)
	}
}

func TestUpdateClusters(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	clsrepo := new(MockClustersRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewClustersUsecase(
		clsrepo,
		hgrepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Cluster{
		{Id: 1, Name: "name space"},
		{Id: 1, Name: "nameUpper"},
		{Id: 1, Name: "0-name"},
		{Id: 1, Name: "-name"},
		{Id: 1, Name: "name-"},
		{Id: 1, Name: "name_upperline"},
		{Name: "name"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.UpdateClusters(ctx, []*biz.Cluster{bc})
		assert.Error(t, err)
	}

	// repo error
	prd := []*biz.Cluster{
		{Id: 1, Name: "name"},
	}
	call := clsrepo.On("UpdateClusters", ctx, mock.Anything).Return(errors.New("repo error"))
	err := usecase.UpdateClusters(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Cluster{
		{Id: 1, Name: "name"},
		{Id: 1, Name: "name-1"},
	}
	clsrepo.On("UpdateClusters", ctx, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.UpdateClusters(ctx, []*biz.Cluster{gc})
		assert.NoError(t, err)
	}
}

func TestDeleteClusters(t *testing.T) {

	ctx := context.Background()
	txm := new(MockTXManager)
	clsrepo := new(MockClustersRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewClustersUsecase(
		clsrepo,
		hgrepo,
		nil,
		txm,
	)

	// Test case: Validation fails
	ids := []uint32{}
	err := usecase.DeleteClusters(ctx, ids)
	assert.Error(t, err)

	err = usecase.DeleteClusters(ctx, nil)
	assert.Error(t, err)

	// Test case: failed on hostgroup need check fail
	ids = []uint32{1, 2}
	hgCall := hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireCluster, ids).Return(int64(1), nil)

	dccall := clsrepo.On("DeleteClusters", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteClusters(ctx, ids)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	hgCall.Unset()
	dccall.Unset()

	// repo fail
	ids = []uint32{1, 2}
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireCluster, ids).Return(int64(0), nil)

	rerr := errors.New("mock repo fail")
	dccall = clsrepo.On("DeleteClusters", ctx, mock.Anything, mock.Anything).
		Return(rerr)
	err = usecase.DeleteClusters(ctx, ids)
	assert.Equal(t, err, rerr)
	t.Logf("error. %v", rerr)
	hgCall.Unset()
	dccall.Unset()

	// Test case: success
	ids = []uint32{1, 2}
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireCluster, ids).Return(int64(0), nil)

	dccall = clsrepo.On("DeleteClusters", ctx, mock.Anything, mock.Anything).
		Return(nil)
	err = usecase.DeleteClusters(ctx, ids)
	assert.NoError(t, err)
	hgCall.Unset()
	dccall.Unset()
}

func TestGetClusters(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	clsrepo := new(MockClustersRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewClustersUsecase(
		clsrepo,
		hgrepo,
		nil,
		txm,
	)
	// id == 0
	_, err := usecase.GetClusters(ctx, 0)
	t.Logf("error is %v", err)
	assert.Error(t, err)

	// repo error
	rerr := errors.New("repo error")
	call := clsrepo.On("GetClusters", ctx, uint32(1)).Return(nil, rerr)
	_, err = usecase.GetClusters(ctx, 1)
	assert.Equal(t, rerr, err)
	call.Unset()

	// success
	db_prds := repo.Cluster{
		ID:   1,
		Name: "prd1",
	}

	biz_prds := biz.Cluster{
		Id:   1,
		Name: "prd1",
	}
	clsrepo.On("GetClusters", ctx, uint32(1)).Return(&db_prds, nil)
	prd, err := usecase.GetClusters(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, *prd)
}

func TestListClusters(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	clsrepo := new(MockClustersRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewClustersUsecase(
		clsrepo,
		hgrepo,
		nil,
		txm,
	)

	bad_filter := []biz.ListClustersFilter{
		{Names: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}},
		{Ids: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{Page: 1, PageSize: 201},
		{Page: 1, PageSize: 0},
		{Page: 0, PageSize: 10},
	}
	for _, bc := range bad_filter {
		_, err := usecase.ListClusters(ctx, &bc)
		t.Logf("error. %v", err)
		assert.Error(t, err)
	}

	// repo error
	filter := biz.ListClustersFilter{
		Page:     1,
		PageSize: 10,
	}
	call := clsrepo.On("ListClusters", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("repo error"))
	_, err := usecase.ListClusters(ctx, &filter)
	assert.Error(t, err)
	call.Unset()

	// success
	db_prds := []*repo.Cluster{
		{
			ID:   1,
			Name: "name1",
		},
		{
			ID:   2,
			Name: "name2",
		},
	}

	biz_prds := []*biz.Cluster{
		{
			Id:   1,
			Name: "name1",
		},
		{
			Id:   2,
			Name: "name2",
		},
	}
	clsrepo.On("ListClusters", mock.Anything, mock.Anything, mock.Anything).Return(db_prds, nil)
	prds, err := usecase.ListClusters(ctx, &filter)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, prds)
}
