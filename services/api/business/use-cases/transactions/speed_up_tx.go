package transactions

import (
	"context"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/types/entities"
	usecases "github.com/consensys/orchestrate/services/api/business/use-cases"
	"github.com/consensys/orchestrate/services/api/store"
)

const speedUpTxComponent = "use-cases.speed-up-tx"

// speedUpTxUseCase is a use case to get a transaction request
type speedUpTxUseCase struct {
	db           store.DB
	getTxUC      usecases.GetTxUseCase
	retryJobTxUC usecases.RetryJobTxUseCase
	logger       *log.Logger
}

func NewSpeedUpTxUseCase(db store.DB, getTxUC usecases.GetTxUseCase, retryJobTxUC usecases.RetryJobTxUseCase) usecases.SpeedUpTxUseCase {
	return &speedUpTxUseCase{
		db:           db,
		getTxUC:      getTxUC,
		retryJobTxUC: retryJobTxUC,
		logger:       log.NewLogger().SetComponent(speedUpTxComponent),
	}
}

// Execute gets a transaction request
func (uc *speedUpTxUseCase) Execute(ctx context.Context, scheduleUUID string, gasIncrement float64, userInfo *multitenancy.UserInfo) (*entities.TxRequest, error) {
	ctx = log.WithFields(
		ctx,
		log.Field("schedule", scheduleUUID),
	)
	logger := uc.logger.WithContext(ctx)
	logger.Debug("speeding up transaction")

	tx, err := uc.getTxUC.Execute(ctx, scheduleUUID, userInfo)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(sendTxComponent)
	}

	if tx.Params.Protocol != "" {
		errMsg := "speed up transaction is not supported"
		logger.Error(errMsg)
		return nil, errors.FeatureNotSupportedError(errMsg).ExtendComponent(sendTxComponent)
	}

	job := tx.Schedule.Jobs[len(tx.Schedule.Jobs)-1]
	err = uc.retryJobTxUC.Execute(ctx, job.UUID, gasIncrement, job.Transaction.Data, userInfo)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(sendTxComponent)
	}

	txRequest, err := uc.getTxUC.Execute(ctx, scheduleUUID, userInfo)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(sendTxComponent)
	}

	logger.WithField("schedule", txRequest.Schedule.UUID).Info("speed-up transaction was sent successfully")
	return txRequest, nil
}
