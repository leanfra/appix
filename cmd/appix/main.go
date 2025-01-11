package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"appix/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	Branch  string
	// flagconf is the config flag.
	flagconf    string
	showVersion bool

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.BoolVar(&showVersion, "version", false, "show version")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Printf("name: %s\nbranch: %s\nversion: %s\n", Name, Branch, Version)
		os.Exit(0)
	}

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	if bc.Server.Tracer != "" {
		logger = log.With(logger,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		)
	}

	if err := validateAdminConfig(bc.Admin); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Admin, bc.Authz, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func validateAdminConfig(conf *conf.Admin) error {
	if conf == nil {
		return errors.New("admin config is nil")
	}
	if conf.AdminUser == "" || conf.AdminPassword == "" {
		return errors.New("admin user or password is empty")
	}
	if conf.JwtSecret == "" {
		return errors.New("jwt secret is empty")
	}
	if conf.JwtExpireHours == 0 {
		return errors.New("jwt expire hours is empty")
	}
	if conf.EmergencyHeader == "" {
		return errors.New("emergency header is empty")
	}
	return nil
}
