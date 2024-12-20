package data

import "appix/internal/biz"

const applicationType = "application"
const applicationTable = "applications"

type Application struct {
	Id           uint32 `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"type:varchar(255);index:idx_app_name_env,unique"`
	Description  string `gorm:"type:varchar(255);"`
	Owner        string `gorm:"type:varchar(255);"`
	IsStateful   bool   `gorm:"type:tinyint(1);"`
	ClusterId    uint32
	DatacenterId uint32
	EnvId        uint32 `gorm:"index:idx_app_name_env,unique"`
	ProductId    uint32
	TeamId       uint32
}

func NewApplication(app *biz.Application) (*Application, error) {
	if app == nil {
		return nil, nil
	}
	return &Application{
		Id:           app.Id,
		Name:         app.Name,
		Description:  app.Description,
		Owner:        app.Owner,
		IsStateful:   app.IsStateful,
		ClusterId:    app.ClusterId,
		DatacenterId: app.DatacenterId,
		EnvId:        app.EnvId,
		ProductId:    app.ProductId,
		TeamId:       app.TeamId,
	}, nil
}

func NewApplications(apps []*biz.Application) ([]*Application, error) {
	db_apps := make([]*Application, len(apps))
	for i, a := range apps {
		db_app, e := NewApplication(a)
		if e != nil {
			return nil, e
		}
		db_apps[i] = db_app
	}
	return db_apps, nil
}

func NewBizApplication(t *Application) (*biz.Application, error) {
	if t == nil {
		return nil, nil
	}
	return &biz.Application{
		Id:           t.Id,
		Name:         t.Name,
		Description:  t.Description,
		Owner:        t.Owner,
		IsStateful:   t.IsStateful,
		ClusterId:    t.ClusterId,
		DatacenterId: t.DatacenterId,
		EnvId:        t.EnvId,
		ProductId:    t.ProductId,
		TeamId:       t.TeamId,
	}, nil
}

func NewBizApplications(es []Application) ([]biz.Application, error) {
	apps := make([]biz.Application, len(es))
	for i, e := range es {
		app, err := NewBizApplication(&e)
		if err != nil {
			return nil, err
		}
		apps[i] = *app
	}
	return apps, nil
}
