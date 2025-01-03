package main

import (
	"context"
	"database/sql"

	pgtypeV4 "github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/replicase/pgcapture"
	"google.golang.org/grpc"
)

// MyTable implements pgcapture.Model interface
// and will be decoded from change that matching the TableName()
type MyTable struct {
	ID     pgtype.Int4    `pg:"id"`       // the pgtype v5 are supported,
	Value1 pgtypeV4.Text  `pg:"value1"`   // the pgtype v4 are also supported,
	Value2 sql.NullString `pg:"value2"`   // the field which implement sql.Scanner interface are also supported,
	Value3 string         `pg:"value3"`   // the golang built-in types are also supported
	Value4 string         `pg:"my_value"` // can use 'pg' tag to specify the name mapping explicitly
}

func (t *MyTable) TableName() (schema, table string) {
	return "public", "my_table"
}

func (t MyTable) MarshalJSON() ([]byte, error) {
	return pgcapture.MarshalJSON(&t) // ignore unchanged TOAST field for pgtype(v4)
}

func main() {
	ctx := context.Background()

	conn, _ := grpc.Dial("127.0.0.1:1000", grpc.WithInsecure())
	defer conn.Close()

	consumer := pgcapture.NewDBLogConsumer(ctx, conn, pgcapture.ConsumerOption{
		// the uri identify which change stream you want.
		// you can implement dblog.SourceResolver to customize gateway behavior based on uri
		URI: "my_subscription_id",
	})
	defer consumer.Stop()

	consumer.Consume(map[pgcapture.Model]pgcapture.ModelHandlerFunc{
		&MyTable{}: func(change pgcapture.Change) error {
			row := change.New.(*MyTable)
			// and then handle the decoded change event

			if row.ID.Valid {
				// handle the changed field
			}

			if row.Value1.Status == pgtypeV4.Undefined {
				// handle the unchanged toast field
			}

			return nil
		},
	})
}
