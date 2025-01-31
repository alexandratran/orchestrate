// +build unit

package usecases

import (
	"context"
	"testing"

	"github.com/consensys/orchestrate/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/consensys/orchestrate/pkg/sdk/client/mock"
	"github.com/consensys/orchestrate/src/api/service/types"
	"github.com/consensys/orchestrate/src/entities"
	"github.com/consensys/orchestrate/src/entities/testdata"
	apitestdata "github.com/consensys/orchestrate/src/api/service/types/testdata"
)

func TestCreateChildJob_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	initialGasPrice := utils.BigIntStringToHex("1000000000")
	ctx := context.Background()

	mockClient := mock.NewMockOrchestrateClient(ctrl)

	usecase := NewRetrySessionJobUseCase(mockClient)

	t.Run("should do nothing if status of the job is not PENDING", func(t *testing.T) {
		parentJob := testdata.FakeJob()
		parentJobResponse := apitestdata.FakeJobResponse()

		mockClient.EXPECT().GetJob(gomock.Any(), parentJob.UUID).Return(parentJobResponse, nil)
		childJobUUID, err := usecase.Execute(ctx, parentJob.UUID, parentJob.UUID, 0)
		assert.NoError(t, err)
		assert.Empty(t, childJobUUID)
	})

	t.Run("should create a new child job if the parent job status is PENDING", func(t *testing.T) {
		parentJob := testdata.FakeJob()
		childJobResponse := apitestdata.FakeJobResponse()
		parentJobResponse := apitestdata.FakeJobResponse()
		parentJobResponse.Status = entities.StatusPending
		parentJobResponse.Transaction.GasPrice = initialGasPrice
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Increment = 0.1
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Limit = 0.2

		mockClient.EXPECT().GetJob(gomock.Any(), parentJob.UUID).Return(parentJobResponse, nil)
		mockClient.EXPECT().CreateJob(gomock.Any(), gomock.Any()).Return(childJobResponse, nil)
		mockClient.EXPECT().StartJob(gomock.Any(), childJobResponse.UUID).Return(nil)

		childJobUUID, err := usecase.Execute(ctx, parentJob.UUID, parentJob.UUID, 0)
		assert.NoError(t, err)
		assert.NotEmpty(t, childJobUUID)
	})

	t.Run("should resend job transaction if the parent job status is PENDING with not gas increment", func(t *testing.T) {
		parentJob := testdata.FakeJob()
		parentJobResponse := apitestdata.FakeJobResponse()
		parentJobResponse.Status = entities.StatusPending
		parentJobResponse.Transaction.GasPrice = initialGasPrice

		mockClient.EXPECT().GetJob(gomock.Any(), parentJob.UUID).Return(parentJobResponse, nil)
		mockClient.EXPECT().ResendJobTx(gomock.Any(), parentJob.UUID).Return(nil)

		childJobUUID, err := usecase.Execute(ctx, parentJob.UUID, parentJob.UUID, 0)
		assert.NoError(t, err)
		assert.Equal(t, childJobUUID, parentJobResponse.UUID)
	})

	t.Run("should resend job transaction last job if gas limit was reached", func(t *testing.T) {
		parentJob := testdata.FakeJob()
		childJob := testdata.FakeJob()
		parentJobResponse := apitestdata.FakeJobResponse()
		parentJobResponse.Status = entities.StatusPending
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Increment = 0.1
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Limit = 0.2
		parentJobResponse.Transaction.GasPrice = initialGasPrice

		mockClient.EXPECT().GetJob(gomock.Any(), parentJob.UUID).Return(parentJobResponse, nil)
		mockClient.EXPECT().ResendJobTx(gomock.Any(), childJob.UUID).Return(nil)

		childJobUUID, err := usecase.Execute(ctx, parentJob.UUID, childJob.UUID, 3)
		assert.NoError(t, err)
		assert.Equal(t, childJobUUID, parentJobResponse.UUID)
	})

	t.Run("should send the same job if job is a raw transaction", func(t *testing.T) {
		parentJob := testdata.FakeJob()
		parentJobResponse := apitestdata.FakeJobResponse()
		parentJobResponse.Transaction.Raw = utils.StringToHexBytes("0xAB")
		parentJobResponse.Type = entities.EthereumRawTransaction
		parentJobResponse.Status = entities.StatusPending
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Increment = 0.1
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Limit = 0.2

		mockClient.EXPECT().GetJob(gomock.Any(), parentJob.UUID).Return(parentJobResponse, nil)
		mockClient.EXPECT().ResendJobTx(gomock.Any(), parentJob.UUID).Return(nil)

		childJobUUID, err := usecase.Execute(ctx, parentJob.UUID, parentJob.UUID, 0)
		assert.NoError(t, err)
		assert.NotEmpty(t, childJobUUID)
	})

	t.Run("should create a new child job by increasing the gasPrice by Increment (legacyTx)", func(t *testing.T) {
		parentJob := testdata.FakeJob()
		childJob := testdata.FakeJob()
		childJobResponse := apitestdata.FakeJobResponse()

		parentJobResponse := apitestdata.FakeJobResponse()
		parentJobResponse.Status = entities.StatusPending
		parentJobResponse.Transaction.GasPrice = initialGasPrice
		parentJobResponse.Transaction.Nonce = utils.ToPtr(uint64(1)).(*uint64)
		parentJobResponse.Transaction.TransactionType = entities.LegacyTxType
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Increment = 0.06
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Limit = 0.12

		mockClient.EXPECT().GetJob(gomock.Any(), parentJob.UUID).Return(parentJobResponse, nil)
		mockClient.EXPECT().CreateJob(gomock.Any(), gomock.Any()).
			DoAndReturn(func(timeoutCtx context.Context, req *types.CreateJobRequest) (*types.JobResponse, error) {
				assert.Equal(t, "1120000000", req.Transaction.GasPrice.ToInt().String())
				assert.Equal(t, parentJobResponse.Transaction.Nonce, req.Transaction.Nonce)
				return childJobResponse, nil
			})
		mockClient.EXPECT().StartJob(gomock.Any(), childJobResponse.UUID).Return(nil)

		childJobUUID, err := usecase.Execute(ctx, parentJob.UUID, childJob.UUID, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, childJobUUID)
	})
	
	t.Run("should create a new child job by increasing the gasPrice by Increment (dynamicTx)", func(t *testing.T) {
		parentJob := testdata.FakeJob()
		childJob := testdata.FakeJob()
		childJobResponse := apitestdata.FakeJobResponse()

		parentJobResponse := apitestdata.FakeJobResponse()
		parentJobResponse.Status = entities.StatusPending
		parentJobResponse.Transaction.GasTipCap = initialGasPrice
		parentJobResponse.Transaction.Nonce = utils.ToPtr(uint64(1)).(*uint64)
		parentJobResponse.Transaction.TransactionType = entities.DynamicFeeTxType
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Increment = 0.06
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Limit = 0.12

		mockClient.EXPECT().GetJob(gomock.Any(), parentJob.UUID).Return(parentJobResponse, nil)
		mockClient.EXPECT().CreateJob(gomock.Any(), gomock.Any()).
			DoAndReturn(func(timeoutCtx context.Context, req *types.CreateJobRequest) (*types.JobResponse, error) {
				assert.Equal(t, "1120000000", req.Transaction.GasTipCap.ToInt().String())
				assert.Equal(t, parentJobResponse.Transaction.Nonce, req.Transaction.Nonce)
				return childJobResponse, nil
			})
		mockClient.EXPECT().StartJob(gomock.Any(), childJobResponse.UUID).Return(nil)

		childJobUUID, err := usecase.Execute(ctx, parentJob.UUID, childJob.UUID, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, childJobUUID)
	})

	t.Run("should create a new child job by increasing the gasPrice and not exceed the limit (legacyTx)", func(t *testing.T) {
		parentJob := testdata.FakeJob()
		childJob := testdata.FakeJob()
		childJobResponse := apitestdata.FakeJobResponse()

		parentJobResponse := apitestdata.FakeJobResponse()
		parentJobResponse.Status = entities.StatusPending
		parentJobResponse.Transaction.GasPrice = initialGasPrice
		parentJobResponse.Transaction.Nonce = utils.ToPtr(uint64(1)).(*uint64)
		parentJobResponse.Transaction.TransactionType = entities.LegacyTxType
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Increment = 0.06
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Limit = 0.05

		mockClient.EXPECT().GetJob(gomock.Any(), parentJob.UUID).Return(parentJobResponse, nil)
		mockClient.EXPECT().CreateJob(gomock.Any(), gomock.Any()).
			DoAndReturn(func(timeoutCtx context.Context, req *types.CreateJobRequest) (*types.JobResponse, error) {
				assert.Equal(t, "1050000000", req.Transaction.GasPrice.ToInt().String())
				assert.Equal(t, parentJobResponse.Transaction.Nonce, req.Transaction.Nonce)
				return childJobResponse, nil
			})
		mockClient.EXPECT().StartJob(gomock.Any(), childJobResponse.UUID).Return(nil)

		childJobUUID, err := usecase.Execute(ctx, parentJob.UUID, childJob.UUID, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, childJobUUID)
	})
	
	t.Run("should create a new child job by increasing the gasPrice and not exceed the limit (dynamicTx)", func(t *testing.T) {
		parentJob := testdata.FakeJob()
		childJob := testdata.FakeJob()
		childJobResponse := apitestdata.FakeJobResponse()

		parentJobResponse := apitestdata.FakeJobResponse()
		parentJobResponse.Status = entities.StatusPending
		parentJobResponse.Transaction.GasTipCap = initialGasPrice
		parentJobResponse.Transaction.Nonce = utils.ToPtr(uint64(1)).(*uint64)
		parentJobResponse.Transaction.TransactionType = entities.DynamicFeeTxType
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Increment = 0.06
		parentJobResponse.Annotations.GasPricePolicy.RetryPolicy.Limit = 0.05

		mockClient.EXPECT().GetJob(gomock.Any(), parentJob.UUID).Return(parentJobResponse, nil)
		mockClient.EXPECT().CreateJob(gomock.Any(), gomock.Any()).
			DoAndReturn(func(timeoutCtx context.Context, req *types.CreateJobRequest) (*types.JobResponse, error) {
				assert.Equal(t, "1050000000", req.Transaction.GasTipCap.ToInt().String())
				assert.Equal(t, parentJobResponse.Transaction.Nonce, req.Transaction.Nonce)
				return childJobResponse, nil
			})
		mockClient.EXPECT().StartJob(gomock.Any(), childJobResponse.UUID).Return(nil)

		childJobUUID, err := usecase.Execute(ctx, parentJob.UUID, childJob.UUID, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, childJobUUID)
	})
}
