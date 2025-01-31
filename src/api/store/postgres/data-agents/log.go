package dataagents

import (
	"context"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/src/api/store"

	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/consensys/orchestrate/src/api/store/models"
	pg "github.com/consensys/orchestrate/src/infra/database/postgres"
	"github.com/gofrs/uuid"
)

const logDAComponent = "data-agents.log"

// PGLog is a log data agent for PostgreSQL
type PGLog struct {
	db     pg.DB
	logger *log.Logger
}

// NewPGLog creates a new PGLog
func NewPGLog(db pg.DB) store.LogAgent {
	return &PGLog{db: db, logger: log.NewLogger().SetComponent(logDAComponent)}
}

// Insert Inserts a new log in DB
func (agent *PGLog) Insert(ctx context.Context, logModel *models.Log) error {
	if logModel.UUID == "" {
		logModel.UUID = uuid.Must(uuid.NewV4()).String()
	}

	if logModel.JobID == nil && logModel.Job != nil {
		logModel.JobID = &logModel.Job.ID
	}

	err := pg.Insert(ctx, agent.db, logModel)
	if err != nil {
		agent.logger.WithError(err).Error("failed to insert job log")
		return errors.FromError(err).ExtendComponent(logDAComponent)
	}

	return nil
}
