package types

import (
	"time"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/ethereum"
)

type Job struct {
	UUID         string
	ChainUUID    string
	ScheduleUUID string
	Type         string
	Labels       map[string]string
	Transaction  *ETHTransaction
	Receipt      *ethereum.Receipt
	Logs         []*Log
	CreatedAt    time.Time
}

// GetStatus Computes the status of a Job by checking its logs
func (job *Job) GetStatus() string {
	var status string
	var logCreatedAt *time.Time
	for idx := range job.Logs {
		if logCreatedAt == nil || job.Logs[idx].CreatedAt.After(*logCreatedAt) {
			status = job.Logs[idx].Status
			logCreatedAt = &job.Logs[idx].CreatedAt
		}
	}

	return status
}