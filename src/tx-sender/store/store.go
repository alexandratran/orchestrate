package store

//go:generate mockgen -source=store.go -destination=mock/store.go -package=mock

type NonceSender interface {
	// GetLastSent retrieves last sent nonce
	GetLastSent(key string) (uint64, error)

	// IncrLastSent increment last sent nonce
	IncrLastSent(key string) error

	// DeleteLastSent last sent nonce
	DeleteLastSent(key string) error

	// SetLastSent sets last sent nonce
	SetLastSent(key string, value uint64) error
}

type RecoveryTracker interface {
	Recovering(key string) (count uint64)
	Recover(key string)
	Recovered(key string)
}
