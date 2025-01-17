package sql_util

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

func isEmpty(s string) bool {
	return s == "" || len(s) == 0
}
func SetString(s string) sql.NullString {
	if !isEmpty(s) {
		return sql.NullString{
			String: s,
			Valid:  true,
		}
	} else {
		return sql.NullString{Valid: false}
	}
}

func SetInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

func SetInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

func SetDecimal(d decimal.Decimal) decimal.NullDecimal {
	return decimal.NullDecimal{
		Decimal: d,
		Valid:   true,
	}
}

func Settime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}
