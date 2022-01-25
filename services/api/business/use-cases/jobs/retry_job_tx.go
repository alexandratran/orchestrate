package jobs

import (
	"context"
	"math/big"

	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/utils"
	usecases "github.com/consensys/orchestrate/services/api/business/use-cases"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/services/api/business/parsers"
	"github.com/consensys/orchestrate/services/api/store"
)

const retryJobTxComponent = "use-cases.retry-job-tx"

type retryJobTxUseCase struct {
	db            store.DB
	createJobTxUC usecases.CreateJobUseCase
	startJobUC    usecases.StartJobUseCase
	logger        *log.Logger
}

func NewRetryJobTxUseCase(db store.DB, createJobTxUC usecases.CreateJobUseCase, startJobUC usecases.StartJobUseCase) usecases.RetryJobTxUseCase {
	return &retryJobTxUseCase{
		db:            db,
		createJobTxUC: createJobTxUC,
		startJobUC:    startJobUC,
		logger:        log.NewLogger().SetComponent(retryJobTxComponent),
	}
}

// Execute sends a job to the Kafka topic
func (uc *retryJobTxUseCase) Execute(ctx context.Context, jobUUID string, gasIncrement float64, txData hexutil.Bytes, userInfo *multitenancy.UserInfo) error {
	ctx = log.WithFields(ctx, log.Field("job", jobUUID))
	logger := uc.logger.WithContext(ctx)
	logger.Debug("retrying job transaction")

	jobModel, err := uc.db.Job().FindOneByUUID(ctx, jobUUID, userInfo.AllowedTenants, userInfo.Username, false)
	if err != nil {
		return errors.FromError(err).ExtendComponent(retryJobTxComponent)
	}

	job := parsers.NewJobEntityFromModels(jobModel)
	if job.Status != entities.StatusPending {
		errMessage := "cannot retry job transaction at the current status"
		logger.WithField("status", job.Status).Error(errMessage)
		return errors.InvalidStateError(errMessage)
	}

	job.InternalData.ParentJobUUID = jobUUID
	job.Transaction.Data = txData
	if job.Transaction.TransactionType == entities.LegacyTxType {
		gasPrice := job.Transaction.GasPrice.ToInt()
		txGasPrice := gasPrice.Mul(gasPrice, big.NewInt(10)).Div(gasPrice, big.NewInt(100))
		job.Transaction.GasPrice = utils.ToPtr(hexutil.Big(*txGasPrice)).(*hexutil.Big)
	} else {
		gasTipCap := job.Transaction.GasTipCap.ToInt()
		txGasTipCap := gasTipCap.Mul(gasTipCap, big.NewInt(10)).Div(gasTipCap, big.NewInt(100))
		job.Transaction.GasTipCap = utils.ToPtr(hexutil.Big(*txGasTipCap)).(*hexutil.Big)
	}

	retriedJob, err := uc.createJobTxUC.Execute(ctx, job, userInfo)
	if err != nil {
		return errors.FromError(err).ExtendComponent(retryJobTxComponent)
	}

	if err = uc.startJobUC.Execute(ctx, retriedJob.UUID, userInfo); err != nil {
		return errors.FromError(err).ExtendComponent(retryJobTxComponent)
	}

	logger.WithField("job", retriedJob.UUID).Info("job retried successfully")
	return nil
}
