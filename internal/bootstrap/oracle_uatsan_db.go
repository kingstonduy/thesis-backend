package configuration

import (
	"github.com/kingstonduy/go-core/database"
	"github.com/kingstonduy/go-core/database/sqlx"
	go_ora "github.com/sijms/go-ora/v2"
)

type OracleUatsanCon struct {
	DB *database.Gdbc
}

func GetOracleUatsanCon(cfg *Configuration) *OracleUatsanCon {
	e := cfg.OracleUatsanConfig
	connStr := go_ora.BuildUrl(e.Host, e.Port, e.Database, e.Username, e.Password, nil)

	db, err := sqlx.NewSqlxGdbc("oracle", connStr,
		database.WithMaxIdleCount(e.IdleConnection),
		database.WithMaxOpen(e.MaxConnection),
		database.WithMaxIdleTime(e.MaxIdleTimeConnection),
		database.WithMaxLifetime(e.MaxLifeTimeConnection),
	)

	if err != nil {
		panic(err)
	}

	return &OracleUatsanCon{
		DB: db,
	}
}
