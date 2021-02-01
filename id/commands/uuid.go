package commands

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/usk81/go-toolbox/shared/cli"
)

var (
	uuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "generate UUIDs",
		Long:  "generate UUIDs. generated UUID are based on RFC 4122 and DCE 1.1: Authentication and Security Services.",
		Run:   uuidCommand,
	}
)

func init() {
	uuidCmd.Flags().IntVarP(&runTimes, "times", "t", 1, "times to generate UUIDs")
	RootCmd.AddCommand(uuidCmd)
}

func uuidCommand(cmd *cobra.Command, args []string) {
	if runTimes <= 0 {
		cli.Exit(fmt.Errorf("`times` must be greater than 0"))
	}

	ids := make([]string, runTimes)
	for n := 0; n < runTimes; n++ {
		i, err := uuidAction()
		if err != nil {
			cli.Exit(fmt.Errorf("fail to generate: %w", err), 1)
		}
		ids[n] = i
	}
	for _, v := range ids {
		fmt.Println(v)
	}
}

func uuidAction() (result string, err error) {
	i, err := uuid.NewRandom()
	if err != nil {
		return
	}
	result = i.String()
	return
}
