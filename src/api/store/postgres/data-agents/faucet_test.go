// +build unit
// +build !race
// +build !integration

package dataagents

import (
	"context"
	"testing"

	"github.com/consensys/orchestrate/pkg/errors"
	"github.com/consensys/orchestrate/src/entities"
	"github.com/consensys/orchestrate/src/api/store/models/testdata"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"

	pgTestUtils "github.com/consensys/orchestrate/src/infra/database/postgres/testutils"
	"github.com/consensys/orchestrate/src/api/store/postgres/migrations"
	"github.com/stretchr/testify/suite"
)

type faucetTestSuite struct {
	suite.Suite
	agents         *PGAgents
	pg             *pgTestUtils.PGTestHelper
	tenantID       string
	allowedTenants []string
}

func TestPGFaucet(t *testing.T) {
	s := new(faucetTestSuite)
	suite.Run(t, s)
}

func (s *faucetTestSuite) SetupSuite() {
	s.pg, _ = pgTestUtils.NewPGTestHelper(nil, migrations.Collection)
	s.tenantID = "tenantID"
	s.allowedTenants = []string{s.tenantID, "_"}
	s.pg.InitTestDB(s.T())
}

func (s *faucetTestSuite) SetupTest() {
	s.pg.UpgradeTestDB(s.T())
	s.agents = New(s.pg.DB)
}

func (s *faucetTestSuite) TearDownTest() {
	s.pg.DowngradeTestDB(s.T())
}

func (s *faucetTestSuite) TearDownSuite() {
	s.pg.DropTestDB(s.T())
}

func (s *faucetTestSuite) TestPGFaucet_Insert() {
	ctx := context.Background()

	s.T().Run("should insert model successfully", func(t *testing.T) {
		faucet := testdata.FakeFaucetModel()
		faucet.TenantID = s.tenantID
		err := s.agents.Faucet().Insert(ctx, faucet)
		assert.NoError(s.T(), err)

		assert.NoError(t, err)
		assert.NotEmpty(t, faucet.UUID)
	})

	s.T().Run("should insert model without UUID successfully", func(t *testing.T) {
		faucet := testdata.FakeFaucetModel()
		faucet.TenantID = s.tenantID
		faucet.Name = "faucet-insert-2"
		faucet.UUID = ""
		err := s.agents.Faucet().Insert(ctx, faucet)
		assert.NoError(s.T(), err)

		assert.NoError(t, err)
		assert.NotEmpty(t, faucet.UUID)
	})
}

func (s *faucetTestSuite) TestPGFaucet_Update() {
	ctx := context.Background()
	faucet := testdata.FakeFaucetModel()
	faucet.TenantID = s.tenantID
	err := s.agents.Faucet().Insert(ctx, faucet)
	assert.NoError(s.T(), err)

	s.T().Run("should update model successfully", func(t *testing.T) {
		newFaucet := testdata.FakeFaucetModel()
		newFaucet.TenantID = s.tenantID
		newFaucet.UUID = faucet.UUID

		err = s.agents.Faucet().Update(ctx, newFaucet, s.allowedTenants)
		assert.NoError(t, err)

		faucetRetrieved, _ := s.agents.Faucet().FindOneByUUID(ctx, faucet.UUID, s.allowedTenants)
		assert.Equal(t, newFaucet.ChainRule, faucetRetrieved.ChainRule)
	})

	s.T().Run("should update model successfully with tenant", func(t *testing.T) {
		newFaucet := testdata.FakeFaucetModel()
		newFaucet.UUID = faucet.UUID

		err = s.agents.Faucet().Update(ctx, newFaucet, []string{faucet.TenantID})

		assert.NoError(t, err)

		faucetRetrieved, _ := s.agents.Faucet().FindOneByUUID(ctx, faucet.UUID, s.allowedTenants)
		assert.Equal(t, newFaucet.ChainRule, faucetRetrieved.ChainRule)
	})
}

func (s *faucetTestSuite) TestPGFaucet_FindOneByUUID() {
	ctx := context.Background()
	faucet := testdata.FakeFaucetModel()
	faucet.TenantID = s.tenantID
	err := s.agents.Faucet().Insert(ctx, faucet)
	assert.NoError(s.T(), err)

	s.T().Run("should get model successfully", func(t *testing.T) {
		faucetRetrieved, err := s.agents.Faucet().FindOneByUUID(ctx, faucet.UUID, s.allowedTenants)

		assert.NoError(t, err)
		assert.Equal(t, faucet.UUID, faucetRetrieved.UUID)
	})

	s.T().Run("should get model successfully as tenant", func(t *testing.T) {
		faucetRetrieved, err := s.agents.Faucet().FindOneByUUID(ctx, faucet.UUID, s.allowedTenants)

		assert.NoError(t, err)
		assert.NotEmpty(t, faucetRetrieved.UUID)
	})

	s.T().Run("should return NotFoundError if select fails", func(t *testing.T) {
		_, err := s.agents.Faucet().FindOneByUUID(ctx, "b6fe7a2a-1a4d-49ca-99d8-8a34aa495ef0", s.allowedTenants)
		assert.True(t, errors.IsNotFoundError(err))
	})
}

func (s *faucetTestSuite) TestPGFaucet_Search() {
	ctx := context.Background()

	faucet0 := testdata.FakeFaucetModel()
	faucet0.TenantID = s.tenantID
	err := s.agents.Faucet().Insert(ctx, faucet0)
	assert.NoError(s.T(), err)

	faucet1 := testdata.FakeFaucetModel()
	faucet1.TenantID = s.tenantID
	faucet1.Name = "faucet-mainnet-2"
	err = s.agents.Faucet().Insert(ctx, faucet1)
	assert.NoError(s.T(), err)

	s.T().Run("should find models successfully with filters", func(t *testing.T) {
		filters := &entities.FaucetFilters{
			Names:     []string{faucet0.Name},
			ChainRule: faucet0.ChainRule,
		}

		retrievedFaucets, err := s.agents.Faucet().Search(ctx, filters, s.allowedTenants)

		assert.NoError(t, err)
		assert.Equal(t, faucet0.UUID, retrievedFaucets[0].UUID)
		assert.Len(t, retrievedFaucets, 1)
	})

	s.T().Run("should not find any model by names", func(t *testing.T) {
		filters := &entities.FaucetFilters{
			Names:     []string{"0x3"},
			ChainRule: faucet0.ChainRule,
		}

		retrievedJobs, err := s.agents.Faucet().Search(ctx, filters, s.allowedTenants)

		assert.NoError(t, err)
		assert.Empty(t, retrievedJobs)
	})

	s.T().Run("should not find any model by chainRule", func(t *testing.T) {
		filters := &entities.FaucetFilters{
			Names:     []string{faucet0.Name},
			ChainRule: uuid.Must(uuid.NewV4()).String(),
		}

		retrievedJobs, err := s.agents.Faucet().Search(ctx, filters, s.allowedTenants)

		assert.NoError(t, err)
		assert.Empty(t, retrievedJobs)
	})

	s.T().Run("should find every inserted model successfully", func(t *testing.T) {
		filters := &entities.FaucetFilters{}
		retrievedJobs, err := s.agents.Faucet().Search(ctx, filters, s.allowedTenants)

		assert.NoError(t, err)
		assert.Equal(t, len(retrievedJobs), 2)
	})
}

func (s *faucetTestSuite) TestPGFaucet_ConnectionErr() {
	ctx := context.Background()

	// We drop the DB to make the test fail
	s.pg.DropTestDB(s.T())
	faucet := testdata.FakeFaucetModel()
	faucet.TenantID = s.tenantID
	s.T().Run("should return PostgresConnectionError if insert fails", func(t *testing.T) {
		err := s.agents.Faucet().Insert(ctx, faucet)
		assert.True(t, errors.IsInternalError(err))
	})

	s.T().Run("should return PostgresConnectionError if update fails", func(t *testing.T) {
		err := s.agents.Faucet().Update(ctx, faucet, s.allowedTenants)
		assert.True(t, errors.IsInternalError(err))
	})

	s.T().Run("should return PostgresConnectionError if findOne fails", func(t *testing.T) {
		_, err := s.agents.Faucet().FindOneByUUID(ctx, faucet.UUID, s.allowedTenants)
		assert.True(t, errors.IsInternalError(err))
	})

	// We bring it back up
	s.pg.InitTestDB(s.T())
}
