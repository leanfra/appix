package server

import (
	apiv1 "appix/api/appix/v1"
	v1 "appix/api/helloworld/v1"
	"appix/internal/conf"
	"appix/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server,
	greeter *service.GreeterService,
	tags *service.TagsService,
	features *service.FeaturesService,
	teams *service.TeamsService,
	products *service.ProductsService,
	envs *service.EnvsService,
	clusters *service.ClustersService,
	datacenters *service.DatacentersService,
	hostgroups *service.HostgroupsService,
	logger log.Logger) *grpc.Server {

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	apiv1.RegisterTagsServer(srv, tags)
	apiv1.RegisterFeaturesServer(srv, features)
	apiv1.RegisterTeamsServer(srv, teams)
	apiv1.RegisterProductsServer(srv, products)
	apiv1.RegisterEnvsServer(srv, envs)
	apiv1.RegisterClustersServer(srv, clusters)
	apiv1.RegisterDatacentersServer(srv, datacenters)
	apiv1.RegisterHostgroupsServer(srv, hostgroups)
	return srv
}
