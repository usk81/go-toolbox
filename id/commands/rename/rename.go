package rename

import "github.com/spf13/cobra"

var (
	// RootCmd sets command config
	RootCmd = &cobra.Command{
		Use:   "rename",
		Short: "rename files",
		Long:  "rename files",
	}
)
