package data

import (
	"appix/internal/data/sqldb"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

type greeterRepo struct {
	data *sqldb.DataGorm
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *sqldb.DataGorm, logger log.Logger) GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *Greeter) (*Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *Greeter) (*Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*Greeter, error) {
	return nil, nil
}
