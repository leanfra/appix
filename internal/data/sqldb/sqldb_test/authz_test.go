package sqldb_test

import (
	"context"
	"testing"

	"appix/internal/conf"
	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
)

var authzRepo repo.AuthzRepo

func initAuthzRepo() {
	dataMem := getDataMem()
	config := &conf.Authz{
		ModelFile: "../../../../configs/rbac_model.conf",
	}
	authzRepo, _ = sqldb.NewAuthzRepoGorm(config, dataMem, logger)
}

func Test__AuthzRepoGorm(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{"CreateRule_Success", TestAuthzRepoGorm_CreateRule},
		{"DeleteRule_Success", TestAuthzRepoGorm_DeleteRule},
		{"ListRule_Success", TestAuthzRepoGorm_ListRule},
		{"Enforce_Success", TestAuthzRepoGorm_Enforce},
		{"CreateGroup_Success", TestAuthzRepoGorm_CreateGroup},
		{"DeleteGroup_Success", TestAuthzRepoGorm_DeleteGroup},
		{"ListGroup_Success", TestAuthzRepoGorm_ListGroup},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func TestAuthzRepoGorm_CreateRule(t *testing.T) {
	initAuthzRepo()

	ires := repo.NewResource4Sv1("data", "team1", "data1", "alice")
	rule := &repo.Rule{Sub: "alice", Resource: ires, Action: "read"}
	err := authzRepo.CreateRule(context.Background(), nil, rule)
	assert.NoError(t, err)

}

func TestAuthzRepoGorm_DeleteRule(t *testing.T) {
	initAuthzRepo()

	ires := repo.NewResource4Sv1("data", "team1", "data1", "alice")
	rule := &repo.Rule{Sub: "alice", Resource: ires, Action: "read"}
	err := authzRepo.CreateRule(context.Background(), nil, rule)
	assert.NoError(t, err)

	err = authzRepo.DeleteRule(context.Background(), nil, rule)
	assert.NoError(t, err)

	err = authzRepo.DeleteRule(context.Background(), nil, rule)
	assert.Error(t, err)
}

func TestAuthzRepoGorm_ListRule(t *testing.T) {
	initAuthzRepo()
	ires := repo.NewResource4Sv1("data", "team1", "data1", "alice")
	rule := &repo.Rule{Sub: "alice", Resource: ires, Action: "read"}
	err := authzRepo.CreateRule(context.Background(), nil, rule)
	assert.NoError(t, err)

	rules, err := authzRepo.ListRule(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, rules[0], rule)
}

func TestAuthzRepoGorm_Enforce(t *testing.T) {
	initAuthzRepo()

	ires := repo.NewResource4Sv1("data", "", "", "alice")
	rule := &repo.Rule{Sub: "alice", Resource: ires, Action: "read"}
	err := authzRepo.CreateRule(context.Background(), nil, rule)
	assert.NoError(t, err)

	ires = repo.NewResource4Sv1("data", "team1", "data1", "alice")
	request := &repo.AuthenRequest{Sub: "alice", Resource: ires, Action: "read"}
	allowed, err := authzRepo.Enforce(context.Background(), nil, request)
	assert.NoError(t, err)
	assert.True(t, allowed)
}

func TestAuthzRepoGorm_CreateGroup(t *testing.T) {
	initAuthzRepo()
	group := &repo.Group{User: "alice", Role: "admin"}
	err := authzRepo.CreateGroup(context.Background(), nil, group)
	assert.NoError(t, err)

}

func TestAuthzRepoGorm_DeleteGroup(t *testing.T) {
	initAuthzRepo()
	group := &repo.Group{User: "alice", Role: "admin"}
	err := authzRepo.CreateGroup(context.Background(), nil, group)
	assert.NoError(t, err)

	err = authzRepo.DeleteGroup(context.Background(), nil, group)
	assert.NoError(t, err)

	err = authzRepo.DeleteGroup(context.Background(), nil, group)
	assert.Error(t, err)
}

func TestAuthzRepoGorm_ListGroup(t *testing.T) {

	initAuthzRepo()
	group := &repo.Group{User: "alice", Role: "admin"}
	err := authzRepo.CreateGroup(context.Background(), nil, group)
	assert.NoError(t, err)

	groups, err := authzRepo.ListGroup(context.Background(), nil, nil)
	assert.NoError(t, err)
	assert.Len(t, groups, 1)
}
