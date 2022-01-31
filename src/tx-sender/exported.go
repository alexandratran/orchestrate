package txsender

import (
	"context"

	"github.com/consensys/orchestrate/src/infra/redis"
	"github.com/consensys/orchestrate/src/infra/redis/redigo"

	sarama2 "github.com/Shopify/sarama"
	orchestrateClient "github.com/consensys/orchestrate/pkg/sdk/client"
	"github.com/consensys/orchestrate/pkg/toolkit/app"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	ethclient "github.com/consensys/orchestrate/src/infra/ethclient/rpc"

	"github.com/consensys/orchestrate/src/infra/broker/sarama"
	qkm "github.com/consensys/orchestrate/src/infra/quorum-key-manager"
	"github.com/spf13/viper"
)

// New Utility function used to initialize a new service
func New(ctx context.Context) (*app.App, error) {
	logger := log.FromContext(ctx)
	config := NewConfig(viper.GetViper())
	var redisClient redis.Client
	var err error

	sarama.InitSyncProducer(ctx)
	qkm.Init()
	orchestrateClient.Init()
	ethclient.Init(ctx)

	if config.NonceManagerType == NonceManagerTypeRedis {
		redisClient, err = redigo.New(config.RedisCfg)
		if err != nil {
			return nil, err
		}
	}

	consumerGroups := make([]sarama2.ConsumerGroup, config.NConsumer)
	hostnames := viper.GetStringSlice(sarama.KafkaURLViperKey)
	for idx := 0; idx < config.NConsumer; idx++ {
		consumerGroups[idx], err = NewSaramaConsumer(hostnames, config.GroupName)
		if err != nil {
			return nil, err
		}
		logger.WithField("host", hostnames).WithField("group_name", config.GroupName).
			Info("consumer client ready")
	}

	return NewTxSender(
		config,
		consumerGroups,
		sarama.GlobalSyncProducer(),
		qkm.GlobalClient(),
		orchestrateClient.GlobalClient(),
		ethclient.GlobalClient(),
		redisClient,
	)
}

func NewSaramaConsumer(hostnames []string, groupName string) (sarama2.ConsumerGroup, error) {
	config, err := sarama.NewSaramaConfig()
	if err != nil {
		return nil, err
	}

	client, err := sarama.NewClient(hostnames, config)
	if err != nil {
		return nil, err
	}

	return sarama.NewConsumerGroupFromClient(groupName, client)
}
