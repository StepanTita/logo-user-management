package config

import (
	"database/sql"
	"github.com/logo-user-management/app/data/migrate"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

type Databaser interface {
	DB() *sql.DB
}

type databaser struct {
	url    string
	method string

	cache struct {
		db *sql.DB
	}

	log *logrus.Logger
	sync.Once
}

func NewDatabaser(url, method string, log *logrus.Logger) Databaser {
	return &databaser{
		url:    url,
		method: method,
		log:    log,
	}
}

func (d *databaser) DB() *sql.DB {
	d.Once.Do(func() {
		var err error
		d.cache.db, err = sql.Open("postgres", d.url)
		if err != nil {
			panic(err)
		}

		switch d.method {
		case migrate.Up:
			applied, err := migrate.MigrateUp(d.cache.db)
			if err != nil {
				panic(err)
			}
			d.log.WithField("applied", applied).Info("Migrations up applied")
		case migrate.Down:
			applied, err := migrate.MigrateDown(d.cache.db)
			if err != nil {
				panic(err)
			}
			d.log.WithField("applied", applied).Info("Migrations down applied")
		default:
			panic("Unknown migration method")
		}

		if err := d.cache.db.Ping(); err != nil {
			panic(errors.Wrap(err, "database unavailable"))
		}
	})
	return d.cache.db
}
