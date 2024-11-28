package migrations

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	DB *sql.DB
}

type Migration struct {
	db *sql.DB
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
		return fmt.Errorf("migration up failed: %w", err)
	}

	err = sourceDriver.Close()
	if err != nil {
		//
	}

	err = dbDriver.Close()
	if err != nil {
		//
	}

	return nil

}
