package redigo

import (
	"github.com/consensys/orchestrate/pkg/toolkit/app/log"
	"github.com/gomodule/redigo/redis"
)

type Client struct {
	pool   *redis.Pool
	logger *log.Logger
}

func New(cfg *Config) (*Client, error) {
	redisOptions, err := cfg.ToRedisOptions()
	if err != nil {
		return nil, err
	}

	pool := &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		IdleTimeout: cfg.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.URL(), redisOptions...)
		},
	}

	return &Client{
		pool:   pool,
		logger: log.NewLogger().WithField("url", cfg.URL()),
	}, nil
}

func (nm *Client) LoadUint64(key string) (uint64, error) {
	conn := nm.pool.Get()
	defer closeConn(conn)

	reply, err := conn.Do("GET", key)
	if err != nil {
		return 0, parseRedisError(err)
	}

	value, err := redis.Uint64(reply, nil)
	if err != nil {
		return 0, parseRedisError(err)
	}

	return value, nil
}

func (nm *Client) Set(key string, expiration int, value interface{}) error {
	conn := nm.pool.Get()
	defer closeConn(conn)

	// Set value with expiration
	_, err := conn.Do("PSETEX", key, expiration, value)
	if err != nil {
		return parseRedisError(err)
	}

	return nil
}

func (nm *Client) Delete(key string) error {
	conn := nm.pool.Get()
	defer closeConn(conn)

	// Delete value
	_, err := conn.Do("DEL", key)
	if err != nil {
		return parseRedisError(err)
	}

	return nil
}

func (nm *Client) Incr(key string) error {
	conn := nm.pool.Get()
	defer closeConn(conn)

	_, err := conn.Do("INCR", key)
	if err != nil {
		return parseRedisError(err)
	}

	return nil
}

func (nm *Client) Ping() error {
	conn := nm.pool.Get()
	defer closeConn(conn)

	_, err := conn.Do("PING")
	if err != nil {
		return parseRedisError(err)
	}

	return nil
}

func closeConn(conn redis.Conn) {
	// There is nothing we can do if the connection fails to close
	_ = conn.Close()
}
