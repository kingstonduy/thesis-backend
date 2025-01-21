package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
	redix "github.com/kingstonduy/cart-service/internal/pkg/redis_broker"
	utils_transport "github.com/kingstonduy/cart-service/internal/pkg/transport"
	cmd_pipeline "github.com/kingstonduy/go-core/comman-pipeline"
	"github.com/kingstonduy/go-core/database"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/trace"
	"github.com/kingstonduy/go-core/transport"
	"github.com/kingstonduy/go-core/transport/broker"
)

type handler struct {
	cfg         *configuration.Configuration
	cartRepo    domain.ICartRepo
	outboxRepo  domain.IOutboxRepo
	db          *configuration.PostgresCon
	redisPubSub redix.PubSubBroker
	kafka       broker.Broker
}

func NewExecuteTransactionHandler(
	cfg *configuration.Configuration,
	cartRepo domain.ICartRepo,
	outboxRepo domain.IOutboxRepo,
	db *configuration.PostgresCon,
	redisPubSub redix.PubSubBroker,
	kafka broker.Broker,
) domain.IExecuteTransactionHandler {
	return &handler{
		cfg:         cfg,
		cartRepo:    cartRepo,
		outboxRepo:  outboxRepo,
		db:          db,
		redisPubSub: redisPubSub,
		kafka:       kafka,
	}
}

// Handle1 implements domain.IExecuteTransactionHandler.
func (h *handler) Handle1(ctx context.Context, outbox cmd_pipeline.OutboxWithTrace) (err error) {
	logger.Info(ctx, "Get products handler start")
	defer logger.Info(ctx, "Get products handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
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
			if item.CartItemID == "TEST_CART_FAILED" {
				err = fmt.Errorf("simulate cart failed")
				return err
			}

			err = h.cartRepo.DeleteById(ctx, item.CartItemID)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		outbox := domain.OutboxEntity{
			AggregateID: outbox.AggregateID,
			CommandID:   uuid.New().String(),
			CommandType: domain.CART_COMPLETED_TRANSACTION_COMMAND,
			Payload:     outbox.Payload,
			Trace:       utils_transport.GenRequestTraceString(outbox.Trace, "", ""),
			ReplyTo:     outbox.ReplyTo,
			TraceParent: trace.ExtractTraceparent(ctx),
		}
		err = h.outboxRepo.Insert(ctx, outbox)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	}, database.WithIsolationLevelOptions(sql.LevelReadUncommitted))

	if err != nil || err1 != nil {
		if err == nil {
			err = err1
		}
		var outboxEntity domain.OutboxEntity = domain.OutboxEntity{
			AggregateID: outbox.AggregateID,
			CommandID:   uuid.New().String(),
			CommandType: domain.CART_FAILED_TRANSACTION_COMMAND,
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

	err = h.redisPubSub.Publish(ctx, redix.NewMessage(
		outbox.AggregateID,
		transport.Response[any]{Result: transport.DefaultSuccessResponse.Result},
		outbox.ReplyTo,
	))
	if err != nil {
		errx := errorx.SuspendedErrorWithDetails(err.Error(), "cart-service")
		logger.Error(ctx, errx)
		return errx
	}
	return nil
}

// Handle implements domain.IExecuteTransactionHandler.
func (h *handler) Handle(ctx context.Context, cmd *domain.Command[transport.Request[domain.ExecuteTransactionRequest]]) (res *domain.ExecuteTransactionResponse, err error) {
	logger.Info(ctx, "Get products handler start")
	defer logger.Info(ctx, "Get products handler end")

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC Get products handler %v", r)
			logger.Errorf(ctx, err.Error())
		}
	}()
	req := cmd.Payload.Data

	err1 := h.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, item := range req.Details {
			if item.CartItemID == "TEST_CART_FAILED" {
				err = fmt.Errorf("simulate cart failed")
				return err
			}

			err = h.cartRepo.DeleteById(ctx, item.CartItemID)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		result := transport.Response[any]{
			Result: transport.DefaultSuccessResponse.Result,
			Trace:  utils_transport.GenRequestTrace(cmd.Payload.Trace, "", ""),
		}
		resultStr, _ := json.Marshal(result)
		outbox := domain.OutboxEntity{
			AggregateID: cmd.AggregateID,
			CommandID:   uuid.New().String(),
			CommandType: domain.CART_COMPLETED_TRANSACTION_COMMAND,
			Payload:     string(resultStr),
			ReplyTo:     cmd.ReplyTo,
		}
		err = h.outboxRepo.Insert(ctx, outbox)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	}, database.WithIsolationLevelOptions(sql.LevelReadUncommitted))

	if err != nil || err1 != nil {
		if err == nil {
			err = err1
		}
		reqTypeStr, _ := json.Marshal(cmd.Payload)
		outbox := domain.OutboxEntity{
			AggregateID: cmd.AggregateID,
			CommandID:   uuid.New().String(),
			CommandType: domain.CART_FAILED_TRANSACTION_COMMAND,
			Payload:     string(reqTypeStr),
			ReplyTo:     "",
		}

		err1 = h.outboxRepo.Insert(ctx, outbox)
		if err1 != nil {
			errx := errorx.SuspendedErrorWithDetails(err1.Error(), "")
			logger.Error(ctx, errx.Error())
			return nil, errx
		}

		errx := errorx.FailedWithDetails(err.Error(), "")
		logger.Error(ctx, errx.Error())
		return nil, errx
	}

	err = h.redisPubSub.Publish(ctx, redix.NewMessage(
		cmd.AggregateID,
		transport.Response[any]{Result: transport.DefaultSuccessResponse.Result},
		cmd.ReplyTo,
	))
	if err != nil {
		errx := errorx.SuspendedErrorWithDetails(err.Error(), "cart-service")
		logger.Error(ctx, errx)
		return nil, errx
	}

	return res, nil
}
