package sqldb_test

import (
	"context"
	"testing"

	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var appRepo repo.ApplicationsRepo

var fakeApps = []*repo.Application{
	{
		Name:         "App1",
		Description:  "First Application",
		ProductId:    101,
		DatacenterId: 201,
		ClusterId:    301,
		TeamId:       401,
		IsStateful:   true,
	},
	{
		Name:         "App2",
		Description:  "Second Application",
		ProductId:    101,
		DatacenterId: 201,
		ClusterId:    301,
		TeamId:       401,
		IsStateful:   false,
	},
	{
		Name:         "App3",
		Description:  "Third Application",
		ProductId:    103,
		DatacenterId: 203,
		ClusterId:    303,
		TeamId:       403,
		IsStateful:   true,
	},
}

func getFakeApps() []*repo.Application {
	data := make([]*repo.Application, len(fakeApps))
	for i, app := range fakeApps {
		data[i] = &repo.Application{
			Name:         app.Name,
			Description:  app.Description,
			ProductId:    app.ProductId,
			DatacenterId: app.DatacenterId,
			ClusterId:    app.ClusterId,
			TeamId:       app.TeamId,
			IsStateful:   app.IsStateful,
		}
	}
	return data
}

func initAppsRepo() {
	dataMem := getDataMem()
	appRepo, _ = sqldb.NewApplicationsRepoGorm(dataMem, logger)
}

func createBaseApps(t *testing.T, data []*repo.Application) {
	initAppsRepo()
	if data == nil {
		data = fakeApps
	}
	if err := appRepo.CreateApplications(context.Background(), nil, data); err != nil {
		t.Fatal(err)
	}

}

func TestApplicationsRepoGorm(t *testing.T) {

	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"CreateApplications_Success", testCreateApplicationsSuccess},
		{"CreateApplications_Error", testCreateApplicationsError},
		{"UpdateApplications_Success", testUpdateApplicationsSuccess},
		{"UpdateApplications_Error", testUpdateApplicationsError},
		{"DeleteApplications_Success", testDeleteApplicationsSuccess},
		{"DeleteApplications_Error", testDeleteApplicationsError},
		{"GetApplications_Success", testGetApplicationsSuccess},
		{"GetApplications_Error", testGetApplicationsError},
		{"ListApplications_emptyFilter_all", testListApplications_emptyFilter_all},
		{"ListApplications_id_partial", testListApplications_id_partial},
		{"ListApplications_page_partial", testListApplications_page_partial},
		{"ListApplications_name_partial", testListApplications_name_partial},
		{"ListApplications_productId_partial", testListApplications_productid_partial},
		{"ListApplications_datacenterId_partial", testListApplications_dcId_partial},
		{"ListApplications_clusterId_partial", testListApplications_clsId_partial},
		{"ListApplications_teamId_partial", testListApplications_teamId_partial},
		{"ListApplications_statefulTrue_partial", testListApplications_stTrue_partial},
		{"ListApplications_statefulFalse_partial", testListApplications_stFalse_partial},
		{"ListApplications_statefulNone_all", testListApplications_stNone_all},
		{"ListApplications_nil_all", testListApplications_nil_all},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func testCreateApplicationsSuccess(t *testing.T) {

	initAppsRepo()

	err := appRepo.CreateApplications(context.Background(), nil, fakeApps)
	assert.NoError(t, err)
}

func testCreateApplicationsError(t *testing.T) {
	createBaseApps(t, nil)
	err := appRepo.CreateApplications(context.Background(), nil, fakeApps)
	assert.Error(t, err)
}

func testUpdateApplicationsSuccess(t *testing.T) {
	data := getFakeApps()
	createBaseApps(t, data)

	data[0].Name = "App1modify"
	err := appRepo.UpdateApplications(context.Background(), nil, data)
	assert.NoError(t, err)
}

func testUpdateApplicationsError(t *testing.T) {
	data := getFakeApps()
	createBaseApps(t, data)
	data[1].Name = "App1"
	err := appRepo.UpdateApplications(context.Background(), nil, data)
	assert.Error(t, err)
}

func testDeleteApplicationsSuccess(t *testing.T) {
	createBaseApps(t, nil)
	err := appRepo.DeleteApplications(context.Background(), nil, []uint32{1})
	assert.NoError(t, err)
}

func testDeleteApplicationsError(t *testing.T) {

	createBaseApps(t, nil)
	err := appRepo.DeleteApplications(context.Background(), nil, []uint32{99})
	assert.Error(t, err)
}

func testGetApplicationsSuccess(t *testing.T) {
	data := getFakeApps()
	createBaseApps(t, data)

	_data, err := appRepo.GetApplications(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, _data)
	assert.Equal(t, data[0], _data)
}

func testGetApplicationsError(t *testing.T) {
	data := getFakeApps()
	createBaseApps(t, data)
	_data, err := appRepo.GetApplications(context.Background(), 99)
	assert.Error(t, err)
	assert.Nil(t, _data)
}

func testListApplications_emptyFilter_all(t *testing.T) {
	data := getFakeApps()
	createBaseApps(t, data)
	_data, err := appRepo.ListApplications(
		context.Background(), nil, &repo.ApplicationsFilter{})
	assert.NoError(t, err)
	assert.Len(t, _data, 3)
}
func testListApplications_id_partial(t *testing.T) {
	data := getFakeApps()
	createBaseApps(t, data)
	_data, err := appRepo.ListApplications(
		context.Background(), nil, &repo.ApplicationsFilter{Ids: []uint32{1}})
	assert.NoError(t, err)
	assert.Equal(t, data[0:1], _data)
}

func testListApplications_page_partial(t *testing.T) {
	data := getFakeApps()
	createBaseApps(t, data)
	filter := &repo.ApplicationsFilter{
		Page:     2,
		PageSize: 1,
	}
	_data, err := appRepo.ListApplications(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data[1:2], _data)
}

func testListApplications_name_partial(t *testing.T) {
	data := getFakeApps()
	createBaseApps(t, data)
	filter := &repo.ApplicationsFilter{
		Names: []string{"pp1"},
	}
	_data, err := appRepo.ListApplications(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, data[:1], _data)
}

func testListApplications_productid_partial(t *testing.T) {
	createBaseApps(t, nil)
	_data, err := appRepo.ListApplications(
		context.Background(), nil, &repo.ApplicationsFilter{ProductsId: []uint32{101}})
	assert.Nil(t, err)
	assert.Equal(t, fakeApps[:2], _data)
}

func testListApplications_dcId_partial(t *testing.T) {
	createBaseApps(t, nil)
	_data, err := appRepo.ListApplications(
		context.Background(), nil, &repo.ApplicationsFilter{DatacentersId: []uint32{201}})
	assert.Nil(t, err)
	assert.Equal(t, fakeApps[:2], _data)
}

func testListApplications_clsId_partial(t *testing.T) {
	createBaseApps(t, nil)
	_data, err := appRepo.ListApplications(
		context.Background(), nil, &repo.ApplicationsFilter{ClustersId: []uint32{301}})
	assert.Nil(t, err)
	assert.Equal(t, fakeApps[:2], _data)
}

func testListApplications_teamId_partial(t *testing.T) {
	createBaseApps(t, nil)
	_data, err := appRepo.ListApplications(
		context.Background(), nil, &repo.ApplicationsFilter{TeamsId: []uint32{401}})
	assert.Nil(t, err)
	assert.Equal(t, fakeApps[:2], _data)
}

func testListApplications_stTrue_partial(t *testing.T) {
	createBaseApps(t, nil)
	_data, err := appRepo.ListApplications(
		context.Background(), nil, &repo.ApplicationsFilter{IsStateful: repo.IsStatefulTrue})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(_data))
}

func testListApplications_stFalse_partial(t *testing.T) {
	createBaseApps(t, nil)
	_data, err := appRepo.ListApplications(
		context.Background(), nil, &repo.ApplicationsFilter{IsStateful: repo.IsStatefulFalse})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(_data))
}

func testListApplications_stNone_all(t *testing.T) {
	createBaseApps(t, nil)
	_data, err := appRepo.ListApplications(
		context.Background(), nil, &repo.ApplicationsFilter{IsStateful: repo.IsStatefulNone})
	assert.Nil(t, err)
	assert.Equal(t, 3, len(_data))
}

func testListApplications_nil_all(t *testing.T) {
	createBaseApps(t, nil)
	_data, err := appRepo.ListApplications(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, fakeApps, _data)
}
