package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/usk81/go-toolbox/id/commands/rename"
	"github.com/usk81/go-toolbox/shared/cli"
)

var (
	// RootCmd sets command config
	RootCmd = &cobra.Command{
		Use: "geneid",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage() // nolint
		},
	}
	runTimes int
)

func init() {
	RootCmd.AddCommand(rename.RootCmd)
}

// Run runs CLI action
func Run() {
	if err := RootCmd.Execute(); err != nil {
		cli.Exit(fmt.Errorf("failed to run: %s", err.Error()), 1)
	}
}
