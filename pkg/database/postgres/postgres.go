package postgres

import (
	"deep-ocean-api/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewClient(cfg *config.Config) (*sqlx.DB, error) {
	source := getConnectionString(cfg)

	db, err := sqlx.Connect("postgres", source)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getConnectionString(cfg *config.Config) string {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s",
		cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseUser, cfg.DatabaseName, cfg.DatabaseSSLMode)

	return connectionString
}
