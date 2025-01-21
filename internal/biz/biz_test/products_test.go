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

func TestCreateProducts(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	prdrepo := new(MockProductsRepo)
	apprepo := new(MockApplicationsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewProductsUsecase(
		prdrepo,
		authzrepo,
		hgrepo,
		apprepo,
		hprepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Product{
		{Name: "name space", Code: "code"},
		{Name: "nameUpper", Code: "code"},
		{Name: "0-name", Code: "code"},
		{Name: "-name", Code: "code"},
		{Name: "name-", Code: "code"},
		{Name: "name_upperline", Code: "code"},
		{Name: "name", Code: "code-"},
		{Name: "name", Code: "-code"},
		{Name: "name", Code: "code_1"},
		{Name: "name", Code: "code 1"},
		{Name: "name", Code: "Code"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.CreateProducts(ctx, []*biz.Product{bc})
		assert.Error(t, err)
	}

	// enforce fail
	prd := []*biz.Product{
		{Name: "name", Code: "code"},
	}
	authzcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, errors.New("PermissionDenied"))
	err := usecase.CreateProducts(ctx, prd)
	assert.Error(t, err)
	authzcall.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// repo error
	prd = []*biz.Product{
		{Name: "name", Code: "code"},
	}
	call := prdrepo.On("CreateProducts", ctx, mock.Anything, mock.Anything).Return(errors.New("repo error"))
	err = usecase.CreateProducts(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Product{
		{Name: "name", Code: "code"},
		{Name: "name-1", Code: "code"},
		{Name: "name", Code: "code-1"},
		{Name: "name", Code: "1-code-1"},
	}
	prdrepo.On("CreateProducts", ctx, mock.Anything, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.CreateProducts(ctx, []*biz.Product{gc})
		assert.NoError(t, err)
	}
}

func TestUpdateProducts(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	prdrepo := new(MockProductsRepo)
	apprepo := new(MockApplicationsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewProductsUsecase(
		prdrepo,
		authzrepo,
		hgrepo,
		apprepo,
		hprepo,
		nil,
		txm,
	)

	// bad field
	bad_filter_cases := []*biz.Product{
		{Id: 1, Name: "name space", Code: "code"},
		{Id: 1, Name: "nameUpper", Code: "code"},
		{Id: 1, Name: "0-name", Code: "code"},
		{Id: 1, Name: "-name", Code: "code"},
		{Id: 1, Name: "name-", Code: "code"},
		{Id: 1, Name: "name_upperline", Code: "code"},
		{Id: 1, Name: "name", Code: "code-"},
		{Id: 1, Name: "name", Code: "-code"},
		{Id: 1, Name: "name", Code: "code_1"},
		{Id: 1, Name: "name", Code: "code 1"},
		{Id: 1, Name: "name", Code: "Code"},
		{Name: "name", Code: "Code"},
	}
	for _, bc := range bad_filter_cases {
		err := usecase.UpdateProducts(ctx, []*biz.Product{bc})
		assert.Error(t, err)
	}

	// enforce fail
	prd := []*biz.Product{
		{Id: 1, Name: "name", Code: "code"},
	}
	authzcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, errors.New("PermissionDenied"))
	err := usecase.UpdateProducts(ctx, prd)
	assert.Error(t, err)
	authzcall.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// repo error
	prd = []*biz.Product{
		{Id: 1, Name: "name", Code: "code"},
	}
	call := prdrepo.On("UpdateProducts", ctx, mock.Anything, mock.Anything).Return(errors.New("repo error"))
	err = usecase.UpdateProducts(ctx, prd)
	assert.Error(t, err)
	call.Unset()

	// good case
	good_cases := []*biz.Product{
		{Id: 1, Name: "name", Code: "code"},
		{Id: 1, Name: "name-1", Code: "code"},
		{Id: 1, Name: "name", Code: "code-1"},
		{Id: 1, Name: "name", Code: "1-code-1"},
	}
	prdrepo.On("UpdateProducts", ctx, mock.Anything, mock.Anything).Return(nil)
	for _, gc := range good_cases {
		err := usecase.UpdateProducts(ctx, []*biz.Product{gc})
		assert.NoError(t, err)
	}
}

func TestDeleteProducts(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	prdrepo := new(MockProductsRepo)
	apprepo := new(MockApplicationsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewProductsUsecase(
		prdrepo,
		authzrepo,
		hgrepo,
		apprepo,
		hprepo,
		nil,
		txm,
	)

	// Test case: Validation fails
	ids := []uint32{}
	err := usecase.DeleteProducts(ctx, ids)
	assert.Error(t, err)

	err = usecase.DeleteProducts(ctx, nil)
	assert.Error(t, err)

	// enforce fail
	ids = []uint32{1, 2}
	authzcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, errors.New("PermissionDenied"))
	err = usecase.DeleteProducts(ctx, ids)
	assert.Error(t, err)
	authzcall.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// Test case: failed on hostgroup need check fail
	ids = []uint32{1, 2}
	hgCall := hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(1), nil)
	appCall := apprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)
	hprepoCall := hprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)

	prdCall := prdrepo.On("DeleteProducts", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteProducts(ctx, ids)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	appCall.Unset()
	hgCall.Unset()
	hprepoCall.Unset()
	prdCall.Unset()

	// Test case: failed on app need check fail
	ids = []uint32{1, 2}
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)
	appCall = apprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(1), nil)
	hprepoCall = hprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)

	prdCall = prdrepo.On("DeleteProducts", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteProducts(ctx, ids)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	appCall.Unset()
	hgCall.Unset()
	hprepoCall.Unset()
	prdCall.Unset()

	// Test case: failed on hostgroup-product repo delete
	ids = []uint32{1, 2}
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)
	appCall = apprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)
	hprepoCall = hprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(1), nil)

	prdCall = prdrepo.On("DeleteProducts", ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteProducts(ctx, ids)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	appCall.Unset()
	hgCall.Unset()
	hprepoCall.Unset()
	prdCall.Unset()

	// repo fail
	ids = []uint32{1, 2}
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)
	appCall = apprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)
	hprepoCall = hprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)

	rerr := errors.New("mock repo fail")
	prdCall = prdrepo.On("DeleteProducts", ctx, mock.Anything, mock.Anything).
		Return(rerr)
	err = usecase.DeleteProducts(ctx, ids)
	assert.Equal(t, err, rerr)
	t.Logf("error. %v", rerr)
	appCall.Unset()
	hgCall.Unset()
	hprepoCall.Unset()
	prdCall.Unset()

	// Test case: success
	ids = []uint32{1, 2}
	hgCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)
	appCall = apprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)
	hprepoCall = hprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireProduct, ids).Return(int64(0), nil)

	prdCall = prdrepo.On("DeleteProducts", ctx, mock.Anything, mock.Anything).
		Return(nil)
	err = usecase.DeleteProducts(ctx, ids)
	assert.NoError(t, err)
	appCall.Unset()
	hgCall.Unset()
	hprepoCall.Unset()
	prdCall.Unset()
}

func TestGetProducts(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	prdrepo := new(MockProductsRepo)
	apprepo := new(MockApplicationsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewProductsUsecase(
		prdrepo,
		authzrepo,
		hgrepo,
		apprepo,
		hprepo,
		nil,
		txm,
	)

	// id == 0
	_, err := usecase.GetProducts(ctx, 0)
	t.Logf("error is %v", err)
	assert.Error(t, err)

	// repo error
	rerr := errors.New("repo error")
	call := prdrepo.On("GetProducts", ctx, uint32(1)).Return(nil, rerr)
	_, err = usecase.GetProducts(ctx, 1)
	assert.Equal(t, rerr, err)
	call.Unset()

	// success
	db_prds := repo.Product{
		ID:          1,
		Name:        "prd1",
		Code:        "prd1",
		Description: "prd1 description",
	}

	biz_prds := biz.Product{
		Id:          1,
		Name:        "prd1",
		Code:        "prd1",
		Description: "prd1 description",
	}
	prdrepo.On("GetProducts", ctx, uint32(1)).Return(&db_prds, nil)
	prd, err := usecase.GetProducts(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, *prd)
}

func TestListProducts(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "admin")
	authzrepo := new(MockAuthzRepo)
	txm := new(MockTXManager)
	prdrepo := new(MockProductsRepo)
	apprepo := new(MockApplicationsRepo)
	hprepo := new(MockHostgroupProductsRepo)
	hgrepo := new(MockHostgroupsRepo)
	usecase := biz.NewProductsUsecase(
		prdrepo,
		authzrepo,
		hgrepo,
		apprepo,
		hprepo,
		nil,
		txm,
	)

	bad_filter := []biz.ListProductsFilter{
		{Codes: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}},
		{Names: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}},
		{Ids: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{Page: 1, PageSize: 201},
		{Page: 1, PageSize: 0},
		{Page: 0, PageSize: 10},
	}
	for _, bc := range bad_filter {
		_, err := usecase.ListProducts(ctx, &bc)
		t.Logf("error. %v", err)
		assert.Error(t, err)
	}

	// repo error
	filter := biz.ListProductsFilter{
		Page:     1,
		PageSize: 10,
	}
	call := prdrepo.On("ListProducts", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("repo error"))
	_, err := usecase.ListProducts(ctx, &filter)
	assert.Error(t, err)
	call.Unset()

	// success
	db_prds := []*repo.Product{
		{
			ID:          1,
			Code:        "code1",
			Name:        "name1",
			Description: "description1",
		},
		{
			ID:          2,
			Code:        "code2",
			Name:        "name2",
			Description: "description2",
		},
	}

	biz_prds := []*biz.Product{
		{
			Id:          1,
			Code:        "code1",
			Name:        "name1",
			Description: "description1",
		},
		{
			Id:          2,
			Code:        "code2",
			Name:        "name2",
			Description: "description2",
		},
	}
	prdrepo.On("ListProducts", mock.Anything, mock.Anything, mock.Anything).Return(db_prds, nil)
	prds, err := usecase.ListProducts(ctx, &filter)
	assert.NoError(t, err)
	assert.Equal(t, biz_prds, prds)
}
