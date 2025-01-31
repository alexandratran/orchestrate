// +build unit

package accounts

import (
	"context"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	parsers2 "github.com/consensys/orchestrate/src/api/store/parsers"
	models2 "github.com/consensys/orchestrate/src/api/store/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/src/entities"
	"github.com/consensys/orchestrate/src/api/store/mocks"
	"github.com/consensys/orchestrate/src/api/store/models/testdata"
)

func TestSearchAccounts_Execute(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)
	accountAgent := mocks.NewMockAccountAgent(ctrl)
	mockDB.EXPECT().Account().Return(accountAgent).AnyTimes()

	userInfo := multitenancy.NewUserInfo("tenantOne", "username")
	usecase := NewSearchAccountsUseCase(mockDB)

	t.Run("should execute use case successfully", func(t *testing.T) {
		acc := testdata.FakeAccountModel()

		filter := &entities.AccountFilters{
			Aliases: []string{"alias1"},
		}

		accountAgent.EXPECT().Search(gomock.Any(), filter, userInfo.AllowedTenants, userInfo.Username).Return([]*models2.Account{acc}, nil)

		resp, err := usecase.Execute(ctx, filter, userInfo)

		assert.NoError(t, err)
		assert.Equal(t, parsers2.NewAccountEntityFromModels(acc), resp[0])
	})

	t.Run("should fail with same error if search identities fails", func(t *testing.T) {
		expectedErr := errors.NotFoundError("error")

		filter := &entities.AccountFilters{
			Aliases: []string{"alias1"},
		}

		accountAgent.EXPECT().Search(gomock.Any(), filter, userInfo.AllowedTenants, userInfo.Username).Return(nil, expectedErr)

		_, err := usecase.Execute(ctx, filter, userInfo)

		assert.Error(t, err)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(createAccountComponent), err)
	})
}
