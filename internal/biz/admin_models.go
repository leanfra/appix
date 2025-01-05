package biz

import (
	"appix/internal/data/repo"
	"errors"
	"strings"
)

type User struct {
	Id       uint32
	UserName string
	Password string
	Email    string
	Phone    string
	Token    string
}

type ListUsersFilter struct {
	Page      uint32
	PageSize  uint32
	Ids       []uint32
	UserNames []string
	Emails    []string
	Phones    []string
}

func (lf *ListUsersFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if lf.Page == 0 {
		return ErrFilterInvalidPage
	}
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	if len(lf.Ids) > MaxFilterValues ||
		len(lf.UserNames) > MaxFilterValues ||
		len(lf.Emails) > MaxFilterValues ||
		len(lf.Phones) > MaxFilterValues {
		return ErrFilterValuesExceedMax
	}
	return nil
}

func ToDBUsersFilter(filter *ListUsersFilter) *repo.UsersFilter {
	if filter == nil {
		return nil
	}
	return &repo.UsersFilter{
		Ids:      filter.Ids,
		UserName: filter.UserNames,
		Email:    filter.Emails,
		Phone:    filter.Phones,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}

type UserOptions struct {
	IsNew                bool
	StrictPasswordPolicy bool
}

func (m *User) Validate(opts *UserOptions) error {
	if opts.IsNew {
		if m.Id == 0 {
			return errors.New("invalid user id")
		}
	}
	if m.Email != "" {
		if !strings.Contains(m.Email, "@") || !strings.Contains(m.Email, ".") {
			return errors.New("invalid email format")
		}
	}
	if opts.StrictPasswordPolicy {
		// Check password policy
		hasNumber := strings.ContainsAny(m.Password, "0123456789")
		hasLetter := strings.ContainsAny(m.Password, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		hasSpecial := strings.ContainsAny(m.Password, "!@#$%^&*()_+-=[]{}|;:,.<>?")
		if !hasNumber || !hasLetter || !hasSpecial {
			return errors.New("password must contain at least one number, one letter and one special character")
		}
		if len(m.Password) < 8 {
			return errors.New("password must be at least 8 characters long")
		}
	}
	return nil
}

func ToRepoUsers(users []*User) []*repo.User {
	repoUsers := make([]*repo.User, 0, len(users))
	for _, user := range users {
		repoUsers = append(repoUsers, ToRepoUser(user))
	}
	return repoUsers
}

func ToRepoUser(m *User) *repo.User {
	return &repo.User{
		UserName: m.UserName,
		Password: m.Password,
		Email:    m.Email,
		Phone:    m.Phone,
	}
}

func ToBizUser(user *repo.User) *User {
	return &User{
		Id:       user.Id,
		UserName: user.UserName,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}

func ToBizUsers(users []*repo.User) []*User {
	bizUsers := make([]*User, 0, len(users))
	for _, user := range users {
		bizUsers = append(bizUsers, ToBizUser(user))
	}
	return bizUsers
}
