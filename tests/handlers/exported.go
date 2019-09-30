package handlers

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/common"
	registryClient "gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/services/contract-registry/client"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/tests/handlers/dispatcher"
)

// Init handlers
func Init(ctx context.Context) {
	common.InParallel(
		func() {
			dispatcher.Init(ctx)
		},
		// Initialize the registryClient
		func() {
			registryClient.Init(ctx)
		},
	)
}