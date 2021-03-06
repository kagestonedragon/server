package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/kagestonedragon/server/configs"
	"github.com/kagestonedragon/server/internal/db/postgres"
	e "github.com/kagestonedragon/server/internal/repository/echo"
	"github.com/kagestonedragon/server/internal/server"
	u "github.com/kagestonedragon/server/internal/user"
	"github.com/kagestonedragon/server/pkg/echo"
	"github.com/kagestonedragon/server/pkg/health"
	"github.com/kagestonedragon/server/pkg/user"
	"github.com/kagestonedragon/server/tools/logging"
	"github.com/kagestonedragon/server/tools/metrics"
	"github.com/kagestonedragon/server/tools/sentry"
	"github.com/kagestonedragon/server/tools/tracing"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Load config
	cfg := configs.NewConfig()
	if err := cfg.Read(); err != nil {
		fmt.Fprintf(os.Stderr, "read config: %s", err)
		os.Exit(1)
	}
	// Print config
	if err := cfg.Print(); err != nil {
		fmt.Fprintf(os.Stderr, "read config: %s", err)
		os.Exit(1)
	}

	logger, err := logging.NewLogger(cfg.Logger.Level, cfg.Logger.TimeFormat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init logger: %s", err)
		os.Exit(1)
	}
	ctx = logging.WithContext(ctx, logger)

	if cfg.Tracer.Enabled {
		tracer, closer, err := tracing.NewJaegerTracer(
			ctx,
			fmt.Sprintf("%s:%d", cfg.Tracer.Host, cfg.Tracer.Port),
			cfg.Tracer.Name,
		)
		if err != nil {
			level.Error(logger).Log("err", err, "msg", "failed to init tracer")
		}
		defer closer.Close()
		ctx = tracing.WithContext(ctx, tracer)
	}
	if cfg.Sentry.Enabled {
		if err := sentry.NewSentry(cfg); err != nil {
			level.Error(logger).Log("err", err, "msg", "failed to init sentry")
		}
	}

	if cfg.Metrics.Enabled {
		ctx = metrics.WithContext(ctx)
	}

	conn, err := postgres.NewConnection(ctx, cfg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	echoRepository := e.NewEcho()
	userRepository, err := u.NewCacheableRepository(
		ctx,
		u.NewPostgreSqlRepository(conn),
		time.Duration(cfg.Cache.Lifetime)*time.Second,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	{
		healthService := initHealthService(ctx, cfg)
		echoService := initEchoService(ctx, cfg, echoRepository)
		userService := initUserService(ctx, cfg, userRepository)
		s, err := server.NewServer(
			server.SetConfig(cfg),
			server.SetLogger(logger),
			server.SetHandler(
				map[string]http.Handler{
					"":     health.MakeHTTPHandler(ctx, healthService),
					"user": user.MakeHTTPHandler(ctx, userService),
				}),
			server.SetGRPC(
				health.JoinGRPC(ctx, healthService),
				echo.JoinGRPC(ctx, echoService),
				user.JoinGRPC(ctx, userService),
			),
		)
		if err != nil {
			level.Error(logger).Log("init", "server", "err", err)
			os.Exit(1)
		}
		defer s.Close()

		if err := s.AddHTTP(); err != nil {
			level.Error(logger).Log("err", err)
			os.Exit(1)
		}

		if err = s.AddGRPC(); err != nil {
			level.Error(logger).Log("err", err)
			os.Exit(1)
		}

		if err = s.AddMetrics(); err != nil {
			level.Error(logger).Log("err", err)
			os.Exit(1)
		}

		s.AddSignalHandler()
		s.Run()
	}
}

func initHealthService(ctx context.Context, cfg *configs.Config) health.Service {
	healthService := health.NewHealthService()
	if cfg.Metrics.Enabled {
		healthService = health.NewMetricsService(ctx, healthService)
	}
	healthService = health.NewLoggingService(ctx, healthService)

	if cfg.Sentry.Enabled {
		healthService = health.NewSentryService(healthService)
	}
	return healthService
}

func initEchoService(_ context.Context, cfg *configs.Config, repo e.Echo) echo.Service {
	echoService := echo.NewEchoService(repo)
	if cfg.Sentry.Enabled {
		echoService = echo.NewSentryService(echoService)
	}
	return echoService
}

func initUserService(_ context.Context, _ *configs.Config, r u.CacheableRepository) user.Service {
	return user.NewUserService(r)
}
