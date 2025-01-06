package server

import (
	appv1 "appix/api/appix/v1"
	"appix/internal/conf"
	"appix/internal/middleware"
	"appix/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	adminConf *conf.Admin,
	tags *service.TagsService,
	features *service.FeaturesService,
	teams *service.TeamsService,
	products *service.ProductsService,
	envs *service.EnvsService,
	clusters *service.ClustersService,
	datacenters *service.DatacentersService,
	hostgroups *service.HostgroupsService,
	applications *service.ApplicationsService,
	adminService *service.AdminService,
	logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.Middleware(middleware.JWTMiddleware(
			middleware.JWTMiddlewareOption{
				Secret:          adminConf.GetJwtSecret(),
				EmergencyHeader: adminConf.GetEmergencyHeader(),
				DefaultSecret:   DefaultSecret,
			},
		),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	appv1.RegisterTagsHTTPServer(srv, tags)
	appv1.RegisterFeaturesHTTPServer(srv, features)
	appv1.RegisterTeamsHTTPServer(srv, teams)
	appv1.RegisterProductsHTTPServer(srv, products)
	appv1.RegisterEnvsHTTPServer(srv, envs)
	appv1.RegisterClustersHTTPServer(srv, clusters)
	appv1.RegisterDatacentersHTTPServer(srv, datacenters)
	appv1.RegisterHostgroupsHTTPServer(srv, hostgroups)
	appv1.RegisterApplicationsHTTPServer(srv, applications)
	appv1.RegisterAdminHTTPServer(srv, adminService)
	return srv
}
