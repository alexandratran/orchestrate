package accounts

import (
	"context"

	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	usecases "github.com/consensys/orchestrate/src/api/business/use-cases"
	parsers2 "github.com/consensys/orchestrate/src/api/store/parsers"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/src/api/store"
	"github.com/consensys/orchestrate/src/entities"
)

const updateAccountComponent = "use-cases.update-account"

type updateAccountUseCase struct {
	db     store.DB
	logger *log.Logger
}

func NewUpdateAccountUseCase(db store.DB) usecases.UpdateAccountUseCase {
	return &updateAccountUseCase{
		db:     db,
		logger: log.NewLogger().SetComponent(updateAccountComponent),
	}
}

func (uc *updateAccountUseCase) Execute(ctx context.Context, account *entities.Account, userInfo *multitenancy.UserInfo) (*entities.Account, error) {
	ctx = log.WithFields(ctx, log.Field("address", account.Address))
	logger := uc.logger.WithContext(ctx)

	model, err := uc.db.Account().FindOneByAddress(ctx, account.Address.Hex(), userInfo.AllowedTenants, userInfo.Username)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(updateAccountComponent)
	}

	if account.Attributes != nil {
		model.Attributes = account.Attributes
	}
	if account.Alias != "" {
		model.Alias = account.Alias
	}
	if account.StoreID != "" {
		model.StoreID = account.StoreID
	}

	err = uc.db.Account().Update(ctx, model)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(updateAccountComponent)
	}

	resp := parsers2.NewAccountEntityFromModels(model)

	logger.Info("account updated successfully")
	return resp, nil
}
