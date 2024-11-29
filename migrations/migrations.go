package migrations

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/fx"
	"log/slog"
)

type Params struct {
	fx.In

	DB     *sql.DB
	Logger slog.Logger
}

type Migration struct {
	db  *sql.DB
	log *slog.Logger
}

//go:embed schema/*.sql
var migrationFiles embed.FS

func RunMirgations(p Params) error {
	sourceDriver, err := iofs.New(migrationFiles, "schema")
	if err != nil {
		return fmt.Errorf("failed to initialize migrations source driver: %w", err)
	}

	dbDriver, err := postgres.WithInstance(p.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to initialize postgres driver: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return fmt.Errorf("failed to initialize migrations: %w", err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		p.Logger.Error("failed to run migrations: ", "error", err)
		return fmt.Errorf("migration up failed: %w", err)
	}

	err = sourceDriver.Close()
	if err != nil {
		p.Logger.Error("failed to close migrations sourceDriver", "error", err)
		return err
	}

	err = dbDriver.Close()
	if err != nil {
		p.Logger.Error("failed to close migrations dbDriver", "error", err)
		return err
	}

	return nil

}
