package session

import (
	"context"

	"github.com/consensys/orchestrate/src/tx-listener/dynamic"
)

type Session interface {
	Run(ctx context.Context) error
}

type Builder interface {
	NewSession(chain *dynamic.Chain) (Session, error)
}

type SManager interface {
	Run(ctx context.Context) error
}
