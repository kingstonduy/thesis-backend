package configuration

import (
	"context"
	"sync"
	"time"

	"github.com/kingstonduy/go-core/config"
	"github.com/kingstonduy/go-core/config/viper"
	"github.com/kingstonduy/go-core/logger"
)

func GetConfigure() config.Configure {
	cfg, err := viper.NewViperConfig(
		config.WithConfigFile("./resources/local.env"),
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
	ServerConfig               ServerConfig               `config:",squash"`
	BrokerConfig               KafkaBrokerConfig          `config:",squash"`
	HealthCheckConfig          HealthcCheckConfig         `config:",squash"`
	LoggerConfig               LoggerConfig               `config:",squash"`
	TraceConfig                TracerConfig               `config:",squash"`
	KafKaTopic                 KafKaTopic                 `config:",squash"`
	HttpConfig                 HttpConfig                 `config:",squash"`
	OracleUatsanConfig         OracleUatsanConfig         `config:",squash"`
	YugabyteMcsAssetMgmtConfig YugabyteMcsAssetMgmtConfig `config:",squash"`
	OracleOsbr20Cofig          OracleOsbr20Cofig          `config:",squash"`
}

type YugabyteMcsAssetMgmtConfig struct {
	Host                  string `config:"YUGABYTE_MCS_ASSET_MGMT_HOST"`
	Port                  int    `config:"YUGABYTE_MCS_ASSET_MGMT_PORT"`
	Username              string `config:"YUGABYTE_MCS_ASSET_MGMT_USER"`
	Password              string `config:"YUGABYTE_MCS_ASSET_MGMT_PASSWORD"`
	Database              string `config:"YUGABYTE_MCS_ASSET_MGMT_DBNAME"`
	IdleConnection        int    `config:"YUGABYTE_MCS_ASSET_MGMT_POOL_IDLE_CONNECTION"`
	MaxConnection         int    `config:"YUGABYTE_MCS_ASSET_MGMT_MAX_POOL_SIZE"`
	MaxLifeIdleConnection int    `config:"YUGABYTE_MCS_ASSET_MGMT_IDLE_TIMEOUT"`  //seconds
	MaxIdleTimeConnection int    `config:"YUGABYTE_MCS_ASSET_MGMT_MAX_LIFE_TIME"` // seconds
	SslMode               string `config:""`
}

type OracleOsbr20Cofig struct {
	Host                  string        `config:"ORACLE_OSBR20_DB_HOST"`
	Port                  int           `config:"ORACLE_OSBR20_DB_PORT"`
	Username              string        `config:"ORACLE_OSBR20_DB_USER"`
	Password              string        `config:"ORACLE_OSBR20_DB_PASSWORD"`
	Database              string        `config:"ORACLE_OSBR20_DB_DBNAME"`
	MaxConnection         int           `config:"ORACLE_OSBR20_DB_MAX_POOL_SIZE"`
	IdleConnection        int           `config:"ORACLE_OSBR20_DB_POOL_IDLE_CONNECTION"`
	MaxIdleTimeConnection time.Duration `config:"ORACLE_OSBR20_DB_MAX_IDLE_TIME"` //seconds
	MaxLifeTimeConnection time.Duration `config:"ORACLE_OSBR20_DB_MAX_LIFE_TIME"` // seconds
}

type KafKaTopic struct {
}

type OracleUatsanConfig struct {
	Host                  string        `config:"ORACLE_UATSAN_DB_HOST"`
	Port                  int           `config:"ORACLE_UATSAN_DB_PORT"`
	Username              string        `config:"ORACLE_UATSAN_DB_USER"`
	Password              string        `config:"ORACLE_UATSAN_DB_PASSWORD"`
	Database              string        `config:"ORACLE_UATSAN_DB_DBNAME"`
	MaxConnection         int           `config:"ORACLE_UATSAN_DB_MAX_POOL_SIZE"`
	IdleConnection        int           `config:"ORACLE_UATSAN_DB_POOL_IDLE_CONNECTION"`
	MaxIdleTimeConnection time.Duration `config:"ORACLE_UATSAN_DB_MAX_IDLE_TIME"` //seconds
	MaxLifeTimeConnection time.Duration `config:"ORACLE_UATSAN_DB_MAX_LIFE_TIME"` // seconds
}

type HttpConfig struct {
	BaseUrl    string `config:"NEW_MCS_URL"`
	ExecuteUrl string `config:"FUND_TRANSFER_EXECUTE_URL"`
}

type ServerConfig struct {
	Name         string `config:"SERVER_NAME"`
	AppVersion   string `config:"SERVER_VERSION"`
	HttpPort     int    `config:"SERVER_HTTP_PORT"`
	HttpBasePath string `config:"SERVER_HTTP_BASE_PATH"`
}

type KafkaBrokerConfig struct {
	Addresses                  string `config:"KAFKA_BROKERS"`
	SASLEnabled                bool   `config:"KAFKA_SASL_ENABLED"`
	SASLUser                   string `config:"KAFKA_SASL_USER"`
	SASLPassword               string `config:"KAFKA_SASL_PASSWORD"`
	SASLAlgorithm              string `config:"KAFKA_SASL_ALGORITHM"`
	TLSEnabled                 bool   `config:"KAFKA_TLS_ENABLED"`
	TLSSkipVerify              bool   `config:"KAFKA_TLS_SKIP_VERIFY"`
	TLSClientCertFile          string `config:"KAFKA_CLIENT_CERT_FILE"`
	TLSClientKeyFile           string `config:"KAFKA_CLIENT_KEY_FILE"`
	TLSCaCertFile              string `config:"KAFKA_CA_CERT_FILE"`
	MessageTimeout             int    `config:"KAFKA_MESSAGE_TIMEOUT"`
	ConsumerGroup              string `config:"KAFKA_CONSUMERGROUP"`
	RollbackErrorConsumerGroup string `config:"KAFKA_ROLLBACK_ERROR_CONSUMER_GROUP"`
	HandlerPool                int    `config:"KAFKA_HANDLER_PUBLISHER"`
}

type HealthcCheckConfig struct {
	GrRunningThreshold    int `config:"HEALTH_CHECK_GR_RUNNING_THRESHOLD"`
	GcMaxPauseThresholdms int `config:"HEALTH_CHECK_GR_RUNNING_THRESHOLD"`
}

type LoggerConfig struct {
	LogLevel string `config:"LOG_LEVEL"`
}

type TracerConfig struct {
	ExporterEndpoint string `config:"TRACE_ENDPOINT"`
}
