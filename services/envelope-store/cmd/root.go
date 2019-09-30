package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/database/postgres"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/handlers/logger"
)

// NewCommand create root command
func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:              "app",
		TraverseChildren: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// This is executed before each run (included on children command run)
			logger.InitLogger()
		},
	}

	// Set Persistent flags
	logger.LogLevel(rootCmd.PersistentFlags())
	logger.LogFormat(rootCmd.PersistentFlags())
	postgres.PGFlags(rootCmd.PersistentFlags())

	// Add Run command
	rootCmd.AddCommand(newRunCommand())

	// Add Migrate command
	rootCmd.AddCommand(mewMigrateCmd())

	return rootCmd
}