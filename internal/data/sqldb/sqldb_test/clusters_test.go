package sqldb_test

import (
	"context"
	"testing"

	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var clusterRepo repo.ClustersRepo

func initClustersRepo() {
	dataMem := getDataMem()
	clusterRepo, _ = sqldb.NewClustersRepoGorm(dataMem, logger)
}

func createBaseClusters(t *testing.T, data []*repo.Cluster) {
	initClustersRepo()
	if data == nil {
		data = []*repo.Cluster{
			{Name: "k8s-0", Description: "k8s-0 cluster"},
			{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
		}
	}
	if err := clusterRepo.CreateClusters(context.Background(), nil, data); err != nil {
		t.Fatal(err)
	}

}

func TestClustersRepoGorm(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"CreateClusters_Success", testCreateClustersSuccess},
		{"CreateClusters_Error", testCreateClustersError},
		{"UpdateClusters_Success", testUpdateClustersSuccess},
		{"UpdateClusters_Error", testUpdateClustersError},
		{"DeleteClusters_Success", testDeleteClustersSuccess},
		{"DeleteClusters_Error", testDeleteClustersError},
		{"GetClusters_Success", testGetClustersSuccess},
		{"GetClusters_Error", testGetClustersError},
		{"ListClusters_emptyFilter_all", testListClusters_emptyFilter_all},
		{"ListClusters_id_partial", testListClusters_id_partial},
		{"ListClusters_page_partial", testListClusters_page_partial},
		{"ListClusters_name_partial", testListClusters_name_partial},
		{"ListClusters_nil_all", testListClusters_nil_all},
		{"CountClusters_partial", testCountClusters_partial},
		{"CountClusters_subpartial", testCountClusters_subpartial},
		{"CountClusters_all", testCountClusters_all},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func testCreateClustersSuccess(t *testing.T) {

	initClustersRepo()

	envs := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	err := clusterRepo.CreateClusters(context.Background(), nil, envs)
	assert.NoError(t, err)
}

func testCreateClustersError(t *testing.T) {
	createBaseClusters(t, nil)
	envs := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	err := clusterRepo.CreateClusters(context.Background(), nil, envs)
	assert.Error(t, err)
}

func testUpdateClustersSuccess(t *testing.T) {
	envs := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, envs)

	envs[0].Name = "prod"
	err := clusterRepo.UpdateClusters(context.Background(), nil, envs)
	assert.NoError(t, err)
}

func testUpdateClustersError(t *testing.T) {
	envs := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, envs)
	envs[1].Name = "k8s-0"
	err := clusterRepo.UpdateClusters(context.Background(), nil, envs)
	assert.Error(t, err)
}

func testDeleteClustersSuccess(t *testing.T) {
	createBaseClusters(t, nil)
	err := clusterRepo.DeleteClusters(context.Background(), nil, []uint32{1})
	assert.NoError(t, err)
}

func testDeleteClustersError(t *testing.T) {

	createBaseClusters(t, nil)
	err := clusterRepo.DeleteClusters(context.Background(), nil, []uint32{99})
	assert.Error(t, err)
}

func testGetClustersSuccess(t *testing.T) {
	Clusters := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, Clusters)

	env, err := clusterRepo.GetClusters(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, env)
	assert.Equal(t, Clusters[0], env)
}

func testGetClustersError(t *testing.T) {
	Clusters := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, Clusters)
	env, err := clusterRepo.GetClusters(context.Background(), 99)
	assert.Error(t, err)
	assert.Nil(t, env)
}

func testListClusters_emptyFilter_all(t *testing.T) {
	Clusters := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, Clusters)
	Clusters, err := clusterRepo.ListClusters(context.Background(), nil, &repo.ClustersFilter{})
	assert.NoError(t, err)
	assert.Len(t, Clusters, 2)
}
func testListClusters_id_partial(t *testing.T) {
	Clusters := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, Clusters)
	_Clusters, err := clusterRepo.ListClusters(context.Background(), nil, &repo.ClustersFilter{Ids: []uint32{1}})
	assert.NoError(t, err)
	assert.Len(t, _Clusters, 1)
}

func testListClusters_page_partial(t *testing.T) {
	Clusters := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, Clusters)
	filter := &repo.ClustersFilter{
		Page:     2,
		PageSize: 1,
	}
	_Clusters, err := clusterRepo.ListClusters(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, Clusters[1:], _Clusters)
}

func testListClusters_name_partial(t *testing.T) {
	Clusters := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, Clusters)
	filter := &repo.ClustersFilter{
		Names: []string{"nk8s"},
	}
	_Clusters, err := clusterRepo.ListClusters(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, Clusters[1:], _Clusters)
}

func testListClusters_nil_all(t *testing.T) {
	createBaseClusters(t, nil)
	_Clusters, err := clusterRepo.ListClusters(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(_Clusters))
}

func testCountClusters_partial(t *testing.T) {
	Clusters := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, Clusters)
	filter := &repo.ClustersFilter{
		Ids: []uint32{2, 3},
	}
	count, err := clusterRepo.CountClusters(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}
func testCountClusters_subpartial(t *testing.T) {
	Clusters := []*repo.Cluster{
		{Name: "k8s-0", Description: "k8s-0 cluster"},
		{Name: "nk8s-1", Description: "non-k8s-1 cluster"},
	}
	createBaseClusters(t, Clusters)
	filter := &repo.ClustersFilter{
		Ids: []uint32{2, 99},
	}
	count, err := clusterRepo.CountClusters(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func testCountClusters_all(t *testing.T) {
	createBaseClusters(t, nil)
	count, err := clusterRepo.CountClusters(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}
