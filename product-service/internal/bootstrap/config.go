package configuration

import (
	"context"
	"sync"

	"github.com/kingstonduy/go-core/config"
	"github.com/kingstonduy/go-core/config/viper"
	"github.com/kingstonduy/go-core/logger"
)

func GetConfigure() config.Configure {
	cfg, err := viper.NewViperConfig(
		config.WithConfigFile("./resources/dev.env"),
		config.WithTagName("config"),
		config.WithAutomaticEnv(true),
	)
	if err != nil {
		panic(err)
	}
	return cfg
}

var (
	instance *Configuration
	once     sync.Once
)

// singleton
func GetConfigurationInstance() *Configuration {
	once.Do(
		func() {
			c := GetConfigure()
			var configuration Configuration
			if err := c.Unmarshal(&configuration); err != nil {
				panic(err)
			}
			instance = &configuration
			logger.Infof(context.Background(), "%v", instance)
		})
	return instance
}

type Configuration struct {
	BrokerConfig      KafkaBrokerConfig `config:",squash"`
	RedisConfig       RedisConfig       `config:",squash"`
	ServerConfig      ServerConfig      `config:",squash"`
	HealthCheckConfig HealthCheckConfig `config:",squash"`
	LoggerConfig      LoggerConfig      `config:",squash"`
	TraceConfig       TracerConfig      `config:",squash"`
	HttpConfig        HttpConfig        `config:",squash"`
	PostgresConfig    PostgresConfig    `config:",squash"`
}

type PostgresConfig struct {
	Host                  string `config:"POSTGRES_HOST"`
	Port                  int    `config:"POSTGRES_PORT"`
	Username              string `config:"POSTGRES_USER"`
	Password              string `config:"POSTGRES_PASSWORD"`
	Database              string `config:"POSTGRES_DBNAME"`
	IdleConnection        int    `config:"POSTGRES_POOL_IDLE_CONNECTION"`
	MaxConnection         int    `config:"POSTGRES_MAX_POOL_SIZE"`
	MaxLifeIdleConnection int    `config:"POSTGRES_IDLE_TIMEOUT"`  //seconds
	MaxIdleTimeConnection int    `config:"POSTGRES_MAX_LIFE_TIME"` // seconds
	SslMode               string `config:""`
}

type HttpConfig struct {
	BaseUrl string `config:"BASE_URL"`
}

type ServerConfig struct {
	Name         string `config:"SERVER_NAME"`
	AppVersion   string `config:"SERVER_VERSION"`
	HttpPort     int    `config:"SERVER_HTTP_PORT"`
	HttpBasePath string `config:"SERVER_HTTP_BASE_PATH"`
}
type HealthCheckConfig struct {
	GrRunningThreshold    int `config:"HEALTH_CHECK_GR_RUNNING_THRESHOLD"`
	GcMaxPauseThresholdms int `config:"HEALTH_CHECK_GR_RUNNING_THRESHOLD"`
}

type TracerConfig struct {
	ExporterEndpoint string `config:"TRACE_ENDPOINT"`
}

type LoggerConfig struct {
	LogLevel string `config:"LOG_LEVEL"`
}

type KafkaBrokerConfig struct {
	Addresses         string `config:"KAFKA_BROKERS"`
	SASLEnabled       bool   `config:"KAFKA_SASL_ENABLED"`
	SASLUser          string `config:"KAFKA_SASL_USER"`
	SASLPassword      string `config:"KAFKA_SASL_PASSWORD"`
	SASLAlgorithm     string `config:"KAFKA_SASL_ALGORITHM"`
	TLSEnabled        bool   `config:"KAFKA_TLS_ENABLED"`
	TLSSkipVerify     bool   `config:"KAFKA_TLS_SKIP_VERIFY"`
	TLSClientCertFile string `config:"KAFKA_CLIENT_CERT_FILE"`
	TLSClientKeyFile  string `config:"KAFKA_CLIENT_KEY_FILE"`
	TLSCaCertFile     string `config:"KAFKA_CA_CERT_FILE"`
	ConsumerGroup     string `config:"KAFKA_CONSUMERGROUP"`
	HandlerPool       int    `config:"KAFKA_HANDLER_PUBLISHER"`
	ProductCDCTopic   string `config:"PRODUCT_DB_CDC_TOPIC"`
}

type RedisConfig struct {
	Addresses []string `config:"REDIS_ADDRESSES"`
	Username  string   `config:"REDIS_USERNAME"`
	Password  string   `config:"REDIS_PASSWORD"`
}
