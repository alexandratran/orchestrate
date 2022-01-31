package nonce

import (
	"context"

	"github.com/consensys/orchestrate/src/entities"
)

//go:generate mockgen -source=manager.go -destination=mocks/manager.go -package=mocks

type Manager interface {
	GetNonce(ctx context.Context, job *entities.Job) (uint64, error)
	CleanNonce(ctx context.Context, job *entities.Job, jobErr error) error
	IncrementNonce(ctx context.Context, job *entities.Job) error
}
