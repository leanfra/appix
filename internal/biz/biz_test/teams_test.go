package biz_test

import (
	"appix/internal/biz"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateTeams tests the CreateTeams method of the TeamsUsecase.
func TestCreateTeams(t *testing.T) {
	ctx := context.Background()
	teamRepo := new(MockTeamsRepo)
	hgrepo := new(MockHostgroupsRepo)
	hgteamrepo := new(MockHostgroupTeamsRepo)
	apprepo := new(MockApplicationsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTeamsUsecase(
		teamRepo,
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

	// Test case: Creation fails
	teams = []*biz.Team{{Name: "valid", Code: "validcode", Leader: "leader1"}}
	repoCall := teamRepo.On("CreateTeams", ctx, mock.Anything).Return(errors.New("creation failed"))
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)
	repoCall.Unset()

	// Test case: Successful creation
	teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	teams = []*biz.Team{{Name: "name", Code: "validcode", Leader: "leader1"}}
	err = usecase.CreateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Name: "name0", Code: "validcode0", Leader: "leader1"}}
	err = usecase.CreateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Name: "name-0", Code: "validcode-0", Leader: "leader1"}}
	err = usecase.CreateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Name: "name-0", Code: "1-validcode-0", Leader: "leader1"}}
	err = usecase.CreateTeams(ctx, teams)
	assert.NoError(t, err)
}

func TestUpdateTeams(t *testing.T) {
	ctx := context.Background()
	teamRepo := new(MockTeamsRepo)
	hgrepo := new(MockHostgroupsRepo)
	hgteamrepo := new(MockHostgroupTeamsRepo)
	apprepo := new(MockApplicationsRepo)
	txm := new(MockTXManager)
	usecase := biz.NewTeamsUsecase(
		teamRepo,
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
	teams = []*biz.Team{{Name: "name", Code: "validcode", Leader: "leader1"}}
	// not call repo
	// repoCall = teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Creation fails
	teams = []*biz.Team{{Id: 1, Name: "valid", Code: "validcode", Leader: "leader1"}}
	repoCall := teamRepo.On("UpdateTeams", ctx, mock.Anything).Return(errors.New("creation failed"))
	err = usecase.UpdateTeams(ctx, teams)
	assert.Error(t, err)
	repoCall.Unset()

	// Test case: Successful creation
	teamRepo.On("UpdateTeams", ctx, mock.Anything).Return(nil)
	teams = []*biz.Team{{Id: 1, Name: "name", Code: "validcode", Leader: "leader1"}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Id: 1, Name: "name0", Code: "validcode0", Leader: "leader1"}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Id: 1, Name: "name-0", Code: "validcode-0", Leader: "leader1"}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.NoError(t, err)

	teams = []*biz.Team{{Id: 1, Name: "name-0", Code: "1-validcode-0", Leader: "leader1"}}
	err = usecase.UpdateTeams(ctx, teams)
	assert.NoError(t, err)
}
