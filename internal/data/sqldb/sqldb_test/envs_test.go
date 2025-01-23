package sqldb_test

import (
	"context"
	"testing"

	"opspillar/internal/data/repo"
	"opspillar/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var envsRepo repo.EnvsRepo

func initEnvsRepo() {
	dataMem := getDataMem()
	envsRepo, _ = sqldb.NewEnvsRepoGorm(dataMem, logger)
}

func createBaseEnvs(t *testing.T, data []*repo.Env) {
	initEnvsRepo()
	if data == nil {
		data = []*repo.Env{
			{Name: "prd", Description: "production"},
			{Name: "stg", Description: "senving"},
		}
	}
	if err := envsRepo.CreateEnvs(context.Background(), nil, data); err != nil {
		t.Fatal(err)
	}

}

func TestEnvsRepoGorm(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"CreateEnvs_Success", testCreateEnvsSuccess},
		{"CreateEnvs_Error", testCreateEnvsError},
		{"UpdateEnvs_Success", testUpdateEnvsSuccess},
		{"UpdateEnvs_Error", testUpdateEnvsError},
		{"DeleteEnvs_Success", testDeleteEnvsSuccess},
		{"DeleteEnvs_Error", testDeleteEnvsError},
		{"GetEnvs_Success", testGetEnvsSuccess},
		{"GetEnvs_Error", testGetEnvsError},
		{"ListEnvs_emptyFilter_all", testListEnvs_emptyFilter_all},
		{"ListEnvs_id_partial", testListEnvs_id_partial},
		{"ListEnvs_page_partial", testListEnvs_page_partial},
		{"ListEnvs_name_partial", testListEnvs_name_partial},
		{"ListEnvs_nil_all", testListEnvs_nil_all},
		{"CountEnvs_partial", testCountEnvs_partial},
		{"CountEnvs_subpartial", testCountEnvs_subpartial},
		{"CountEnvs_all", testCountEnvs_all},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func testCreateEnvsSuccess(t *testing.T) {

	initEnvsRepo()

	envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	err := envsRepo.CreateEnvs(context.Background(), nil, envs)
	assert.NoError(t, err)
}

func testCreateEnvsError(t *testing.T) {
	createBaseEnvs(t, nil)
	envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	err := envsRepo.CreateEnvs(context.Background(), nil, envs)
	assert.Error(t, err)
}

func testUpdateEnvsSuccess(t *testing.T) {
	envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, envs)

	envs[0].Name = "prod"
	err := envsRepo.UpdateEnvs(context.Background(), nil, envs)
	assert.NoError(t, err)
}

func testUpdateEnvsError(t *testing.T) {
	envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, envs)
	envs[0].Name = "stg"
	err := envsRepo.UpdateEnvs(context.Background(), nil, envs)
	assert.Error(t, err)
}

func testDeleteEnvsSuccess(t *testing.T) {
	createBaseEnvs(t, nil)
	err := envsRepo.DeleteEnvs(context.Background(), nil, []uint32{1})
	assert.NoError(t, err)
}

func testDeleteEnvsError(t *testing.T) {

	createBaseEnvs(t, nil)
	err := envsRepo.DeleteEnvs(context.Background(), nil, []uint32{99})
	assert.Error(t, err)
}

func testGetEnvsSuccess(t *testing.T) {
	Envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, Envs)

	env, err := envsRepo.GetEnvs(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, env)
	assert.Equal(t, Envs[0], env)
}

func testGetEnvsError(t *testing.T) {
	Envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, Envs)
	env, err := envsRepo.GetEnvs(context.Background(), 99)
	assert.Error(t, err)
	assert.Nil(t, env)
}

func testListEnvs_emptyFilter_all(t *testing.T) {
	Envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, Envs)
	Envs, err := envsRepo.ListEnvs(context.Background(), nil, &repo.EnvsFilter{})
	assert.NoError(t, err)
	assert.Len(t, Envs, 2)
}
func testListEnvs_id_partial(t *testing.T) {
	Envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, Envs)
	_Envs, err := envsRepo.ListEnvs(context.Background(), nil, &repo.EnvsFilter{Ids: []uint32{1}})
	assert.NoError(t, err)
	assert.Len(t, _Envs, 1)
}

func testListEnvs_page_partial(t *testing.T) {
	Envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, Envs)
	filter := &repo.EnvsFilter{
		Page:     2,
		PageSize: 1,
	}
	_Envs, err := envsRepo.ListEnvs(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, Envs[1:], _Envs)
}

func testListEnvs_name_partial(t *testing.T) {
	Envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, Envs)
	filter := &repo.EnvsFilter{
		Names: []string{"tg"},
	}
	_Envs, err := envsRepo.ListEnvs(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, Envs[1:], _Envs)
}

func testListEnvs_nil_all(t *testing.T) {
	createBaseEnvs(t, nil)
	_Envs, err := envsRepo.ListEnvs(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(_Envs))
}

func testCountEnvs_partial(t *testing.T) {
	Envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, Envs)
	filter := &repo.EnvsFilter{
		Ids: []uint32{2, 3},
	}
	count, err := envsRepo.CountEnvs(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}
func testCountEnvs_subpartial(t *testing.T) {
	Envs := []*repo.Env{
		{Name: "prd", Description: "production"},
		{Name: "stg", Description: "senving"},
	}
	createBaseEnvs(t, Envs)
	filter := &repo.EnvsFilter{
		Ids: []uint32{2, 99},
	}
	count, err := envsRepo.CountEnvs(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func testCountEnvs_all(t *testing.T) {
	createBaseEnvs(t, nil)
	count, err := envsRepo.CountEnvs(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}
