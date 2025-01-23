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

func TestCreateEnvs(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	txm := new(MockTXManager)
	envrepo := new(MockEnvsRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewEnvsUsecase(
		envrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Env{
		{Name: "name space"},
		{Name: "nameUpper"},
		{Name: "0-name"},
		{Name: "-name"},
		{Name: "name-"},
		{Name: "name_upperline"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.CreateEnvs(ctx, []*biz.Env{bc})
		assert.Error(t, err)
	}

	prd := []*biz.Env{
		{Name: "name"}}
	// enforce error
	call_authz := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).
		Return(false, nil)
	err := usecase.CreateEnvs(ctx, prd)
	assert.Error(t, err)
	t.Log(err)
	call_authz.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// repo error
	call := envrepo.On("CreateEnvs", ctx, mock.Anything, mock.Anything).Return(errors.New("repo error"))
	err = usecase.CreateEnvs(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Env{
		{Name: "name"},
		{Name: "name-1"},
		{Name: "name"},
		{Name: "name"},
	}
	envrepo.On("CreateEnvs", ctx, mock.Anything, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.CreateEnvs(ctx, []*biz.Env{gc})
		assert.NoError(t, err)
	}
}

func TestUpdateEnvs(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	txm := new(MockTXManager)
	envrepo := new(MockEnvsRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewEnvsUsecase(
		envrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Env{
		{Id: 1, Name: "name space"},
		{Id: 1, Name: "nameUpper"},
		{Id: 1, Name: "0-name"},
		{Id: 1, Name: "-name"},
		{Id: 1, Name: "name-"},
		{Id: 1, Name: "name_upperline"},
		{Name: "name"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.UpdateEnvs(ctx, []*biz.Env{bc})
		assert.Error(t, err)
	}

	prd := []*biz.Env{
		{Id: 1, Name: "name"},
	}
	// enforce error
	call_authz := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).
		Return(false, nil)
	err := usecase.UpdateEnvs(ctx, prd)
	assert.Error(t, err)
	t.Log(err)
	call_authz.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// repo error
	call := envrepo.On("UpdateEnvs", ctx, mock.Anything, mock.Anything).Return(errors.New("repo error"))
	err = usecase.UpdateEnvs(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Env{
		{Id: 1, Name: "name"},
		{Id: 1, Name: "name-1"},
	}
	envrepo.On("UpdateEnvs", ctx, mock.Anything, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.UpdateEnvs(ctx, []*biz.Env{gc})
		assert.NoError(t, err)
	}
}

func TestDeleteEnvs(t *testing.T) {

	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	txm := new(MockTXManager)
	envrepo := new(MockEnvsRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewEnvsUsecase(
		envrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)

	// Test case: Validation fails
	ids := []uint32{}
	err := usecase.DeleteEnvs(ctx, ids)
	assert.Error(t, err)

	err = usecase.DeleteEnvs(ctx, nil)
	assert.Error(t, err)

	// enforce error
	call_authz := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).
		Return(false, nil)
	err = usecase.DeleteEnvs(ctx, ids)
	assert.Error(t, err)
	t.Log(err)
	call_authz.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// Test case: failed on hostgroup need check fail
	ids = []uint32{1, 2}
	hgCall := hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireEnv, ids).Return(int64(1), nil)

	envCall := envrepo.On("DeleteEnvs", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteEnvs(ctx, ids)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	hgCall.Unset()
	envCall.Unset()

	ids = []uint32{1, 2}

	// repo fail
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireEnv, ids).Return(int64(0), nil)

	rerr := errors.New("mock repo fail")
	envCall = envrepo.On("DeleteEnvs", ctx, mock.Anything, mock.Anything).
		Return(rerr)
	err = usecase.DeleteEnvs(ctx, ids)
	assert.Equal(t, err, rerr)
	t.Logf("error. %v", rerr)
	hgCall.Unset()
	envCall.Unset()

	// Test case: success
	ids = []uint32{1, 2}
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireEnv, ids).Return(int64(0), nil)

	envCall = envrepo.On("DeleteEnvs", ctx, mock.Anything, mock.Anything).
		Return(nil)
	err = usecase.DeleteEnvs(ctx, ids)
	assert.NoError(t, err)
	hgCall.Unset()
	envCall.Unset()
}

func TestGetEnvs(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	txm := new(MockTXManager)
	envrepo := new(MockEnvsRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewEnvsUsecase(
		envrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)
	// id == 0
	_, err := usecase.GetEnvs(ctx, 0)
	t.Logf("error is %v", err)
	assert.Error(t, err)

	// repo error
	rerr := errors.New("repo error")
	call := envrepo.On("GetEnvs", ctx, uint32(1)).Return(nil, rerr)
	_, err = usecase.GetEnvs(ctx, 1)
	assert.Equal(t, rerr, err)
	call.Unset()

	// success
	db_prds := repo.Env{
		ID:   1,
		Name: "prd1",
	}

	biz_prds := biz.Env{
		Id:   1,
		Name: "prd1",
	}
	envrepo.On("GetEnvs", ctx, uint32(1)).Return(&db_prds, nil)
	prd, err := usecase.GetEnvs(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, *prd)
}

func TestListEnvs(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	txm := new(MockTXManager)
	envrepo := new(MockEnvsRepo)
	hgrepo := new(MockHostgroupsRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewEnvsUsecase(
		envrepo,
		authzrepo,
		hgrepo,
		nil,
		txm,
	)

	bad_filter := []biz.ListEnvsFilter{
		{Names: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}},
		{Ids: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{Page: 1, PageSize: 201},
		{Page: 1, PageSize: 0},
		{Page: 0, PageSize: 10},
	}
	for _, bc := range bad_filter {
		_, err := usecase.ListEnvs(ctx, &bc)
		t.Logf("error. %v", err)
		assert.Error(t, err)
	}

	// repo error
	filter := biz.ListEnvsFilter{
		Page:     1,
		PageSize: 10,
	}
	call := envrepo.On("ListEnvs", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("repo error"))
	_, err := usecase.ListEnvs(ctx, &filter)
	assert.Error(t, err)
	call.Unset()

	// success
	db_prds := []*repo.Env{
		{
			ID:   1,
			Name: "name1",
		},
		{
			ID:   2,
			Name: "name2",
		},
	}

	biz_prds := []*biz.Env{
		{
			Id:   1,
			Name: "name1",
		},
		{
			Id:   2,
			Name: "name2",
		},
	}
	envrepo.On("ListEnvs", mock.Anything, mock.Anything, mock.Anything).Return(db_prds, nil)
	prds, err := usecase.ListEnvs(ctx, &filter)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, prds)
}
