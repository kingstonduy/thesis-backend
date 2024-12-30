package configuration

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kingstonduy/go-core/database"
	"github.com/kingstonduy/go-core/database/sqlx"
	// _ "github.com/yugabyte/pgx/v5/stdlib"
)

type PostgresCon struct {
	DB *database.Gdbc
}

func NewYugabyteCon(cfg *Configuration) *PostgresCon {
	c := cfg.PostgresConfig

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)

	db, err := sqlx.NewSqlxGdbc("pgx", dsn,
		database.WithMaxIdleCount(10),
		// database.WithMaxIdleTime(maxConnIdleTime),
		// database.WithMaxLifetime(maxConnLifetime),
		database.WithMaxIdleTime(time.Duration(c.MaxIdleTimeConnection)*time.Hour),
		database.WithMaxLifetime(time.Duration(c.MaxLifeIdleConnection)*time.Hour),
		database.WithMaxOpen(30),
	)

	if err != nil {
		panic(err)
	}

	return &PostgresCon{
		DB: db,
	}
}
