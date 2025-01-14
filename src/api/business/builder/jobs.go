package builder

import (
	"github.com/Shopify/sarama"
	pkgsarama "github.com/consensys/orchestrate/src/infra/broker/sarama"
	usecases "github.com/consensys/orchestrate/src/api/business/use-cases"
	"github.com/consensys/orchestrate/src/api/business/use-cases/jobs"
	"github.com/consensys/orchestrate/src/api/metrics"
	"github.com/consensys/orchestrate/src/api/store"
)

type jobUseCases struct {
	createJob   usecases.CreateJobUseCase
	getJob      usecases.GetJobUseCase
	startJob    usecases.StartJobUseCase
	resendJobTx usecases.ResendJobTxUseCase
	retryJobTx  usecases.RetryJobTxUseCase
	updateJob   usecases.UpdateJobUseCase
	searchJobs  usecases.SearchJobsUseCase
}

func newJobUseCases(
	db store.DB,
	appMetrics metrics.TransactionSchedulerMetrics,
	producer sarama.SyncProducer,
	topicsCfg *pkgsarama.KafkaTopicConfig,
	getChainUC usecases.GetChainUseCase,
	qkmStoreID string,
) *jobUseCases {
	startJobUC := jobs.NewStartJobUseCase(db, producer, topicsCfg, appMetrics)
	updateChildrenUC := jobs.NewUpdateChildrenUseCase(db)
	startNextJobUC := jobs.NewStartNextJobUseCase(db, startJobUC)
	createJobUC := jobs.NewCreateJobUseCase(db, getChainUC, qkmStoreID)

	return &jobUseCases{
		createJob:   createJobUC,
		getJob:      jobs.NewGetJobUseCase(db),
		searchJobs:  jobs.NewSearchJobsUseCase(db),
		updateJob:   jobs.NewUpdateJobUseCase(db, updateChildrenUC, startNextJobUC, appMetrics),
		startJob:    startJobUC,
		resendJobTx: jobs.NewResendJobTxUseCase(db, producer, topicsCfg),
		retryJobTx:  jobs.NewRetryJobTxUseCase(db, createJobUC, startJobUC),
	}
}

func (u *jobUseCases) CreateJob() usecases.CreateJobUseCase {
	return u.createJob
}

func (u *jobUseCases) GetJob() usecases.GetJobUseCase {
	return u.getJob
}

func (u *jobUseCases) SearchJobs() usecases.SearchJobsUseCase {
	return u.searchJobs
}

func (u *jobUseCases) StartJob() usecases.StartJobUseCase {
	return u.startJob
}

func (u *jobUseCases) ResendJobTx() usecases.ResendJobTxUseCase {
	return u.resendJobTx
}

func (u *jobUseCases) RetryJobTx() usecases.RetryJobTxUseCase {
	return u.retryJobTx
}

func (u *jobUseCases) UpdateJob() usecases.UpdateJobUseCase {
	return u.updateJob
}
