package storage

import (
	"database/sql"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(c CompletedConfig, logger *log.Helper) (*gorm.DB, error) {
	var opener func(string) gorm.Dialector
	var db *gorm.DB

	logger.Info("Persistence disabled: ", c.Options.DisablePersistence)

	if c.Options.DisablePersistence {
		logger.Info("Persistence disabled, skipping database connection...")
		// Return nil database connection
		return nil, nil
	}

	switch c.Options.Database {
	case "postgres":
		opener = postgres.Open
	case "sqlite3":
		opener = sqlite.Open
	default:
		return nil, fmt.Errorf("unrecognized database type: %s", c.Options.Database)
	}

	logger.Infof("Using backing storage: %s", c.Options.Database)
	db, err := gorm.Open(opener(c.DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %s", err.Error())
	}

	return db, nil
}

func NewBunDB(c CompletedConfig, logger *log.Helper) (*bun.DB, error) {
	logger.Info("Persistence disabled: ", c.Options.DisablePersistence)

	if c.Options.DisablePersistence {
		logger.Info("Persistence disabled, skipping database connection...")
		return nil, nil
	}

	if c.Options.Database != "postgres" {
		return nil, fmt.Errorf("unrecognized database type: %s", c.Options.Database)
	}

	// TODO: config for SSL and non-SSL based on environment
	pgconn := pgdriver.NewConnector(
		pgdriver.WithAddr(c.Options.Postgres.Host+":"+c.Options.Postgres.Port),
		pgdriver.WithDatabase(c.Options.Postgres.DbName),
		pgdriver.WithUser(c.Options.Postgres.User),
		pgdriver.WithPassword(c.Options.Postgres.Password),
		// pgdriver.WithConnParams(map[string]interface{}{
		// "sslmode":     c.Options.Postgres.SSLMode,
		// "sslrootcert": c.Options.Postgres.SSLRootCert,
		// }),
		pgdriver.WithTLSConfig(nil), // Explicitly needed to bypass pgdriver's SSL validation in nonprod-like envs
	)
	sqldb := sql.OpenDB(pgconn)
	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging DB via bun: %v", err)
	}

	return db, nil
}
