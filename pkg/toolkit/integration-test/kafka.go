package integrationtest

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	loader "github.com/consensys/orchestrate/handlers/loader/sarama"
	"github.com/consensys/orchestrate/handlers/offset"
	sarama2 "github.com/consensys/orchestrate/pkg/broker/sarama"
	"github.com/consensys/orchestrate/pkg/engine"
	"github.com/consensys/orchestrate/pkg/types/tx"
	"github.com/consensys/orchestrate/tests/utils/chanregistry"
	log "github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	broker       sarama.ConsumerGroup
	handler      *embeddingConsumerGroupHandler
	chanRegistry *chanregistry.ChanRegistry
	topics       []string
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewKafkaTestConsumer(ctx context.Context, groupID string, client sarama.Client, topics []string) (*KafkaConsumer, error) {
	broker, err := sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		return nil, err
	}

	chanRegistry := chanregistry.NewChanRegistry()
	engine.Init(ctx)

	engine.Register(loader.Loader)
	engine.Register(offset.Marker)
	engine.Register(msgHandler(chanRegistry))

	return &KafkaConsumer{
		broker:       broker,
		topics:       topics,
		chanRegistry: chanRegistry,
		handler: &embeddingConsumerGroupHandler{
			engine:  sarama2.NewEngineConsumerGroupHandler(engine.GlobalEngine()),
			isReady: make(chan bool, 1),
		},
	}, nil
}

func (c *KafkaConsumer) Start(ctx context.Context) error {
	cerr := make(chan error, 1)
	c.ctx, c.cancel = context.WithCancel(ctx)
	go func() {
		log.WithFields(log.Fields{
			"topics": c.topics,
		}).Info("connecting")

		err := c.broker.Consume(
			c.ctx,
			c.topics,
			c.handler,
		)

		if err != nil {
			log.WithError(err).Error("error on consumer")
		}

		cerr <- err
	}()

	select {
	case <-c.handler.isReady:
		return nil
	case err := <-cerr:
		return err
	}
}

func (c *KafkaConsumer) Stop(_ context.Context) error {
	c.cancel()
	return nil
}

func (c *KafkaConsumer) WaitForEnvelope(id, topic string, timeout time.Duration) (*tx.Envelope, error) {
	log.Debugf("waiting for envelope %s on topic %s. Timeout in %ds", id, topic, timeout/time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var ch = make(chan *tx.Envelope, 1)
	go func(chx chan *tx.Envelope) {
		msgKey := keyGenOf(id, topic)
		if !c.chanRegistry.HasChan(msgKey) {
			c.chanRegistry.Register(msgKey, make(chan *tx.Envelope, 1))
		}

		e := <-c.chanRegistry.GetChan(msgKey)
		chx <- e
	}(ch)

	select {
	case e := <-ch:
		log.Debugf("envelope %s found in topic %s", id, topic)
		return e, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-c.ctx.Done():
		return nil, c.ctx.Err()
	}
}

type embeddingConsumerGroupHandler struct {
	engine  *sarama2.EngineConsumerGroupHandler
	isReady chan bool
}

func (h *embeddingConsumerGroupHandler) Setup(s sarama.ConsumerGroupSession) error {
	err := h.engine.Setup(s)
	h.isReady <- true
	return err
}

func (h *embeddingConsumerGroupHandler) ConsumeClaim(s sarama.ConsumerGroupSession, c sarama.ConsumerGroupClaim) error {
	return h.engine.ConsumeClaim(s, c)
}

func (h *embeddingConsumerGroupHandler) Cleanup(s sarama.ConsumerGroupSession) error {
	return h.engine.Cleanup(s)
}

// Dispatcher dispatch envelopes into a channel registry
func msgHandler(reg *chanregistry.ChanRegistry) engine.HandlerFunc {
	return func(txctx *engine.TxContext) {
		if txctx.In == nil {
			panic("input message is nil")
		}

		txctx.Logger = txctx.Logger.WithFields(log.Fields{
			"id":        txctx.Envelope.GetID(),
			"msg.topic": txctx.In.Entrypoint(),
		})

		// Copy envelope before dispatching (it ensures that envelope can de manipulated in a concurrent safe manner once dispatched)
		envelope := *txctx.Envelope

		msgKey := keyGenOf(txctx.Envelope.GetID(), txctx.In.Entrypoint())
		if !reg.HasChan(msgKey) {
			reg.Register(msgKey, make(chan *tx.Envelope, 1))
		}

		// Dispatch envelope
		err := reg.Send(msgKey, &envelope)
		if err != nil {
			txctx.Logger.WithError(err).Error("dispatcher: envelope dispatched with errors")
		}

		txctx.Logger.Info("dispatcher: envelope dispatched")
	}
}

func keyGenOf(key, topic string) string {
	return fmt.Sprintf("%s/%s", topic, key)
}
