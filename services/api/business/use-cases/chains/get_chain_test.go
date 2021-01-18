package chains

import (
	"context"
	"testing"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/multitenancy"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/parsers"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/mocks"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/models/testutils"
)

func TestGetChain_Execute(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)
	chainAgent := mocks.NewMockChainAgent(ctrl)

	mockDB.EXPECT().Chain().Return(chainAgent).AnyTimes()

	usecase := NewGetChainUseCase(mockDB)

	tenantID := multitenancy.DefaultTenant
	tenants := []string{tenantID}

	t.Run("should execute use case successfully", func(t *testing.T) {
		chainModel := testutils.FakeChainModel()

		chainAgent.EXPECT().FindOneByUUID(ctx, chainModel.UUID, tenants).Return(chainModel, nil)

		resp, err := usecase.Execute(ctx, chainModel.UUID, tenants)

		assert.NoError(t, err)
		assert.Equal(t, parsers.NewChainFromModel(chainModel), resp)
	})

	t.Run("should fail with same error if get chain fails", func(t *testing.T) {
		expectedErr := errors.NotFoundError("error")

		chainAgent.EXPECT().FindOneByUUID(ctx, "uuid", tenants).Return(nil, expectedErr)

		resp, err := usecase.Execute(ctx, "uuid", tenants)

		assert.Nil(t, resp)
		assert.Error(t, err)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(getChainComponent), err)
	})
}