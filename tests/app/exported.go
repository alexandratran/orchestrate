package app

import (
	"context"
	"fmt"
	"sync"

	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/common"

	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/handlers/loader"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/handlers/logger"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/handlers/offset"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	broker "gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/broker/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/engine"
	server "gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/http"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/http/healthcheck"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/tests/handlers"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/tests/handlers/dispatcher"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/tests/service/cucumber"
)

var (
	app         *common.App
	readyToTest chan bool
	startOnce   = &sync.Once{}
)

func init() {
	// Create app
	app = common.NewApp()

	// Set Kafka Group value
	viper.Set("kafka.group", "group-e2e")
}

func startServer(ctx context.Context) {
	// Initialize server
	server.Init(ctx)

	// Register Healthcheck
	server.Enhance(healthcheck.HealthCheck(app))

	// Start Listening
	_ = server.ListenAndServe()
}

func initComponents(ctx context.Context) {
	common.InParallel(
		// Initialize Engine
		func() {
			engine.Init(ctx)
		},
		// Initialize ConsumerGroup
		func() {
			broker.InitConsumerGroup(ctx)
		},
		// Initialize Handlers
		func() {
			handlers.Init(ctx)
		},
		// Initialize cucumber registry
		func() {
			cucumber.Init(ctx)
		},
	)
}

func registerHandlers() {
	// Generic handlers on every worker
	engine.Register(logger.Logger)
	engine.Register(loader.Loader)
	engine.Register(offset.Marker)
	engine.Register(dispatcher.GlobalHandler())
}

// Start starts application
func Start(ctx context.Context) {
	startOnce.Do(func() {

		cancelCtx, cancel := context.WithCancel(ctx)
		go func() {
			// Start Server
			startServer(ctx)
			cancel()
		}()

		// Initialize ConsumerGroup
		initComponents(cancelCtx)

		// Register all Handlers
		registerHandlers()

		// Indicate that application is ready
		app.SetReady(true)

		// Start consuming on every topics
		// Initialize Topics list by chain
		topics := []string{
			viper.GetString("kafka.topic.crafter"),
			viper.GetString("kafka.topic.nonce"),
			viper.GetString("kafka.topic.signer"),
			viper.GetString("kafka.topic.sender"),
			viper.GetString("kafka.topic.decoded"),
			viper.GetString("kafka.topic.wallet.generator"),
			viper.GetString("kafka.topic.wallet.generated"),
		}
		if primary := viper.GetString("cucumber.chainid.primary"); primary != "" {
			topics = append(topics, fmt.Sprintf("%s-%s", viper.GetString("kafka.topic.decoder"), primary))
		}
		if secondary := viper.GetString("cucumber.chainid.secondary"); secondary != "" {
			topics = append(topics, fmt.Sprintf("%s-%s", viper.GetString("kafka.topic.decoder"), secondary))
		}

		readyToTest = make(chan bool, 1)

		go func() {
			<-readyToTest
			cucumber.Run(cancel, cucumber.GlobalOptions())
		}()

		cg := &EmbeddingConsumerGroupHandler{
			engine: broker.NewEngineConsumerGroupHandler(engine.GlobalEngine()),
		}

		log.Debugf("worker: start consuming on %q", topics)
		err := broker.Consume(
			cancelCtx,
			topics,
			cg,
		)
		if err != nil {
			log.WithError(err).Fatal("worker: error on consumer")
		}

	})
}

type EmbeddingConsumerGroupHandler struct {
	engine *broker.EngineConsumerGroupHandler
}

func (h *EmbeddingConsumerGroupHandler) Setup(s sarama.ConsumerGroupSession) error {
	err := h.engine.Setup(s)
	readyToTest <- true
	return err
}

func (h *EmbeddingConsumerGroupHandler) ConsumeClaim(s sarama.ConsumerGroupSession, c sarama.ConsumerGroupClaim) error {
	return h.engine.ConsumeClaim(s, c)
}

func (h *EmbeddingConsumerGroupHandler) Cleanup(s sarama.ConsumerGroupSession) error {
	return h.engine.Cleanup(s)
}