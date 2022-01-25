// +build unit

package transactions

import (
	"context"
	"fmt"
	"testing"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/types/testutils"
	mocks2 "github.com/consensys/orchestrate/services/api/business/use-cases/mocks"

	"github.com/consensys/orchestrate/services/api/store/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCallOffTx_Execute(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)
	mockTransactionRequestDA := mocks.NewMockTransactionRequestAgent(ctrl)
	retryJobUC := mocks2.NewMockRetryJobTxUseCase(ctrl)
	getTxUC := mocks2.NewMockGetTxUseCase(ctrl)
	gasIncrement := 0.1
	mockDB.EXPECT().TransactionRequest().Return(mockTransactionRequestDA).AnyTimes()

	userInfo := multitenancy.NewUserInfo("tenantOne", "username")
	usecase := NewCallOffTxUseCase(mockDB, getTxUC, retryJobUC)

	t.Run("should execute successfully", func(t *testing.T) {
		txRequest := testutils.FakeTxRequest()
		job := txRequest.Schedule.Jobs[0] 
		
		getTxUC.EXPECT().Execute(gomock.Any(), txRequest.Schedule.UUID, userInfo).Times(2).Return(txRequest, nil)
		retryJobUC.EXPECT().Execute(gomock.Any(), job.UUID, gasIncrement, nil, userInfo).Return(nil)
		_, err := usecase.Execute(ctx, txRequest.Schedule.UUID, userInfo)
		assert.NoError(t, err)
	})
	
	t.Run("should fail to if it is a private transaction", func(t *testing.T) {
		txRequest := testutils.FakeTxRequest()
		txRequest.Params.Protocol = entities.TesseraChainType
		
		getTxUC.EXPECT().Execute(gomock.Any(), txRequest.Schedule.UUID, userInfo).Return(txRequest, nil)
		_, err := usecase.Execute(ctx, txRequest.Schedule.UUID, userInfo)
		assert.Error(t, err)
		assert.True(t, errors.IsInvalidParameterError(err))
	})
	
	t.Run("should fail to execute if getTxUC fails", func(t *testing.T) {
		expectedErr := fmt.Errorf("err")
		txRequest := testutils.FakeTxRequest()
		
		getTxUC.EXPECT().Execute(gomock.Any(), txRequest.Schedule.UUID, userInfo).Return(nil, expectedErr)
		_, err := usecase.Execute(ctx, txRequest.Schedule.UUID, userInfo)
		assert.Error(t, err)
	})
	
	t.Run("should fail to execute if retryJobUC fails", func(t *testing.T) {
		expectedErr := fmt.Errorf("err")
		txRequest := testutils.FakeTxRequest()
		job := txRequest.Schedule.Jobs[0] 
		
		getTxUC.EXPECT().Execute(gomock.Any(), txRequest.Schedule.UUID, userInfo).Return(txRequest, nil)
		retryJobUC.EXPECT().Execute(gomock.Any(), job.UUID, gasIncrement, nil, userInfo).Return(expectedErr)
		_, err := usecase.Execute(ctx, txRequest.Schedule.UUID, userInfo)
		assert.Error(t, err)
	})
}
