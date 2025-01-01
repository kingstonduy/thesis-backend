package broker_server

import (
	"context"

	"github.com/gammazero/workerpool"
	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport/broker"
	configuration "github.com/kingstonduy/product-service/internal/bootstrap"
)

type BrokerServer struct {
	cfg           *configuration.Configuration
	Broker        broker.Broker
	logger        logger.Logger
	connected     bool
	consumerGroup string
	quit          chan struct{}
	workerpool    *workerpool.WorkerPool
}

func (s *BrokerServer) GetStartOptions() []BrokerServerStartOption {
	return []BrokerServerStartOption{
		WithSubscriptions(),
	}
}

func NewBrokerServer(
	cfg *configuration.Configuration,
	broker broker.Broker,
	logger logger.Logger,
	config *configuration.Configuration,
) *BrokerServer {

	bConfig := config.BrokerConfig

	BrokerServer := BrokerServer{
		cfg:           cfg,
		Broker:        broker,
		logger:        logger,
		consumerGroup: bConfig.ConsumerGroup,
		quit:          make(chan struct{}),
		workerpool:    workerpool.New(100),
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
	logger.Info(ctx, "stopped redis server")
	return nil
}

func (s *BrokerServer) Connected() bool {
	return s.connected
}

type BrokerServerStartOption func(*BrokerServer) error
