package configuration

import (
	"context"
	"strings"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/transport/broker"
	"github.com/kingstonduy/go-core/transport/broker/kafka"
)

func GetKafkaBroker(cfg *Configuration, logger logger.Logger) broker.Broker {
	var (
		bConfig = cfg.BrokerConfig
	)

	var config = &kafka.KafkaBrokerConfig{
		Addresses:         strings.Split(bConfig.Addresses, ","),
		TLSEnabled:        bConfig.TLSEnabled,
		TLSSkipVerify:     bConfig.TLSSkipVerify,
		TLSCaCertFile:     bConfig.TLSCaCertFile,
		TLSClientCertFile: bConfig.TLSClientCertFile,
		TLSClientKeyFile:  bConfig.TLSClientKeyFile,
		SASLEnabled:       bConfig.SASLEnabled,
		SASLAlgorithm:     bConfig.SASLAlgorithm,
		SASLUser:          bConfig.SASLUser,
		SASLPassword:      bConfig.SASLPassword,
	}

	br, err := kafka.GetKafkaBroker(
		config,
		broker.WithLogger(logger),
	)

	ctx := context.Background()

	if err != nil {
		logger.Error(ctx, "Failted to create kafka broker")
		panic(err)
	}

	if err := br.Connect(); err != nil {
		logger.Error(ctx, "Failted to connect to kafka broker")
		panic(err)
	} else {
		logger.Info(ctx, "Connected to kafka broker")
	}

	return br
}
