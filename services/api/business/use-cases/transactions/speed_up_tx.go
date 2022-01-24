package transactions

import (
	"context"

	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/types/entities"
	usecases "github.com/consensys/orchestrate/services/api/business/use-cases"

	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/services/api/store"
)

const speedUpTxComponent = "use-cases.speed-up-tx"

// speedUpTxUseCase is a use case to get a transaction request
type speedUpTxUseCase struct {
	db      store.DB
	getTxUC usecases.GetTxUseCase
	logger  *log.Logger
}

// NewSpeedUpTxUseCase creates a new SpeedUpTxUseCase
func NewSpeedUpTxUseCase(db store.DB, getTxUC usecases.GetTxUseCase) usecases.SpeedUpTxUseCase {
	return &speedUpTxUseCase{
		db:      db,
		getTxUC: getTxUC,
		logger:  log.NewLogger().SetComponent(speedUpTxComponent),
	}
}

// Execute gets a transaction request
func (uc *speedUpTxUseCase) Execute(ctx context.Context, scheduleUUID string, gasIncrement float64, userInfo *multitenancy.UserInfo) (*entities.TxRequest, error) {
	return nil, nil
}
