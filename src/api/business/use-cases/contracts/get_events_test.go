// +build unit

package contracts

import (
	"context"
	"testing"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/utils"
	"github.com/consensys/orchestrate/src/api/store/mocks"
	"github.com/consensys/orchestrate/src/api/store/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetEvents_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	sigHash := utils.StringToHexBytes("0x0123")
	indexedInputCount := uint32(1)
	eventModel := &models.EventModel{
		ABI: "eventABI",
	}

	eventAgent := mocks.NewMockEventAgent(ctrl)
	usecase := NewGetEventsUseCase(eventAgent)

	t.Run("should execute use case successfully if event is found", func(t *testing.T) {
		eventAgent.EXPECT().
			FindOneByAccountAndSigHash(gomock.Any(), chainID, contractAddress.Hex(), sigHash.String(), indexedInputCount).
			Return(eventModel, nil)

		responseABI, eventsABI, err := usecase.Execute(ctx, chainID, contractAddress, sigHash, indexedInputCount)

		assert.Equal(t, responseABI, eventModel.ABI)
		assert.Nil(t, eventsABI)
		assert.NoError(t, err)
	})

	t.Run("should fail if data agent returns connection error", func(t *testing.T) {
		pgError := errors.PostgresConnectionError("error")
		eventAgent.EXPECT().
			FindOneByAccountAndSigHash(gomock.Any(), chainID, contractAddress.Hex(), sigHash.String(), indexedInputCount).
			Return(nil, pgError)

		responseABI, eventsABI, err := usecase.Execute(ctx, chainID, contractAddress, sigHash, indexedInputCount)

		assert.Equal(t, errors.FromError(pgError).ExtendComponent(getEventsComponent), err)
		assert.Empty(t, responseABI)
		assert.Nil(t, eventsABI)
	})

	t.Run("should execute use case successfully if event is not found", func(t *testing.T) {
		eventAgent.EXPECT().
			FindOneByAccountAndSigHash(gomock.Any(), chainID, contractAddress.Hex(), sigHash.String(), indexedInputCount).
			Return(nil, nil)

		eventAgent.EXPECT().
			FindDefaultBySigHash(gomock.Any(), sigHash.String(), indexedInputCount).
			Return([]*models.EventModel{eventModel, eventModel}, nil)

		responseABI, eventsABI, err := usecase.Execute(ctx, chainID, contractAddress, sigHash, indexedInputCount)

		assert.Equal(t, eventsABI, []string{eventModel.ABI, eventModel.ABI})
		assert.Empty(t, responseABI)
		assert.NoError(t, err)
	})

	t.Run("should fail if data agent returns error on find default", func(t *testing.T) {
		pgError := errors.PostgresConnectionError("error")
		eventAgent.EXPECT().
			FindOneByAccountAndSigHash(gomock.Any(), chainID, contractAddress.Hex(), sigHash.String(), indexedInputCount).
			Return(nil, nil)
		eventAgent.EXPECT().FindDefaultBySigHash(gomock.Any(), sigHash.String(), indexedInputCount).
			Return(nil, pgError)

		responseABI, eventsABI, err := usecase.Execute(ctx, chainID, contractAddress, sigHash, indexedInputCount)

		assert.Equal(t, errors.FromError(pgError).ExtendComponent(getEventsComponent), err)
		assert.Empty(t, responseABI)
		assert.Nil(t, eventsABI)
	})
}
