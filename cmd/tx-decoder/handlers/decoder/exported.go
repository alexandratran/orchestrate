package decoder

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/engine"
	registryclient "gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/services/contract-registry/client"
)

var (
	component = "handler.decoder"
	handler   engine.HandlerFunc
	initOnce  = &sync.Once{}
)

// Init initialize Gas Estimator Handler
func Init(ctx context.Context) {
	initOnce.Do(func() {
		if handler != nil {
			return
		}

		// Initialize Registry Client
		registryclient.Init(ctx)

		// Create Handler
		handler = Decoder(registryclient.GlobalContractRegistryClient())

		log.Infof("decoder: handler ready")
	})
}

// SetGlobalHandler sets global Gas Estimator Handler
func SetGlobalHandler(h engine.HandlerFunc) {
	handler = h
}

// GlobalHandler returns global Gas Estimator handler
func GlobalHandler() engine.HandlerFunc {
	return handler
}