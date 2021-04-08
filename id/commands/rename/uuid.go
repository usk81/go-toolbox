package rename

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/usk81/go-toolbox/shared/cli"
)

var (
	uuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "rename files to UUID format",
		Long:  "rename files to UUID format",
		Run:   uuidCommand,
	}
)

func init() {
	RootCmd.AddCommand(uuidCmd)
}

func uuidCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cli.Exit(fmt.Errorf("request arguments is not enough"))
	}
	if err := uuidAction(afero.NewOsFs(), args); err != nil {
		cli.Exit(err)
	}
}

func uuidAction(fs afero.Fs, ps []string) (err error) {
	for _, p := range ps {
		var exist bool
		exist, err = afero.Exists(fs, p)
		if !exist {
			return fmt.Errorf("file does not exist :%s", p)
		}
		if err != nil {
			return
		}
		if err = renameUUID(fs, p); err != nil {
			return
		}
	}
	return
}

func renameUUID(fs afero.Fs, p string) (err error) {
	dp, fn := filepath.Split(p)
	ext := filepath.Ext(fn)
	u, err := uuid.NewRandom()
	if err != nil {
		return
	}
	return fs.Rename(p, filepath.Join(dp, strings.Join([]string{u.String(), ext}, ``)))
}
