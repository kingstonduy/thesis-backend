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
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
	"github.com/kingstonduy/product-service/internal/domain"
	redix "github.com/kingstonduy/product-service/internal/pkg/redis_broker"
	utils_transport "github.com/kingstonduy/product-service/internal/pkg/transport"
)

type handler struct {
	repo        domain.IProductRepo
	outboxRepo  domain.IOutboxRepo
	db          *configuration.PostgresCon
	redisPubSub redix.PubSubBroker
}

func NewExecuteTransactionHandler(
	repo domain.IProductRepo,
	outboxRepo domain.IOutboxRepo,
	db *configuration.PostgresCon,
	redisPubSub redix.PubSubBroker,
) domain.IExecuteTransactionHandler {
	return &handler{
		repo:        repo,
		outboxRepo:  outboxRepo,
		db:          db,
		redisPubSub: redisPubSub,
	}
}

// Handle implements domain.IExecuteTransactionHandler.
func (h *handler) Handle(ctx context.Context, cmd *domain.Command[domain.ExecuteTransactionRequest]) (res *domain.ExecuteTransactionResponse, err error) {
	logger.Info(ctx, "Get products handler start")
	defer logger.Info(ctx, "Get products handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()

	req := cmd.Payload

	var product domain.ProductEntity
	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, item := range req.Details {
			product, err = h.repo.GetProductByID(ctx, item.ProductID)
			if err != nil {
				return err
			}

			if product.ProductQuantity == 0 {
				err = fmt.Errorf("Product is out of stock")
				return err
			}

			if product.ProductQuantity < item.CartItemQuantity {
				err = fmt.Errorf("insufficient stock")
				return err
			}

			product.ProductQuantity -= item.CartItemQuantity

			conditions := map[string]interface{}{
				"PRODUCT_ID": product.ProductID,
				"UPDATED_AT": product.UpdatedAt,
			}
			columns := map[string]interface{}{
				"PRODUCT_QUANTITY": product.ProductQuantity,
				"UPDATED_AT":       time.Now(),
			}

			err = h.repo.Update(ctx, columns, conditions)
			if err != nil {
				logger.Error(ctx, err.Error())
				return err
			}
		}

		payloadStr, _ := json.Marshal(req)
		// if update all products successfully
		var outbox domain.OutboxEntity = domain.OutboxEntity{
			AggregateID: cmd.AggregateID,
			CommandID:   uuid.New().String(),
			CommandType: domain.PRODUCT_COMPLETED_TRANSACTION_COMMAND,
			Payload:     string(payloadStr),
			ReplyTo:     cmd.ReplyTo,
		}
		err = h.outboxRepo.Insert(ctx, outbox)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	}, database.WithIsolationLevelOptions(sql.LevelReadCommitted))
	if err1 != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())

		result := utils_transport.GenResultFromErrorx(ctx, errx)
		resultStr, _ := json.Marshal(result)
		err2 := h.redisPubSub.Publish(ctx, redix.RedisMessage{
			Key:     cmd.AggregateID,
			Value:   string(resultStr),
			Channel: cmd.ReplyTo,
		})
		if err2 != nil {
			logger.Error(ctx, err2)
		}
		return nil, errx
	}

	if err != nil {
		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		result := utils_transport.GenResultFromErrorx(ctx, errx)
		resultStr, _ := json.Marshal(result)
		err2 := h.redisPubSub.Publish(ctx, redix.RedisMessage{
			Key:     cmd.AggregateID,
			Value:   string(resultStr),
			Channel: cmd.ReplyTo,
		})
		if err2 != nil {
			logger.Error(ctx, err2)
		}
		return nil, errx
	}

	return res, nil
}
