package biz

import (
	"appix/internal/conf"
	"appix/internal/data"
	"appix/internal/data/repo"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase struct {
	adminRepo repo.AdminRepo
	tokenRepo repo.TokenRepo
	authzRepo repo.AuthzRepo
	teamsRepo repo.TeamsRepo
	txm       repo.TxManager
	log       *log.Helper
	conf      *conf.Admin
	adminUser *repo.User
}

const AdminUser = "admin"
const AdminTeam = "admin-team"

const BcryptPrefix = "{bcrypt}"

func HashPassword(password string) (string, error) {
	// 生成盐并对密码进行加密，默认成本为 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", BcryptPrefix, string(hashedPassword)), nil
}

func CheckPassword(hashedPassword, password string) bool {
	// 比较输入的密码和存储的哈希密码
	if !strings.HasPrefix(hashedPassword, BcryptPrefix) {
		return false
	}
	hash := strings.TrimLeft(hashedPassword, BcryptPrefix)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewAdminUsecase(
	conf *conf.Admin,
	adminRepo repo.AdminRepo,
	tokenRepo repo.TokenRepo,
	authzRepo repo.AuthzRepo,
	teamsRepo repo.TeamsRepo,
	txm repo.TxManager,
	logger log.Logger,
) *AdminUsecase {

	uc := &AdminUsecase{
		adminRepo: adminRepo,
		tokenRepo: tokenRepo,
		authzRepo: authzRepo,
		teamsRepo: teamsRepo,
		txm:       txm,
		log:       log.NewHelper(logger),
		conf:      conf,
	}

	// get admin user
	admin_users, err := adminRepo.ListUsers(context.Background(), nil, &repo.UsersFilter{
		UserName: []string{AdminUser},
	})
	if err != nil {
		panic(err)
	}
	if len(admin_users) == 0 {
		// create admin user
		md5pass, err := HashPassword(conf.AdminPassword)
		if err != nil {
			panic(err)
		}
		admin_user := &repo.User{
			UserName: AdminUser,
			Password: md5pass,
		}
		err = adminRepo.CreateUsers(context.Background(), nil, []*repo.User{admin_user})
		if err != nil {
			panic(err)
		}
		uc.adminUser = admin_user
	} else {
		// update admin user
		admin_user := admin_users[0]
		md5pass, err := HashPassword(conf.AdminPassword)
		if err != nil {
			panic(err)
		}
		if admin_user.Password != md5pass {
			admin_user.Password = md5pass
			err = adminRepo.UpdateUsers(context.Background(), nil, []*repo.User{admin_user})
			if err != nil {
				panic(err)
			}
		}
		uc.adminUser = admin_user
	}

	// create admin team
	admin_team, err := teamsRepo.ListTeams(context.Background(), nil, &repo.TeamsFilter{
		Names: []string{AdminTeam},
	})
	if err != nil {
		panic(err)
	}
	if len(admin_team) == 0 {
		// create admin team
		_admin_team := &repo.Team{
			Name:        AdminTeam,
			Code:        AdminTeam,
			Leader:      AdminUser,
			Description: "admin team",
		}
		err = teamsRepo.CreateTeams(context.Background(), []*repo.Team{_admin_team})
		if err != nil {
			panic(err)
		}
	}

	// create rules
	err = authzRepo.CreateGroup(context.Background(), nil, &repo.Group{
		User: AdminUser,
		Role: AdminTeam,
	})
	if err != nil {
		panic(err)
	}

	resourceAll := repo.NewResource4Sv1("", "", "", "")

	err = authzRepo.CreateRule(context.Background(), nil, &repo.Rule{
		Sub:      AdminTeam,
		Resource: resourceAll,
		Action:   repo.ActWrite,
	})
	if err != nil {
		panic(err)
	}

	return uc

}

func (s *AdminUsecase) validate(isNew bool, users []*User) error {
	for _, user := range users {
		if err := user.Validate(&UserOptions{
			IsNew:                isNew,
			StrictPasswordPolicy: s.conf.StrictPasswordPolicy,
		}); err != nil {
			return err
		}
	}
	return nil
}

func getUsername(ctx context.Context) (string, error) {

	uname := ctx.Value(data.UserName)
	if uname == nil {
		return "", errors.New("user name is nil")
	}
	usernameStr, ok := uname.(string)
	if !ok {
		return "", errors.New("user name is not string")
	}
	return usernameStr, nil
}

func (s *AdminUsecase) enforceUserAdmin(team, user, sub string) error {
	ires := repo.NewResource4Sv1("users", team, user, sub)
	can, e := s.authzRepo.Enforce(context.Background(), nil, &repo.AuthenRequest{
		Sub:      sub,
		Resource: ires,
		Action:   repo.ActWrite,
	})
	if e != nil {
		return e
	}
	if !can {
		return errors.New("no permission")
	}
	return nil
}

// CreateUsers is
func (s *AdminUsecase) CreateUsers(ctx context.Context, users []*User) error {
	if err := s.validate(true, users); err != nil {
		return errors.Join(errors.New("CreateUsers failed"), err)
	}
	if len(users) == 0 {
		return errors.Join(errors.New("CreateUsers failed"), errors.New("no user to create"))
	}

	usernameStr, err := getUsername(ctx)
	if err != nil {
		return errors.Join(errors.New("CreateUsers failed"), err)
	}
	repoUsers, err := ToRepoUsers(users)
	if err != nil {
		return errors.Join(errors.New("CreateUsers failed"), err)
	}

	err = s.txm.RunInTX(func(tx repo.TX) error {
		if e := s.enforceUserAdmin("", "", usernameStr); err != nil {
			return e
		}
		e := s.adminRepo.CreateUsers(ctx, tx, repoUsers)
		if e != nil {
			return e
		}

		// create user policy
		for _, user := range repoUsers {
			ires := repo.NewResource4Sv1("", "", "", user.UserName)
			e := s.authzRepo.CreateRule(ctx, tx, &repo.Rule{
				Sub:      user.UserName,
				Resource: ires,
				Action:   repo.ActWrite,
			})
			if e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		return errors.Join(errors.New("CreateUsers failed"), err)
	}
	return nil
}

// UpdateUsers is
func (s *AdminUsecase) UpdateUsers(ctx context.Context, users []*User) error {
	if err := s.validate(false, users); err != nil {
		return errors.Join(errors.New("UpdateUsers failed"), err)
	}
	if len(users) == 0 {
		return errors.Join(errors.New("UpdateUsers failed"), errors.New("no user to update"))
	}

	repoUsers, err := ToRepoUsers(users)
	if err != nil {
		return errors.Join(errors.New("UpdateUsers failed"), err)
	}

	usernameStr, err := getUsername(ctx)
	if err != nil {
		return errors.Join(errors.New("UpdateUsers failed"), err)
	}

	err = s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforceUserAdmin("", "", usernameStr); err != nil {
			return err
		}

		err = s.adminRepo.UpdateUsers(ctx, tx, repoUsers)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.Join(errors.New("UpdateUsers failed"), err)
	}
	return nil
}

// DeleteUsers is
func (s *AdminUsecase) DeleteUsers(ctx context.Context, tx repo.TX, ids []uint32) error {
	if len(ids) == 0 {
		return errors.Join(errors.New("DeleteUsers failed"), errors.New("no user to delete"))
	}
	// check admin user
	for _, _id := range ids {
		if _id == s.adminUser.Id {
			return errors.Join(errors.New("DeleteUsers failed"), errors.New("can not delete admin user"))
		}
	}

	usernameStr, err := getUsername(ctx)
	if err != nil {
		return errors.Join(errors.New("DeleteUsers failed"), err)
	}
	// return s.adminRepo.DeleteUsers(ctx, tx, ids)
	err = s.txm.RunInTX(func(tx repo.TX) error {
		if err := s.enforceUserAdmin("", "", usernameStr); err != nil {
			return err
		}
		err = s.adminRepo.DeleteUsers(ctx, tx, ids)
		if err != nil {
			return err
		}
		// delete authz
		for _, _id := range ids {
			user, err := s.adminRepo.GetUsers(ctx, tx, _id)
			if err != nil {
				return err
			}
			ires := repo.NewResource4Sv1("", "", "", user.UserName)
			err = s.authzRepo.DeleteRule(ctx, tx, &repo.Rule{
				Sub:      user.UserName,
				Resource: ires,
				Action:   repo.ActWrite,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return errors.Join(errors.New("DeleteUsers failed"), err)
	}
	return nil
}

// GetUsers is
func (s *AdminUsecase) GetUsers(ctx context.Context, id uint32) (*User, error) {
	user, err := s.adminRepo.GetUsers(ctx, nil, id)
	if err != nil {
		return nil, errors.Join(errors.New("GetUsers failed"), err)
	}
	return ToBizUser(user), nil
}

// ListUsers is
func (s *AdminUsecase) ListUsers(ctx context.Context, tx repo.TX, filter *ListUsersFilter) ([]*User, error) {
	repoUsers, err := s.adminRepo.ListUsers(ctx, tx, ToDBUsersFilter(filter))
	if err != nil {
		return nil, err
	}
	return ToBizUsers(repoUsers), nil
}

// Login is
func (s *AdminUsecase) Login(ctx context.Context, username, password string) (*User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password is empty")
	}
	var user *repo.User
	var err error

	users, err := s.adminRepo.ListUsers(context.Background(), nil, &repo.UsersFilter{
		UserName: []string{username},
	})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 || len(users) > 1 {
		return nil, errors.New("user not found")
	} else {
		user = users[0]
	}
	if !CheckPassword(user.Password, password) {
		return nil, errors.New("password is incorrect")
	}

	_idstr := strconv.Itoa(int(user.Id))

	token, err := s.tokenRepo.CreateToken(ctx, repo.TokenClaims{
		string(data.UserId):   _idstr,
		string(data.UserName): user.UserName,
	})
	if err != nil {
		return nil, err
	}

	bizUser := ToBizUser(user)
	bizUser.Token = token

	return bizUser, nil
}

// Logout is
func (s *AdminUsecase) Logout(ctx context.Context, id uint32) error {
	_idClaim := ctx.Value(data.UserId)
	_id := strconv.Itoa(int(id))
	if _id != _idClaim {
		return errors.Join(errors.New("logout failed"), errors.New("invalid user id"))
	}
	tokenStr := ctx.Value(data.UserTokenKey)

	if tokenStr != nil {
		_tokenStr, ok := tokenStr.(string)
		if ok {
			err := s.tokenRepo.DeleteToken(ctx, _tokenStr)
			if err != nil {
				return errors.Join(errors.New("logout failed"), err)
			}
		}
	}

	return s.adminRepo.Logout(ctx, id)
}
