package handlers

import (
	"context"

	"github.com/spf13/viper"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/cmd/tx-signer/handlers/producer"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/cmd/tx-signer/handlers/vault"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/handlers/opentracing"
)

type serviceName string

// Init inialize handlers
func Init(ctx context.Context) {
	common.InParallel(
		// Initialize Jaeger tracer
		func() {
			ctx = context.WithValue(ctx, serviceName("service-name"), viper.GetString("jaeger.service.name"))
			opentracing.Init(ctx)
		},
		// Initialize Vault
		func() { vault.Init(ctx) },
		// Initialize Producer
		func() { producer.Init(ctx) },
	)
}