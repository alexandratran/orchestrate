package jobs

import (
	"context"

	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/utils/envelope"
	usecases "github.com/consensys/orchestrate/src/api/business/use-cases"
	"github.com/consensys/orchestrate/src/entities"

	"github.com/Shopify/sarama"
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/src/api/store"
	"github.com/consensys/orchestrate/src/api/store/parsers"
	pkgsarama "github.com/consensys/orchestrate/src/infra/broker/sarama"
)

const resendJobTxComponent = "use-cases.resend-job-tx"

type resendJobTxUseCase struct {
	db            store.DB
	kafkaProducer sarama.SyncProducer
	topicsCfg     *pkgsarama.KafkaTopicConfig
	logger        *log.Logger
}

func NewResendJobTxUseCase(db store.DB, kafkaProducer sarama.SyncProducer, topicsCfg *pkgsarama.KafkaTopicConfig) usecases.ResendJobTxUseCase {
	return &resendJobTxUseCase{
		db:            db,
		kafkaProducer: kafkaProducer,
		topicsCfg:     topicsCfg,
		logger:        log.NewLogger().SetComponent(resendJobTxComponent),
	}
}

// Execute sends a job to the Kafka topic
func (uc *resendJobTxUseCase) Execute(ctx context.Context, jobUUID string, userInfo *multitenancy.UserInfo) error {
	ctx = log.WithFields(ctx, log.Field("job", jobUUID))
	logger := uc.logger.WithContext(ctx)
	logger.Debug("resending job transaction")

	jobModel, err := uc.db.Job().FindOneByUUID(ctx, jobUUID, userInfo.AllowedTenants, userInfo.Username, false)
	if err != nil {
		return errors.FromError(err).ExtendComponent(resendJobTxComponent)
	}

	jobModel.InternalData.ParentJobUUID = jobUUID
	jobEntity := parsers.NewJobEntityFromModels(jobModel)
	if jobEntity.Status != entities.StatusPending {
		errMessage := "cannot resend job transaction at the current status"
		logger.WithField("status", jobEntity.Status).Error(errMessage)
		return errors.InvalidStateError(errMessage)
	}

	partition, offset, err := envelope.SendJobMessage(jobEntity, uc.kafkaProducer, uc.topicsCfg.Sender)
	if err != nil {
		logger.WithError(err).Error("failed to send job message")
		return errors.FromError(err).ExtendComponent(resendJobTxComponent)
	}

	logger.WithField("partition", partition).WithField("offset", offset).Info("job resend successfully")
	return nil
}
