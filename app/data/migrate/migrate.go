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

var migrations = &migrate.PackrMigrationSource{
	Box: assets.Migrations,
}

func MigrateUp(db *sql.DB) (int, error) {
	applied, err := migrate.Exec(db, "postgres", migrations, migrate.Up)

	if err != nil {
		return 0, errors.Wrap(err, "failed to apply migrations")
	}


	return applied, nil
}

func MigrateDown(db *sql.DB) (int, error) {
	applied, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		return 0, errors.Wrap(err, "failed to apply migrations")
	}
	return applied, nil
}
