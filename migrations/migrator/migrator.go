package migrator

import (
	"embed"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func DoMigrate(fs embed.FS, db string) (err error) {
	src, err := iofs.New(fs, ".")
	if err != nil {
		return errors.Join(errors.New("embed.FS init failed"), err)
	}
	defer func() {
		_ = src.Close()
	}()

	instance, err := migrate.NewWithSourceInstance("iofs", src, db)
	if err != nil {
		return errors.Join(errors.New("db instance init failed"), err)
	}
	defer func() {
		_, _ = instance.Close()
	}()

	if me := instance.Up(); me != nil && !errors.Is(me, migrate.ErrNoChange) {
		return errors.Join(errors.New("migrate-up failed"), me)
	}

	return nil
}
