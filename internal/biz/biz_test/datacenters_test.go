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

func TestCreateDatacenters(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "user1")
	txm := new(MockTXManager)
	dcrepo := new(MockDatacentersRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewDatacentersUsecase(
		dcrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Datacenter{
		{Name: "name space"},
		{Name: "nameUpper"},
		{Name: "0-name"},
		{Name: "-name"},
		{Name: "name-"},
		{Name: "name_upperline"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.CreateDatacenters(ctx, []*biz.Datacenter{bc})
		assert.Error(t, err)
	}

	prd := []*biz.Datacenter{
		{Name: "name"}}
	// enforce error
	authz_call := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	err := usecase.CreateDatacenters(ctx, prd)
	assert.Error(t, err)
	authz_call.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything, mock.Anything).Return(true, nil)

	// repo error

	call := dcrepo.On("CreateDatacenters", ctx, mock.Anything, mock.Anything).Return(errors.New("repo error"))
	err = usecase.CreateDatacenters(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Datacenter{
		{Name: "name"},
		{Name: "name-1"},
		{Name: "name"},
		{Name: "name"},
	}
	dcrepo.On("CreateDatacenters", ctx, mock.Anything, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.CreateDatacenters(ctx, []*biz.Datacenter{gc})
		assert.NoError(t, err)
	}
}

func TestUpdateDatacenters(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "user1")
	txm := new(MockTXManager)
	dcrepo := new(MockDatacentersRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewDatacentersUsecase(
		dcrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Datacenter{
		{Id: 1, Name: "name space"},
		{Id: 1, Name: "nameUpper"},
		{Id: 1, Name: "0-name"},
		{Id: 1, Name: "-name"},
		{Id: 1, Name: "name-"},
		{Id: 1, Name: "name_upperline"},
		{Name: "name"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.UpdateDatacenters(ctx, []*biz.Datacenter{bc})
		assert.Error(t, err)
	}

	prd := []*biz.Datacenter{
		{Id: 1, Name: "name"},
	}
	// enforce error
	call_authz := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	err := usecase.UpdateDatacenters(ctx, prd)
	assert.Error(t, err)
	call_authz.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything, mock.Anything).Return(true, nil)

	// repo error
	call := dcrepo.On("UpdateDatacenters", ctx, mock.Anything, mock.Anything).Return(errors.New("repo error"))
	err = usecase.UpdateDatacenters(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Datacenter{
		{Id: 1, Name: "name"},
		{Id: 1, Name: "name-1"},
	}
	dcrepo.On("UpdateDatacenters", ctx, mock.Anything, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.UpdateDatacenters(ctx, []*biz.Datacenter{gc})
		assert.NoError(t, err)
	}
}

func TestDeleteDatacenters(t *testing.T) {

	ctx := context.WithValue(context.Background(), data.CtxUserName, "user1")
	txm := new(MockTXManager)
	dcrepo := new(MockDatacentersRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewDatacentersUsecase(
		dcrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)

	// Test case: Validation fails
	ids := []uint32{}
	err := usecase.DeleteDatacenters(ctx, ids)
	assert.Error(t, err)

	err = usecase.DeleteDatacenters(ctx, nil)
	assert.Error(t, err)

	ids = []uint32{1, 2}
	// enforce error
	call_authz := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	err = usecase.DeleteDatacenters(ctx, ids)
	assert.Error(t, err)
	call_authz.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything, mock.Anything).Return(true, nil)

	// Test case: failed on hostgroup need check fail
	hgCall := hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireDatacenter, ids).Return(int64(1), nil)

	dccall := dcrepo.On("DeleteDatacenters", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteDatacenters(ctx, ids)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	hgCall.Unset()
	dccall.Unset()

	// repo fail
	ids = []uint32{1, 2}
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireDatacenter, ids).Return(int64(0), nil)

	rerr := errors.New("mock repo fail")
	dccall = dcrepo.On("DeleteDatacenters", ctx, mock.Anything, mock.Anything).
		Return(rerr)
	err = usecase.DeleteDatacenters(ctx, ids)
	assert.Equal(t, err, rerr)
	t.Logf("error. %v", rerr)
	hgCall.Unset()
	dccall.Unset()

	// Test case: success
	ids = []uint32{1, 2}
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireDatacenter, ids).Return(int64(0), nil)

	dccall = dcrepo.On("DeleteDatacenters", ctx, mock.Anything, mock.Anything).
		Return(nil)
	err = usecase.DeleteDatacenters(ctx, ids)
	assert.NoError(t, err)
	hgCall.Unset()
	dccall.Unset()
}

func TestGetDatacenters(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	dcrepo := new(MockDatacentersRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewDatacentersUsecase(
		dcrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)
	// id == 0
	_, err := usecase.GetDatacenters(ctx, 0)
	t.Logf("error is %v", err)
	assert.Error(t, err)

	// repo error
	rerr := errors.New("repo error")
	call := dcrepo.On("GetDatacenters", ctx, uint32(1)).Return(nil, rerr)
	_, err = usecase.GetDatacenters(ctx, 1)
	assert.Equal(t, rerr, err)
	call.Unset()

	// success
	db_prds := repo.Datacenter{
		ID:   1,
		Name: "prd1",
	}

	biz_prds := biz.Datacenter{
		Id:   1,
		Name: "prd1",
	}
	dcrepo.On("GetDatacenters", ctx, uint32(1)).Return(&db_prds, nil)
	prd, err := usecase.GetDatacenters(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, *prd)
}

func TestListDatacenters(t *testing.T) {
	ctx := context.Background()
	txm := new(MockTXManager)
	dcrepo := new(MockDatacentersRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewDatacentersUsecase(
		dcrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)

	bad_filter := []biz.ListDatacentersFilter{
		{Names: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}},
		{Ids: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{Page: 1, PageSize: 201},
		{Page: 1, PageSize: 0},
		{Page: 0, PageSize: 10},
	}
	for _, bc := range bad_filter {
		_, err := usecase.ListDatacenters(ctx, &bc)
		t.Logf("error. %v", err)
		assert.Error(t, err)
	}

	// repo error
	filter := biz.ListDatacentersFilter{
		Page:     1,
		PageSize: 10,
	}
	call := dcrepo.On("ListDatacenters", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("repo error"))
	_, err := usecase.ListDatacenters(ctx, &filter)
	assert.Error(t, err)
	call.Unset()

	// success
	db_prds := []*repo.Datacenter{
		{
			ID:   1,
			Name: "name1",
		},
		{
			ID:   2,
			Name: "name2",
		},
	}

	biz_prds := []*biz.Datacenter{
		{
			Id:   1,
			Name: "name1",
		},
		{
			Id:   2,
			Name: "name2",
		},
	}
	dcrepo.On("ListDatacenters", mock.Anything, mock.Anything, mock.Anything).Return(db_prds, nil)
	prds, err := usecase.ListDatacenters(ctx, &filter)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, prds)
}
