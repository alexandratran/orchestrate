package redis

import (
	"github.com/consensys/orchestrate/src/infra/redis"
)

const recoverTrackerSuf = "recover-tracker"

type NonceRecoveryTracker struct {
	redisCli redis.Client
}

func NewNonceRecoveryTracker(redisCli redis.Client) *NonceRecoveryTracker {
	return &NonceRecoveryTracker{
		redisCli: redisCli,
	}
}

func (t *NonceRecoveryTracker) Recovering(key string) uint64 {
	v, err := t.redisCli.LoadUint64(computeKey(key, recoverTrackerSuf))
	if err != nil {
		return 0
	}

	return v
}

func (t *NonceRecoveryTracker) Recover(key string) {
	_ = t.redisCli.Incr(computeKey(key, recoverTrackerSuf))
}

func (t *NonceRecoveryTracker) Recovered(key string) {
	_ = t.redisCli.Delete(computeKey(key, recoverTrackerSuf))
}
