package cmd

import (
	"github.com/spf13/cobra"

	"github.com/thockin/go-build-template/pkg/myapp"
)

func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "TODO",
		Long: `TODO
TODO
TODO`,
		SilenceUsage: false,
		RunE:         rootRun,
	}
	return rootCmd
}

func rootRun(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	return myapp.Foo(cmd.Context())
}
