package sql_util

import (
	"database/sql"
	"time"
)

func SetString(s string) sql.NullString {

	return sql.NullString{String: s, Valid: true}
}

func SetTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

func SetInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}
