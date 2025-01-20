package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kingstonduy/go-core/database"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
	utils_transport "github.com/kingstonduy/product-service/internal/pkg/transport"
)

type handler struct {
	inventoryRepo domain.IInventoryRepo
	outboxRepo    domain.IOutboxRepo
	db            *configuration.PostgresCon
}

func NewRevertTransactionHandler(
	inventoryRepo domain.IInventoryRepo,
	outboxRepo domain.IOutboxRepo,
	db *configuration.PostgresCon,
) domain.IRevertTransactionHandler {
	return &handler{
		inventoryRepo: inventoryRepo,
		outboxRepo:    outboxRepo,
		db:            db,
	}
}

// Handle implements domain.IRevertTransactionHandler.
func (h *handler) Handle(ctx context.Context, cmd domain.Command[transport.Request[domain.RevertTransactionRequest]]) (res *domain.RevertTransactionResponse, err error) {
	logger.Info(ctx, "RevertTransaction handler start")
	defer logger.Info(ctx, "RevertTransaction handler end")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC RevertTransaction handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	req := cmd.Payload.Data

	reqType := transport.Request[domain.RevertTransactionRequest]{
		Trace: utils_transport.GenRequestTrace(cmd.Payload.Trace, "", cmd.ReplyTo),
		Data:  req,
	}
	reqTypeStr, _ := json.Marshal(reqType)

	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, item := range req.Details {
			inventory, err := h.inventoryRepo.SelectByProductID(ctx, item.ProductID)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			inventory.InventoryQuantity += item.CartItemQuantity

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

		// if update all products successfully
		var outbox domain.OutboxEntity = domain.OutboxEntity{
			AggregateID: cmd.AggregateID,
			CommandID:   uuid.New().String(),
			CommandType: domain.PRODUCT_FAILED_TRANSACTION_COMMAND,
			Payload:     string(reqTypeStr),
			ReplyTo:     cmd.ReplyTo,
		}

		err = h.outboxRepo.Insert(ctx, outbox)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	}, database.WithIsolationLevelOptions(sql.LevelRepeatableRead))

	if err != nil || err1 != nil {
		// co the persist vao db transaction_revert || kafka
		if err == nil {
			err = err1
		}
		errx := errorx.SuspendedErrorWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
	}

	return nil, nil
}
