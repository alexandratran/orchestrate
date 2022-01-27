package testdata

import (
	"github.com/consensys/orchestrate/src/entities"
)

func FakeLog() *entities.Log {
	return &entities.Log{
		Status:  entities.StatusCreated,
		Message: "job message",
	}
}
