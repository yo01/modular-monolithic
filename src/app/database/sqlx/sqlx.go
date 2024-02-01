package sqlx

import (
	"modular-monolithic/config"

	"git.motiolabs.com/library/motiolibs/msqlx"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func InitPostgreConnection(config *config.Config) *sqlx.DB {
	sqlxConfig := msqlx.SqlxConfig{
		Host:     config.PgHostname,
		Username: config.PgUsername,
		Password: config.PgPassword,
		DBName:   config.PgDatabase,
		SSLMode:  config.PgSslMode,
		TimeZone: config.PgTimezone,
		Port:     config.PgPort,
	}

	db, err := msqlx.InitPostgreConnection(sqlxConfig)
	if err.Error != nil {
		panic(err.Error)
	}

	return db
}
