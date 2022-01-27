package parsers

import (
	"github.com/consensys/orchestrate/src/api/store/models"
	"github.com/consensys/orchestrate/src/entities"
)

func NewLogEntityFromModels(logModel *models.Log) *entities.Log {
	return &entities.Log{
		Status:    logModel.Status,
		Message:   logModel.Message,
		CreatedAt: logModel.CreatedAt,
	}
}

func NewLogModelFromEntity(log *entities.Log) *models.Log {
	return &models.Log{
		Status:    log.Status,
		Message:   log.Message,
		CreatedAt: log.CreatedAt,
	}
}
