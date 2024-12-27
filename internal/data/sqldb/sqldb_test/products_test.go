package sqldb_test

import (
	"context"
	"testing"

	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var prodsRepo repo.ProductsRepo

func initProdsRepo() {
	dataMem := getDataMem()
	prodsRepo, _ = sqldb.NewProductsRepoGorm(dataMem, logger)
}

func createBaseProds(t *testing.T, data []*repo.Product) {
	initProdsRepo()
	if data == nil {
		data = []*repo.Product{
			{Name: "test", Code: "test"},
			{Name: "test1", Code: "test1"},
		}
	}
	if err := prodsRepo.CreateProducts(context.Background(), data); err != nil {
		t.Fatal(err)
	}

}

func TestProdsRepoGorm(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"CreateProds_Success", testCreateProdsSuccess},
		{"CreateProds_Error", testCreateProdsError},
		{"UpdateProds_Success", testUpdateProdsSuccess},
		{"UpdateProds_Error", testUpdateProdsError},
		{"DeleteProds_Success", testDeleteProdsSuccess},
		{"DeleteProds_Error", testDeleteProdsError},
		{"GetProds_Success", testGetProdsSuccess},
		{"GetProds_Error", testGetProdsError},
		{"ListProds_emptyFilter_all", testListProds_emptyFilter_all},
		{"ListProds_id_partial", testListProds_id_partial},
		{"ListProds_page_partial", testListProds_page_partial},
		{"ListProds_keys_partial", testListProds_keys_partial},
		{"ListProds_code_partial", testListProds_code_partial},
		{"ListProds_nil_all", testListProds_nil_all},
		{"CountProds_partial", testCountProds_partial},
		{"CountProds_subpartial", testCountProds_subpartial},
		{"CountProds_all", testCountProds_all},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func testCreateProdsSuccess(t *testing.T) {

	initProdsRepo()

	prods := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	err := prodsRepo.CreateProducts(context.Background(), prods)
	assert.NoError(t, err)
}

func testCreateProdsError(t *testing.T) {
	createBaseProds(t, nil)
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test", Code: "test"},
	}
	err := prodsRepo.CreateProducts(context.Background(), data)
	assert.Error(t, err)
}

func testUpdateProdsSuccess(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	createBaseProds(t, data)

	data[0].Name = "test2"
	err := prodsRepo.UpdateProducts(context.Background(), data)
	assert.NoError(t, err)
}

func testUpdateProdsError(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	createBaseProds(t, data)
	data[0].Code = "test1"
	err := prodsRepo.UpdateProducts(context.Background(), data)
	assert.Error(t, err)
}

func testDeleteProdsSuccess(t *testing.T) {
	createBaseProds(t, nil)
	err := prodsRepo.DeleteProducts(context.Background(), []uint32{1})
	assert.NoError(t, err)
}

func testDeleteProdsError(t *testing.T) {

	createBaseProds(t, nil)
	err := prodsRepo.DeleteProducts(context.Background(), []uint32{99})
	assert.Error(t, err)
}

func testGetProdsSuccess(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	createBaseProds(t, data)

	prod, err := prodsRepo.GetProducts(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, prod)
	assert.Equal(t, uint32(1), prod.ID)
	assert.Equal(t, "test", prod.Name)
	assert.Equal(t, "test", prod.Code)
}

func testGetProdsError(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	createBaseProds(t, data)
	tag, err := prodsRepo.GetProducts(context.Background(), 99)
	assert.Error(t, err)
	assert.Nil(t, tag)
}

func testListProds_emptyFilter_all(t *testing.T) {

	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	createBaseProds(t, data)
	tags, err := prodsRepo.ListProducts(context.Background(), nil, &repo.ProductsFilter{})
	assert.NoError(t, err)
	assert.Len(t, tags, 2)
}
func testListProds_id_partial(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	createBaseProds(t, data)
	_tags, err := prodsRepo.ListProducts(context.Background(), nil, &repo.ProductsFilter{Ids: []uint32{1}})
	assert.NoError(t, err)
	assert.Len(t, _tags, 1)
}

func testListProds_page_partial(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	createBaseProds(t, data)
	filter := &repo.ProductsFilter{
		Page:     2,
		PageSize: 1,
	}
	_data, err := prodsRepo.ListProducts(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data[1:], _data)
}

func testListProds_keys_partial(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	createBaseProds(t, data)
	filter := &repo.ProductsFilter{
		Names: []string{"test"},
	}
	_tags, err := prodsRepo.ListProducts(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data[:1], _tags)
}
func testListProds_code_partial(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
	}
	createBaseProds(t, data)
	filter := &repo.ProductsFilter{
		Codes: []string{"test", "test1"},
	}
	_data, err := prodsRepo.ListProducts(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data, _data)
}
func testListProds_nil_all(t *testing.T) {
	createBaseProds(t, nil)
	_tags, err := prodsRepo.ListProducts(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(_tags))
}

func testCountProds_partial(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
		{Name: "test2", Code: "test2"},
		{Name: "test3", Code: "test3"},
	}
	createBaseProds(t, data)
	filter := &repo.ProductsFilter{
		Ids: []uint32{2, 3},
	}
	count, err := prodsRepo.CountProducts(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}
func testCountProds_subpartial(t *testing.T) {
	data := []*repo.Product{
		{Name: "test", Code: "test"},
		{Name: "test1", Code: "test1"},
		{Name: "test2", Code: "test2"},
		{Name: "test3", Code: "test3"},
	}
	createBaseProds(t, data)
	filter := &repo.ProductsFilter{
		Ids: []uint32{2, 99},
	}
	count, err := prodsRepo.CountProducts(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func testCountProds_all(t *testing.T) {
	createBaseProds(t, nil)
	count, err := prodsRepo.CountProducts(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}
