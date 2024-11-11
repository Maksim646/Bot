package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

type Migrator struct {
	databaseURL  string
	migrationDir string
}

type migratorLogger struct {
	l *zap.Logger
}

func (m *migratorLogger) Printf(format string, args ...interface{}) {
	m.l.Info(fmt.Sprintf("migrations: "+format, args...))
}

func (m *migratorLogger) Verbose() bool { return false }

func NewMigrator(databaseURL string, migrationDir string) *Migrator {
	return &Migrator{
		databaseURL:  databaseURL,
		migrationDir: migrationDir,
	}
}

func (p *Migrator) Apply() error {
	m, err := migrate.New("file://"+p.migrationDir, p.databaseURL)
	if err != nil {
		return err
	}

	m.Log = &migratorLogger{zap.L()}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return err
	}

	sourceErr, dbErr := m.Close()
	if sourceErr != nil {
		return sourceErr
	}

	return dbErr
}

// Revert will revert all migrations.
func (p *Migrator) Revert() error {
	m, err := migrate.New("file://"+p.migrationDir, p.databaseURL)
	if err != nil {
		return err
	}

	m.Log = &migratorLogger{zap.L()}

	if err = m.Down(); err != nil {
		return err
	}

	sourceErr, dbErr := m.Close()
	if sourceErr != nil {
		return sourceErr
	}

	return dbErr
}
