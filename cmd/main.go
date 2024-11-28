package main

import (
	"context"
	"employees/internal/pkg/config"
	"employees/internal/pkg/db"
	"employees/internal/pkg/employee"
	employeeHttp "employees/internal/pkg/employee/delivery/http"
	"employees/internal/pkg/employee/repo"
	"employees/internal/pkg/employee/usecase"
	"employees/internal/pkg/logger"
	"employees/internal/pkg/server"
	"employees/migrations"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// @title Swagger Employees service
// @version 1.0
// @description This is a sample server employees service.
// @host localhost:8080
// @schemes http
// @BasePath /api/v1
func main() {
	app := fx.New(
		// конструкторы
		fx.Provide(
			logger.SetupLogger,
			server.NewRouter,

			config.MustLoad,

			db.NewPostgresConn,
			db.NewPostgresPool,

			employeeHttp.New,
			fx.Annotate(usecase.New, fx.As(new(employee.Usecase))),
			fx.Annotate(repo.New, fx.As(new(employee.Repository))),
		),

		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: logger}
		}),

		fx.Invoke(
			server.RunServer,
			migrations.RunMirgations,
		),
	)

	ctx := context.Background()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	if err := app.Start(ctx); err != nil {
		panic(err)
	}

	<-stop
	app.Stop(ctx)
}
