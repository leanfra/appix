package sqldb_test

import (
	"context"
	"errors"
	"testing"

	"appix/internal/data/repo"
	"appix/internal/data/sqldb"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var adminRepo repo.AdminRepo

func initAdminRepo() {
	dataMem = getDataMem()
	adminRepo, _ = sqldb.NewAdminRepoGorm(dataMem, logger)
}

func TestAdminRepoGorm_CreateUsers(t *testing.T) {

	initAdminRepo()

	db := dataMem.DB
	ctx := context.Background()

	users := []*repo.User{
		{UserName: "user1", Email: "user1@example.com", Password: "password123", Phone: "1234567890"},
		{UserName: "user2", Email: "user2@example.com", Password: "password456", Phone: "0987654321"},
	}

	err := adminRepo.CreateUsers(ctx, nil, users)
	assert.NoError(t, err)

	var createdUsers []repo.User
	db.Find(&createdUsers)
	assert.Len(t, createdUsers, 2)
}

func TestAdminRepoGorm_UpdateUsers(t *testing.T) {

	initAdminRepo()
	db := dataMem.DB

	ctx := context.Background()

	user := &repo.User{UserName: "user1", Email: "user1@example.com"}
	db.Create(user)

	updatedUser := &repo.User{Id: user.Id, UserName: "user1_updated", Email: "user1_updated@example.com"}
	err := adminRepo.UpdateUsers(ctx, nil, []*repo.User{updatedUser})
	assert.NoError(t, err)

	var updatedUserFromDB repo.User
	db.First(&updatedUserFromDB, user.Id)
	assert.Equal(t, "user1_updated", updatedUserFromDB.UserName)
	assert.Equal(t, "user1_updated@example.com", updatedUserFromDB.Email)
}

func TestAdminRepoGorm_DeleteUsers(t *testing.T) {

	initAdminRepo()
	db := dataMem.DB

	ctx := context.Background()

	user := &repo.User{UserName: "user1", Email: "user1@example.com"}
	db.Create(user)

	err := adminRepo.DeleteUsers(ctx, nil, []uint32{user.Id})
	assert.NoError(t, err)

	var deletedUser repo.User
	err = db.First(&deletedUser, user.Id).Error
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestAdminRepoGorm_GetUsers(t *testing.T) {

	initAdminRepo()
	db := dataMem.DB

	ctx := context.Background()

	user := &repo.User{UserName: "user1", Email: "user1@example.com"}
	db.Create(user)

	retrievedUser, err := adminRepo.GetUsers(ctx, nil, user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, retrievedUser.Id)
	assert.Equal(t, user.UserName, retrievedUser.UserName)
	assert.Equal(t, user.Email, retrievedUser.Email)
}

func TestAdminRepoGorm_ListUsers(t *testing.T) {

	initAdminRepo()
	db := dataMem.DB

	ctx := context.Background()

	users := []*repo.User{
		{UserName: "user1", Email: "user1@example.com", Phone: "1234567890", Password: "password123"},
		{UserName: "user2", Email: "user2@example.com", Phone: "0987654321", Password: "password456"},
	}
	db.Create(users)

	filter := &repo.UsersFilter{UserName: []string{"user1", "user2"}}
	retrievedUsers, err := adminRepo.ListUsers(ctx, nil, filter)
	assert.NoError(t, err)
	assert.Len(t, retrievedUsers, 2)
}

func TestAdminRepoGorm_Logout(t *testing.T) {

	initAdminRepo()

	ctx := context.Background()

	err := adminRepo.Logout(ctx, 1)
	assert.NoError(t, err)
}
