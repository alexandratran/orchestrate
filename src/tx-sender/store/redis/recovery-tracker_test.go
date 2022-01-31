// +build unit

package redis

import (
	"github.com/consensys/orchestrate/src/infra/redis/mocks"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoveryTracker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testKey := "recovery-tracker-redis"
	expectedKey := computeKey(testKey, recoverTrackerSuf)

	mockRedisClient := mocks.NewMockClient(ctrl)

	rt := NewNonceRecoveryTracker(mockRedisClient)

	t.Run("should call recovering nonce successfully", func(t *testing.T) {
		mockRedisClient.EXPECT().LoadUint64(expectedKey).Return(uint64(0), nil)

		n := rt.Recovering(testKey)
		assert.Equal(t, uint64(0), n)
	})

	t.Run("should call recover nonce successfully", func(t *testing.T) {
		mockRedisClient.EXPECT().Incr(expectedKey).Return(nil)

		rt.Recover(testKey)
	})

	t.Run("should call recovered nonce successfully", func(t *testing.T) {
		mockRedisClient.EXPECT().Delete(expectedKey).Return(nil)

		rt.Recovered(testKey)
	})
}
