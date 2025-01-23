package repo

import (
	"context"
)

const ApplicationTable = "applications"

type Application struct {
	ChangeInfo
	Id          uint32 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_app_name_env,unique"`
	Description string `gorm:"type:varchar(255);"`
	OwnerId     uint32 `gorm:"column:owner_id"`
	IsStateful  bool   `gorm:"type:tinyint(1);"`
	ProductId   uint32
	TeamId      uint32
}

type ApplicationsFilter struct {
	Page       uint32
	PageSize   uint32
	Ids        []uint32
	Names      []string
	IsStateful string
	ProductsId []uint32
	TeamsId    []uint32
}

const IsStatefulTrue = "true"
const IsStatefulFalse = "false"
const IsStatefulNone = ""

type ApplicationsRepo interface {
	RequireCounter
	CreateApplications(ctx context.Context, tx TX, apps []*Application) error
	UpdateApplications(ctx context.Context, tx TX, apps []*Application) error
	DeleteApplications(ctx context.Context, tx TX, ids []uint32) error
	GetApplications(ctx context.Context, id uint32) (*Application, error)
	ListApplications(ctx context.Context, tx TX, filter *ApplicationsFilter) ([]*Application, error)
	//CountApplications(ctx context.Context, tx TX, filter *ApplicationsFilter) (int64, error)
}
