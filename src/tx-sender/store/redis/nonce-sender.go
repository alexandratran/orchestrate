package redis

import (
	"time"

	"github.com/consensys/orchestrate/src/infra/redis"
)

const lastSentSuf = "last-sent"

type NonceSender struct {
	redis      redis.Client
	expiration int
}

// NewNonceSender creates a new mock NonceManager
func NewNonceSender(client redis.Client, expiration time.Duration) *NonceSender {
	return &NonceSender{
		redis:      client,
		expiration: int(expiration.Milliseconds()),
	}
}

func (ns *NonceSender) GetLastSent(key string) (uint64, error) {
	return ns.redis.LoadUint64(computeKey(key, lastSentSuf))
}

func (ns *NonceSender) SetLastSent(key string, value uint64) error {
	return ns.redis.Set(computeKey(key, lastSentSuf), ns.expiration, value)
}

func (ns *NonceSender) IncrLastSent(key string) error {
	return ns.redis.Incr(computeKey(key, lastSentSuf))
}

func (ns *NonceSender) DeleteLastSent(key string) error {
	return ns.redis.Delete(computeKey(key, lastSentSuf))
}
