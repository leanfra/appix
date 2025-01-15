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

func TestCreateFeatures(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "user")
	txm := new(MockTXManager)
	ftrepo := new(MockFeaturesRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	afrepo := new(MockAppFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewFeaturesUsecase(
		ftrepo,
		authzrepo,
		hfrepo,
		afrepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Feature{
		{Name: "name space", Value: "code"},
		{Name: "nameUpper", Value: "code"},
		{Name: "0-name", Value: "code"},
		{Name: "-name", Value: "code"},
		{Name: "name-", Value: "code"},
		{Name: "name_upperline", Value: "code"},
		{Name: "name", Value: "code-"},
		{Name: "name", Value: "-code"},
		{Name: "name", Value: "code_1"},
		{Name: "name", Value: "code 1"},
		{Name: "name", Value: "Code"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.CreateFeatures(ctx, []*biz.Feature{bc})
		assert.Error(t, err)
	}

	prd := []*biz.Feature{
		{Name: "name", Value: "code"},
	}
	// enforce error
	enforceCall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).
		Return(false, nil)
	err := usecase.CreateFeatures(ctx, prd)
	assert.Error(t, err)
	enforceCall.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// repo error
	call := ftrepo.On("CreateFeatures", ctx, mock.Anything, mock.Anything).Return(errors.New("repo error"))
	err = usecase.CreateFeatures(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Feature{
		{Name: "name", Value: "code"},
		{Name: "name-1", Value: "code"},
		{Name: "name", Value: "code-1"},
		{Name: "name", Value: "1-code-1"},
	}
	ftrepo.On("CreateFeatures", ctx, mock.Anything, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.CreateFeatures(ctx, []*biz.Feature{gc})
		assert.NoError(t, err)
	}
}

func TestUpdateFeatures(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "user")
	txm := new(MockTXManager)
	ftrepo := new(MockFeaturesRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	afrepo := new(MockAppFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewFeaturesUsecase(
		ftrepo,
		authzrepo,
		hfrepo,
		afrepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Feature{
		{Id: 1, Name: "name space", Value: "code"},
		{Id: 1, Name: "nameUpper", Value: "code"},
		{Id: 1, Name: "0-name", Value: "code"},
		{Id: 1, Name: "-name", Value: "code"},
		{Id: 1, Name: "name-", Value: "code"},
		{Id: 1, Name: "name_upperline", Value: "code"},
		{Id: 1, Name: "name", Value: "code-"},
		{Id: 1, Name: "name", Value: "-code"},
		{Id: 1, Name: "name", Value: "code_1"},
		{Id: 1, Name: "name", Value: "code 1"},
		{Id: 1, Name: "name", Value: "Code"},
		{Name: "name", Value: "Code"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.UpdateFeatures(ctx, []*biz.Feature{bc})
		assert.Error(t, err)
	}

	prd := []*biz.Feature{
		{Id: 1, Name: "name", Value: "code"},
	}
	// enforce error
	enforceCall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).
		Return(false, nil)
	err := usecase.UpdateFeatures(ctx, prd)
	assert.Error(t, err)
	enforceCall.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// repo error
	call := ftrepo.On("UpdateFeatures", ctx, mock.Anything, mock.Anything).Return(errors.New("repo error"))
	err = usecase.UpdateFeatures(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Feature{
		{Id: 1, Name: "name", Value: "code"},
		{Id: 1, Name: "name-1", Value: "code"},
		{Id: 1, Name: "name", Value: "code-1"},
		{Id: 1, Name: "name", Value: "1-code-1"},
	}
	ftrepo.On("UpdateFeatures", ctx, mock.Anything, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.UpdateFeatures(ctx, []*biz.Feature{gc})
		assert.NoError(t, err)
	}
}

func TestDeleteFeatures(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "user")
	txm := new(MockTXManager)
	ftrepo := new(MockFeaturesRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	afrepo := new(MockAppFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewFeaturesUsecase(
		ftrepo,
		authzrepo,
		hfrepo,
		afrepo,
		nil,
		txm,
	)
	// Test case: Validation fails
	ids := []uint32{}
	err := usecase.DeleteFeatures(ctx, ids)
	assert.Error(t, err)

	err = usecase.DeleteFeatures(ctx, nil)
	assert.Error(t, err)

	ids = []uint32{1, 2}
	// enforce error
	enforceCall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).
		Return(false, nil)
	err = usecase.DeleteFeatures(ctx, ids)
	assert.Error(t, err)
	enforceCall.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// Test case: failed on hostgroup need check fail
	hfCall := hfrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireFeature, ids).Return(int64(1), nil)
	afCall := afrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireFeature, ids).Return(int64(0), nil)

	ftCall := ftrepo.On("DeleteFeatures", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteFeatures(ctx, ids)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	afCall.Unset()
	hfCall.Unset()
	ftCall.Unset()

	// Test case: failed on app need check fail
	ids = []uint32{1, 2}
	hfCall = hfrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireFeature, ids).Return(int64(0), nil)
	afCall = afrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireFeature, ids).Return(int64(1), nil)

	ftCall = ftrepo.On("DeleteFeatures",
		ctx, mock.Anything).Return(nil)
	err = usecase.DeleteFeatures(ctx, ids)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	afCall.Unset()
	hfCall.Unset()
	ftCall.Unset()

	// repo fail
	ids = []uint32{1, 2}
	hfCall = hfrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireFeature, ids).Return(int64(0), nil)
	afCall = afrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireFeature, ids).Return(int64(0), nil)

	rerr := errors.New("mock repo fail")
	ftCall = ftrepo.On("DeleteFeatures", ctx, mock.Anything).
		Return(rerr)
	err = usecase.DeleteFeatures(ctx, ids)
	assert.Equal(t, err, rerr)
	t.Logf("error. %v", rerr)
	afCall.Unset()
	hfCall.Unset()
	ftCall.Unset()

	// Test case: success
	ids = []uint32{1, 2}
	hfCall = hfrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireFeature, ids).Return(int64(0), nil)
	afCall = afrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireFeature, ids).Return(int64(0), nil)

	ftCall = ftrepo.On("DeleteFeatures", ctx, mock.Anything).
		Return(nil)
	err = usecase.DeleteFeatures(ctx, ids)
	assert.NoError(t, err)
	afCall.Unset()
	hfCall.Unset()
	ftCall.Unset()
}

func TestGetFeatures(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "user")
	txm := new(MockTXManager)
	ftrepo := new(MockFeaturesRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	afrepo := new(MockAppFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewFeaturesUsecase(
		ftrepo,
		authzrepo,
		hfrepo,
		afrepo,
		nil,
		txm,
	)

	// id == 0
	_, err := usecase.GetFeatures(ctx, 0)
	t.Logf("error is %v", err)
	assert.Error(t, err)

	// repo error
	rerr := errors.New("repo error")
	call := ftrepo.On("GetFeatures", ctx, uint32(1)).Return(nil, rerr)
	_, err = usecase.GetFeatures(ctx, 1)
	assert.Equal(t, rerr, err)
	call.Unset()

	// success
	db_prds := repo.Feature{
		Id:    1,
		Name:  "prd1",
		Value: "prd1",
	}

	biz_prds := biz.Feature{
		Id:    1,
		Name:  "prd1",
		Value: "prd1",
	}
	ftrepo.On("GetFeatures", ctx, uint32(1)).Return(&db_prds, nil)
	prd, err := usecase.GetFeatures(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, *prd)
}

func TestListFeatures(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.UserName, "user")
	txm := new(MockTXManager)
	ftrepo := new(MockFeaturesRepo)
	hfrepo := new(MockHostgroupFeaturesRepo)
	afrepo := new(MockAppFeaturesRepo)
	authzrepo := new(MockAuthzRepo)
	usecase := biz.NewFeaturesUsecase(
		ftrepo,
		authzrepo,
		hfrepo,
		afrepo,
		nil,
		txm,
	)

	bad_filter := []biz.ListFeaturesFilter{
		{Names: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}},
		{Ids: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{Kvs: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}},
		{Kvs: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}},
		{Page: 1, PageSize: 201},
		{Page: 1, PageSize: 0},
		{Page: 0, PageSize: 10},
	}
	for _, bc := range bad_filter {
		_, err := usecase.ListFeatures(ctx, &bc)
		t.Logf("error. %v", err)
		assert.Error(t, err)
	}

	// repo error
	filter := biz.ListFeaturesFilter{
		Page:     1,
		PageSize: 10,
	}
	call := ftrepo.On("ListFeatures", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("repo error"))
	_, err := usecase.ListFeatures(ctx, &filter)
	assert.Error(t, err)
	call.Unset()

	// success
	db_prds := []*repo.Feature{
		{
			Id:    1,
			Value: "code1",
			Name:  "name1",
		},
		{
			Id:    2,
			Value: "code2",
			Name:  "name2",
		},
	}

	biz_prds := []*biz.Feature{
		{
			Id:    1,
			Value: "code1",
			Name:  "name1",
		},
		{
			Id:    2,
			Value: "code2",
			Name:  "name2",
		},
	}
	ftrepo.On("ListFeatures", mock.Anything, mock.Anything, mock.Anything).Return(db_prds, nil)
	prds, err := usecase.ListFeatures(ctx, &filter)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, prds)
}
