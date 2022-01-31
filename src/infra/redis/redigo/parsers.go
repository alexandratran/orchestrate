package redigo

import (
	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/gomodule/redigo/redis"
)

func parseRedisError(err error) error {
	switch {
	case err == redis.ErrNil:
		return errors.NotFoundError(err.Error())
	default:
		return errors.RedisConnectionError(err.Error())
	}
}
