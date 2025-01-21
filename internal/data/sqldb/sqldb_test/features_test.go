package sqldb_test

import (
	"context"
	"testing"

	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var ftRepo repo.FeaturesRepo

func initFeaturesRepo() {
	dataMem := getDataMem()
	ftRepo, _ = sqldb.NewFeaturesRepoGorm(dataMem, logger)
}

func createBaseFeatures(t *testing.T, data []*repo.Feature) {
	initFeaturesRepo()
	if data == nil {
		data = []*repo.Feature{
			{Name: "cpu", Value: "amd"},
			{Name: "cpu", Value: "intel"},
			{Name: "net", Value: "private"},
			{Name: "net", Value: "public"},
		}
	}
	if err := ftRepo.CreateFeatures(context.Background(), nil, data); err != nil {
		t.Fatal(err)
	}

}

func TestFeaturesRepoGorm(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"CreateFeatures_Success", testCreateFeaturesSuccess},
		{"CreateFeatures_Error", testCreateFeaturesError},
		{"UpdateFeatures_Success", testUpdateFeaturesSuccess},
		{"UpdateFeatures_Error", testUpdateFeaturesError},
		{"DeleteFeatures_Success", testDeleteFeaturesSuccess},
		{"DeleteFeatures_Error", testDeleteFeaturesError},
		{"GetFeatures_Success", testGetFeaturesSuccess},
		{"GetFeatures_Error", testGetFeaturesError},
		{"ListFeatures_emptyFilter_all", testListFeatures_emptyFilter_all},
		{"ListFeatures_id_partial", testListFeatures_id_partial},
		{"ListFeatures_page_partial", testListFeatures_page_partial},
		{"ListFeatures_name_partial", testListFeatures_name_partial},
		{"ListFeatures_kvs_partial", testListFeatures_kvs_partial},
		{"ListFeatures_nil_all", testListFeatures_nil_all},
		{"CountFeatures_partial", testCountFeatures_partial},
		{"CountFeatures_subpartial", testCountFeatures_subpartial},
		{"CountFeatures_all", testCountFeatures_all},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func testCreateFeaturesSuccess(t *testing.T) {

	initFeaturesRepo()

	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	err := ftRepo.CreateFeatures(context.Background(), nil, data)
	assert.NoError(t, err)
}

func testCreateFeaturesError(t *testing.T) {
	createBaseFeatures(t, nil)
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	err := ftRepo.CreateFeatures(context.Background(), nil, data)
	assert.Error(t, err)
}

func testUpdateFeaturesSuccess(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)

	data[0].Value = "arm"
	err := ftRepo.UpdateFeatures(context.Background(), nil, data)
	assert.NoError(t, err)
}

func testUpdateFeaturesError(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)
	data[0].Value = "intel"
	err := ftRepo.UpdateFeatures(context.Background(), nil, data)
	assert.Error(t, err)
}

func testDeleteFeaturesSuccess(t *testing.T) {
	createBaseFeatures(t, nil)
	err := ftRepo.DeleteFeatures(context.Background(), nil, []uint32{1})
	assert.NoError(t, err)
}

func testDeleteFeaturesError(t *testing.T) {

	createBaseFeatures(t, nil)
	err := ftRepo.DeleteFeatures(context.Background(), nil, []uint32{99})
	assert.Error(t, err)
}

func testGetFeaturesSuccess(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)

	_data, err := ftRepo.GetFeatures(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, _data)
	assert.Equal(t, data[0], _data)
}

func testGetFeaturesError(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)
	tag, err := ftRepo.GetFeatures(context.Background(), 99)
	assert.Error(t, err)
	assert.Nil(t, tag)
}

func testListFeatures_emptyFilter_all(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)
	data, err := ftRepo.ListFeatures(context.Background(), nil, &repo.FeaturesFilter{})
	assert.NoError(t, err)
	assert.Len(t, data, 4)
}
func testListFeatures_id_partial(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)
	_data, err := ftRepo.ListFeatures(context.Background(), nil, &repo.FeaturesFilter{Ids: []uint32{1}})
	assert.NoError(t, err)
	assert.Len(t, _data, 1)
}

func testListFeatures_page_partial(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)
	filter := &repo.FeaturesFilter{
		Page:     2,
		PageSize: 1,
	}
	_data, err := ftRepo.ListFeatures(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data[1:2], _data)
}

func testListFeatures_name_partial(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)
	filter := &repo.FeaturesFilter{
		Names: []string{"pu"},
	}
	_data, err := ftRepo.ListFeatures(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data[:2], _data)
}
func testListFeatures_kvs_partial(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)
	filter := &repo.FeaturesFilter{
		Kvs: []string{"cpu:intel", "net:private"},
	}
	_data, err := ftRepo.ListFeatures(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data[1:3], _data)
}
func testListFeatures_nil_all(t *testing.T) {
	createBaseFeatures(t, nil)
	_data, err := ftRepo.ListFeatures(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(_data))
}

func testCountFeatures_partial(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)
	filter := &repo.FeaturesFilter{
		Ids: []uint32{2, 3},
	}
	count, err := ftRepo.CountFeatures(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}
func testCountFeatures_subpartial(t *testing.T) {
	data := []*repo.Feature{
		{Name: "cpu", Value: "amd"},
		{Name: "cpu", Value: "intel"},
		{Name: "net", Value: "private"},
		{Name: "net", Value: "public"},
	}
	createBaseFeatures(t, data)
	filter := &repo.FeaturesFilter{
		Ids: []uint32{2, 99},
	}
	count, err := ftRepo.CountFeatures(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func testCountFeatures_all(t *testing.T) {
	createBaseFeatures(t, nil)
	count, err := ftRepo.CountFeatures(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(4), count)
}
