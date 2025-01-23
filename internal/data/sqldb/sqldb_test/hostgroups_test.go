package sqldb_test

import (
	"context"
	"testing"

	"opspillar/internal/data/repo"
	"opspillar/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var hostgroupRepo repo.HostgroupsRepo

var fakeHostgroups = []*repo.Hostgroup{
	{
		Name:         "hostgroup1",
		Description:  "hostgroup1 description",
		ClusterId:    301,
		DatacenterId: 201,
		EnvId:        501,
		ProductId:    101,
		TeamId:       401,
		ChangeInfo: repo.ChangeInfo{
			CreatedBy: "test",
			UpdatedBy: "test",
		},
	},
	{
		Name:         "hostgroup2",
		Description:  "hostgroup2 description",
		ClusterId:    301,
		DatacenterId: 201,
		EnvId:        501,
		ProductId:    101,
		TeamId:       401,
		ChangeInfo: repo.ChangeInfo{
			CreatedBy: "test",
			UpdatedBy: "test",
		},
	},
	{
		Name:         "hostgroup3",
		Description:  "hostgroup3 description",
		ClusterId:    103,
		DatacenterId: 203,
		EnvId:        303,
		ProductId:    403,
		TeamId:       503,
		ChangeInfo: repo.ChangeInfo{
			CreatedBy: "test",
			UpdatedBy: "test",
		},
	},
}

func getFakeHostgroups() []*repo.Hostgroup {
	data := make([]*repo.Hostgroup, len(fakeHostgroups))
	for i, app := range fakeHostgroups {
		data[i] = &repo.Hostgroup{
			Name:         app.Name,
			Description:  app.Description,
			ProductId:    app.ProductId,
			DatacenterId: app.DatacenterId,
			ClusterId:    app.ClusterId,
			TeamId:       app.TeamId,
			ChangeInfo: repo.ChangeInfo{
				UpdatedBy: "test",
			},
		}
	}
	return data
}

func initHostgroupsRepo() {
	dataMem := getDataMem()
	hostgroupRepo, _ = sqldb.NewHostgroupsRepoGorm(dataMem, logger)
}

func createBaseHostgroups(t *testing.T, data []*repo.Hostgroup) {
	initHostgroupsRepo()
	if data == nil {
		data = fakeHostgroups
	}
	if err := hostgroupRepo.CreateHostgroups(context.Background(), nil, data); err != nil {
		t.Fatal(err)
	}

}

func TestHostgroupsRepoGorm(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"CreateHostgroups_Success", testCreateHostgroupsSuccess},
		{"CreateHostgroups_Error", testCreateHostgroupsError},
		{"UpdateHostgroups_Success", testUpdateHostgroupsSuccess},
		{"UpdateHostgroups_Error", testUpdateHostgroupsError},
		{"DeleteHostgroups_Success", testDeleteHostgroupsSuccess},
		{"DeleteHostgroups_Error", testDeleteHostgroupsError},
		{"GetHostgroups_Success", testGetHostgroupsSuccess},
		{"GetHostgroups_Error", testGetHostgroupsError},
		{"ListHostgroups_emptyFilter_all", testListHostgroups_emptyFilter_all},
		{"ListHostgroups_id_partial", testListHostgroups_id_partial},
		{"ListHostgroups_page_partial", testListHostgroups_page_partial},
		{"ListHostgroups_name_partial", testListHostgroups_name_partial},
		{"ListHostgroups_name_none", testListHostgroups_name_none},
		{"ListHostgroups_productId_partial", testListHostgroups_productid_partial},
		{"ListHostgroups_datacenterId_partial", testListHostgroups_dcId_partial},
		{"ListHostgroups_clusterId_partial", testListHostgroups_clsId_partial},
		{"ListHostgroups_teamId_partial", testListHostgroups_teamId_partial},
		{"ListHostgroups_envId_partial", testListHostgroups_envId_partial},
		{"ListHostgroups_nil_all", testListHostgroups_nil_all},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func testCreateHostgroupsSuccess(t *testing.T) {

	initHostgroupsRepo()

	err := hostgroupRepo.CreateHostgroups(context.Background(), nil, fakeHostgroups)
	assert.NoError(t, err)
}

func testCreateHostgroupsError(t *testing.T) {
	createBaseHostgroups(t, nil)
	err := hostgroupRepo.CreateHostgroups(context.Background(), nil, fakeHostgroups)
	assert.Error(t, err)
}

func testUpdateHostgroupsSuccess(t *testing.T) {
	data := getFakeHostgroups()
	createBaseHostgroups(t, data)

	data[0].Name = "App1modify"
	err := hostgroupRepo.UpdateHostgroups(context.Background(), nil, data)
	assert.NoError(t, err)
}

func testUpdateHostgroupsError(t *testing.T) {
	data := getFakeHostgroups()
	createBaseHostgroups(t, data)
	data[1].Name = "hostgroup1"
	err := hostgroupRepo.UpdateHostgroups(context.Background(), nil, data)
	assert.Error(t, err)
}

func testDeleteHostgroupsSuccess(t *testing.T) {
	createBaseHostgroups(t, nil)
	err := hostgroupRepo.DeleteHostgroups(context.Background(), nil, []uint32{1})
	assert.NoError(t, err)
}

func testDeleteHostgroupsError(t *testing.T) {

	createBaseHostgroups(t, nil)
	err := hostgroupRepo.DeleteHostgroups(context.Background(), nil, []uint32{99})
	assert.Error(t, err)
}

func testGetHostgroupsSuccess(t *testing.T) {
	data := getFakeHostgroups()
	createBaseHostgroups(t, data)

	_data, err := hostgroupRepo.GetHostgroups(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, _data)
	assert.Equal(t, data[0], _data)
}

func testGetHostgroupsError(t *testing.T) {
	data := getFakeHostgroups()
	createBaseHostgroups(t, data)
	_data, err := hostgroupRepo.GetHostgroups(context.Background(), 99)
	assert.Error(t, err)
	assert.Nil(t, _data)
}

func testListHostgroups_emptyFilter_all(t *testing.T) {
	data := getFakeHostgroups()
	createBaseHostgroups(t, data)
	_data, err := hostgroupRepo.ListHostgroups(
		context.Background(), nil, &repo.HostgroupsFilter{})
	assert.NoError(t, err)
	assert.Len(t, _data, 3)
}
func testListHostgroups_id_partial(t *testing.T) {
	data := getFakeHostgroups()
	createBaseHostgroups(t, data)
	_data, err := hostgroupRepo.ListHostgroups(
		context.Background(), nil, &repo.HostgroupsFilter{Ids: []uint32{1}})
	assert.NoError(t, err)
	assert.Equal(t, data[0:1], _data)
}

func testListHostgroups_page_partial(t *testing.T) {
	data := getFakeHostgroups()
	createBaseHostgroups(t, data)
	filter := &repo.HostgroupsFilter{
		Page:     2,
		PageSize: 1,
	}
	_data, err := hostgroupRepo.ListHostgroups(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data[1:2], _data)
}

func testListHostgroups_name_partial(t *testing.T) {
	data := getFakeHostgroups()
	createBaseHostgroups(t, data)
	filter := &repo.HostgroupsFilter{
		Names: []string{"group1"},
	}
	_data, err := hostgroupRepo.ListHostgroups(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data[:1], _data)
}
func testListHostgroups_name_none(t *testing.T) {
	data := getFakeHostgroups()
	createBaseHostgroups(t, data)
	filter := &repo.HostgroupsFilter{
		Names: []string{"nogroup1"},
	}
	_data, err := hostgroupRepo.ListHostgroups(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, []*repo.Hostgroup{}, _data)
}

func testListHostgroups_productid_partial(t *testing.T) {
	createBaseHostgroups(t, nil)
	_data, err := hostgroupRepo.ListHostgroups(
		context.Background(), nil, &repo.HostgroupsFilter{ProductsId: []uint32{101}})
	assert.Nil(t, err)
	assert.Equal(t, fakeHostgroups[:2], _data)
}

func testListHostgroups_dcId_partial(t *testing.T) {
	createBaseHostgroups(t, nil)
	_data, err := hostgroupRepo.ListHostgroups(
		context.Background(), nil, &repo.HostgroupsFilter{DatacentersId: []uint32{201}})
	assert.Nil(t, err)
	assert.Equal(t, fakeHostgroups[:2], _data)
}

func testListHostgroups_clsId_partial(t *testing.T) {
	createBaseHostgroups(t, nil)
	_data, err := hostgroupRepo.ListHostgroups(
		context.Background(), nil, &repo.HostgroupsFilter{ClustersId: []uint32{301}})
	assert.Nil(t, err)
	assert.Equal(t, fakeHostgroups[:2], _data)
}

func testListHostgroups_teamId_partial(t *testing.T) {
	createBaseHostgroups(t, nil)
	_data, err := hostgroupRepo.ListHostgroups(
		context.Background(), nil, &repo.HostgroupsFilter{TeamsId: []uint32{401}})
	assert.Nil(t, err)
	assert.Equal(t, fakeHostgroups[:2], _data)
}
func testListHostgroups_envId_partial(t *testing.T) {
	createBaseHostgroups(t, nil)
	_data, err := hostgroupRepo.ListHostgroups(
		context.Background(), nil, &repo.HostgroupsFilter{EnvsId: []uint32{501}})
	assert.Nil(t, err)
	assert.Equal(t, fakeHostgroups[:2], _data)
}

func testListHostgroups_nil_all(t *testing.T) {
	createBaseHostgroups(t, nil)
	_data, err := hostgroupRepo.ListHostgroups(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, fakeHostgroups, _data)
}
