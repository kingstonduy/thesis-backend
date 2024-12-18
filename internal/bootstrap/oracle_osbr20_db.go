package configuration

import (
	"github.com/kingstonduy/go-core/database"
	"github.com/kingstonduy/go-core/database/sqlx"
	go_ora "github.com/sijms/go-ora/v2"
)

type OracleOsbr20Con struct {
	DB *database.Gdbc
}

func GetOracleOsbr20Con(cfg *Configuration) *OracleOsbr20Con {
	e := cfg.OracleOsbr20Cofig
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

	return &OracleOsbr20Con{
		DB: db,
	}
}
