package dataagents

import (
	"github.com/consensys/orchestrate/src/api/store"
	pg "github.com/consensys/orchestrate/src/infra/database/postgres"
)

type PGAgents struct {
	tx               store.TransactionAgent
	job              store.JobAgent
	log              store.LogAgent
	schedule         store.ScheduleAgent
	txRequest        store.TransactionRequestAgent
	account          store.AccountAgent
	faucet           store.FaucetAgent
	artifact         store.ArtifactAgent
	codeHash         store.CodeHashAgent
	event            store.EventAgent
	repository       store.RepositoryAgent
	tag              store.TagAgent
	contract         store.ContractAgent
	chain            store.ChainAgent
	privateTxManager store.PrivateTxManagerAgent
}

func New(db pg.DB) *PGAgents {
	return &PGAgents{
		tx:               NewPGTransaction(db),
		job:              NewPGJob(db),
		log:              NewPGLog(db),
		schedule:         NewPGSchedule(db),
		txRequest:        NewPGTransactionRequest(db),
		account:          NewPGAccount(db),
		faucet:           NewPGFaucet(db),
		artifact:         NewPGArtifact(db),
		codeHash:         NewPGCodeHash(db),
		event:            NewPGEvent(db),
		repository:       NewPGRepository(db),
		tag:              NewPGTag(db),
		contract:         NewPGContract(db),
		chain:            NewPGChain(db),
		privateTxManager: NewPGPrivateTxManager(db),
	}
}

func (a *PGAgents) Job() store.JobAgent {
	return a.job
}

func (a *PGAgents) Log() store.LogAgent {
	return a.log
}

func (a *PGAgents) Schedule() store.ScheduleAgent {
	return a.schedule
}

func (a *PGAgents) Transaction() store.TransactionAgent {
	return a.tx
}

func (a *PGAgents) TransactionRequest() store.TransactionRequestAgent {
	return a.txRequest
}

func (a *PGAgents) Account() store.AccountAgent {
	return a.account
}

func (a *PGAgents) Faucet() store.FaucetAgent {
	return a.faucet
}

func (a *PGAgents) Artifact() store.ArtifactAgent {
	return a.artifact
}

func (a *PGAgents) CodeHash() store.CodeHashAgent {
	return a.codeHash
}

func (a *PGAgents) Event() store.EventAgent {
	return a.event
}

func (a *PGAgents) Repository() store.RepositoryAgent {
	return a.repository
}

func (a *PGAgents) Tag() store.TagAgent {
	return a.tag
}

func (a *PGAgents) Contract() store.ContractAgent {
	return a.contract
}

func (a *PGAgents) Chain() store.ChainAgent {
	return a.chain
}

func (a *PGAgents) PrivateTxManager() store.PrivateTxManagerAgent {
	return a.privateTxManager
}
