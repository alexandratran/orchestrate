// +build unit

package parsers

import (
	"testing"

	"github.com/consensys/orchestrate/src/entities/testdata"
	"github.com/stretchr/testify/assert"
)

func TestChainsParser(t *testing.T) {
	chain := testdata.FakeChain()
	chainModel := NewChainModelFromEntity(chain)
	finalChain := NewChainFromModel(chainModel)

	assert.Equal(t, chain, finalChain)
}
