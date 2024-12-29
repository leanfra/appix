package sqldb_test

import (
	"context"
	"testing"

	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var dcsRepo repo.DatacentersRepo

func initDatacentersRepo() {
	dataMem := getDataMem()
	dcsRepo, _ = sqldb.NewDatacentersRepoGorm(dataMem, logger)
}

func createBaseDatacenters(t *testing.T, data []*repo.Datacenter) {
	initDatacentersRepo()
	if data == nil {
		data = []*repo.Datacenter{
			{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
			{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
		}
	}
	if err := dcsRepo.CreateDatacenters(context.Background(), data); err != nil {
		t.Fatal(err)
	}

}

func TestDatacentersRepoGorm(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"CreateDatacenters_Success", testCreateDatacentersSuccess},
		{"CreateDatacenters_Error", testCreateDatacentersError},
		{"UpdateDatacenters_Success", testUpdateDatacentersSuccess},
		{"UpdateDatacenters_Error", testUpdateDatacentersError},
		{"DeleteDatacenters_Success", testDeleteDatacentersSuccess},
		{"DeleteDatacenters_Error", testDeleteDatacentersError},
		{"GetDatacenters_Success", testGetDatacentersSuccess},
		{"GetDatacenters_Error", testGetDatacentersError},
		{"ListDatacenters_emptyFilter_all", testListDatacenters_emptyFilter_all},
		{"ListDatacenters_id_partial", testListDatacenters_id_partial},
		{"ListDatacenters_page_partial", testListDatacenters_page_partial},
		{"ListDatacenters_name_partial", testListDatacenters_name_partial},
		{"ListDatacenters_nil_all", testListDatacenters_nil_all},
		{"CountDatacenters_partial", testCountDatacenters_partial},
		{"CountDatacenters_subpartial", testCountDatacenters_subpartial},
		{"CountDatacenters_all", testCountDatacenters_all},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func testCreateDatacentersSuccess(t *testing.T) {

	initDatacentersRepo()

	envs := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	err := dcsRepo.CreateDatacenters(context.Background(), envs)
	assert.NoError(t, err)
}

func testCreateDatacentersError(t *testing.T) {
	createBaseDatacenters(t, nil)
	envs := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	err := dcsRepo.CreateDatacenters(context.Background(), envs)
	assert.Error(t, err)
}

func testUpdateDatacentersSuccess(t *testing.T) {
	envs := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, envs)

	envs[0].Name = "prod"
	err := dcsRepo.UpdateDatacenters(context.Background(), envs)
	assert.NoError(t, err)
}

func testUpdateDatacentersError(t *testing.T) {
	envs := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, envs)
	envs[1].Name = "ali-cn-bj"
	err := dcsRepo.UpdateDatacenters(context.Background(), envs)
	assert.Error(t, err)
}

func testDeleteDatacentersSuccess(t *testing.T) {
	createBaseDatacenters(t, nil)
	err := dcsRepo.DeleteDatacenters(context.Background(), nil, []uint32{1})
	assert.NoError(t, err)
}

func testDeleteDatacentersError(t *testing.T) {

	createBaseDatacenters(t, nil)
	err := dcsRepo.DeleteDatacenters(context.Background(), nil, []uint32{99})
	assert.Error(t, err)
}

func testGetDatacentersSuccess(t *testing.T) {
	Datacenters := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, Datacenters)

	env, err := dcsRepo.GetDatacenters(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, env)
	assert.Equal(t, Datacenters[0], env)
}

func testGetDatacentersError(t *testing.T) {
	Datacenters := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, Datacenters)
	env, err := dcsRepo.GetDatacenters(context.Background(), 99)
	assert.Error(t, err)
	assert.Nil(t, env)
}

func testListDatacenters_emptyFilter_all(t *testing.T) {
	Datacenters := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, Datacenters)
	Datacenters, err := dcsRepo.ListDatacenters(context.Background(), nil, &repo.DatacentersFilter{})
	assert.NoError(t, err)
	assert.Len(t, Datacenters, 2)
}
func testListDatacenters_id_partial(t *testing.T) {
	Datacenters := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, Datacenters)
	_Datacenters, err := dcsRepo.ListDatacenters(context.Background(), nil, &repo.DatacentersFilter{Ids: []uint32{1}})
	assert.NoError(t, err)
	assert.Len(t, _Datacenters, 1)
}

func testListDatacenters_page_partial(t *testing.T) {
	Datacenters := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, Datacenters)
	filter := &repo.DatacentersFilter{
		Page:     2,
		PageSize: 1,
	}
	_Datacenters, err := dcsRepo.ListDatacenters(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, Datacenters[1:], _Datacenters)
}

func testListDatacenters_name_partial(t *testing.T) {
	Datacenters := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, Datacenters)
	filter := &repo.DatacentersFilter{
		Names: []string{"bj"},
	}
	_Datacenters, err := dcsRepo.ListDatacenters(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, Datacenters[:1], _Datacenters)
}

func testListDatacenters_nil_all(t *testing.T) {
	createBaseDatacenters(t, nil)
	_Datacenters, err := dcsRepo.ListDatacenters(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(_Datacenters))
}

func testCountDatacenters_partial(t *testing.T) {
	Datacenters := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, Datacenters)
	filter := &repo.DatacentersFilter{
		Ids: []uint32{2, 3},
	}
	count, err := dcsRepo.CountDatacenters(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}
func testCountDatacenters_subpartial(t *testing.T) {
	Datacenters := []*repo.Datacenter{
		{Name: "ali-cn-bj", Description: "aliyun cn beijing"},
		{Name: "tc-cn-nj", Description: "tencent cn nanjing"},
	}
	createBaseDatacenters(t, Datacenters)
	filter := &repo.DatacentersFilter{
		Ids: []uint32{2, 99},
	}
	count, err := dcsRepo.CountDatacenters(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func testCountDatacenters_all(t *testing.T) {
	createBaseDatacenters(t, nil)
	count, err := dcsRepo.CountDatacenters(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}
