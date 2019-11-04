package txdecoder

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/handlers/producer"
	encoding "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/encoding/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/engine"
)

// PrepareMsg prepare message to produce from TxContexts
func PrepareMsg(txctx *engine.TxContext, msg *sarama.ProducerMessage) error {
	// Marshal Envelope into sarama Message
	err := encoding.Marshal(txctx.Envelope, msg)
	if err != nil {
		return err
	}

	// Set Topic to Nonce topic
	msg.Topic = viper.GetString("topic.tx.decoded")

	// Set key
	msg.Key = sarama.ByteEncoder(txctx.Envelope.GetChain().ID().Bytes())

	return nil
}

// Producer creates a producer handler that filters in Orchestrate transaction
// NB: If the transaction
func Producer(p sarama.SyncProducer) engine.HandlerFunc {

	classicProducer := producer.Producer(p, PrepareMsg)

	return func(txctx *engine.TxContext) {
		// Test if transaction was matched by Orchestrate.
		// TODO: Have an actual flag to make the check, because there is no guarantee
		// that metadata and tx will be unset in unmatched transaction forever.
		// TODO: Make it possible to filter at the last moment. So that we can produce in multiple topics
		if ExternalTxDisabled() && (txctx.Envelope.Metadata == nil || txctx.Envelope.Tx == nil) {

			// For robustness make sure that receipt and tx hash are set even though tx-listener already guarantee it
			if txctx.Envelope.Receipt != nil && txctx.Envelope.Receipt.TxHash != nil {
				txctx.Logger.WithFields(log.Fields{
					"tx.hash": txctx.Envelope.Receipt.GetTxHash(),
				}).Debugf("External tx disabled, skipping transaction with tx hash")
			}

			// Ignore the following handlers
			txctx.Abort()
			return
		}

		// Else normally call the producer as we would normally do
		classicProducer(txctx)
	}
}