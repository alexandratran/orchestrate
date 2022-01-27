package kafka

import (
	"context"
	"sync"

	txscheduler "github.com/consensys/orchestrate/pkg/sdk/client"

	"github.com/consensys/orchestrate/pkg/utils"
	broker "github.com/consensys/orchestrate/src/infra/broker/sarama"
	ethclient "github.com/consensys/orchestrate/src/infra/ethclient/rpc"
)

var (
	hook     *Hook
	initOnce = &sync.Once{}
)

func initComponent(ctx context.Context) {
	utils.InParallel(
		// Initialize Ethereum Client
		func() { ethclient.Init(ctx) },
		// Initialize Sync Producer
		func() { broker.InitSyncProducer(ctx) },
		// Initialize transaction scheduler client
		func() { txscheduler.Init() },
	)
}

// Init Kafka hook
func Init(ctx context.Context) {
	initOnce.Do(func() {
		initComponent(ctx)

		hook = NewHook(
			NewConfig(),
			ethclient.GlobalClient(),
			broker.GlobalSyncProducer(),
			txscheduler.GlobalClient(),
		)
	})
}

// SetGlobalHook set global Kafka hook
func SetGlobalHook(hk *Hook) {
	hook = hk
}

// GlobalHook return global Kafka hook
func GlobalHook() *Hook {
	return hook
}
