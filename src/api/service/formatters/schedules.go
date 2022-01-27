package formatters

import (
	"github.com/consensys/orchestrate/src/api/service/types"
	"github.com/consensys/orchestrate/src/entities"
)

func FormatScheduleResponse(schedule *entities.Schedule) *types.ScheduleResponse {
	scheduleResponse := &types.ScheduleResponse{
		UUID:      schedule.UUID,
		TenantID:  schedule.TenantID,
		OwnerID:   schedule.OwnerID,
		CreatedAt: schedule.CreatedAt,
		Jobs:      []*types.JobResponse{},
	}

	for idx := range schedule.Jobs {
		scheduleResponse.Jobs = append(scheduleResponse.Jobs, FormatJobResponse(schedule.Jobs[idx]))
	}

	return scheduleResponse
}
