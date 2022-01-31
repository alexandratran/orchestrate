package redis

//go:generate mockgen -source=redis.go -destination=mocks/redis.go -package=mocks

type Client interface {
	LoadUint64(key string) (uint64, error)
	Set(key string, expiration int, value interface{}) error
	Delete(key string) error
	Incr(key string) error
	Ping() error
}
