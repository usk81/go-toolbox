package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/spf13/cobra"
	"github.com/usk81/go-toolbox/shared/cli"
)

var (
	ulidCmd = &cobra.Command{
		Use:   "ulid",
		Short: "generate ULIDs",
		Long:  "generate ULIDs. ULID is Universally Unique Lexicographically Sortable Identifier.",
		Run:   ulidCommand,
	}
	toLower bool
)

func init() {
	ulidCmd.Flags().IntVarP(&runTimes, "times", "t", 1, "times to generate ULIDs")
	ulidCmd.Flags().BoolVarP(&toLower, "lower", "l", false, "convert lowercase")
	RootCmd.AddCommand(ulidCmd)
}

func ulidCommand(cmd *cobra.Command, args []string) {
	if runTimes <= 0 {
		cli.Exit(fmt.Errorf("`times` must be greater than 0"))
	}

	ids := make([]string, runTimes)
	for n := 0; n < runTimes; n++ {
		i, err := ulidAction(toLower)
		if err != nil {
			cli.Exit(fmt.Errorf("fail to generate: %w", err), 1)
		}
		ids[n] = i
	}
	for _, v := range ids {
		fmt.Println(v)
	}
}

func ulidAction(toLower bool) (result string, err error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	i, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		return
	}
	result = i.String()
	if toLower {
		result = strings.ToLower(result)
	}
	return
}
