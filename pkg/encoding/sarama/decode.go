package sarama

import (
	encoding "github.com/consensys/orchestrate/pkg/encoding/proto"
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/src/infra/broker/sarama"
	"google.golang.org/protobuf/proto"
)

// Unmarshal a sarama message into a protobuffer
func Unmarshal(msg *sarama.Msg, pb proto.Message) error {
	// Unmarshal Sarama message to Envelope
	err := encoding.Unmarshal(msg.Value(), pb)
	if err != nil {
		return errors.FromError(err).SetComponent(component)
	}
	return nil
}
