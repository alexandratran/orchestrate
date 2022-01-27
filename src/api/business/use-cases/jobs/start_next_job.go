package jobs

import (
	"context"
	"fmt"
	"strconv"

	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	usecases "github.com/consensys/orchestrate/src/api/business/use-cases"
	"github.com/consensys/orchestrate/src/api/store/models"
	"github.com/consensys/orchestrate/src/api/store/parsers"
	"github.com/consensys/orchestrate/src/entities"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/src/api/store"
)

const startNextJobComponent = "use-cases.next-job-start"

type startNextJobUseCase struct {
	db              store.DB
	startJobUseCase usecases.StartJobUseCase
	logger          *log.Logger
}

func NewStartNextJobUseCase(db store.DB, startJobUC usecases.StartJobUseCase) usecases.StartNextJobUseCase {
	return &startNextJobUseCase{
		db:              db,
		startJobUseCase: startJobUC,
		logger:          log.NewLogger().SetComponent(startNextJobComponent),
	}
}

// Execute gets a job
func (uc *startNextJobUseCase) Execute(ctx context.Context, jobUUID string, userInfo *multitenancy.UserInfo) error {
	ctx = log.WithFields(ctx, log.Field("job", jobUUID))
	logger := uc.logger.WithContext(ctx).WithField("job", jobUUID)
	logger.Debug("starting job")
	jobModel, err := uc.db.Job().FindOneByUUID(ctx, jobUUID, userInfo.AllowedTenants, userInfo.Username, false)
	if err != nil {
		return errors.FromError(err).ExtendComponent(startNextJobComponent)
	}

	if jobModel.NextJobUUID == "" {
		errMsg := fmt.Sprintf("job %s does not have a next job to start", jobModel.NextJobUUID)
		logger.Error(errMsg)
		return errors.DataError(errMsg)
	}

	logger = logger.WithField("next_job", jobModel.NextJobUUID)
	logger.Debug("start next job use-case")

	nextJobModel, err := uc.db.Job().FindOneByUUID(ctx, jobModel.NextJobUUID, userInfo.AllowedTenants, userInfo.Username, false)
	if err != nil {
		return errors.FromError(err).ExtendComponent(startNextJobComponent)
	}

	switch nextJobModel.Type {
	case entities.EEAMarkingTransaction:
		err = uc.handleEEAMarkingTx(ctx, jobModel, nextJobModel)
	case entities.TesseraMarkingTransaction:
		err = uc.handleTesseraMarkingTx(ctx, jobModel, nextJobModel)
	}

	if err != nil {
		logger.WithError(err).Error("failed to validate next transaction data")
		return errors.FromError(err).ExtendComponent(startNextJobComponent)
	}

	return uc.startJobUseCase.Execute(ctx, nextJobModel.UUID, userInfo)
}

func (uc *startNextJobUseCase) handleEEAMarkingTx(ctx context.Context, prevJobModel, jobModel *models.Job) error {
	if prevJobModel.Type != entities.EEAPrivateTransaction {
		return errors.DataError("expected previous job as type: %s", entities.EEAPrivateTransaction)
	}

	prevJobEntity := parsers.NewJobEntityFromModels(prevJobModel)
	if prevJobEntity.Status != entities.StatusStored {
		return errors.DataError("expected previous job status as: STORED")
	}

	jobModel.Transaction.Data = prevJobModel.Transaction.Hash
	return uc.db.Transaction().Update(ctx, jobModel.Transaction)
}

func (uc *startNextJobUseCase) handleTesseraMarkingTx(ctx context.Context, prevJobModel, jobModel *models.Job) error {
	if prevJobModel.Type != entities.TesseraPrivateTransaction {
		return errors.DataError("expected previous job as type: %s", entities.TesseraPrivateTransaction)
	}

	prevJobEntity := parsers.NewJobEntityFromModels(prevJobModel)
	if prevJobEntity.Status != entities.StatusStored {
		return errors.DataError("expected previous job status as: STORED")
	}

	jobModel.Transaction.Data = prevJobModel.Transaction.EnclaveKey
	gas, err := strconv.ParseInt(prevJobModel.Transaction.Gas, 10, 64)
	if err == nil && gas < entities.TesseraGasLimit {
		jobModel.Transaction.Gas = strconv.Itoa(entities.TesseraGasLimit)
	} else {
		jobModel.Transaction.Gas = prevJobModel.Transaction.Gas
	}

	return uc.db.Transaction().Update(ctx, jobModel.Transaction)
}
