package steps

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/auth/jwt/generator"
	broker "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/broker/sarama"
	ethclient "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethclient/rpc"
	chainregistry "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/client"
	contractregistry "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/client"
	noncememory "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/nonce/memory"
	txscheduler "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/client"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/tests/service/e2e/cucumber/alias"
)

// Init initialize handlers
func Init(ctx context.Context) {
	broker.InitSyncProducer(ctx)
	generator.Init(ctx)
	chainregistry.Init(ctx)
	alias.Init(ctx)
	contractregistry.Init(ctx)
	txscheduler.Init()
	noncememory.Init(ctx)
	ethclient.Init(ctx)
}