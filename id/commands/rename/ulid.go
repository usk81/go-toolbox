package rename

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/usk81/go-toolbox/shared/cli"
)

var (
	ulidCmd = &cobra.Command{
		Use:   "ulid",
		Short: "rename files to ULID format",
		Long:  "rename files to ULID format",
		Run:   ulidCommand,
	}
	toLower bool
)

func init() {
	ulidCmd.Flags().BoolVarP(&toLower, "lower", "l", false, "convert lowercase")
	RootCmd.AddCommand(ulidCmd)
}

func ulidCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cli.Exit(fmt.Errorf("request arguments is not enough"))
	}
	if err := ulidAction(afero.NewOsFs(), args); err != nil {
		cli.Exit(err)
	}
}

func ulidAction(fs afero.Fs, ps []string) (err error) {
	for _, p := range ps {
		var exist bool
		exist, err = afero.Exists(fs, p)
		if !exist {
			return fmt.Errorf("file does not exist :%s", p)
		}
		if err != nil {
			return
		}
		if err = renameULID(fs, p); err != nil {
			return
		}
	}
	return
}

func renameULID(fs afero.Fs, p string) (err error) {
	dp, fn := filepath.Split(p)
	ext := filepath.Ext(fn)

	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		return
	}
	is := id.String()
	if toLower {
		is = strings.ToLower(is)
	}

	if err != nil {
		return
	}
	return fs.Rename(p, filepath.Join(dp, strings.Join([]string{is, ext}, ``)))
}
