package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/ethereum/abi/registry"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/ethereum/abi/registry/redis"
	ethclient "gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/ethereum/ethclient/rpc"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/http"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/tracing/opentracing/jaeger"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/services/contract-registry/app"
)

func newRunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run application",
		Run:   run,
	}

	// Register OpenTracing flags
	jaeger.InitFlags(runCmd.Flags())

	// Register HTTP server flags
	http.Hostname(runCmd.Flags())

	// EthClient flag
	ethclient.URLs(runCmd.Flags())

	// ContractRegistry flag
	registry.ContractRegistryType(runCmd.Flags())
	registry.ABIs(runCmd.Flags())

	// Redis ContractRegistry flag
	redis.InitFlags(runCmd.Flags())

	return runCmd
}

func run(cmd *cobra.Command, args []string) {
	// Process signals
	sig := utils.NewSignalListener(func(signal os.Signal) { app.Close(context.Background()) })
	defer sig.Close()

	// Start application
	app.Start(context.Background())
}