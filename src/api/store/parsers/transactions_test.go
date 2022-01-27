// +build unit

package parsers

import (
	"github.com/consensys/orchestrate/src/entities/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	modelstestdata "github.com/consensys/orchestrate/src/api/store/models/testdata"

	"encoding/json"
)

func TestParsersTransaction_NewModelFromEntity(t *testing.T) {
	txEntity := testdata.FakeETHTransaction()
	txModel := NewTransactionModelFromEntities(txEntity)
	finalTxEntity := NewTransactionEntityFromModels(txModel)

	expectedJSON, _ := json.Marshal(txEntity)
	actualJOSN, _ := json.Marshal(finalTxEntity)
	assert.Equal(t, string(expectedJSON), string(actualJOSN))
}

func TestParsersTransaction_NewEntityFromModel(t *testing.T) {
	txModel := modelstestdata.FakeTransaction()
	txEntity := NewTransactionEntityFromModels(txModel)
	finalTxModel := NewTransactionModelFromEntities(txEntity)
	finalTxModel.UUID = txModel.UUID

	expectedJSON, _ := json.Marshal(txModel)
	actualJOSN, _ := json.Marshal(finalTxModel)
	assert.Equal(t, string(expectedJSON), string(actualJOSN))
}

func TestParsersTransaction_UpdateTransactionModel(t *testing.T) {
	txModel := modelstestdata.FakeTransaction()
	txEntity := testdata.FakeETHTransaction()
	UpdateTransactionModelFromEntities(txModel, txEntity)

	expectedTxModel := NewTransactionModelFromEntities(txEntity)
	expectedTxModel.UUID = txModel.UUID
	expectedTxModel.CreatedAt = txModel.CreatedAt
	expectedTxModel.UpdatedAt = txModel.UpdatedAt

	expectedJSON, _ := json.Marshal(txModel)
	actualJOSN, _ := json.Marshal(expectedTxModel)
	assert.Equal(t, string(expectedJSON), string(actualJOSN))
}
