// +build unit

package schedules

import (
	"context"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/consensys/orchestrate/src/entities/testdata"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/src/api/store/parsers"
	"github.com/consensys/orchestrate/src/api/store/mocks"
)

func TestGetSchedule_Execute(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)
	mockScheduleDA := mocks.NewMockScheduleAgent(ctrl)
	mockJobDA := mocks.NewMockJobAgent(ctrl)

	userInfo := multitenancy.NewUserInfo("tenantOne", "username")
	usecase := NewGetScheduleUseCase(mockDB)

	t.Run("should execute use case successfully", func(t *testing.T) {
		scheduleEntity := testdata.FakeSchedule()
		scheduleModel := parsers.NewScheduleModelFromEntities(scheduleEntity)
		expectedResponse := parsers.NewScheduleEntityFromModels(scheduleModel)

		mockDB.EXPECT().Schedule().Return(mockScheduleDA).Times(1)
		mockDB.EXPECT().Job().Return(mockJobDA).Times(1)

		mockScheduleDA.EXPECT().
			FindOneByUUID(gomock.Any(), scheduleEntity.UUID, userInfo.AllowedTenants, userInfo.Username).
			Return(scheduleModel, nil)

		mockJobDA.EXPECT().
			FindOneByUUID(gomock.Any(), scheduleModel.Jobs[0].UUID, userInfo.AllowedTenants, userInfo.Username, false).
			Return(scheduleModel.Jobs[0], nil)

		scheduleResponse, err := usecase.Execute(ctx, scheduleEntity.UUID, userInfo)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, scheduleResponse)
	})

	t.Run("should fail with same error if FindOne fails for schedules", func(t *testing.T) {
		scheduleEntity := testdata.FakeSchedule()
		expectedErr := errors.NotFoundError("error")

		mockDB.EXPECT().Schedule().Return(mockScheduleDA)

		mockScheduleDA.EXPECT().FindOneByUUID(gomock.Any(), scheduleEntity.UUID, userInfo.AllowedTenants, userInfo.Username).Return(nil, expectedErr)

		scheduleResponse, err := usecase.Execute(ctx, scheduleEntity.UUID, userInfo)

		assert.Nil(t, scheduleResponse)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(createScheduleComponent), err)
	})

	t.Run("should fail with same error if FindOne fails for jobs", func(t *testing.T) {
		scheduleEntity := testdata.FakeSchedule()
		scheduleModel := parsers.NewScheduleModelFromEntities(scheduleEntity)
		expectedErr := errors.NotFoundError("error")

		mockDB.EXPECT().Schedule().Return(mockScheduleDA)
		mockDB.EXPECT().Job().Return(mockJobDA)

		mockScheduleDA.EXPECT().
			FindOneByUUID(gomock.Any(), scheduleEntity.UUID, userInfo.AllowedTenants, userInfo.Username).
			Return(scheduleModel, nil)
		mockJobDA.EXPECT().
			FindOneByUUID(gomock.Any(), scheduleModel.Jobs[0].UUID, userInfo.AllowedTenants, userInfo.Username, false).
			Return(nil, expectedErr)

		scheduleResponse, err := usecase.Execute(ctx, scheduleEntity.UUID, userInfo)

		assert.Nil(t, scheduleResponse)
		assert.Equal(t, errors.FromError(expectedErr).ExtendComponent(createScheduleComponent), err)
	})
}
