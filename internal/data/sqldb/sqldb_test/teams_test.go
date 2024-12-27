package sqldb_test

import (
	"context"
	"testing"

	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var teamsRepo repo.TeamsRepo

func init() {
	teamsRepo, _ = sqldb.NewTeamsRepoGorm(dataMem, logger)
}

func createBaseData(data []*repo.Team) error {

	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}
	if data != nil {
		teams = data
	}
	err := teamsRepo.CreateTeams(context.Background(), teams)
	return err
}

func TestCreateTeams_Success(t *testing.T) {
	err := createBaseData(nil)
	assert.NoError(t, err)
}

func TestCreateTeams_Failure(t *testing.T) {

	teams := []*repo.Team{
		{Name: "Team1"},
		{Name: "Team1"},
	}

	err := teamsRepo.CreateTeams(context.Background(), teams)
	assert.Error(t, err)

	teams = []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team1", Leader: "leader2", Description: "description2"},
	}

	err = teamsRepo.CreateTeams(context.Background(), teams)
	assert.Error(t, err)
}

func TestUpdateTeams_SuccessfulUpdate_ReturnsNil(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}

	err := teamsRepo.CreateTeams(context.Background(), teams)
	assert.NoError(t, err)

	teams[0].Name = "Team1_updated"
	teams[1].Name = "Team2_updated"

	t.Logf("Updating teams: %v", *teams[0])
	t.Logf("Updating teams: %v", *teams[1])

	err = teamsRepo.UpdateTeams(context.Background(), teams)
	assert.NoError(t, err)
}

func TestUpdateTeams_UpdateFails_ReturnsError(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}

	err := teamsRepo.CreateTeams(context.Background(), teams)
	assert.NoError(t, err)

	teams[0].Name = "Team2"

	t.Logf("Updating teams: %v", *teams[0])
	t.Logf("Updating teams: %v", *teams[1])

	err = teamsRepo.UpdateTeams(context.Background(), teams)
	assert.Error(t, err)
}

func TestDeleteTeams_Success_ReturnNil(t *testing.T) {
	createBaseData(nil)

	err := teamsRepo.DeleteTeams(context.Background(), []uint32{})
	assert.NoError(t, err)

	err = teamsRepo.DeleteTeams(context.Background(), []uint32{1, 2})
	assert.NoError(t, err)
}

func TestDeleteTeams_Fail_ReturnError(t *testing.T) {
	createBaseData(nil)

	err := teamsRepo.DeleteTeams(context.Background(), []uint32{3})
	t.Log(err)
	assert.Error(t, err)
}

func TestGetTeams_Success(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}

	err := teamsRepo.CreateTeams(context.Background(), teams)
	assert.NoError(t, err)

	team, err := teamsRepo.GetTeams(context.Background(), teams[0].ID)
	t.Logf("%v", *teams[0])
	assert.Equal(t, teams[0], team)
	assert.NoError(t, err)
}

func TestGetTeams_Fail(t *testing.T) {
	createBaseData(nil)

	team, err := teamsRepo.GetTeams(context.Background(), 3)
	assert.Nil(t, team)
	t.Logf("%v", err)
	assert.Error(t, err)
}

func TestListTeams_NoFilter_ReturnsAllTeams(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}

	createBaseData(teams)

	teams_all, err := teamsRepo.ListTeams(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, teams, teams_all)
}

func TestListTeams_WithPagination_ReturnsPaginatedTeams(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}

	createBaseData(teams)

	teams_page_2, err := teamsRepo.ListTeams(context.Background(),
		nil, &repo.TeamsFilter{Page: 2, PageSize: 1})
	assert.NoError(t, err)
	assert.Equal(t, teams[1:], teams_page_2)
}

func TestListTeams_Ids_ReturnsFilteredTeams(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}
	createBaseData(teams)
	filter := &repo.TeamsFilter{Ids: []uint32{1, 2, 3}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams, teams_filtered)
}

func TestListTeams_Codes_ReturnsFilteredTeams(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}
	createBaseData(teams)
	filter := &repo.TeamsFilter{Codes: []string{"team1", "team2"}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams, teams_filtered)
	for _, _t := range teams_filtered {
		t.Log(_t)
	}
}

func TestListTeams_Leaders_ReturnsFilteredTeams(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}
	createBaseData(teams)
	filter := &repo.TeamsFilter{Leaders: []string{"leader1", "leader2"}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams, teams_filtered)
	for _, _t := range teams_filtered {
		t.Log(_t)
	}
}

func TestListTeams_WithNamesFilter_ReturnsFilteredTeams(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}
	createBaseData(teams)
	filter := &repo.TeamsFilter{Names: []string{"Team1", "Team2"}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams, teams_filtered)
	for _, _t := range teams_filtered {
		t.Log(_t)
	}
}

func TestListTeams_Partial_ReturnsFiltered(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}
	createBaseData(teams)
	filter := &repo.TeamsFilter{Names: []string{"Team3", "Team2"}}

	teams_filtered, err := teamsRepo.ListTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, teams[1:], teams_filtered)
	for _, _t := range teams_filtered {
		t.Log(_t)
	}
}

func TestCountTeams_NoFilter_AllTeamsCounted(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}
	createBaseData(teams)
	count, err := teamsRepo.CountTeams(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}

func TestCountTeams_Id_PartialTeamsCounted(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}
	createBaseData(teams)
	filter := &repo.TeamsFilter{Ids: []uint32{1, 3}}
	count, err := teamsRepo.CountTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func TestCountTeams_Id_Empty(t *testing.T) {
	teams := []*repo.Team{
		{Name: "Team1", Code: "team1", Leader: "leader1", Description: "description1"},
		{Name: "Team2", Code: "team2", Leader: "leader2", Description: "description2"},
	}
	createBaseData(teams)
	filter := &repo.TeamsFilter{Ids: []uint32{4, 3}}
	count, err := teamsRepo.CountTeams(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}
