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

const searchAccountsComponent = "use-cases.search-accounts"

type searchAccountsUseCase struct {
	db     store.DB
	logger *log.Logger
}

func NewSearchAccountsUseCase(db store.DB) usecases.SearchAccountsUseCase {
	return &searchAccountsUseCase{
		db:     db,
		logger: log.NewLogger().SetComponent(searchAccountsComponent),
	}
}

func (uc *searchAccountsUseCase) Execute(ctx context.Context, filters *entities.AccountFilters, userInfo *multitenancy.UserInfo) ([]*entities.Account, error) {
	models, err := uc.db.Account().Search(ctx, filters, userInfo.AllowedTenants, userInfo.Username)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(searchAccountsComponent)
	}

	var resp []*entities.Account
	for _, model := range models {
		iden := parsers2.NewAccountEntityFromModels(model)
		resp = append(resp, iden)
	}

	uc.logger.WithContext(ctx).Debug("accounts found successfully")
	return resp, nil
}
