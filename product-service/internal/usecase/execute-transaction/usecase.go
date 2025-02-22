package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	cmd_pipeline "github.com/kingstonduy/go-core/comman-pipeline"
	"github.com/kingstonduy/go-core/database"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/trace"
	"github.com/kingstonduy/go-core/transport"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
	redix "github.com/kingstonduy/product-service/internal/pkg/redis_broker"
	utils_transport "github.com/kingstonduy/product-service/internal/pkg/transport"
)

type handler struct {
	inventoryRepo domain.IWriteInventoryRepo
	outboxRepo    domain.IOutboxRepo
	db            *configuration.PostgresCon
	redisPubSub   redix.PubSubBroker
}

func NewExecuteTransactionHandler(
	inventoryRepo domain.IWriteInventoryRepo,
	outboxRepo domain.IOutboxRepo,
	db *configuration.PostgresCon,
	redisPubSub redix.PubSubBroker,
) domain.IExecuteTransactionHandler {
	return &handler{
		inventoryRepo: inventoryRepo,
		outboxRepo:    outboxRepo,
		db:            db,
		redisPubSub:   redisPubSub,
	}
}

// Handle1 implements domain.IExecuteTransactionHandler.
func (h *handler) Handle1(ctx context.Context, outbox cmd_pipeline.OutboxWithTrace) (err error) {
	logger.Info(ctx, "ExecuteTransaction handler start")
	defer logger.Info(ctx, "ExecuteTransaction handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC ExecuteTransaction handler %v", r)
			logger.Errorf(ctx, err.Error())
		}

		// cast err to errorx
		errx, ok := err.(*errorx.Error)
		if !ok {
			errx = errorx.InternalServerErrorWithDetails(err.Error(), "")
		}

		resType := transport.Response[any]{
			Trace:  utils_transport.GenResponseTrace(outbox.Trace),
			Result: *utils_transport.GenResultFromErrorx(ctx, errx),
		}
		h.redisPubSub.Publish(ctx, redix.NewMessage(outbox.AggregateID, resType, outbox.ReplyTo))
	}()

	var req domain.ExecuteTransactionRequest
	json.Unmarshal([]byte(outbox.Payload), &req)

	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, item := range req.Details {
			inventory, err := h.inventoryRepo.SelectByProductID(ctx, item.ProductID)
			if err != nil {
				return err
			}

			if inventory.InventoryQuantity == 0 {
				err = fmt.Errorf("Product is out of stock")
				return err
			}

			if inventory.InventoryQuantity < item.CartItemQuantity {
				err = fmt.Errorf("insufficient stock")
				return err
			}

			inventory.InventoryQuantity -= item.CartItemQuantity

			conditions := map[string]interface{}{
				"PRODUCT_ID": inventory.ProductID,
				"UPDATED_AT": inventory.UpdatedAt,
			}
			columns := map[string]interface{}{
				"INVENTORY_QUANTITY": inventory.InventoryQuantity,
				"UPDATED_AT":         time.Now(),
			}

			err = h.inventoryRepo.Update(ctx, columns, conditions)
			if err != nil {
				logger.Error(ctx, err.Error())
				return err
			}
		}
		var outboxEntity domain.WriteOutboxEntity = domain.WriteOutboxEntity{
			AggregateID: outbox.AggregateID,
			CommandID:   uuid.New().String(),
			CommandType: domain.PRODUCT_COMPLETED_TRANSACTION_COMMAND,
			Payload:     outbox.Payload,
			Trace:       utils_transport.GenRequestTraceString(outbox.Trace, "", ""),
			ReplyTo:     outbox.ReplyTo,
			TraceParent: trace.ExtractTraceparent(ctx),
		}

		err = h.outboxRepo.Insert(ctx, outboxEntity)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	}, database.WithIsolationLevelOptions(sql.LevelRepeatableRead))
	if err != nil || err1 != nil {
		if err == nil {
			err = err1
		}

		var outboxEntity domain.WriteOutboxEntity = domain.WriteOutboxEntity{
			AggregateID: outbox.AggregateID,
			CommandID:   uuid.New().String(),
			CommandType: domain.PRODUCT_FAILED_TRANSACTION_COMMAND,
			Payload:     outbox.Payload,
			Trace:       utils_transport.GenRequestTraceString(outbox.Trace, "", ""),
			ReplyTo:     outbox.ReplyTo,
			TraceParent: trace.ExtractTraceparent(ctx),
		}

		err1 = h.outboxRepo.Insert(ctx, outboxEntity)
		if err1 != nil {
			errx := errorx.SuspendedErrorWithDetails(err1.Error(), "")
			logger.Error(ctx, errx.Error())
			return errx
		}

		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return errx
	}

	return nil
}
