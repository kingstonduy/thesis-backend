package postgres

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/trace"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
	gensql "github.com/kingstonduy/product-service/internal/pkg/gen_sql"
)

type outboxRepoImpl struct {
	db *configuration.PostgresCon
}

func NewOutboxRepo(db *configuration.PostgresCon) domain.IOutboxRepo {
	return &outboxRepoImpl{
		db: db,
	}
}

// Insert implements domain.IOutboxRepo.
func (repo *outboxRepoImpl) Insert(ctx context.Context, entity domain.WriteOutboxEntity) error {
	logger.Info(ctx, "Insert OUTBOX starts")
	defer logger.Info(ctx, "Insert OUTBOX ends")
	entity.TraceParent = trace.ExtractTraceparent(ctx)

	sqlQuery, err := gensql.GenInsertSql("OUTBOX", entity)
	if err != nil {
		return err
	}

	logger.Info(ctx, sqlQuery)

	res, err := repo.db.DB.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return fmt.Errorf(errorx.ErrorMessageNoRowAffected)
	}

	return nil
}
