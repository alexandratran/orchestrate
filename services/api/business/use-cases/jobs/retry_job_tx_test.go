// +build unit

package jobs

import (
	"context"
	"fmt"
	"testing"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/pkg/types/entities"
	"github.com/consensys/orchestrate/pkg/utils"
	mocks3 "github.com/consensys/orchestrate/services/api/business/use-cases/mocks"
	"github.com/consensys/orchestrate/services/api/store/mocks"
	"github.com/consensys/orchestrate/services/api/store/models/testutils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	tenantID = "tenant_id"
	username = "username"
)

func TestRetryJobTx_Execute(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockJobDA := mocks.NewMockJobAgent(ctrl)
	createJobUC := mocks3.NewMockCreateJobUseCase(ctrl)
	startJobUC := mocks3.NewMockStartJobUseCase(ctrl)
	mockDB := mocks.NewMockDB(ctrl)
	mockDB.EXPECT().Job().Return(mockJobDA).AnyTimes()

	userInfo := multitenancy.NewUserInfo(tenantID, username)
	usecase := NewRetryJobTxUseCase(mockDB, createJobUC, startJobUC)

	t.Run("should execute successfully", func(t *testing.T) {
		job := testutils.FakeJobModel(1)
		job.Transaction.TxType = string(entities.DynamicFeeTxType)
		job.Transaction.GasTipCap = "10000000000"
		job.Status = entities.StatusPending
		gasIncrement := 0.1
		nextJobUUID := "uuid"
		mockJobDA.EXPECT().FindOneByUUID(gomock.Any(), job.UUID, userInfo.AllowedTenants, userInfo.Username, false).Return(job, nil)
		createJobUC.EXPECT().Execute(gomock.Any(), gomock.Any(), userInfo).DoAndReturn(func(ctx context.Context, nextJob *entities.Job, ui *multitenancy.UserInfo) (*entities.Job, error) {
			assert.Equal(t, job.UUID, nextJob.InternalData.ParentJobUUID)
			assert.Equal(t, "0x28fa6ae00", nextJob.Transaction.GasTipCap.String())
			assert.Empty(t, nextJob.Transaction.GasFeeCap)
			nextJob.UUID = nextJobUUID
			return nextJob, nil
		})
		startJobUC.EXPECT().Execute(gomock.Any(), nextJobUUID, userInfo)
		
		err := usecase.Execute(ctx, job.UUID, gasIncrement, nil, userInfo)
		assert.NoError(t, err)
	})
	
	t.Run("should execute for legacy tx successfully", func(t *testing.T) {
		job := testutils.FakeJobModel(1)
		job.Transaction.TxType = string(entities.LegacyTxType)
		job.Transaction.GasPrice = "10000000000"
		job.Status = entities.StatusPending
		gasIncrement := 0.2
		nextJobUUID := "uuid"
		nextJobTxData := utils.StringToHexBytes("0xac")
		mockJobDA.EXPECT().FindOneByUUID(gomock.Any(), job.UUID, userInfo.AllowedTenants, userInfo.Username, false).Return(job, nil)
		createJobUC.EXPECT().Execute(gomock.Any(), gomock.Any(), userInfo).DoAndReturn(func(ctx context.Context, nextJob *entities.Job, ui *multitenancy.UserInfo) (*entities.Job, error) {
			assert.Equal(t, job.UUID, nextJob.InternalData.ParentJobUUID)
			assert.Equal(t, nextJobTxData, nextJob.Transaction.Data)
			assert.Equal(t, "0x2cb417800", nextJob.Transaction.GasPrice.String())
			nextJob.UUID = nextJobUUID
			return nextJob, nil
		})
		startJobUC.EXPECT().Execute(gomock.Any(), nextJobUUID, userInfo)
		
		err := usecase.Execute(ctx, job.UUID, gasIncrement, nextJobTxData, userInfo)
		assert.NoError(t, err)
	})
	
	t.Run("should fail to execute if status is not pending", func(t *testing.T) {
		job := testutils.FakeJobModel(1)
		job.Transaction.TxType = string(entities.LegacyTxType)
		job.Transaction.GasPrice = "10000000000"
		job.Status = entities.StatusCreated
		mockJobDA.EXPECT().FindOneByUUID(gomock.Any(), job.UUID, userInfo.AllowedTenants, userInfo.Username, false).Return(job, nil)
		err := usecase.Execute(ctx, job.UUID, 0.1, nil, userInfo)
		require.Error(t, err)
		assert.True(t, errors.IsInvalidStateError(err))
	})
	
	t.Run("should fail to execute it fails to get job from DB", func(t *testing.T) {
		job := testutils.FakeJobModel(1)
		expectedErr := fmt.Errorf("err")
		mockJobDA.EXPECT().FindOneByUUID(gomock.Any(), job.UUID, userInfo.AllowedTenants, userInfo.Username, false).Return(nil, expectedErr)
		err := usecase.Execute(ctx, job.UUID, 0.1, nil, userInfo)
		require.Error(t, err)
	})
	
	t.Run("should fail to execute it createJobUC fails", func(t *testing.T) {
		job := testutils.FakeJobModel(1)
		job.Status = entities.StatusPending
		job.Transaction.TxType = string(entities.LegacyTxType)
		job.Transaction.GasPrice = "10000000000"
		expectedErr := fmt.Errorf("err")
		mockJobDA.EXPECT().FindOneByUUID(gomock.Any(), job.UUID, userInfo.AllowedTenants, userInfo.Username, false).Return(job, nil)
		createJobUC.EXPECT().Execute(gomock.Any(), gomock.Any(), userInfo).Return(nil, expectedErr)
		err := usecase.Execute(ctx, job.UUID, 0.1, nil, userInfo)
		require.Error(t, err)
	})
}
