package gensql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestGenSql(t *testing.T) {
	checkUpdate()
	checkInset()
}

func checkInset() {
	// Example usage
	type t struct {
		StrVal         string              `json:"id" db:"StrVal"`
		StrPtr         *string             `db:"StrPtr"`
		NullPtr        *string             `db:"NullPtr"`
		ValidNullStr   sql.NullString      `db:"ValidNullStr"`
		InvalidNullStr sql.NullString      `db:"InvalidNullStr"`
		ValidNullInt   sql.NullInt64       `db:"ValidNullInt64"`
		InvalidNullInt sql.NullInt64       `db:"InvalidNullInt"`
		ValidNullDec   decimal.NullDecimal `db:"ValidNullDec"`
		InvalidNullDec decimal.NullDecimal `db:"InvalidNullDec"`
	}

	strPtr := "string pointer"
	tr := t{
		StrVal:         "stringVal",
		StrPtr:         &strPtr,
		NullPtr:        nil,
		ValidNullStr:   sql.NullString{String: "nullString", Valid: true},
		InvalidNullStr: sql.NullString{String: "nullString", Valid: false},
		ValidNullInt:   sql.NullInt64{Int64: 1, Valid: true},
		InvalidNullInt: sql.NullInt64{Int64: 1, Valid: false},
		ValidNullDec:   decimal.NullDecimal{Decimal: decimal.NewFromInt(123), Valid: true},
		InvalidNullDec: decimal.NullDecimal{Decimal: decimal.NewFromInt(123), Valid: false},
	}

	s, _ := GenInsertSql("TABLE_NAME", tr)
	fmt.Println(s)

	s, _ = GenInsertSql("TABLE_NAME", &tr)
	fmt.Println(s)
}

func checkUpdate() {
	// Example usage
	type t struct {
		StrVal         string              `json:"id" db:"StrVal"`
		StrPtr         *string             `db:"StrPtr"`
		NullPtr        *string             `db:"NullPtr"`
		ValidNullStr   sql.NullString      `db:"ValidNullStr"`
		InvalidNullStr sql.NullString      `db:"InvalidNullStr"`
		ValidNullInt   sql.NullInt64       `db:"ValidNullInt64"`
		InvalidNullInt sql.NullInt64       `db:"InvalidNullInt"`
		ValidNullDec   decimal.NullDecimal `db:"ValidNullDec"`
		InvalidNullDec decimal.NullDecimal `db:"InvalidNullDec"`
	}

	strPtr := "string pointer"
	tr := t{
		StrVal:         "stringVal",
		StrPtr:         &strPtr,
		NullPtr:        nil,
		ValidNullStr:   sql.NullString{String: "nullString", Valid: true},
		InvalidNullStr: sql.NullString{String: "nullString", Valid: false},
		ValidNullInt:   sql.NullInt64{Int64: 1, Valid: true},
		InvalidNullInt: sql.NullInt64{Int64: 1, Valid: false},
		ValidNullDec:   decimal.NullDecimal{Decimal: decimal.NewFromInt(123), Valid: true},
		InvalidNullDec: decimal.NullDecimal{Decimal: decimal.NewFromInt(123), Valid: false},
	}

	conditions := map[string]interface{}{
		"id": tr.StrVal,
	}
	columns := map[string]interface{}{
		"StrVal":         tr.StrVal,
		"StrPtr":         tr.StrPtr,
		"NullPtr":        tr.NullPtr,
		"ValidNullStr":   tr.ValidNullStr,
		"InvalidNullStr": tr.InvalidNullStr,
		"ValidNullInt":   tr.ValidNullInt,
		"InvalidNullInt": tr.InvalidNullInt,
		"ValidNullDec":   tr.ValidNullDec,
		"InvalidNullDec": tr.InvalidNullDec,
	}

	s, _ := GenUpdateSql("TABLE_NAME", columns, conditions)
	fmt.Println(s)

	s, _ = GenUpdateSql("TABLE_NAME", columns, conditions)
	fmt.Println(s)
}
