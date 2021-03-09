GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*" | grep -v pkg/http/handler/dashboard/genstatic/gen.go | grep -v pkg/http/handler/swagger/genstatic/gen.go | egrep -v "^\./\.go" | grep -v _test.go)
PACKAGES ?= $(shell go list ./... | grep -Fv -e e2e -e examples -e genstatic -e mock )
INTEGRATION_TEST_PACKAGES ?= $(shell go list ./... | grep integration-tests )
CMD_RUN = tx-crafter tx-signer tx-sender tx-listener contract-registry chain-registry transaction-scheduler
CMD_PERSISTENT = redis postgres-chain-registry postgres-contract-registry postgres-transaction-scheduler jaeger
CMD_KAFKA = zookeeper kafka
CMD_MIGRATE = contract-registry chain-registry transaction-scheduler
DEPS_VAULT = vault vault-init vault-agent

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OPEN = xdg-open
endif
ifeq ($(UNAME_S),Darwin)
	OPEN = open
endif

.PHONY: all run-coverage coverage fmt fmt-check vet lint misspell-check misspell race tools help

# Linters
run-coverage: ## Generate global code coverage report
	@sh scripts/coverage.sh $(PACKAGES)

coverage: postgres run-coverage down-postgres
	@$(OPEN) build/coverage/coverage.html 2>/dev/null

race: ## Run data race detector
	@go test -count=1 -race -tags unit -short ${PACKAGES}

run-integration:
	@go test -v -tags integration ${INTEGRATION_TEST_PACKAGES}

mod-tidy: ## Run deps cleanup
	@go mod tidy

lint: ## Run linter to fix issues
	@misspell -w $(GOFILES)
	@golangci-lint run --fix

lint-ci: ## Check linting
	@misspell -error $(GOFILES)
	@golangci-lint run

run-e2e: gobuild-e2e
	@docker-compose up -V e2e

run-stress: gobuild-e2e
	@docker-compose up -V stress

e2e: run-e2e
	@docker-compose -f scripts/report/docker-compose.yml up --build
	@$(OPEN) build/report/report.html 2>/dev/null
	@exit $(docker inspect orchestrategit_e2e_1 --format='{{.State.ExitCode}}')

e2e-ci:
	@docker-compose up e2e
	@docker-compose -f scripts/report/docker-compose.yml up
	@exit $(docker inspect orchestrate_e2e_1 --format='{{.State.ExitCode}}')

stress: run-stress
	@exit $(docker inspect orchestrategit_stress_1 --format='{{.State.ExitCode}}')

stress-ci:
	@docker-compose up stress
	@exit $(docker inspect orchestrate_stress_1 --format='{{.State.ExitCode}}')

clean: protobuf gen-swagger gen-mocks mod-tidy lint coverage ## Run all clean-up tasks

gen-mocks:
	@go generate -run mockgen ./...

gen-swagger:
	@go generate gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/service/controllers
	@go generate gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/service/controllers

gen-deepcopy:
	@bash scripts/deepcopy/generate.sh

# Tools
lint-tools: ## Install linting tools
	@GO111MODULE=on go get github.com/client9/misspell/cmd/misspell@v0.3.4
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.27.0

tools: lint-tools ## Install test tools
	@GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.3
	@GO111MODULE=on go get github.com/swaggo/swag/cmd/swag@v1.6.7

# Help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gen-help: gobuild ## Generate Command Help file
	@mkdir -p build/cmd
	@./build/bin/orchestrate help tx-crafter | grep -A 9999 "Global Flags:" | head -n -2 > build/cmd/global.txt
	@for cmd in $(CMD_RUN); do \
		./build/bin/orchestrate help $$cmd run | grep -B 9999 "Global Flags:" | tail -n +3 | head -n -2 > build/cmd/$$cmd-run.txt; \
	done
	@for cmd in $(CMD_MIGRATE); do \
		./build/bin/orchestrate help $$cmd migrate | grep -B 9999 "Global Flags:" | tail -n +3 | head -n -2 > build/cmd/$$cmd-migrate.txt; \
	done

gen-help-docker: docker-build ## Generate Command Help file using docker
	@mkdir -p build/cmd
	@docker run orchestrate help tx-crafter | grep -A 9999 "Global Flags:" | head -n -3 > build/cmd/global.txt
	@for cmd in $(CMD_RUN); do \
		docker run orchestrate help $$cmd run | grep -B 9999 "Global Flags:" | tail -n +3 | head -n -3 > build/cmd/$$cmd-run.txt; \
	done
	@for cmd in $(CMD_MIGRATE); do \
		docker run orchestrate help $$cmd migrate | grep -B 9999 "Global Flags:" | tail -n +3 | head -n -3 > build/cmd/$$cmd-migrate.txt; \
	done

# Protobuf
protobuf: ## Generate protobuf stubs
	@docker-compose -f scripts/protobuf/docker-compose.yml up --build

topics: ## Create kafka topics
	@bash scripts/deps/kafka/initTopics.sh

gobuild: ## Build Orchestrate Go binary
	@GOOS=linux GOARCH=amd64 go build -i -o ./build/bin/orchestrate

docker-build: ## Build Orchestrate Docker image
	@DOCKER_BUILDKIT=1 docker build -t orchestrate .

bootstrap: ## Wait for dependencies to be ready
	@bash scripts/bootstrap.sh

bootstrap-deps: bootstrap ## Wait for dependencies to be ready
	@bash scripts/bootstrap-deps.sh

gobuild-e2e: ## Build Orchestrate e2e Docker image
	@GOOS=linux GOARCH=amd64 go build -i -o ./build/bin/test ./tests/cmd

orchestrate: gobuild ## Start Orchestrate
	@docker-compose up -d $(CMD_RUN)

ci-orchestrate:
	@docker-compose up -d $(CMD_RUN)

stop-orchestrate: ## Stop Orchestrate
	@docker-compose stop $(CMD_RUN)

down-orchestrate:## Down Orchestrate
	@docker-compose down --volumes --timeout 0

deps-persistent:
	@docker-compose -f scripts/deps/docker-compose.yml up -d $(CMD_PERSISTENT)

deps-vault:
	@docker-compose -f scripts/deps/docker-compose.yml up -d $(DEPS_VAULT)

deps: deps-vault deps-persistent
	@docker-compose -f scripts/deps/docker-compose.yml up -d $(CMD_KAFKA)

down-deps:
	@docker-compose -f scripts/deps/docker-compose.yml down --volumes --timeout 0

geth:
	@docker-compose -f scripts/geth/docker-compose.yml up -d

stop-geth:
	@docker-compose -f scripts/geth/docker-compose.yml stop

down-geth:
	@docker-compose -f scripts/geth/docker-compose.yml down  --volumes --timeout 0

quorum:
	@docker-compose -f scripts/quorum/docker-compose.yml up -d

stop-quorum:
	@docker-compose -f scripts/quorum/docker-compose.yml stop

down-quorum:
	@docker-compose -f scripts/quorum/docker-compose.yml down --volumes --timeout 0

besu:
	@docker-compose -f scripts/besu/docker-compose.yml up -d

stop-besu:
	@docker-compose -f scripts/besu/docker-compose.yml stop

down-besu:
	@docker-compose -f scripts/besu/docker-compose.yml down --volumes --timeout 0

postgres:
	@docker-compose -f scripts/deps/docker-compose.yml up -d postgres-unit

down-postgres:
	@docker-compose -f scripts/deps/docker-compose.yml rm --force -s -v postgres-unit

up: deps geth besu quorum bootstrap-deps topics orchestrate ## Start Orchestrate and deps

dev: deps orchestrate ## Start Orchestrate and light deps

geth-dev: deps geth orchestrate ## Start Orchestrate and light deps

besu-dev: deps besu orchestrate ## Start Orchestrate and light besu deps

quorum-dev: deps quorum orchestrate ## Start Orchestrate and light quorum deps

remote-dev: deps-persistent orchestrate

down: down-orchestrate down-quorum down-geth down-besu down-deps  ## Down Orchestrate and deps

up-ci: deps geth besu quorum bootstrap-deps ci-orchestrate ## Start Orchestrate and deps

up-azure: deps-persistent geth besu quorum bootstrap orchestrate ## Start Blockchain and Orchestrate to be connect to Azure Event Hub

hashicorp-accounts:
	@bash scripts/deps/config/hashicorp/vault.sh kv list secret/default

hashicorp-token-lookup:
	@bash scripts/deps/config/hashicorp/vault.sh token lookup

hashicorp-vault:
	@bash scripts/deps/config/hashicorp/vault.sh $(COMMAND)

pgadmin:
	@docker-compose -f scripts/deps/docker-compose-tools.yml up -d pgadmin

down-pgadmin:
	@docker-compose -f scripts/deps/docker-compose-tools.yml rm --force -s -v pgadmin

efk:
	@docker-compose -f scripts/deps/docker-compose-efk.yml up -d

down-efk:
	@docker-compose -f scripts/deps/docker-compose-efk.yml down --volumes --timeout 0

ifeq (restart,$(firstword $(MAKECMDGOALS)))
  CMD_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  ifeq ($(CMD_ARGS),)
  	CMD_ARGS := $(CMD_RUN)
  endif
  $(eval $(CMD_ARGS):;@:)
endif

restart: gobuild
	@docker-compose stop $(CMD_ARGS) && docker-compose start $(CMD_ARGS)

redisinsight:
	@docker-compose -f scripts/deps/docker-compose-tools.yml up -d redisinsight

down-redisinsight:
	@docker-compose -f scripts/deps/docker-compose-tools.yml rm --force -s -v redisinsight

up-all: efk pgadmin redisinsight up

down-all: down-efk down-pgadmin down-redisinsight down

observability:
	@docker-compose -f scripts/deps/docker-compose-tools.yml up -d prometheus grafana

down-observability:
	@docker-compose -f scripts/deps/docker-compose-tools.yml rm --force -s -v prometheus grafana

nginx:
	@docker-compose -f scripts/deps/docker-compose-tools.yml up -d nginx nginx-prometheus-exporter

down-nginx:
	@docker-compose -f scripts/deps/docker-compose-tools.yml rm --force -s -v nginx nginx-prometheus-exporter

vegeta:
	@mkdir -p build/vegeta
	@cat scripts/vegeta/test | vegeta attack -format=http -duration=90s -rate=150/s | tee build/vegeta/results.bin | vegeta report
	@vegeta report -type=json build/vegeta/results.bin > build/vegeta/metrics.json
	@cat build/vegeta/results.bin | vegeta plot > build/vegeta/plot.html
	@cat build/vegeta/results.bin | vegeta report -type="hist[0,100ms,200ms,300ms,500ms]"
	@$(OPEN) build/vegeta/plot.html 2>/dev/null
