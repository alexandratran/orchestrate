package transactions

import (
	"context"

	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/types/entities"
	usecases "github.com/consensys/orchestrate/services/api/business/use-cases"

	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/services/api/store"
)

const callOffTxComponent = "use-cases.call-off-tx"

// callOffTxUseCase is a use case to get a transaction request
type callOffTxUseCase struct {
	db      store.DB
	getTxUC usecases.GetTxUseCase
	logger  *log.Logger
}

func NewCallOffTxUseCase(db store.DB, getTxUC usecases.GetTxUseCase) usecases.CallOffTxUseCase {
	return &callOffTxUseCase{
		db:      db,
		getTxUC: getTxUC,
		logger:  log.NewLogger().SetComponent(callOffTxComponent),
	}
}

// Execute gets a transaction request
func (uc *callOffTxUseCase) Execute(ctx context.Context, scheduleUUID string, userInfo *multitenancy.UserInfo) (*entities.TxRequest, error) {
	return nil, nil
}
