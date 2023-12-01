package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresqlDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s "+
		"sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.Username, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, errors.Wrap(err, "error open connection to PostgresDB")
	}

	err = db.Ping()

	if err != nil {
		return nil, errors.Wrap(err, "data base is not active")
	}

	return db, nil
}
