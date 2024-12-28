package biz_test

import (
	"appix/internal/biz"
	"appix/internal/data/repo"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTeamRepo is a mock implementation of the TeamRepository interface.
type MockTeamRepo struct {
	mock.Mock
}

// CreateTeams is a mock implementation of the CreateTeams method.
func (m *MockTeamRepo) CreateTeams(ctx context.Context, teams []*repo.TeamsRepo) error {
	args := m.Called(ctx, teams)
	return args.Error(0)
}

// TestCreateTeams tests the CreateTeams method of the TeamsUsecase.
func TestCreateTeams(t *testing.T) {
	ctx := context.Background()
	teamRepo := new(MockTeamRepo)
	usecase := &biz.TeamsUsecase{teamRepo: teamRepo}

	// Test case: Validation fails
	teams := []*biz.Team{{Name: "Invalid Team"}}
	teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err := usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Conversion fails
	teams = []*biz.Team{{Name: "Valid Team"}}
	teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)

	// Test case: Successful creation
	teams = []*biz.Team{{Name: "Valid Team"}}
	teamRepo.On("CreateTeams", ctx, mock.Anything).Return(nil)
	err = usecase.CreateTeams(ctx, teams)
	assert.NoError(t, err)

	// Test case: Creation fails
	teams = []*biz.Team{{Name: "Valid Team"}}
	teamRepo.On("CreateTeams", ctx, mock.Anything).Return(errors.New("creation failed"))
	err = usecase.CreateTeams(ctx, teams)
	assert.Error(t, err)
}
