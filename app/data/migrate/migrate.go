package migrate

import (
	"database/sql"
	"github.com/logo-user-management/app/assets"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	Up   = "up"
	Down = "down"
)

// package for database migration
// stores migrations to binaries
var migrations = &migrate.PackrMigrationSource{
	Box: assets.Migrations,
}

// Migrates database up
func MigrateUp(db *sql.DB) (int, error) {
	applied, err := migrate.Exec(db, "postgres", migrations, migrate.Up)

	if err != nil {
		return 0, errors.Wrap(err, "failed to apply migrations")
	}

	return applied, nil
}

// Migrates database down
func MigrateDown(db *sql.DB) (int, error) {
	applied, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		return 0, errors.Wrap(err, "failed to apply migrations")
	}
	return applied, nil
}
