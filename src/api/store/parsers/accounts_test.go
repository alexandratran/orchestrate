// +build unit

package parsers

import (
	"testing"

	"github.com/consensys/orchestrate/src/entities/testdata"
	"github.com/stretchr/testify/assert"
)

func TestAccountsParser(t *testing.T) {
	account := testdata.FakeAccount()
	accountModel := NewAccountModelFromEntities(account)
	finalAccount := NewAccountEntityFromModels(accountModel)

	assert.Equal(t, account, finalAccount)
}
