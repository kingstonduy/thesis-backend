package broker_server

import (
	"context"

	"github.com/gammazero/workerpool"
	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
	"github.com/kingstonduy/go-core/cache"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport/broker"
)

type BrokerServer struct {
	redisClient   cache.CacheClient
	cfg           *configuration.Configuration
	Broker        broker.Broker
	logger        logger.Logger
	connected     bool
	consumerGroup string
	quit          chan struct{}
	workerpool    *workerpool.WorkerPool
	productRepo   domain.IProductRepo
}

func (s *BrokerServer) GetStartOptions() []BrokerServerStartOption {
	return []BrokerServerStartOption{
		WithSubscriptions(),
	}
}

func NewBrokerServer(
	redisClient cache.CacheClient,
	cfg *configuration.Configuration,
	broker broker.Broker,
	logger logger.Logger,
	config *configuration.Configuration,
	productRepo domain.IProductRepo,
) *BrokerServer {

	bConfig := config.BrokerConfig

	BrokerServer := BrokerServer{
		redisClient:   redisClient,
		cfg:           cfg,
		Broker:        broker,
		logger:        logger,
		consumerGroup: bConfig.ConsumerGroup,
		quit:          make(chan struct{}),
		workerpool:    workerpool.New(100),
		productRepo:   productRepo,
	}

	go func() {
		defer func() {
			ctx := context.TODO()

			if err := BrokerServer.Broker.Disconnect(); err != nil {
				logger.Errorf(ctx, "Failed to shutdown broker server: %v", err)
			}

			BrokerServer.connected = false
			BrokerServer.workerpool.Stop()

			logger.Info(ctx, "Stopped Broker Server")
		}()

		// waiting util channel has event
		<-BrokerServer.quit
	}()

	return &BrokerServer
}

func (s *BrokerServer) Start(ctx context.Context) error {
	opts := s.GetStartOptions()
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return err
		}
	}

	if err := s.Broker.Connect(); err != nil {
		s.logger.Error(ctx, "Failted to connect to kafka broker")
		return err
	} else {
		s.logger.Info(ctx, "Connected to kafka broker")
	}

	s.connected = true
	s.logger.Info(ctx, "Started broker server")

	return nil
}

func (s *BrokerServer) Stop(ctx context.Context) error {
	logger.Info(ctx, "stopped broker server")
	return nil
}

func (s *BrokerServer) Connected() bool {
	return s.connected
}

type BrokerServerStartOption func(*BrokerServer) error
