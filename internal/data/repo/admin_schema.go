package repo

import "context"

const UserTable = "users"

type User struct {
	Id       uint32 `gorm:"column:id;primaryKey;autoIncrement"`
	UserName string `gorm:"column:user_name;not null;unique"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email;unique"`
	Phone    string `gorm:"column:phone;unique"`
}

type UsersFilter struct {
	Ids      []uint32
	UserName []string
	Email    []string
	Phone    []string
	Page     uint32
	PageSize uint32
}

func (f *UsersFilter) GetIds() []uint32 {
	return f.Ids
}

type AdminRepo interface {
	CreateUsers(ctx context.Context, tx TX, users []*User) error
	UpdateUsers(ctx context.Context, tx TX, users []*User) error
	DeleteUsers(ctx context.Context, tx TX, ids []uint32) error
	GetUsers(ctx context.Context, tx TX, id uint32) (*User, error)
	ListUsers(ctx context.Context, tx TX, filter *UsersFilter) ([]*User, error)
	Logout(ctx context.Context, id uint32) error
	CountUsers(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
