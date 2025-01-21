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

// TestCreateTeams tests the CreateTeams method of the TeamsUsecase.
func TestCreateTeams(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "user")
	teamRepo := new(MockTeamsRepo)
	authzrepo := new(MockAuthzRepo)
	hgrepo := new(MockHostgroupsRepo)
	hgteamrepo := new(MockHostgroupTeamsRepo)
	apprepo := new(MockApplicationsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTeamsUsecase(
		teamRepo,
		authzrepo,
		hgrepo,
		hgteamrepo,
		apprepo,
		nil,
		txm,
	)

	// Test case: Validation fails
	teams := []*biz.Team{{Name: "Name", Code: "validcode"}}
	err := usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Validation fails
	teams = []*biz.Team{{Name: "0-name", Code: "validcode"}}
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	teams = []*biz.Team{{Name: "-name", Code: "validcode"}}
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	teams = []*biz.Team{{Name: "name-", Code: "validcode"}}
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	teams = []*biz.Team{{Name: "nAme", Code: "validcode"}}
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Validation fails
	teams = []*biz.Team{{Name: "Invalid Name", Code: "validcode"}}
	// should not call repo
	// repoCall := teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Conversion fails
	teams = []*biz.Team{{Name: "ValidTeam", Code: "invalid code"}}
	// should not call repo
	// repoCall = teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Invalid Leader
	teams = []*biz.Team{{Name: "ValidTeam", Code: "validcode"}}
	// not call repo
	// repoCall = teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	// enforce fail
	teams = []*biz.Team{{Name: "valid", Code: "validcode", LeaderId: 1}}
	authcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, errors.New("enforce failed"))
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)
	authcall.Unset()

	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// Test case: Creation fails
	teams = []*biz.Team{{Name: "valid", Code: "validcode", LeaderId: 1}}
	repoCall := teamRepo.On("CreateTeams", ctx, mock.Anything, mock.Anything).Return(errors.New("creation failed"))
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)
	repoCall.Unset()

	// Test case: Successful creation
	teamRepo.On("CreateTeams", ctx, mock.Anything, mock.Anything).Return(nil)
	teams = []*biz.Team{{Name: "name", Code: "validcode", LeaderId: 1}}
	err = usecase.CreateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Name: "name0", Code: "validcode0", LeaderId: 1}}
	err = usecase.CreateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Name: "name-0", Code: "validcode-0", LeaderId: 1}}
	err = usecase.CreateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Name: "name-0", Code: "1-validcode-0", LeaderId: 1}}
	err = usecase.CreateTeams(ctx, teams)
	assert.NoError(t, err)
}

func TestUpdateTeams(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "user")
	teamRepo := new(MockTeamsRepo)
	authzrepo := new(MockAuthzRepo)
	hgrepo := new(MockHostgroupsRepo)
	hgteamrepo := new(MockHostgroupTeamsRepo)
	apprepo := new(MockApplicationsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTeamsUsecase(
		teamRepo,
		authzrepo,
		hgrepo,
		hgteamrepo,
		apprepo,
		nil,
		txm,
	)

	// Test case: Validation fails
	teams := []*biz.Team{{Id: 1, Name: "Name", Code: "validcode"}}
	err := usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Validation fails
	teams = []*biz.Team{{Id: 1, Name: "0-name", Code: "validcode"}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	teams = []*biz.Team{{Id: 1, Name: "-name", Code: "validcode"}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	teams = []*biz.Team{{Id: 1, Name: "name-", Code: "validcode"}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	teams = []*biz.Team{{Id: 1, Name: "nAme", Code: "validcode"}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Validation fails
	teams = []*biz.Team{{Id: 1, Name: "Invalid Name", Code: "validcode"}}
	// should not call repo
	// repoCall := teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Conversion fails
	teams = []*biz.Team{{Id: 1, Name: "ValidTeam", Code: "invalid code"}}
	// should not call repo
	// repoCall = teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Invalid Leader
	teams = []*biz.Team{{Id: 1, Name: "ValidTeam", Code: "validcode"}}
	// not call repo
	// repoCall = teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	// no id
	teams = []*biz.Team{{Name: "name", Code: "validcode", LeaderId: 1}}
	// not call repo
	// repoCall = teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	// enforce fail
	teams = []*biz.Team{{Id: 1, Name: "valid", Code: "validcode", LeaderId: 1}}
	authcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, errors.New("enforce failed"))
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)
	authcall.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// Test case: Creation fails
	teams = []*biz.Team{{Id: 1, Name: "valid", Code: "validcode", LeaderId: 1}}
	repoCall := teamRepo.On("UpdateTeams", ctx, mock.Anything, mock.Anything).Return(errors.New("creation failed"))
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)
	repoCall.Unset()

	// Test case: Successful creation
	teamRepo.On("UpdateTeams", ctx, mock.Anything, mock.Anything).Return(nil)
	teams = []*biz.Team{{Id: 1, Name: "name", Code: "validcode", LeaderId: 1}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Id: 1, Name: "name0", Code: "validcode0", LeaderId: 1}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Id: 1, Name: "name-0", Code: "validcode-0", LeaderId: 1}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Id: 1, Name: "name-0", Code: "1-validcode-0", LeaderId: 1}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.NoError(t, err)
}

func TestDeleteTeams(t *testing.T) {
	ctx := context.WithValue(context.Background(), data.CtxUserName, "user")
	teamRepo := new(MockTeamsRepo)
	authzrepo := new(MockAuthzRepo)
	hgrepo := new(MockHostgroupsRepo)
	hgteamrepo := new(MockHostgroupTeamsRepo)
	apprepo := new(MockApplicationsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTeamsUsecase(
		teamRepo,
		authzrepo,
		hgrepo,
		hgteamrepo,
		apprepo,
		nil,
		txm,
	)

	// Test case: Validation fails
	teams := []uint32{}
	err := usecase.DeleteTeams(ctx, teams)
	assert.Error(t, err)

	err = usecase.DeleteTeams(ctx, nil)
	assert.Error(t, err)

	// enforce fail
	authcall := authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(false, errors.New("enforce failed"))
	err = usecase.DeleteTeams(ctx, teams)
	assert.Error(t, err)
	authcall.Unset()
	authzrepo.On("Enforce", ctx, mock.Anything, mock.Anything).Return(true, nil)

	// Test case: failed on hostgroup need check fail
	teams = []uint32{1, 2}
	hgrepoCall := hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(1), nil)
	apprepoCall := apprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(0), nil)
	htrepoCall := hgteamrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(0), nil)
	teamrepoCall := teamRepo.On("DeleteTeams", ctx, mock.Anything).Return(nil)
	err = usecase.DeleteTeams(ctx, teams)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	hgrepoCall.Unset()
	apprepoCall.Unset()
	htrepoCall.Unset()
	teamrepoCall.Unset()

	// Test case: failed on app need check fail
	teams = []uint32{1, 2}
	hgrepoCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(0), nil)
	apprepoCall = apprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(1), nil)
	htrepoCall = hgteamrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(0), nil)
	teamrepoCall = teamRepo.On("DeleteTeams",
		ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteTeams(ctx, teams)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	hgrepoCall.Unset()
	apprepoCall.Unset()
	htrepoCall.Unset()
	teamrepoCall.Unset()

	// Test case: failed on hostgroup-team need check fail
	teams = []uint32{1, 2}
	hgrepoCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(0), nil)
	apprepoCall = apprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(0), nil)
	htrepoCall = hgteamrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(1), nil)
	teamrepoCall = teamRepo.On("DeleteTeams",
		ctx, mock.Anything, mock.Anything).Return(nil)
	err = usecase.DeleteTeams(ctx, teams)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	hgrepoCall.Unset()
	apprepoCall.Unset()
	htrepoCall.Unset()
	teamrepoCall.Unset()

	// Test case: failed on delete
	teams = []uint32{1, 2}
	hgrepoCall = hgrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(0), nil)
	apprepoCall = apprepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(0), nil)
	htrepoCall = hgteamrepo.On("CountRequire",
		ctx, mock.Anything, repo.RequireTeam, teams).Return(int64(0), nil)
	teamrepoCall = teamRepo.On("DeleteTeams",
		ctx, mock.Anything, teams).Return(errors.New("delete mock fail"))
	err = usecase.DeleteTeams(ctx, teams)
	assert.Error(t, err)
	t.Logf("error. %v", err)
	hgrepoCall.Unset()
	apprepoCall.Unset()
	htrepoCall.Unset()
	teamrepoCall.Unset()
}

func TestGetTeams(t *testing.T) {
	ctx := context.Background()
	teamRepo := new(MockTeamsRepo)
	authzrepo := new(MockAuthzRepo)
	hgrepo := new(MockHostgroupsRepo)
	hgteamrepo := new(MockHostgroupTeamsRepo)
	apprepo := new(MockApplicationsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTeamsUsecase(
		teamRepo,
		authzrepo,
		hgrepo,
		hgteamrepo,
		apprepo,
		nil,
		txm,
	)

	// id == 0
	team_id := uint32(0)
	_, err := usecase.GetTeams(ctx, team_id)
	t.Logf("err. %s", err)
	assert.Error(t, err)

	// repo error
	call := teamRepo.On("GetTeams", ctx, team_id).Return(nil, errors.New("mock repo error"))
	_, err = usecase.GetTeams(ctx, team_id)
	t.Logf("err. %s", err)
	assert.Error(t, err)
	call.Unset()

	// success
	team_id = uint32(1)
	db_team := repo.Team{
		ID:          team_id,
		Name:        "team1",
		Code:        "code1",
		LeaderId:    1,
		Description: "team1 description",
	}
	biz_team := &biz.Team{
		Id:          team_id,
		Name:        "team1",
		Code:        "code1",
		LeaderId:    1,
		Description: "team1 description",
	}
	call = teamRepo.On("GetTeams", ctx, team_id).Return(&db_team, nil)
	team, err := usecase.GetTeams(ctx, team_id)
	t.Logf("team. %+v", team)
	assert.NoError(t, err)
	assert.Equal(t, biz_team, team)
	call.Unset()
}

func TestListTeams(t *testing.T) {
	ctx := context.Background()
	teamRepo := new(MockTeamsRepo)
	authzrepo := new(MockAuthzRepo)
	hgrepo := new(MockHostgroupsRepo)
	hgteamrepo := new(MockHostgroupTeamsRepo)
	apprepo := new(MockApplicationsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTeamsUsecase(
		teamRepo,
		authzrepo,
		hgrepo,
		hgteamrepo,
		apprepo,
		nil,
		txm,
	)

	// filter error. codes exceeds
	filter := biz.ListTeamsFilter{
		Codes: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
	}
	_, e := usecase.ListTeams(context.Background(), &filter)
	assert.Equal(t, e, biz.ErrFilterValuesExceedMax)
	// filter error. ids exceeds
	filter = biz.ListTeamsFilter{
		Ids: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34},
	}
	_, e = usecase.ListTeams(context.Background(), &filter)
	assert.Equal(t, e, biz.ErrFilterValuesExceedMax)
	// filter error. name exceeds
	filter = biz.ListTeamsFilter{
		Names: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34"},
	}
	_, e = usecase.ListTeams(context.Background(), &filter)
	assert.Equal(t, e, biz.ErrFilterValuesExceedMax)
	// filter error. leaders exceeds
	filter = biz.ListTeamsFilter{
		LeadersId: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34},
	}
	_, e = usecase.ListTeams(context.Background(), &filter)
	assert.Equal(t, e, biz.ErrFilterValuesExceedMax)
	// filter error. pagesize 0
	filter = biz.ListTeamsFilter{
		PageSize: 0,
	}
	_, e = usecase.ListTeams(context.Background(), &filter)
	assert.Equal(t, e, biz.ErrFilterInvalidPagesize)
	// filter error. pagesize exceeds
	filter = biz.ListTeamsFilter{
		PageSize: 201,
	}
	_, e = usecase.ListTeams(context.Background(), &filter)
	assert.Equal(t, e, biz.ErrFilterInvalidPagesize)
	// repo error
	filter = *biz.DefaultTeamsFilter()
	call := teamRepo.On("ListTeams",
		ctx, mock.Anything, mock.Anything).Return(nil, errors.New("repo error"))
	_, e = usecase.ListTeams(context.Background(), &filter)
	assert.Equal(t, e, errors.New("repo error"))
	call.Unset()

	// success
	db_teams := []*repo.Team{
		{
			ID:       1,
			Code:     "team1",
			LeaderId: 1,
			Name:     "team1",
		},
		{
			ID:       2,
			Name:     "team2",
			Code:     "team2",
			LeaderId: 2,
		},
	}
	biz_teams := []*biz.Team{
		{
			Id:          1,
			Code:        "team1",
			Description: "",
			LeaderId:    1,
			Name:        "team1",
		},
		{
			Id:          2,
			Code:        "team2",
			Description: "",
			LeaderId:    2,
			Name:        "team2",
		},
	}
	teamRepo.On("ListTeams", mock.Anything, mock.Anything, mock.Anything).Return(db_teams, nil)

	ts, e := usecase.ListTeams(ctx, &filter)
	assert.Nil(t, e)
	assert.Equal(t, biz_teams, ts)
}
