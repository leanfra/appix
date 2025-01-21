package sqldb_test

import (
	"context"
	"testing"

	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var teamsRepo repo.TeamsRepo

func initTeamsRepo() {
	dataMem := getDataMem()
	teamsRepo, _ = sqldb.NewTeamsRepoGorm(dataMem, logger)
}

func createBaseTeamsData(data []*repo.Team) error {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}
	if data != nil {
		teams = data
	}
	err := teamsRepo.CreateTeams(context.Background(), nil, teams)
	return err
}

func TestCreateTeams_Success(t *testing.T) {
	initTeamsRepo()
	err := createBaseTeamsData(nil)
	assert.NoError(t, err)
}

func TestCreateTeams_Failure(t *testing.T) {

	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1"},
		{Name: "Team1"},
	}

	err := teamsRepo.CreateTeams(context.Background(), nil, teams)
	assert.Error(t, err)

	teams = []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team1", LeaderId: 2, Description: "description2"},
	}

	err = teamsRepo.CreateTeams(context.Background(), nil, teams)
	assert.Error(t, err)
}

func TestUpdateTeams_SuccessfulUpdate_ReturnsNil(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}

	err := teamsRepo.CreateTeams(context.Background(), nil, teams)
	assert.NoError(t, err)

	teams[0].Name = "Team1_updated"
	teams[1].Name = "Team2_updated"

	t.Logf("Updating teams: %v", *teams[0])
	t.Logf("Updating teams: %v", *teams[1])

	err = teamsRepo.UpdateTeams(context.Background(), nil, teams)
	assert.NoError(t, err)
}

func TestUpdateTeams_UpdateFails_ReturnsError(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}

	err := teamsRepo.CreateTeams(context.Background(), nil, teams)
	assert.NoError(t, err)

	teams[0].Name = "Team2"

	t.Logf("Updating teams: %v", *teams[0])
	t.Logf("Updating teams: %v", *teams[1])

	err = teamsRepo.UpdateTeams(context.Background(), nil, teams)
	assert.Error(t, err)
}

func TestDeleteTeams_Success_ReturnNil(t *testing.T) {
	initTeamsRepo()
	createBaseTeamsData(nil)

	err := teamsRepo.DeleteTeams(context.Background(), nil, []uint32{})
	assert.NoError(t, err)

	err = teamsRepo.DeleteTeams(context.Background(), nil, []uint32{1, 2})
	assert.NoError(t, err)
}

func TestDeleteTeams_Fail_ReturnError(t *testing.T) {
	initTeamsRepo()
	createBaseTeamsData(nil)

	err := teamsRepo.DeleteTeams(context.Background(), nil, []uint32{3})
	t.Log(err)
	assert.Error(t, err)
}

func TestGetTeams_Success(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}

	err := teamsRepo.CreateTeams(context.Background(), nil, teams)
	assert.NoError(t, err)

	team, err := teamsRepo.GetTeams(context.Background(), teams[0].ID)
	t.Logf("%v", *teams[0])
	assert.Equal(t, teams[0], team)
	assert.NoError(t, err)
}

func TestGetTeams_Fail(t *testing.T) {
	initTeamsRepo()
	createBaseTeamsData(nil)

	team, err := teamsRepo.GetTeams(context.Background(), 3)
	assert.Nil(t, team)
	t.Logf("%v", err)
	assert.Error(t, err)
}

func TestListTeams_NoFilter_ReturnsAllTeams(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}

	createBaseTeamsData(teams)

	teams_all, err := teamsRepo.ListTeams(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, teams, teams_all)
}

func TestListTeams_WithPagination_ReturnsPaginatedTeams(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}

	createBaseTeamsData(teams)

	teams_page_2, err := teamsRepo.ListTeams(context.Background(),
		nil, &repo.TeamsFilter{Page: 2, PageSize: 1})
	assert.NoError(t, err)
	assert.Equal(t, teams[1:], teams_page_2)
}

func TestListTeams_Ids_ReturnsFilteredTeams(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}
	createBaseTeamsData(teams)
	filter := &repo.TeamsFilter{Ids: []uint32{1, 2, 3}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams, teams_filtered)
}

func TestListTeams_Codes_ReturnsFilteredTeams(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}
	createBaseTeamsData(teams)
	filter := &repo.TeamsFilter{Codes: []string{"team1", "team2"}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams, teams_filtered)
	for _, _t := range teams_filtered {
		t.Log(_t)
	}
}

func TestListTeams_Leaders_ReturnsFilteredTeams(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}
	createBaseTeamsData(teams)
	filter := &repo.TeamsFilter{LeadersId: []uint32{1, 2}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams, teams_filtered)
	for _, _t := range teams_filtered {
		t.Log(_t)
	}
}

func TestListTeams_WithNamesFilter_ReturnsFilteredTeams(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}
	createBaseTeamsData(teams)
	filter := &repo.TeamsFilter{Names: []string{"am1"}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams[:1], teams_filtered)
	for _, _t := range teams_filtered {
		t.Log(_t)
	}
}

func TestListTeams_Partial_ReturnsFiltered(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}
	createBaseTeamsData(teams)
	filter := &repo.TeamsFilter{Names: []string{"Team3", "Team2"}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams[1:], teams_filtered)
	for _, _t := range teams_filtered {
		t.Log(_t)
	}
}

func TestCountTeams_NoFilter_AllTeamsCounted(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}
	createBaseTeamsData(teams)
	count, err := teamsRepo.CountTeams(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}

func TestCountTeams_Id_PartialTeamsCounted(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}
	createBaseTeamsData(teams)
	filter := &repo.TeamsFilter{Ids: []uint32{1, 3}}
	count, err := teamsRepo.CountTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func TestCountTeams_Id_Empty(t *testing.T) {
	initTeamsRepo()
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", LeaderId: 1, Description: "description1"},
		{Name: "Team2", Code: "team2", LeaderId: 2, Description: "description2"},
	}
	createBaseTeamsData(teams)
	filter := &repo.TeamsFilter{Ids: []uint32{4, 3}}
	count, err := teamsRepo.CountTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}
