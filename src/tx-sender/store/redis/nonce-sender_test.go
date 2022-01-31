// +build unit

package redis

import (
	"github.com/consensys/orchestrate/src/infra/redis/mocks"
	"github.com/golang/mock/gomock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNonceSender(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testKey := "nonce-sender-redis"
	expectedKey := computeKey(testKey, lastSentSuf)
	expiration := 100 * time.Millisecond

	mockRedisClient := mocks.NewMockClient(ctrl)

	ns := NewNonceSender(mockRedisClient, expiration)

	t.Run("should set nonce successfully ", func(t *testing.T) {
		mockRedisClient.EXPECT().Set(expectedKey, 100, uint64(10)).Return(nil)

		err := ns.SetLastSent(testKey, 10)
		assert.NoError(t, err)
	})

	t.Run("should get last sent successfully ", func(t *testing.T) {
		expectedValue := uint64(10)
		mockRedisClient.EXPECT().LoadUint64(expectedKey).Return(expectedValue, nil)

		n, err := ns.GetLastSent(testKey)
		assert.NoError(t, err)
		assert.Equal(t, expectedValue, n)
	})

	t.Run("should increment nonce successfully ", func(t *testing.T) {
		mockRedisClient.EXPECT().Incr(expectedKey).Return(nil)

		err := ns.IncrLastSent(testKey)
		assert.NoError(t, err)
	})

	t.Run("should delete nonce successfully ", func(t *testing.T) {
		mockRedisClient.EXPECT().Delete(expectedKey).Return(nil)

		err := ns.DeleteLastSent(testKey)
		assert.NoError(t, err)
	})
}
