package repo

const UserTable = "users"

type User struct {
	Id       uint32 `gorm:"column:id;primaryKey;autoIncrement"`
	UserName string `gorm:"column:user_name;not null;unique"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email;unique"`
	Phone    string `gorm:"column:phone;unique"`
	Token    string `gorm:"column:token;not null;unique"`
}

type UsersFilter struct {
	Ids      []uint32
	UserName []string
	Email    []string
	Phone    []string
	Page     uint32
	PageSize uint32
}
